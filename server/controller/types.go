package controller

import "github.com/benoitkugler/intendance/server/models"

type Recette struct {
	models.Recette
	Ingredients []models.RecetteIngredient `json:"ingredients"`
}

type Menu struct {
	models.Menu

	Recettes    []models.MenuRecette    `json:"recettes"`
	Ingredients []models.MenuIngredient `json:"ingredients"`
}

// type Items struct {
// 	Ingredients models.Ingredients `json:"ingredients"`
// 	Recettes    models.Recettes    `json:"recettes"`
// 	Menus       models.Menus       `json:"menus"`
// }

type Repas struct {
	IdMenu      int64          `json:"id_menu"`
	NbPersonnes int64          `json:"nb_personnes"`
	Horaire     models.Horaire `json:"horaire"`
}

type Journee struct {
	JourOffset int64   `json:"jour_offset"`
	Menus      []Repas `json:"menus"`
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
