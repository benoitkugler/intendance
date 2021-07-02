package controller

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

// ProduitsPossibles indique les produits disponibles pour un ingrédient
type ProduitsPossibles map[int64][]models.Produit // id_ingredient -> produits

// CommandeCompleteItem représente la commande d'un produit.
type CommandeCompleteItem struct {
	Produit models.Produit `json:"produit"`

	// jour conseillé de commande, prenant en compte les délais de livraison
	JourCommande time.Time `json:"jour_commande"`

	Quantite int64 `json:"quantite"`

	// ingrédients liés à ce produit
	Origines []TimedIngredientQuantite `json:"origines"`
}

type cacheIngredientProduits struct {
	produits     models.Produits
	associations map[int64][]models.Produit // id ingredient -> produits
	defauts      map[int64][]models.Produit // id ingredient -> produits par défaut
}

// charge en (seulement) 3 requêtes les infos nécessaires
// à l'établissement d'une commande
// se restreint aux produits associés aux livraisons données ou demandé explicitement
func (ct RequeteContext) newCacheIngredientProduits(idsIngredients models.Ids, livraisons models.Livraisons, idProduits models.Ids) (cacheIngredientProduits, error) {
	out := cacheIngredientProduits{ // initialization des map
		produits:     make(models.Produits),
		associations: make(map[int64][]models.Produit),
		defauts:      make(map[int64][]models.Produit),
	}
	rows, err := ct.DB.Query(`SELECT ingredient_produits.* FROM ingredient_produits 
	JOIN produits ON ingredient_produits.id_produit = produits.id 
	WHERE ingredient_produits.id_ingredient = ANY($1)
	AND produits.id_livraison = ANY($2)`,
		idsIngredients.AsSQL(), livraisons.Ids().AsSQL())
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}
	ingredientsProduits, err := models.ScanIngredientProduits(rows)
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}

	rows, err = ct.DB.Query(`SELECT produits.* FROM produits 
	JOIN ingredient_produits ON ingredient_produits.id_produit = produits.id 
	WHERE ingredient_produits.id_ingredient = ANY($1) AND produits.id_livraison = ANY($2)
	UNION 
	SELECT * FROM produits WHERE id = ANY($3)
	`,
		idsIngredients.AsSQL(), livraisons.Ids().AsSQL(), idProduits.AsSQL())
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}
	out.produits, err = models.ScanProduits(rows)
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}

	// on regroupe les produits par ingrédients
	for _, ingProd := range ingredientsProduits {
		out.associations[ingProd.IdIngredient] = append(out.associations[ingProd.IdIngredient], out.produits[ingProd.IdProduit])
	}

	defauts, err := models.SelectDefautProduitsByIdProduits(ct.DB, out.produits.Ids()...)
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}
	for _, def := range defauts {
		out.defauts[def.IdIngredient] = append(out.defauts[def.IdIngredient], out.produits[def.IdProduit])
	}
	return out, nil
}

// ProposeLienIngredientProduit renvoie une association possible
// pour les ingrédients donnés, en général incomplète.
// Le client doit la complèter avant d'utiliser `EtablitCommandeComplete`.
// Les produits possibles sont restreint au séjour demandé.
func (ct RequeteContext) ProposeLienIngredientProduit(params InAssocieIngredients) (ProduitsPossibles, error) {
	dc, err := ct.fetchDataCommande(params)
	if err != nil {
		return nil, err
	}

	// on récupére les données des produits
	data, err := ct.newCacheIngredientProduits(dc.ingredients.Ids(), dc.livraisons, nil)
	if err != nil {
		return nil, err
	}

	targetProduits := make(ProduitsPossibles)
	for idIngredient := range dc.ingredients {
		// on récupère les produits associés
		produits := data.associations[idIngredient]
		if len(produits) > 1 {
			// on essaye de résoudre les ambiguités
			if defauts := data.defauts[idIngredient]; len(defauts) >= 1 {
				// utilise les produits par défaut
				produits = defauts
			}
		}
		sort.Slice(produits, func(i, j int) bool {
			return produits[i].Nom < produits[j].Nom
		})
		targetProduits[idIngredient] = produits
	}

	return targetProduits, nil
}

