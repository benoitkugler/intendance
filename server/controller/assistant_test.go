package controller

import (
	"fmt"
	"testing"

	"github.com/benoitkugler/intendance/server/models"
)

func TestAssistant(t *testing.T) {
	s, ct := setupTest(t)
	defer s.db.Close()

	sej, err := s.CreateSejour(ct)
	if err != nil {
		t.Fatal(err)
	}

	groupe1, err := s.CreateGroupe(ct, sej.Id)
	if err != nil {
		t.Fatal(err)
	}
	groupe2, err := s.CreateGroupe(ct, sej.Id)
	if err != nil {
		t.Fatal(err)
	}
	groupe3, err := s.CreateGroupe(ct, sej.Id)
	if err != nil {
		t.Fatal(err)
	}
	err = s.InitiateRepas(ct, InAssistantCreateRepass{
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
	a, err := s.LoadSejoursUtilisateur(ct)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}
