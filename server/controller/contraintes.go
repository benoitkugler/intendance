package controller

import (
	"fmt"

	"github.com/benoitkugler/intendance/server/models"
)

// Ce fichier défini les invariants du modèle.
// Certains sont assurés par la base de données, mais d'autres sont plus fins.

type ContrainteIngredientProduit struct {
	ingredient models.Ingredient
	produit    models.Produit
}

func (c ContrainteIngredientProduit) Check() error {
	// pour les ingrédients en vrac, l'unité de l'ingrédient et du produit
	// doivent être égales
	if c.ingredient.Unite != models.Piece && c.ingredient.Unite != c.produit.Conditionnement.Unite {
		return fmt.Errorf(`L'unité %s de l'ingrédient <b>%s</b> n'est pas compatible 
		avec le conditionnement du produit <b>%s</b> (unité %s)`, c.ingredient.Unite, c.ingredient.Nom,
			c.produit.Nom, c.produit.Conditionnement.Unite)
	}

	// si l'ingrédient impose un conditionnement, le produit doit le respecter
	if !c.ingredient.Conditionnement.IsNull() && c.ingredient.Conditionnement != c.produit.Conditionnement {
		return fmt.Errorf(`Le conditionnement %s de l'ingrédient <b>%s</b> n'est pas compatible 
		avec le conditionnement du produit <b>%s</b> : %s`, c.ingredient.Conditionnement,
			c.ingredient.Nom, c.produit.Nom, c.produit.Conditionnement)
	}
	return nil
}

type ContrainteIngredient struct {
	ingredient models.Ingredient
}

func (c ContrainteIngredient) Check() error {
	// Le conditionnement n'est valide que pour les ingrédients à la pièce
	if c.ingredient.Unite != models.Piece && !c.ingredient.Conditionnement.IsNull() {
		return fmt.Errorf(`Le conditionnement n'est supporté que pour les ingrédients à la pièce.
		Conditionnement de %s : %s`, c.ingredient.Nom, c.ingredient.Conditionnement)
	}
	return nil
}
