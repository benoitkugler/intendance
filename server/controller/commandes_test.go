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
				{Ingredient: models.Ingredient{Id: 57}, Quantite: 44.5},
				{Ingredient: models.Ingredient{Id: 17}, Quantite: 5},
			},
		},
		{
			Date: time.Now().Add(jourDuration),
			Ingredients: []IngredientQuantite{
				{Ingredient: models.Ingredient{Id: 57}, Quantite: 44.5},
				{Ingredient: models.Ingredient{Id: 17}, Quantite: 2},
				{Ingredient: models.Ingredient{Id: 55}, Quantite: 44.5},
			},
		},
	}
	contraintes := CommandeContraintes{
		ContrainteProduits: map[int64]int64{
			57: 66,
			17: 66,
			55: 71,
		},
	}
	out, err := s.EtablitCommande(ct, ingredients, contraintes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}
