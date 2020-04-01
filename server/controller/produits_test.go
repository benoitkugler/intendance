package controller

import (
	"fmt"
	"testing"
)

func TestGetProduits(t *testing.T) {
	s, ct := setupTest(t)
	defer s.db.Close()
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

func TestDefault(t *testing.T) {
	s, ct := setupTest(t)
	defer s.db.Close()
	if err := s.SetDefautProduit(ct, 2, 2, true); err != nil {
		t.Fatal(err)
	}
}
