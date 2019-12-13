package controller

import "github.com/benoitkugler/intendance/server/datamodel"

type IngredientRecette struct {
	datamodel.Ingredient
	datamodel.RecetteIngredient
}

type Recette struct {
	datamodel.Recette

	Ingredients []IngredientRecette `json:"ingredients"`
}

type Menu struct {
	datamodel.Menu

	Recettes    []Recette           `json:"recettes"`
	Ingredients []IngredientRecette `json:"ingredients"`
	NbPersonnes int64               `json:"nb_personnes"`
	Horaire     datamodel.Horaire   `json:"horaire"`
}

type Journee struct {
	JourOffset int64  `json:"jour_offset"`
	Menus      []Menu `json:"menus"`
}

type Sejour struct {
	datamodel.Sejour
	Journees map[int64]Journee `json:"journees"`
}

// AgendaUtilisateur rassemble toutes les données
// relative aux séjours, menus, recettes, etc...
// d'un utilisateur.
type AgendaUtilisateur struct {
	Sejours []Sejour `json:"sejours"`
}
