package controller

import (
	"fmt"
	"testing"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

func TestCommande(t *testing.T) {
	s, ct := setupTest(t)
	defer s.db.Close()

	ingredients := []DateIngredientQuantites{
		{
			Date: time.Now(),
			Ingredients: []IngredientQuantite{
				{Ingredient: models.Ingredient{Id: 1, Nom: "seze"}, Quantite: 44.5},
				{Ingredient: models.Ingredient{Id: 2, Nom: "zezezze"}, Quantite: 5},
			},
		},
		{
			Date: time.Now().Add(jourDuration),
			Ingredients: []IngredientQuantite{
				{Ingredient: models.Ingredient{Id: 1, Nom: "seze"}, Quantite: 44.5},
				{Ingredient: models.Ingredient{Id: 2, Nom: "zezezze"}, Quantite: 2},
				{Ingredient: models.Ingredient{Id: 3, Nom: "aa"}, Quantite: 44.5},
			},
		},
	}
	contraintes := CommandeContraintes{
		ContrainteProduits: map[int64]int64{
			1: 2,
		},
	}
	out, ambs, err := s.EtablitCommande(ct, ingredients, contraintes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
	fmt.Println(ambs)
}
