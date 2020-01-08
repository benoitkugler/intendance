package controller

import "time"

// Ce fichier implémente la génération d'une commande à
// partir d'une liste d'ingrédients.
// La commande peut prendre en compte plusieurs contraintes (ingrédients, coût, etc...)

const (
	MeilleurPrix         CommandeOptimisation = "prix"
	PlusFaibleGaspillage CommandeOptimisation = "gaspillage"
)

type CommandeOptimisation string

type ContrainteProduit struct {
	Date         time.Time `json:"date"`
	IdIngredient int64     `json:"id_ingredient"`
	IdProduit    int64     `json:"id_produit"`
}

type CommandeContraintes struct {
	// Si `false`, autorise à anticiper les livraisons
	// pour regrouper les commandes.
	// Permet de diminuer le nombre de jours final.
	// Sinon, les commandes respectent exactement le jour de livraison.
	RespectJour bool `json:"respect_jour"`

	// Choisit automatiquement les produits pour favoriser
	// le critère choisit.
	Optimisation CommandeOptimisation `json:"optimisation"`

	// Force l'utilisation du produit donné pour l'ingrédient
	// à la date donnée.
	ContrainteProduits []ContrainteProduit `json:"contrainte_produits"`
}

func (s Server) EtablitCommandes(ingredients []DateIngredientQuantites, contraintes CommandeContraintes) ([]Commande, error) {
	//TODO:
	return nil, nil
}
