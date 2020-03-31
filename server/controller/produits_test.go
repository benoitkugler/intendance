package controller

import (
	"fmt"
	"testing"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
)

func TestGetProduits(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	s := Server{db: db}
	ct := RequeteContext{idProprietaire: 2}
	out, err := s.GetIngredientProduits(ct, 56)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}

func TestDelete(t *testing.T) {
	s, ct := setupTest(t)
	defer s.db.Close()
	if err := s.DeleteProduit(ct, 67); err != nil {
		t.Fatal(err)
	}
}
