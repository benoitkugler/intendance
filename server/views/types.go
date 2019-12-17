package views

import (
	"github.com/benoitkugler/intendance/server/controller"
	"github.com/benoitkugler/intendance/server/models"
)

type OutAgenda struct {
	Token  string                       `json:"token"`
	Agenda controller.AgendaUtilisateur `json:"agenda"`
}

type OutCreateIngredient struct {
	Token      string            `json:"token"`
	Ingredient models.Ingredient `json:"ingredient"`
}

type OutUpdateIngredient struct {
	Token       string             `json:"token"`
	Ingredients models.Ingredients `json:"ingredients"`
}
