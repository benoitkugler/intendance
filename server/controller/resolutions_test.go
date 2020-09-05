package controller

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
)

func TestResoudRepas(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	s := Server{DB: db}
	row := s.DB.QueryRow("SELECT * FROM repass LIMIT 1")
	repas, err := models.ScanRepas(row)
	if err != nil {
		t.Fatal(err)
	}
	res, err := s.ResoudIngredientsRepas(repas.Id, -1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestJsonNil(t *testing.T) {
	var a []int64
	b, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	if err = json.Unmarshal(b, &a); err != nil {
		t.Fatal(err)
	}
	if a != nil {
		t.Errorf("expected nil, got %v", a)
	}
}

func TestResoudSejour(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	s := Server{DB: db}
	row := s.DB.QueryRow("SELECT * FROM repass LIMIT 1")
	repas, err := models.ScanRepas(row)
	if err != nil {
		t.Fatal(err)
	}
	res, err := s.ResoudIngredientsJournees(repas.IdSejour, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
