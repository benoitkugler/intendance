package controller

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

type CommandeCompleteContraintes struct {
	// Force l'utilisation du produit pour l'ingrédient (idIngredient -> idProduit)
	ContrainteProduits map[int64]int64 `json:"contrainte_produits"`

	// Si `true`, regroupe toutes les commandes
	// à la date courante (prototype)
	Regroupe bool `json:"regroupe"`
}

// Ambiguites indique les ingrédients pour lesquelles plusieurs produits sont disponibles
type Ambiguites map[int64][]models.Produit // id_ingredient -> produits

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
// se restreint aux produits associés aux livraisons données
func (ct RequeteContext) resoudProduits(idsIngredients models.Ids, livraisons models.Livraisons) (cacheIngredientProduits, error) {
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
	WHERE ingredient_produits.id_ingredient = ANY($1)
	AND produits.id_livraison = ANY($2)`,
		idsIngredients.AsSQL(), livraisons.Ids().AsSQL())
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

type produitResolver struct {
	targets    map[int64]models.Produit // id-ingredient -> produit
	livraisons models.Livraisons
}

func (p produitResolver) resolve(idIngredient int64) (idTarget int64, livraison models.Livraison) {
	targetProduit := p.targets[idIngredient]
	livraison = p.livraisons[targetProduit.IdLivraison]
	return targetProduit.Id, livraison
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
// Les produits sont résolus en prenant en compte les associations enregistrées puis les contraintes fournies en arguments.
func (ct RequeteContext) EtablitCommandeComplete(ingredients []DateIngredientQuantites, contraintes CommandeCompleteContraintes) (OutCommandeComplete, error) {
	// TODO: vérifier les associations ing -> produit,
	// où au moins les contraintes d'unité, etc..

	livraisons, allIngredients, err := ct.fetchDataCommande(ingredients)
	if err != nil {
		return OutCommandeComplete{}, err
	}

	// on récupére les données des produits
	data, err := ct.resoudProduits(allIngredients.Ids(), livraisons)
	if err != nil {
		return OutCommandeComplete{}, err
	}

	// on commence par associer à chaque ingrédient un produit
	// indépendement de la date d'utilisation
	ambiguites := make(Ambiguites)
	targetProduits := produitResolver{targets: make(map[int64]models.Produit), livraisons: livraisons}
	for idIngredient, ingredient := range allIngredients {
		var (
			ambs          []models.Produit
			targetProduit models.Produit
		)
		targetIdProduit, hasContrainte := contraintes.ContrainteProduits[idIngredient]
		if hasContrainte {
			var has bool
			targetProduit, has = data.produits[targetIdProduit]
			if !has { // le produit imposé n'est pas conforme
				return OutCommandeComplete{}, fmt.Errorf("Le produit (%d) n'est pas associé à l'ingrédient <b>%s</b> !",
					targetIdProduit, ingredient.Nom)
			}
		} else {
			// on récupère les produits associés
			prods := data.associations[idIngredient]

			switch {
			case len(prods) == 0: // l'absence de produit est fatale
				return OutCommandeComplete{}, fmt.Errorf("L'ingrédient %s n'est associé à aucun produit !", ingredient.Nom)
			case len(prods) > 1: // on essaye de résoudre les ambiguités
				defauts := data.defauts[idIngredient]
				switch {
				case len(defauts) == 0:
					// pas de valeur par défaut : on fait un choix arbitraire
					ambs = prods
					targetProduit = prods[0]
				case len(defauts) > 1:
					// on restreint l'ambiguité aux défauts et on fait un choix arbitraire
					ambs = defauts
					targetProduit = defauts[0]
				default:
					// OK : on utilise l'unique produit par défaut
					targetProduit = defauts[0]
				}
			default:
				// OK : on utilise le seul produit
				targetProduit = prods[0]
			}
		}
		sort.Slice(ambs, func(i, j int) bool {
			return ambs[i].Nom < ambs[j].Nom
		})
		targetProduits.targets[idIngredient] = targetProduit
		if len(ambs) > 0 {
			ambiguites[idIngredient] = ambs
		}
	}

	// puis on utilise les produits trouvés pour
	// ajuster les dates de commandes

	accu := calculeDateCommande(targetProduits, ingredients)

	if contraintes.Regroupe {
		accu = accu.groupe()
	}

	var out []CommandeCompleteItem
	// toutes les demandes ont étés regroupées en produit,
	// on peut maitenant calculer le nombre de produit nécessaire
	for key, value := range accu {
		total, err := aggregeIngredients(value)
		if err != nil {
			return OutCommandeComplete{}, err
		}
		prod := data.produits[key.idTarget]
		if prod.Conditionnement.Quantite <= 0 {
			var chunks []string
			for _, ing := range value {
				chunks = append(chunks, "<b>"+ing.Ingredient.Nom+"</b>")
			}
			return OutCommandeComplete{}, fmt.Errorf(`Le conditionnement du produit <b>%s</b> est invalide : <i>%0.3f</i> <br/>
			Ingrédients liés : %s
			`, prod.Nom, prod.Conditionnement.Quantite, strings.Join(chunks, ", "))
		}
		colisage := prod.ColisageNeeded(total)
		out = append(out, CommandeCompleteItem{Produit: prod, JourCommande: key.dateCommande, Quantite: colisage, Origines: value})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Quantite > out[j].Quantite
	})
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].JourCommande.Before(out[j].JourCommande)
	})

	return OutCommandeComplete{Commande: out, Ambiguites: ambiguites}, nil
}
