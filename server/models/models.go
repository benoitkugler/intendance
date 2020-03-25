// Définit les structures et types utilisés par le serveur.
package models

import (
	"database/sql"
	"time"
)

// sql:UNIQUE(mail)
type Utilisateur struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
	Mail     string `json:"mail"`

	PrenomNom string `json:"prenom_nom"`
}

// sql:UNIQUE(nom)
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
	Id            int64         `json:"id"`
	IdUtilisateur sql.NullInt64 `json:"id_utilisateur"`
	Nom           string        `json:"nom"`

	ModeEmploi string `json:"mode_emploi"`
}

// sql:UNIQUE(id_recette, id_ingredient)
type RecetteIngredient struct {
	IdRecette    int64   `json:"id_recette"`
	IdIngredient int64   `json:"id_ingredient"`
	Quantite     float64 `json:"quantite"`
	Cuisson      string  `json:"cuisson"`
}

// Menu définit un raccourci pour organiser recettes et ingrédients
// Voir `Repas` pour un repas effectif.
type Menu struct {
	Id            int64         `json:"id"`
	IdUtilisateur sql.NullInt64 `json:"id_utilisateur"`

	Commentaire string `json:"commentaire"`
}

// sql:UNIQUE(id_menu, id_ingredient)
type MenuIngredient struct {
	IdMenu       int64 `json:"id_menu"`
	IdIngredient int64 `json:"id_ingredient"`

	Quantite float64 `json:"quantite"`
	Cuisson  string  `json:"cuisson"`
}

// sql:UNIQUE(id_menu, id_recette)
type MenuRecette struct {
	IdMenu    int64 `json:"id_menu"`
	IdRecette int64 `json:"id_recette"`
}

type Sejour struct {
	Id            int64 `json:"id"`
	IdUtilisateur int64 `json:"id_utilisateur"`

	// Fixe l'origine du séjour.
	// Une journée est déterminé par un "offset"
	// relatif à cette date.
	DateDebut time.Time `json:"date_debut"`
	Nom       string    `json:"nom"`
}

// Groupe est un groupe de personnes lié à un séjour
type Groupe struct {
	Id          int64  `json:"id"`
	IdSejour    int64  `json:"id_sejour"`
	Nom         string `json:"nom"`
	NbPersonnes int64  `json:"nb_personnes"`
	Couleur     string `json:"couleur"`
}

// Repas représente un repas effectif, lié à un séjour.
// Il est constitué de recettes et d'ingrédients (de la même manière qu'un menu)
type Repas struct {
	Id              int64   `json:"id"`
	IdSejour        int64   `json:"id_sejour"`
	OffsetPersonnes int64   `json:"offset_personnes"`
	JourOffset      int64   `json:"jour_offset"`
	Horaire         Horaire `json:"horaire"`
}

// sql:UNIQUE(id_repas, id_ingredient)
type RepasIngredient struct {
	IdRepas      int64 `json:"id_repas"`
	IdIngredient int64 `json:"id_ingredient"`

	Quantite float64 `json:"quantite"`
	Cuisson  string  `json:"cuisson"`
}

// sql:UNIQUE(id_repas, id_recette)
type RepasRecette struct {
	IdRepas   int64 `json:"id_repas"`
	IdRecette int64 `json:"id_recette"`
}

// sql:UNIQUE(id_repas, id_groupe)
type RepasGroupe struct {
	IdRepas  int64 `json:"id_repas"`
	IdGroupe int64 `json:"id_groupe"`
}

// sql:UNIQUE(nom)
type Fournisseur struct {
	Id  int64  `json:"id"`
	Nom string `json:"nom"`

	DelaiCommande  int64          `json:"delai_commande"`
	JoursLivraison JoursLivraison `json:"jours_livraison"`
}

// Enregistre les fournisseurs associés à chaque utilisateur
// sql:UNIQUE(id_utilisateur,id_fournisseur)
type UtilisateurFournisseur struct {
	IdUtilisateur int64 `json:"id_utilisateur"`
	IdFournisseur int64 `json:"id_fournisseur"`
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

// sql:UNIQUE(id_ingredient, id_produit)
type IngredientProduit struct {
	IdIngredient  int64 `json:"id_ingredient"`
	IdProduit     int64 `json:"id_produit"`
	IdUtilisateur int64 `json:"id_utilisateur"`
}

type Commande struct {
	Id            int64     `json:"id"`
	IdUtilisateur int64     `json:"id_utilisateur"`
	DateEmission  time.Time `json:"date_emission"`
	Tag           string    `json:"tag"`
}

// sql:UNIQUE(id_commande, id_produit)
type CommandeProduit struct {
	IdCommande int64 `json:"id_commande"`
	IdProduit  int64 `json:"id_produit"`
	Quantite   int64 `json:"quantite"`
}
