package controller

import (
	"fmt"
	"strings"
	"testing"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
)

func TestLoadData(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	s := Server{db: db}
	r, err := s.NewRequete(2)
	if err != nil {
		t.Fatal(err)
	}
	a, err := s.loadAgendaUtilisateur(r)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

type agendaFormatter struct {
	s      strings.Builder
	indent int
}

func (s *agendaFormatter) Printf(format string, args ...interface{}) {
	c := fmt.Sprintf(format, args...)
	s.s.WriteString(strings.Repeat("\t", s.indent) + c + "\n")
}

func (a AgendaUtilisateur) String() string {
	var out agendaFormatter
	out.Printf("Séjours :")
	out.indent++
	for _, s := range a.Sejours {
		out.Printf("Séjour %s, début : %s", s.Nom, s.DateDebut)
		out.indent++
		for _, j := range s.Journees {
			out.Printf("Journée %d", j.JourOffset)
		}
		out.indent--
	}
	out.indent--
	return out.s.String()
}

func TestIngredients(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	s := Server{db: db}
	r, err := s.NewRequete(2)
	if err != nil {
		t.Fatal(err)
	}
	ig, err := s.createIngredient(r)
	ig.Nom = "Tomates"
	ig.Unite = models.Kilos
	err = s.updateIngredient(r, ig)
	re, err := s.createRecette(r)
	re.Nom = "Tomates farcies"
	err = s.updateRecette(r, re, []models.RecetteIngredient{
		{IdIngredient: ig.Id, IdRecette: re.Id, Quantite: 4},
	})
	r.tx.Commit()
	err = r.tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}
}
