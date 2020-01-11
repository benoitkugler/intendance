package controller

import (
	"fmt"
	"strings"
	"testing"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
)

func TestLoggin(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	s := Server{db: db}

	out, err := s.Loggin("mldks", "sdsd")
	if err != nil {
		t.Fatal(err)
	}
	if out.Erreur == "" {
		t.Error("expected error")
	}
	_, err = s.Loggin("test@intendance.fr", "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadData(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	s := Server{db: db}
	a, err := s.LoadAgendaUtilisateur(RequeteContext{idProprietaire: 2})
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
		out.Printf("Séjour %s, début : %s", s.Sejour.Nom, s.Sejour.DateDebut)
		out.indent++
		for _, j := range s.Journees {
			out.Printf("Journée %d", j.JourOffset)
			out.indent++
			for _, men := range j.Repas {
				out.Printf("Repas %v", men)
			}
			out.indent--
		}
		out.indent--
	}
	out.indent--
	return out.s.String()
}

func TestCRUD(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	s := Server{db: db}
	r := RequeteContext{idProprietaire: 2}

	ig, err := s.CreateIngredient(r)
	if err != nil {
		t.Fatal(err)
	}
	ig.Nom = "Tom acs tesd sdl sddddds ddsd35"
	ig.Unite = models.Kilos
	ig, err = s.UpdateIngredient(r, ig)
	if err != nil {
		t.Fatal(err)
	}
	re, err := s.CreateRecette(r)
	if err != nil {
		t.Fatal(err)
	}
	re.Nom = "Toma tes f a rcies2"
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.UpdateRecette(r, Recette{Recette: re, Ingredients: []models.RecetteIngredient{
		{IdIngredient: ig.Id, IdRecette: re.Id, Quantite: 4},
	}})
	if err != nil {
		t.Fatal(err)
	}
	m, err := s.CreateMenu(r)
	if err != nil {
		t.Fatal(err)
	}
	m.Commentaire = "Un menu bien équilibré"
	_, err = s.UpdateMenu(r, Menu{Menu: m, Recettes: []models.MenuRecette{
		{IdMenu: m.Id, IdRecette: re.Id},
	}, Ingredients: []models.MenuIngredient{
		{IdMenu: m.Id, IdIngredient: ig.Id},
	}})
	if err != nil {
		t.Fatal(err)
	}
	sej, err := s.CreateSejour(r)
	if err != nil {
		t.Fatal(err)
	}
	sej.Nom = "C2"
	if _, err = s.UpdateSejour(r, sej); err != nil {
		t.Fatal(err)
	}
	rep, err := s.CreateRepas(r, sej.Id, m.Id)
	if err != nil {
		t.Fatal(err)
	}
	rep.OffsetPersonnes = 55
	rep.Horaire = models.Horaire{Heure: 12, Minute: 5}
	if err = s.UpdateManyRepas(r, []models.Repas{rep}); err != nil {
		t.Fatal(err)
	}

	a, err := s.LoadAgendaUtilisateur(RequeteContext{idProprietaire: 2})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
	if err = s.DeleteRepas(r, rep.Id); err != nil {
		t.Fatal(err)
	}
	if err = s.DeleteSejour(r, sej.Id); err != nil {
		t.Fatal(err)
	}
	if err = s.DeleteMenu(r, m.Id); err != nil {
		t.Fatal(err)
	}
	if err = s.DeleteRecette(r, re.Id); err != nil {
		t.Fatal(err)
	}
	if err = s.DeleteIngredient(r, ig.Id, true); err != nil {
		t.Fatal(err)
	}
}
