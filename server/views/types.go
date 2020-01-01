package views

import (
	"github.com/benoitkugler/intendance/server/controller"
	"github.com/benoitkugler/intendance/server/models"
)

type OutAgenda struct {
	Token  string                       `json:"token"`
	Agenda controller.AgendaUtilisateur `json:"agenda"`
}

type OutIngredient struct {
	Token      string            `json:"token"`
	Ingredient models.Ingredient `json:"ingredient"`
}

type OutIngredients struct {
	Token       string             `json:"token"`
	Ingredients models.Ingredients `json:"ingredients"`
}

type OutRecette struct {
	Token   string             `json:"token"`
	Recette controller.Recette `json:"recette"`
}

type OutRecettes struct {
	Token    string                        `json:"token"`
	Recettes map[int64]*controller.Recette `json:"recettes"`
}

type OutMenu struct {
	Token string          `json:"token"`
	Menu  controller.Menu `json:"menu"`
}

type OutMenus struct {
	Token string                     `json:"token"`
	Menus map[int64]*controller.Menu `json:"menus"`
}

type OutSejour struct {
	Token  string        `json:"token"`
	Sejour models.Sejour `json:"sejour"`
}

type OutUtilisateurs struct {
	Token        string                           `json:"token"`
	Utilisateurs map[int64]controller.Utilisateur `json:"utilisateurs"`
}

type InResoudRepas struct {
	Mode string

	IdRepas    int64 // pour Mode == "repas"
	IdSejour   int64 // pour Mode == "sejour" ou "journee"
	JourOffset int   // pour Mode == "journee"
}

type OutResoudRepas struct {
	Token       string                          `json:"token"`
	Ingredients []controller.IngredientQuantite `json:"ingredients"`
}
