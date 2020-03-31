package controller

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

func TestGetFournisseurs(t *testing.T) {
	s, ct := setupTest(t)
	defer s.db.Close()
	four, livr, err := s.LoadFournisseurs(ct)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(four, livr)
}

func TestCRUDFournisseur(t *testing.T) {
	rand.Seed(time.Now().Unix())
	s, ct := setupTest(t)
	defer s.db.Close()
	m, err := s.CreateFournisseur(ct, models.Fournisseur{Nom: fmt.Sprintf("GERLAND %d", rand.Int()), Lieu: "Lyon"})
	if err != nil {
		t.Fatal(err)
	}
	m.Nom = fmt.Sprintf("Intermarché %d", rand.Int())
	m, err = s.UpdateFournisseur(ct, m)
	if err != nil {
		t.Fatal(err)
	}
	err = s.DeleteFournisseur(ct, m.Id)
	if err != nil {
		t.Fatal(err)
	}
}
