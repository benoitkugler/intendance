package controller

import "github.com/benoitkugler/intendance/server/models"

type IngredientRecette struct {
	models.Ingredient
	models.RecetteIngredient
}

type Recette struct {
	models.Recette

	Ingredients []IngredientRecette `json:"ingredients"`
}

type Menu struct {
	models.Menu

	Recettes    []Recette           `json:"recettes"`
	Ingredients []IngredientRecette `json:"ingredients"`
	NbPersonnes int64               `json:"nb_personnes"`
	Horaire     models.Horaire   `json:"horaire"`
}

type Journee struct {
	JourOffset int64  `json:"jour_offset"`
	Menus      []Menu `json:"menus"`
}

type Sejour struct {
	models.Sejour
	Journees map[int64]Journee `json:"journees"`
}

// AgendaUtilisateur rassemble toutes les données
// relative aux séjours, menus, recettes, etc...
// d'un utilisateur.
type AgendaUtilisateur struct {
	Sejours []Sejour `json:"sejours"`
}
