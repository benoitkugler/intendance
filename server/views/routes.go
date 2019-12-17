package views

import (
	"fmt"
	"strconv"

	"github.com/benoitkugler/intendance/server/controller"
	"github.com/benoitkugler/intendance/server/models"
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

func CreateIngredient(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	ing, err := Server.CreateIngredient(ct)
	if err != nil {
		return err
	}
	return c.JSON(200, OutCreateIngredient{Token: ct.Token, Ingredient: ing})
}

func UpdateIngredient(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	var ig models.Ingredient
	if err = c.Bind(&ig); err != nil {
		return err
	}
	if err = Server.UpdateIngredient(ct, ig); err != nil {
		return err
	}
	out, err := Server.LoadIngredients()
	if err != nil {
		return err
	}
	return c.JSON(200, OutUpdateIngredient{Token: ct.Token, Ingredients: out})
}

func DeleteIngredient(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	idS, checkS := c.QueryParam("id"), c.QueryParam("check_produits")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		return fmt.Errorf("Impossible de lire l'ID de l'ingrédient à supprimer.")
	}
	if err = Server.DeleteIngredient(ct, id, checkS != ""); err != nil {
		return err
	}
	out, err := Server.LoadIngredients()
	if err != nil {
		return err
	}
	return c.JSON(200, OutUpdateIngredient{Token: ct.Token, Ingredients: out})
}