type produitResolver struct {
	targets    map[int64]int64 // id-ingredient -> id-produit
	produits   models.Produits
	livraisons models.Livraisons
}

func (p produitResolver) resolve(idIngredient int64) (idTarget int64, livraison models.Livraison) {
	idTargetProduit := p.targets[idIngredient]
	livraison = p.livraisons[p.produits[idTargetProduit].IdLivraison]
	return idTargetProduit, livraison
}

// renvoie la quantité équivalente à la somme
// des ingrédients contenus dans `l`
// les ingrédients doivent avoir tous la même unité
func aggregeIngredients(l []TimedIngredientQuantite) (float64, error) {
	contrainte := ContrainteListeIngredients{ingredients: l}
	if err := contrainte.Check(); err != nil {
		return 0, nil
	}
	total := 0.
	for _, iq := range l {
		total += iq.Quantite
	}
	return total, nil
}

// EtablitCommandeComplete associe à chaque ingrédient un produit (avec une quantité) et un jour de commande
// respectant la date d'utilisation et le délai de livraison.
// Tous les ingrédients doivent être associés à un produit par le client.
func (ct RequeteContext) EtablitCommandeComplete(params InCommandeComplete) (OutCommandeComplete, error) {
	dc, err := ct.fetchDataCommande(params.IngredientsSejour)
	if err != nil {
		return OutCommandeComplete{}, err
	}

	idProduits, err := params.Contraintes.checkAssociations(dc.ingredients)
	if err != nil {
		return OutCommandeComplete{}, err
	}

	// on récupére les données des produits
	data, err := ct.newCacheIngredientProduits(dc.ingredients.Ids(), dc.livraisons, idProduits)
	if err != nil {
		return OutCommandeComplete{}, err
	}

	// on ajuste les dates de commandes

	resolver := produitResolver{targets: params.Contraintes.Associations, produits: data.produits, livraisons: dc.livraisons}
	accu := calculeDateCommande(resolver, params.Ingredients)

	if params.Contraintes.Regroupe {
		accu = accu.groupe()
	}

	var out []CommandeCompleteItem
	// toutes les demandes ont étés regroupées en produit,
	// on peut maitenant calculer le nombre de produit nécessaire
	for key, value := range accu {
		if len(value) == 0 {
			continue
		}
		ingredient := value[0].Ingredient

		total, err := aggregeIngredients(value)
		if err != nil {
			return OutCommandeComplete{}, err
		}
		prod := data.produits[key.idTarget]

		err = ContrainteIngredientProduit{ingredient: ingredient, produit: prod}.Check()
		if err != nil {
			return OutCommandeComplete{}, err
		}

		if prod.Conditionnement.Quantite <= 0 {
			var chunks []string
			for _, ing := range value {
				chunks = append(chunks, "<b>"+ing.Ingredient.Nom+"</b>")
			}
			fourn := dc.getFournisseur(prod)
			return OutCommandeComplete{}, fmt.Errorf(`Le conditionnement du produit <b>%s</b> (%s) est invalide : <i>%0.3f</i> <br/>
			Ingrédients liés : %s
			`, prod.Nom, fourn.Nom, prod.Conditionnement.Quantite, strings.Join(chunks, ", "))
		}
		colisage := prod.ColisageNeeded(total, ingredient.Unite == models.Piece)
		out = append(out, CommandeCompleteItem{Produit: prod, JourCommande: key.dateCommande, Quantite: colisage, Origines: value})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Quantite > out[j].Quantite
	})
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].JourCommande.Before(out[j].JourCommande)
	})

	return OutCommandeComplete{Commande: out}, nil
}
