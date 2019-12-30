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

func getId(c echo.Context) (int64, error) {
	idS := c.QueryParam("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Impossible de lire l'ID de l'ingrédient à supprimer.")
	}
	return id, nil
}

// --------------------------------------------------------------------------
// ------------------------------ Utilisateurs ------------------------------
// --------------------------------------------------------------------------
func GetUtilisateurs(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	out, err := Server.LoadUtilisateurs()
	if err != nil {
		return err
	}
	return c.JSON(200, OutUtilisateurs{Token: ct.Token, Utilisateurs: out})
}

// --------------------------------------------------------------------------
// ------------------------------ Ingredients -------------------------------
// --------------------------------------------------------------------------

func GetIngredients(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	out, err := Server.LoadIngredients()
	if err != nil {
		return err
	}
	return c.JSON(200, OutIngredients{Token: ct.Token, Ingredients: out})
}

func CreateIngredient(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	var ingredientIn models.Ingredient
	if err = c.Bind(&ingredientIn); err != nil {
		return err
	}
	newIngredient, err := Server.CreateIngredient(ct)
	ingredientIn.Id = newIngredient.Id
	if err != nil {
		return err
	}
	ingredientIn, err = Server.UpdateIngredient(ct, ingredientIn)
	if err != nil {
		return err
	}
	return c.JSON(200, OutIngredient{Token: ct.Token, Ingredient: ingredientIn})
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
	ig, err = Server.UpdateIngredient(ct, ig)
	if err != nil {
		return err
	}
	return c.JSON(200, OutIngredient{Token: ct.Token, Ingredient: ig})
}

func DeleteIngredient(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	id, err := getId(c)
	if err != nil {
		return err
	}
	checkS := c.QueryParam("check_produits")
	if err = Server.DeleteIngredient(ct, id, checkS != ""); err != nil {
		return err
	}
	out, err := Server.LoadIngredients()
	if err != nil {
		return err
	}
	return c.JSON(200, OutIngredients{Token: ct.Token, Ingredients: out})
}

// --------------------------------------------------------------------------
// ------------------------------ Recettes ----------------------------------
// --------------------------------------------------------------------------

func GetRecettes(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	out, err := Server.LoadRecettes()
	if err != nil {
		return err
	}
	return c.JSON(200, OutRecettes{Token: ct.Token, Recettes: out})
}

func CreateRecette(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	var recetteIn controller.Recette
	if err = c.Bind(&recetteIn); err != nil {
		return err
	}
	newRecette, err := Server.CreateRecette(ct)
	if err != nil {
		return err
	}
	recetteIn.Id = newRecette.Id // on garde les valeurs d'entrée
	for i := range recetteIn.Ingredients {
		(&recetteIn.Ingredients[i]).IdRecette = newRecette.Id
	}
	recetteIn, err = Server.UpdateRecette(ct, recetteIn)
	if err != nil {
		return err
	}
	return c.JSON(200, OutRecette{Token: ct.Token, Recette: recetteIn})
}

func UpdateRecette(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	var recette controller.Recette
	if err = c.Bind(&recette); err != nil {
		return err
	}
	recette, err = Server.UpdateRecette(ct, recette)
	if err != nil {
		return err
	}
	return c.JSON(200, OutRecette{Token: ct.Token, Recette: recette})
}

func DeleteRecette(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	id, err := getId(c)
	if err != nil {
		return err
	}
	if err = Server.DeleteRecette(ct, id); err != nil {
		return err
	}
	out, err := Server.LoadRecettes()
	if err != nil {
		return err
	}
	return c.JSON(200, OutRecettes{Token: ct.Token, Recettes: out})
}

// --------------------------------------------------------------------------
// -------------------------------- Menus -----------------------------------
// --------------------------------------------------------------------------

func GetMenus(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	out, err := Server.LoadMenus()
	if err != nil {
		return err
	}
	return c.JSON(200, OutMenus{Token: ct.Token, Menus: out})
}

