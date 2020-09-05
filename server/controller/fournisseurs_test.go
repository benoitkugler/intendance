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
	defer s.DB.Close()
	out, err := ct.LoadFournisseurs()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}

func TestCRUDFournisseur(t *testing.T) {
	rand.Seed(time.Now().Unix())
	s, ct := setupTest(t)
	defer s.DB.Close()
	m, err := ct.CreateFournisseur(models.Fournisseur{Nom: fmt.Sprintf("GERLAND %d", rand.Int()), Lieu: "Lyon"})
	if err != nil {
		t.Fatal(err)
	}
	m.Nom = fmt.Sprintf("Intermarché %d", rand.Int())
	m, err = ct.UpdateFournisseur(m)
	if err != nil {
		t.Fatal(err)
	}
	err = ct.DeleteFournisseur(m.Id)
	if err != nil {
		t.Fatal(err)
	}
}
