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
type AgendaUtilisateur struct {
	Sejours map[int64]*SejourJournees `json:"sejours"`
}