func CreateMenu(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	var menuIn controller.Menu
	if err = c.Bind(&menuIn); err != nil {
		return err
	}
	newMenu, err := Server.CreateMenu(ct)
	if err != nil {
		return err
	}
	menuIn.Id = newMenu.Id // on garde les valeurs d'entrée
	for i := range menuIn.Recettes {
		(&menuIn.Recettes[i]).IdMenu = newMenu.Id
	}
	for i := range menuIn.Ingredients {
		(&menuIn.Ingredients[i]).IdMenu = newMenu.Id
	}
	menuIn, err = Server.UpdateMenu(ct, menuIn)
	if err != nil {
		return err
	}
	return c.JSON(200, OutMenu{Token: ct.Token, Menu: menuIn})
}

func UpdateMenu(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	var recette controller.Menu
	if err = c.Bind(&recette); err != nil {
		return err
	}
	recette, err = Server.UpdateMenu(ct, recette)
	if err != nil {
		return err
	}
	return c.JSON(200, OutMenu{Token: ct.Token, Menu: recette})
}

func DeleteMenu(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	id, err := getId(c)
	if err != nil {
		return err
	}
	if err = Server.DeleteMenu(ct, id); err != nil {
		return err
	}
	out, err := Server.LoadMenus()
	if err != nil {
		return err
	}
	return c.JSON(200, OutMenus{Token: ct.Token, Menus: out})
}

// --------------------------------------------------------------------------
// ---------------------------- Sejours et repas ----------------------------
// --------------------------------------------------------------------------

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

func CreateSejour(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	var sejourIn models.Sejour
	if err = c.Bind(&sejourIn); err != nil {
		return err
	}
	newSejour, err := Server.CreateSejour(ct)
	if err != nil {
		return err
	}
	sejourIn.Id = newSejour.Id // on garde les valeurs d'entrée
	sejourIn, err = Server.UpdateSejour(ct, sejourIn)
	if err != nil {
		return err
	}
	return c.JSON(200, OutSejour{Token: ct.Token, Sejour: sejourIn})
}

func UpdateSejour(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	var sejour models.Sejour
	if err = c.Bind(&sejour); err != nil {
		return err
	}
	sejour, err = Server.UpdateSejour(ct, sejour)
	if err != nil {
		return err
	}
	return c.JSON(200, OutSejour{Token: ct.Token, Sejour: sejour})
}

func DeleteSejour(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	id, err := getId(c)
	if err != nil {
		return err
	}
	if err = Server.DeleteSejour(ct, id); err != nil {
		return err
	}
	out, err := Server.LoadAgendaUtilisateur(ct)
	if err != nil {
		return err
	}
	return c.JSON(200, OutAgenda{Token: ct.Token, Agenda: out})
}

func CreateRepas(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	var repasIn models.Repas
	if err = c.Bind(&repasIn); err != nil {
		return err
	}
	newRepas, err := Server.CreateRepas(ct, repasIn.IdSejour, repasIn.IdMenu)
	if err != nil {
		return err
	}
	repasIn.Id = newRepas.Id // on garde les valeurs d'entrée
	err = Server.UpdateManyRepas(ct, []models.Repas{repasIn})
	if err != nil {
		return err
	}
	out, err := Server.LoadAgendaUtilisateur(ct)
	if err != nil {
		return err
	}
	return c.JSON(200, OutAgenda{Token: ct.Token, Agenda: out})
}

func UpdateRepas(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	var repass []models.Repas
	if err = c.Bind(&repass); err != nil {
		return err
	}
	err = Server.UpdateManyRepas(ct, repass)
	if err != nil {
		return err
	}
	out, err := Server.LoadAgendaUtilisateur(ct)
	if err != nil {
		return err
	}
	return c.JSON(200, OutAgenda{Token: ct.Token, Agenda: out})
}

func DeleteRepas(c echo.Context) error {
	ct, err := Server.Authentifie(c.Request())
	if err != nil {
		return err
	}
	id, err := getId(c)
	if err != nil {
		return err
	}
	if err = Server.DeleteRepas(ct, id); err != nil {
		return err
	}
	out, err := Server.LoadAgendaUtilisateur(ct)
	if err != nil {
		return err
	}
	return c.JSON(200, OutAgenda{Token: ct.Token, Agenda: out})
}
