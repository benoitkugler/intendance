package views

import "github.com/benoitkugler/intendance/server/datamodel"

type IngredientRecette struct {
	datamodel.Ingredient

	Quantite float64 `json:"quantite"`
	Cuisson  string  `json:"cuisson"`
}

type Recette struct {
	datamodel.Recette

	Ingredients []IngredientRecette `json:"ingredients"`
}

type Menu struct {
	datamodel.Menu

	Recettes    []Recette           `json:"recettes"`
	Ingredients []IngredientRecette `json:"ingredients"`
	NbPersonnes int                 `json:"nb_personnes"`
	Horaire     datamodel.Horaire   `json:"horaire"`
}

type Journee struct {
	JourOffset int    `json:"jour_offset"`
	Menus      []Menu `json:"menus"`
}

type Sejour struct {
	datamodel.Sejour
	Journees []Journee `json:"journees"`
}

// AgendaUtilisateur rassemble toutes les données
// relative aux séjours, menus, recettes, etc...
// d'un utilisateur.
type AgendaUtilisateur struct {
	Sejours []Sejour `json:"sejours"`
}
