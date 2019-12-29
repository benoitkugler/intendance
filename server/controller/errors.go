package controller

import (
	"fmt"
	"strings"

	"github.com/lib/pq"

	"github.com/benoitkugler/intendance/server/models"
)

func ErrorAuth(err error) error {
	return fmt.Errorf(`Impossible d'authentifier votre requête.<br/>
	Détails : <i>%s</i>`, err)
}

type ErrorIngredientProduitUnite struct {
	ingredient models.Ingredient
	produit    models.Produit
}

func (e ErrorIngredientProduitUnite) Error() string {
	return fmt.Sprintf(`Cet ingrédient est associé au produit <b>%s</b>, dont le conditionnement
	n'est pas compatible avec le nouveau choix d'unité. 
	(produit : %s, unité souhaitée : %s)
	Si vous souhaitez vraiment changer l'unité de cet ingrédient,
	il faudra d'abord <b>enlever</b> %s des produits associés à %s.`,
		e.produit.Nom, e.produit.Conditionnement.Unite, e.ingredient.Unite,
		e.produit.Nom, e.ingredient.Nom)
}

type ErrorIngredientProduitConditionnement struct {
	ingredient models.Ingredient
	produit    models.Produit
}

func (e ErrorIngredientProduitConditionnement) Error() string {
	return fmt.Sprintf(`Cet ingrédient est associé au produit <b>%s</b>, dont le conditionnement
	n'est pas compatible avec la nouvelle contrainte de conditionnement. 
	(produit : %s, conditionnement souhaite : %s)
	Si vous souhaitez vraiment changer le conditionnement de cet ingrédient,
	il faudra d'abord <b>enlever</b> %s des produits associés à %s.`,
		e.produit.Nom, e.produit.Conditionnement, e.ingredient.Conditionnement,
		e.produit.Nom, e.ingredient.Nom)

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

func ErrorSQL(err error) error { return errorSQL{err} }
