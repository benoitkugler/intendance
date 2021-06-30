package controller

import (
	"fmt"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

// Ce fichier implémente la génération d'une commande à
// partir d'une liste d'ingrédients.
// La commande peut prendre en compte plusieurs contraintes (ingrédients, coût, etc...)

// const (
// 	MeilleurPrix         CommandeOptimisation = "prix"
// 	PlusFaibleGaspillage CommandeOptimisation = "gaspillage"
// )

// type CommandeOptimisation string

// type ContrainteProduit struct {
// 	Date         time.Time `json:"date"`
// 	IdIngredient int64     `json:"id_ingredient"`
// 	IdProduit    int64     `json:"id_produit"`
// }

// TODO: ajouter un paramètre de regroupement des jours. Ex :
// un même produit nécessaire sur 2 jour serait commandé groupé pour le premier

const jourDuration = 24 * time.Hour

// TimedIngredientQuantite ajoute la date de demande de l'ingrédient
type TimedIngredientQuantite struct {
	IngredientQuantite
	Date time.Time `json:"date"`
}

// CommandeContraintes paramétrise la requête de commande
type CommandeContraintes struct {
	// Précise la cible à utiliser pour un ingrédient (idIngredient -> idCible)
	Associations map[int64]int64 `json:"associations"`

	// Si `true`, regroupe toutes les commandes
	// à la date courante (prototype)
	Regroupe bool `json:"regroupe"`
}

// vérifie que toutes les correspondances ingrédient -> target
// sont fournies par le client, et renvoie les ids correspondants
func (c CommandeContraintes) checkAssociations(ingredients models.Ingredients) (models.Ids, error) {
	var out models.Ids
	for _, ingredient := range ingredients {
		idProduit, ok := c.Associations[ingredient.Id]
		if !ok {
			return nil, fmt.Errorf("L'ingrédient %s n'est associé à aucune cible !", ingredient.Nom)
		}
		out = append(out, idProduit)
	}
	return out, nil
}

type timedTarget struct {
	idTarget     int64
	dateCommande time.Time
}

type timedTargets map[timedTarget][]TimedIngredientQuantite

func (ts timedTargets) addIngredient(idTarget int64, jour time.Time, ing TimedIngredientQuantite) {
	// on normalise les dates
	jour = jour.Truncate(jourDuration)
	key := timedTarget{idTarget: idTarget, dateCommande: jour}
	ts[key] = append(ts[key], ing)
}

// ajuste la date de tous les items au premier jour de commande
func (ts timedTargets) groupe() timedTargets {
	var first time.Time
	for key := range ts {
		if first.IsZero() || key.dateCommande.Before(first) {
			first = key.dateCommande
		}
	}
	// on regroupe toutes les demandes sur le même jour
	out := timedTargets{}
	for key, v := range ts {
		out[timedTarget{idTarget: key.idTarget, dateCommande: first}] = append(out[timedTarget{idTarget: key.idTarget, dateCommande: first}], v...)
	}
	return out
}

type dataCommande struct {
	fournisseurs models.Fournisseurs
	livraisons   models.Livraisons
	ingredients  models.Ingredients
}

func (ct RequeteContext) fetchDataCommande(ingredients []DateIngredientQuantites) (out dataCommande, err error) {
	out.fournisseurs, err = ct.loadFournisseurs()
	if err != nil {
		return out, err
	}
	out.livraisons, err = ct.loadLivraisons(out.fournisseurs)
	if err != nil {
		return out, err
	}

	out.ingredients = models.Ingredients{}
	for _, iq := range ingredients {
		for _, ing := range iq.Ingredients {
			out.ingredients[ing.Ingredient.Id] = ing.Ingredient
		}
	}
	return out, nil
}

func (dc dataCommande) getFournisseur(produit models.Produit) models.Fournisseur {
	return dc.fournisseurs[dc.livraisons[produit.IdLivraison].IdFournisseur]
}

type targetResolver interface {
	// id target est soit un produit soit une livraison
	resolve(idIngredient int64) (idTarget int64, livraison models.Livraison)
}

var (
	_ targetResolver = livraisonResolver{}
	_ targetResolver = produitResolver{}
)

// Accumule les ingrédients en les regroupants par
// (idTarget, date de commande)
func calculeDateCommande(resolver targetResolver, ingredients []DateIngredientQuantites) timedTargets {
	accu := timedTargets{}

	for _, demande := range ingredients {
		for _, ing := range demande.Ingredients {
			idTarget, livraison := resolver.resolve(ing.Ingredient.Id)

			dateCommande, _ := livraison.DateCommande(demande.Date)

			// on ajoute au timed-produit l'ingrédient et sa quantité
			// avec la date de demande
			timedIngredient := TimedIngredientQuantite{IngredientQuantite: ing, Date: demande.Date}
			accu.addIngredient(idTarget, dateCommande, timedIngredient)
		}
	}
	return accu
}
