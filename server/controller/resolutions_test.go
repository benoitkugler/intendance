package controller

import (
	"fmt"
	"testing"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
)

func TestResoudRepas(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	s := Server{db: db}
	row := s.db.QueryRow("SELECT * FROM repass LIMIT 1")
	repas, err := models.ScanRepas(row)
	if err != nil {
		t.Fatal(err)
	}
	res, err := s.ResoudIngredients(repas.Id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
