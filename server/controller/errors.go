package controller

import (
	"fmt"

	"github.com/benoitkugler/intendance/server/datamodel"
)

type ErrorIngredientProduitUnite struct {
	ingredient datamodel.Ingredient
	produit    datamodel.Produit
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
	ingredient datamodel.Ingredient
	produit    datamodel.Produit
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

func ErrorSQL(err error) error {
	return fmt.Errorf(`La requête SQL correspondant à votre demande a échoué.
		Détails : 
		<i>%s</i>
	`, err)
}
