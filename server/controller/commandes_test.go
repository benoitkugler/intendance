package controller

import (
	"fmt"
	"testing"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

func TestCommandeComplete(t *testing.T) {
	s, ct := setupTest(t)
	defer s.DB.Close()
	ings, err := models.SelectAllIngredients(ct.DB)
	if err != nil {
		t.Fatal(err)
	}
	idIngs := ings.Ids()
	ing1, ing2 := ings[idIngs[0]], ings[idIngs[1]]
	livraison, _ := getLivraison(s.DB)

	produit1, err := ct.AjouteIngredientProduit(ing1.Id, models.Produit{
		IdLivraison: livraison.Id,
		Nom:         fmt.Sprintf("Produit de test%d", time.Now().Unix()), Conditionnement: models.Conditionnement{
			Quantite: 1, Unite: ing1.Unite,
		}})
	if err != nil {
		t.Fatal(err)
	}
	produit2, err := ct.AjouteIngredientProduit(ing2.Id, models.Produit{
		IdLivraison: livraison.Id,
		Nom:         fmt.Sprintf("Produit de test %d", time.Now().Unix()), Conditionnement: models.Conditionnement{
			Quantite: 1, Unite: ing2.Unite,
		}})
	if err != nil {
		t.Fatal(err)
	}

	ingredients := []DateIngredientQuantites{
		{
			Date: time.Now(),
			Ingredients: []IngredientQuantite{
				{Ingredient: ing1, Quantite: 44.5},
				{Ingredient: ing2, Quantite: 5},
			},
		},
		{
			Date: time.Now().Add(jourDuration),
			Ingredients: []IngredientQuantite{
				{Ingredient: ing1, Quantite: 44.5},
				{Ingredient: ing2, Quantite: 2},
			},
		},
	}
	contraintes := CommandeCompleteContraintes{
		ContrainteProduits: map[int64]int64{
			ing1.Id: produit1.Id,
			ing2.Id: produit2.Id,
		},
	}
	out, err := ct.EtablitCommandeComplete(ingredients, contraintes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}

func TestCommandeSimple(t *testing.T) {
	s, ct := setupTest(t)
	defer s.DB.Close()
	ings, err := models.SelectAllIngredients(ct.DB)
	if err != nil {
		t.Fatal(err)
	}
	idIngs := ings.Ids()
	ing1, ing2, ing3 := ings[idIngs[0]], ings[idIngs[1]], ings[idIngs[2]]
	livraison1, livraison2 := getLivraison(s.DB)

	ingredients := []DateIngredientQuantites{
		{
			Date: time.Now(),
			Ingredients: []IngredientQuantite{
				{Ingredient: ing1, Quantite: 44.5},
				{Ingredient: ing2, Quantite: 5},
			},
		},
		{
			Date: time.Now().Add(jourDuration),
			Ingredients: []IngredientQuantite{
				{Ingredient: ing1, Quantite: 44.5},
				{Ingredient: ing2, Quantite: 2},
			},
		},
	}
	contraintes := CommandeSimpleContraintes{
		ContrainteLivraisons: map[int64]int64{
			ing1.Id: livraison1.Id,
			ing2.Id: livraison1.Id,
			ing3.Id: livraison2.Id,
		},
	}
	out, err := ct.EtablitCommandeSimple(ingredients, contraintes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}
