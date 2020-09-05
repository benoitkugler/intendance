package controller

import (
	"fmt"
	"testing"

	"github.com/benoitkugler/intendance/server/models"
)

func TestAssistant(t *testing.T) {
	s, ct := setupTest(t)
	defer s.DB.Close()

	sej, err := ct.CreateSejour()
	if err != nil {
		t.Fatal(err)
	}

	groupe1, err := ct.CreateGroupe(sej.Id)
	if err != nil {
		t.Fatal(err)
	}
	groupe2, err := ct.CreateGroupe(sej.Id)
	if err != nil {
		t.Fatal(err)
	}
	groupe3, err := ct.CreateGroupe(sej.Id)
	if err != nil {
		t.Fatal(err)
	}
	err = ct.InitiateRepas(InAssistantCreateRepass{
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
	a, err := ct.LoadSejoursUtilisateur()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}
