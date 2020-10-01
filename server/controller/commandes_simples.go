package controller

import (
	"sort"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

// LivraisonsPossibles propose une livraison conseillées pour un ingrédient,
// en faisant éventuellement un choix arbitraire
type LivraisonsPossibles map[int64]int64 // id_ingredient -> id_livraison

// ProposeLienIngredientLivraison renvoie une association possible
// pour les ingrédients donnés, en général incomplète.
// Le client doit la complèter avant d'utiliser `EtablitCommandeSimple`.
func (ct RequeteContext) ProposeLienIngredientLivraison(ingredients []DateIngredientQuantites) (LivraisonsPossibles, error) {
	livraisons, allIngredients, err := ct.fetchDataCommande(ingredients)
	if err != nil {
		return nil, err
	}
	data, err := ct.newCacheIngredientProduits(allIngredients.Ids(), livraisons)
	if err != nil {
		return nil, err
	}

	out := make(LivraisonsPossibles)
	for idIngredient := range allIngredients {
		// utilise les produits déjà associés
		if defauts := data.defauts[idIngredient]; len(defauts) > 0 {
			// choisit arbitrairement une livraison par défaut
			out[idIngredient] = defauts[0].IdLivraison
		} else if prods := data.associations[idIngredient]; len(prods) > 0 {
			// choisit arbitrairement une livraison
			out[idIngredient] = prods[0].IdLivraison
		}
	}
	return out, nil
}

// CommandeSimpleItem représente la commande d'un ingrédient (ou plusieurs ingrédient)
// chez un fournisseur. Alternative plus simple à `CommandeItem`
type CommandeSimpleItem struct {
	Livraison models.Livraison `json:"livraison"`

	// jour conseillé de commande, prenant en compte les délais de livraison
	JourCommande time.Time `json:"jour_commande"`

	// liste groupée des ingrédients à commander
	// (un même ingrédient n'y apparait qu'une fois)
	Ingredients []IngredientQuantite

	// ingrédients donnant lieu à cet item
	Origines []TimedIngredientQuantite `json:"origines"`
}

type livraisonResolver struct {
	livraisons models.Livraisons
	targets    map[int64]int64
}

func (l livraisonResolver) resolve(idIngredient int64) (idTarget int64, livraison models.Livraison) {
	idLivraison := l.targets[idIngredient]
	return idLivraison, l.livraisons[idLivraison]
}

// regroupe les ingrédients par id (en sommant les quantités);
// les date de demande sont ignorées
func regroupeIngredients(l []TimedIngredientQuantite) map[int64]float64 {
	out := map[int64]float64{} // id ingredient -> quantité
	for _, ing := range l {
		out[ing.Ingredient.Id] += ing.Quantite
	}
	return out
}

// EtablitCommandeSimple regroupe chaque ingrédient par fournisseur, en précisant
// une date de commande respectant les délais de livraisons.
// Tous les ingrédients doivent être associés à une livraison par le client.
func (ct RequeteContext) EtablitCommandeSimple(ingredients []DateIngredientQuantites, contraintes CommandeContraintes) (OutCommandeSimple, error) {
	livraisons, allIngredients, err := ct.fetchDataCommande(ingredients)
	if err != nil {
		return OutCommandeSimple{}, err
	}

	// on vérifie que toutes les correspondances ingrédient -> livraisons
	// sont fournies par le client
	err = contraintes.checkAssociations(allIngredients)
	if err != nil {
		return OutCommandeSimple{}, err
	}

	resolver := livraisonResolver{livraisons: livraisons, targets: contraintes.Associations}
	accu := calculeDateCommande(resolver, ingredients)

	if contraintes.Regroupe {
		accu = accu.groupe()
	}

	// on ré-organise les données
	var out []CommandeSimpleItem
	for key, value := range accu {
		item := CommandeSimpleItem{Livraison: livraisons[key.idTarget], JourCommande: key.dateCommande, Origines: value}
		for idIngredient, quantite := range regroupeIngredients(value) {
			item.Ingredients = append(item.Ingredients, IngredientQuantite{Ingredient: allIngredients[idIngredient], Quantite: quantite})
		}
		out = append(out, item)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].JourCommande.Before(out[j].JourCommande)
	})
	return OutCommandeSimple{Commande: out}, nil
}
