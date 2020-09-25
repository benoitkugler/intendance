package controller

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

func TestGetProduits(t *testing.T) {
	s, ct := setupTest(t)
	defer s.DB.Close()
	out, err := ct.GetIngredientProduits(56)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}

func getLivraison(db models.DB) (models.Livraison, models.Livraison) {
	livraisons, err := models.SelectAllLivraisons(db)
	if err != nil {
		log.Fatal(err)
	}
	ids := livraisons.Ids()
	if L := len(ids); L < 2 {
		log.Fatalf("besoin de 2 livraisons, seulement %d présentes", L)
	}
	return livraisons[ids[0]], livraisons[ids[1]]
}

func TestDelete(t *testing.T) {
	s, ct := setupTest(t)
	defer s.DB.Close()
	livraison, _ := getLivraison(s.DB)
	produit, err := models.Produit{IdLivraison: livraison.Id}.Insert(ct.DB)
	if err != nil {
		t.Fatal(err)
	}
	if err := ct.DeleteProduit(produit.Id); err != nil {
		t.Fatal(err)
	}
}

func TestDefault(t *testing.T) {
	s, ct := setupTest(t)
	defer s.DB.Close()
	ingredients, err := models.SelectAllIngredients(ct.DB)
	if err != nil {
		t.Fatal(err)
	}
	ingredient := ingredients[ingredients.Ids()[0]]
	livraison, _ := getLivraison(s.DB)
	produit, err := ct.AjouteIngredientProduit(ingredient.Id, models.Produit{
		IdLivraison: livraison.Id,
		Nom:         fmt.Sprintf("Produit superbe %d", time.Now().Unix()), Conditionnement: models.Conditionnement{Quantite: 1, Unite: ingredient.Unite},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := ct.SetDefautProduit(ingredient.Id, produit.Id, true); err != nil {
		t.Fatal(err)
	}
}
