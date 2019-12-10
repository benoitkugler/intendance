// Définit les structures et types utilisés par le serveur.
package datamodel

import (
	"database/sql"

	"github.com/lib/pq"
)

type Utilisateur struct {
	Id       int64
	Password string
	Mail     string

	PrenomNom sql.NullString
}

type Ingredient struct {
	Id    int64
	Nom   string
	Unite string

	Categorie sql.NullString
	Callories Callories
}

type Recette struct {
	Id             int64
	IdProprietaire int64
	Nom            string

	ModeEmploi sql.NullString
}

type Menu struct {
	Id             int64
	IdProprietaire int64
}

type Sejours struct {
	Id             int64
	IdProprietaire int64

	// Fixe l'origine du séjour.
	// Une journée est déterminé par un "offset"
	// relatif à cette date.
	DateDebut pq.NullTime
	Nom       sql.NullString
}

type Fournisseur struct {
	Id  int64
	Nom string

	DelaiCommande  sql.NullInt64
	JoursLivraison JoursLivraison
}

type Produit struct {
	Id              int64
	IdFournisseur   int64
	Nom             string
	Conditionnement Conditionnement
	Prix            float64

	ReferenceFournisseur sql.NullString
}

type Commande struct {
	Id             int64
	IdProprietaire int64
	DateLivraison  pq.NullTime
}
