package views

import "github.com/benoitkugler/intendance/server/controller"

type OutAgenda struct {
	Token  string                       `json:"token"`
	Agenda controller.AgendaUtilisateur `json:"agenda"`
}
