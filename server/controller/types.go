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

type RepasComplet struct {
	models.Repas

	Groupes     []models.RepasGroupe     `json:"groupes"`
	Recettes    []models.RepasRecette    `json:"recettes"`
	Ingredients []models.RepasIngredient `json:"ingredients"`
}
type SejourRepas struct {
	models.Sejour
	Fournisseurs []models.SejourFournisseur `json:"fournisseurs"`
	Repass       []RepasComplet             `json:"repass"`
}

// Sejours contient les séjours, ainsi que les groupes et repas associés.
type Sejours struct {
	Sejours map[int64]SejourRepas `json:"sejours"`
	Groupes models.Groupes        `json:"groupes"`
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

type PreviewCommande struct {
	Produits []CommandeItem
}
