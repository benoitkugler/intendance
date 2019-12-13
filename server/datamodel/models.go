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
	// optionnel, zero signifie pas de contrainte
	Conditionnement Conditionnement `json:"conditionnement,omitempty"`
}

type Recette struct {
	Id             int64  `json:"id"`
	IdProprietaire int64  `json:"id_proprietaire"`
	Nom            string `json:"nom"`

	ModeEmploi string `json:"mode_emploi"`
}

type RecetteIngredient struct {
	IdRecette    int64   `json:"-"`
	IdIngredient int64   `json:"-"`
	Quantite     float64 `json:"quantite"`
	Cuisson      string  `json:"cuisson"`
}

type Menu struct {
	Id             int64 `json:"id"`
	IdProprietaire int64 `json:"id_proprietaire"`

	Commentaire string `json:"commentaire"`
}

type MenuIngredient struct {
	IdMenu       int64
	IdIngredient int64

	Quantite float64
	Cuisson  string
}

type MenuRecette struct {
	IdMenu    int64
	IdRecette int64
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

type SejourMenu struct {
	IdSejour    int64
	IdMenu      int64
	NbPersonnes int64
	JourOffset  int64
	Horaire     Horaire
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

type Commande struct {
	Id             int64     `json:"id"`
	IdProprietaire int64     `json:"id_proprietaire"`
	DateLivraison  time.Time `json:"date_livraison"`
}
