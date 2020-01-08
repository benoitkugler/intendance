// Définit les structures et types utilisés par le serveur.
package models

import (
	"database/sql"
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

	Categorie Categorie `json:"categorie"`
	Callories Callories `json:"callories"`
	// optionnel, zero signifie pas de contrainte
	Conditionnement Conditionnement `json:"conditionnement,omitempty"`
}

type Recette struct {
	Id             int64         `json:"id"`
	IdProprietaire sql.NullInt64 `json:"id_proprietaire"`
	Nom            string        `json:"nom"`

	ModeEmploi string `json:"mode_emploi"`
}

type RecetteIngredient struct {
	IdRecette    int64   `json:"id_recette"`
	IdIngredient int64   `json:"id_ingredient"`
	Quantite     float64 `json:"quantite"`
	Cuisson      string  `json:"cuisson"`
}

type Menu struct {
	Id             int64         `json:"id"`
	IdProprietaire sql.NullInt64 `json:"id_proprietaire"`

	Commentaire string `json:"commentaire"`
}

type MenuIngredient struct {
	IdMenu       int64 `json:"id_menu"`
	IdIngredient int64 `json:"id_ingredient"`

	Quantite float64 `json:"quantite"`
	Cuisson  string  `json:"cuisson"`
}

type MenuRecette struct {
	IdMenu    int64 `json:"id_menu"`
	IdRecette int64 `json:"id_recette"`
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

type Repas struct {
	Id          int64   `json:"id"`
	IdSejour    int64   `json:"id_sejour"`
	IdMenu      int64   `json:"id_menu"`
	NbPersonnes int64   `json:"nb_personnes"`
	JourOffset  int64   `json:"jour_offset"`
	Horaire     Horaire `json:"horaire"`
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
	// zero signifie pas de contrainte
	Colisage int64 `json:"colisage"`
}

type IngredientProduit struct {
	IdIngredient int64 `json:"id_ingredient"`
	IdProduit    int64 `json:"id_produit"`
}

type Commande struct {
	Id             int64     `json:"id"`
	IdProprietaire int64     `json:"id_proprietaire"`
	DateEmission   time.Time `json:"date_emission"`
	Tag            string    `json:"tag"`
}

type CommandeProduit struct {
	IdCommande int64 `json:"id_commande"`
	IdProduit  int64 `json:"id_produit"`
	Quantite   int64 `json:"quantite"`
}
