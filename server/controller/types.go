package controller

import (
	"database/sql"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

type Recette struct {
	models.Recette
	Ingredients []models.RecetteIngredient `json:"ingredients"`
}

type Menu struct {
	models.Menu

	Recettes    []models.MenuRecette    `json:"recettes"`
	Ingredients []models.MenuIngredient `json:"ingredients"`
}

type Journee struct {
	JourOffset int64          `json:"jour_offset"`
	Repas      []models.Repas `json:"menus"`
}

type SejourJournees struct {
	Sejour   models.Sejour     `json:"sejour"`
	Journees map[int64]Journee `json:"journees"` // key is JourOffset
}

// AgendaUtilisateur rassemble toutes les données
// relative aux séjours et repas
// d'un utilisateur.
type Sejours struct {
	Sejours map[int64]*SejourJournees `json:"sejours"`
	Groupes models.Groupes            `json:"groupes"`
}

type Utilisateur struct {
	Id        int64  `json:"id"`
	PrenomNom string `json:"prenom_nom"`
}

type OutLoggin struct {
	Erreur      string      `json:"erreur"`
	Token       string      `json:"token"`
	Utilisateur Utilisateur `json:"utilisateur"`
}

type IngredientQuantite struct {
	Ingredient models.Ingredient `json:"ingredient"`
	Quantite   float64           `json:"quantite"`
}

type DateIngredientQuantites struct {
	Date        time.Time            `json:"date"`
	Ingredients []IngredientQuantite `json:"ingredients"`
}

type IngredientProduits struct {
	Produits  []models.Produit `json:"produits"`
	IdDefault sql.NullInt64    `json:"id_default"`
}

type Commande struct {
	models.Commande
	Produits []models.CommandeProduit
}
