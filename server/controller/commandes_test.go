package controller

import (
	"fmt"
	"testing"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func getSejour(t *testing.T, ct RequeteContext) models.Sejour {
	sejours, err := models.SelectSejoursByIdUtilisateurs(ct.DB, ct.IdProprietaire)
	check(t, err)

	if len(sejours) == 0 {
		t.Fatal("aucun séjour pour l'utilisateur de test")
	}
	return sejours[sejours.Ids()[0]]
}

func TestCommandeComplete(t *testing.T) {
	s, ct := setupTest(t)
	defer s.DB.Close()
	ings, err := models.SelectAllIngredients(ct.DB)
	check(t, err)

	idIngs := ings.Ids()
	ing1, ing2 := ings[idIngs[0]], ings[idIngs[1]]
	livraison, _ := getLivraison(s.DB, ct.IdProprietaire)

	sejour := getSejour(t, ct)

	err = ct.UpdateSejourFournisseurs(sejour.Id, []int64{livraison.IdFournisseur})
	check(t, err)

	cond1 := models.Conditionnement{Quantite: 1, Unite: ing1.Unite}
	if !ing1.Conditionnement.IsNull() {
		cond1 = ing1.Conditionnement
	}
	produit1, err := ct.AjouteIngredientProduit(ing1.Id, models.Produit{
		IdLivraison: livraison.Id,
		Nom:         fmt.Sprintf("Produit de test%d", time.Now().Unix()), Conditionnement: cond1,
	})
	check(t, err)

	cond2 := models.Conditionnement{Quantite: 1, Unite: ing2.Unite}
	if !ing2.Conditionnement.IsNull() {
		cond2 = ing2.Conditionnement
	}
	produit2, err := ct.AjouteIngredientProduit(ing2.Id, models.Produit{
		IdLivraison: livraison.Id,
		Nom:         fmt.Sprintf("Produit de test %d", time.Now().Unix()), Conditionnement: cond2,
	})
	check(t, err)

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

	_, err = ct.ProposeLienIngredientProduit(InAssocieIngredients{Ingredients: ingredients, IdSejour: sejour.Id})
	check(t, err)

	contraintes := CommandeContraintes{
		Associations: map[int64]int64{
			ing1.Id: produit1.Id,
			ing2.Id: produit2.Id,
		},
	}
	out, err := ct.EtablitCommandeComplete(InCommandeComplete{IngredientsSejour: IngredientsSejour{Ingredients: ingredients, IdSejour: sejour.Id}, Contraintes: contraintes})
	check(t, err)

	fmt.Println(out)
}

func TestCommandeSimple(t *testing.T) {
	s, ct := setupTest(t)
	defer s.DB.Close()
	ings, err := models.SelectAllIngredients(ct.DB)
	check(t, err)

	idIngs := ings.Ids()
	ing1, ing2, ing3 := ings[idIngs[0]], ings[idIngs[1]], ings[idIngs[2]]
	livraison1, livraison2 := getLivraison(s.DB, ct.IdProprietaire)

	sejour := getSejour(t, ct)

	err = ct.UpdateSejourFournisseurs(sejour.Id, []int64{livraison1.IdFournisseur, livraison2.IdFournisseur})
	check(t, err)

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

	_, err = ct.ProposeLienIngredientLivraison(IngredientsSejour{Ingredients: ingredients, IdSejour: sejour.Id})
	check(t, err)

	contraintes := CommandeContraintes{
		Associations: map[int64]int64{
			ing1.Id: livraison1.Id,
			ing2.Id: livraison1.Id,
			ing3.Id: livraison2.Id,
		},
	}
	out, err := ct.EtablitCommandeSimple(InCommandeSimple{IngredientsSejour: IngredientsSejour{Ingredients: ingredients, IdSejour: sejour.Id}, Contraintes: contraintes})
	check(t, err)

	fmt.Println(out)
}
