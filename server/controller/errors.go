package controller

import (
	"fmt"
	"strings"

	"github.com/benoitkugler/intendance/server/models"
	"github.com/lib/pq"
)

func ErrorAuth(err error) error {
	return fmt.Errorf(`Impossible d'authentifier votre requête.<br/>
	Détails : <i>%s</i>`, err)
}

type errorSQL struct {
	error
}

func (e errorSQL) Error() string {
	return fmt.Sprintf(`La requête SQL correspondant à votre demande a échoué.<br/>
		Détails : <i>%s</i>
	`, e.parseDetails())
}

func (e errorSQL) parseDetails() string {
	if pqError, ok := e.error.(*pq.Error); ok {
		switch pqError.Constraint {
		case "ingredients_nom_key":
			return "Le nom demandé pour cet ingrédient est <b>déjà pris</b>."
		}
	}
	return e.error.Error()
}

func ErrorSQL(err error) error {
	if err == nil {
		return nil
	}
	return errorSQL{err}
}

type ErrorIngredientUsed struct {
	recettes models.Recettes
	menus    models.Menus
	produits models.Produits
}

func (e ErrorIngredientUsed) Error() string {
	var b strings.Builder
	b.WriteString("Cet ingrédient est associé à :")
	if len(e.recettes) > 0 {
		b.WriteString("<br/>	<b>des recettes</b>:")
		for _, r := range e.recettes {
			b.WriteString("<br/>		" + r.Nom)
		}
	}
	if L := len(e.menus); L > 0 {
		b.WriteString(fmt.Sprintf("<br/>	<b>%d menu(s)</b>", L))
	}
	if len(e.produits) > 0 {
		b.WriteString("<br/>	<b>des produits</b>:")
		for _, r := range e.produits {
			b.WriteString("<br/>		" + r.Nom)
		}
	}
	return b.String()
}
