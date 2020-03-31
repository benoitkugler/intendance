package controller

import (
	"fmt"
	"testing"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
)

func TestAssistant(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	s := Server{db: db}
	r := RequeteContext{idProprietaire: 2}

	sej, err := s.CreateSejour(r)
	if err != nil {
		t.Fatal(err)
	}

	groupe1, err := s.CreateGroupe(r, sej.Id)
	if err != nil {
		t.Fatal(err)
	}
	groupe2, err := s.CreateGroupe(r, sej.Id)
	if err != nil {
		t.Fatal(err)
	}
	groupe3, err := s.CreateGroupe(r, sej.Id)
	if err != nil {
		t.Fatal(err)
	}
	err = s.InitiateRepas(r, InAssistantCreateRepass{
		IdSejour: sej.Id,
		GroupesSorties: map[int][]int64{
			0: {groupe1.Id, groupe2.Id},
			1: {},
			2: {groupe3.Id},
			3: {groupe1.Id, groupe2.Id, groupe3.Id},
		},
		Options: OptionsAssistantCreateRepass{
			Duree:          5,
			Cinquieme:      models.Ids{groupe1.Id, groupe2.Id},
			DeleteExisting: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	a, err := s.LoadSejoursUtilisateur(RequeteContext{idProprietaire: 2})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}
