package views

import (
	"github.com/benoitkugler/intendance/server/controller"
	"github.com/labstack/echo"
)

// Server est partagé entre chaque requête.
// Il est à instancier dans le main
var Server controller.Server

func GetAgenda(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	out, err := Server.LoadAgendaUtilisateur(ct)
	if err != nil {
		return err
	}
	return c.JSON(200, OutAgenda{Token: ct.Token, Agenda: out})
}
