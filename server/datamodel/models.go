// Définit les structures et types utilisés par le serveur.
package datamodel

import (
	"time"
)

type Utilisateur struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
	Mail     string `json:"mail"`

	PrenomNom string `json:"prenom_nom"`
}

type Ingredient struct {
	Id    int64  `json:"id"`
	Nom   string `json:"nom"`
	Unite Unite  `json:"unite"`

	Categorie string    `json:"categorie"`
	Callories Callories `json:"callories"`
}

type Recette struct {
	Id             int64  `json:"id"`
	IdProprietaire int64  `json:"id_proprietaire"`
	Nom            string `json:"nom"`

	ModeEmploi string `json:"mode_emploi"`
}

type Menu struct {
	Id             int64 `json:"id"`
	IdProprietaire int64 `json:"id_proprietaire"`

	Commentaire string `json:"commentaire"`
}

type Sejour struct {
	Id             int64 `json:"id"`
	IdProprietaire int64 `json:"id_proprietaire"`

	// Fixe l'origine du séjour.
	// Une journée est déterminé par un "offset"
	// relatif à cette date.
	DateDebut time.Time `json:"date_debut"`
	Nom       string    `json:"nom"`
}

type Fournisseur struct {
	Id  int64  `json:"id"`
	Nom string `json:"nom"`

	DelaiCommande  int64          `json:"delai_commande"`
	JoursLivraison JoursLivraison `json:"jours_livraison"`
}

type Produit struct {
	Id              int64           `json:"id"`
	IdFournisseur   int64           `json:"id_fournisseur"`
	Nom             string          `json:"nom"`
	Conditionnement Conditionnement `json:"conditionnement"`
	Prix            float64         `json:"prix"`

	ReferenceFournisseur string `json:"reference_fournisseur"`
}

type Commande struct {
	Id             int64     `json:"id"`
	IdProprietaire int64     `json:"id_proprietaire"`
	DateLivraison  time.Time `json:"date_livraison"`
}
