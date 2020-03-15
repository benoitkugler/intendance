package controller

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
)

func TestLoggin(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
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
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	s := Server{db: db}
	a, err := s.LoadSejoursUtilisateur(RequeteContext{idProprietaire: 2})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

type indentFormatter struct {
	s      strings.Builder
	indent int
}

func (s *indentFormatter) Printf(format string, args ...interface{}) {
	c := fmt.Sprintf(format, args...)
	s.s.WriteString(strings.Repeat("\t", s.indent) + c + "\n")
}

func (a Sejours) String() string {
	var out indentFormatter
	out.Printf("Séjours :")
	out.indent++
	for _, s := range a.Sejours {
		out.Printf("Séjour %s, début : %s", s.Sejour.Nom, s.Sejour.DateDebut)
		out.indent++
		for _, j := range s.Repass {
			out.Printf("Repas %v", j)

		}
		out.indent--
	}
	out.indent--
	out.Printf("Groupes :")
	out.indent++
	for _, g := range a.Groupes {
		out.Printf("Groupe %s (id séjour : %d)", g.Nom, g.IdSejour)
	}
	return out.s.String()
}

func TestCRUD(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	s := Server{db: db}
	r := RequeteContext{idProprietaire: 2}

	ig, err := s.CreateIngredient(r)
	if err != nil {
		t.Fatal(err)
	}
	ig.Nom = fmt.Sprintf("Tom acs tesd sdl sddddds ddsd35 %d", time.Now().UnixNano())
	ig.Unite = models.Kilos
	ig, err = s.UpdateIngredient(r, ig)
	if err != nil {
		t.Fatal(err)
	}
	re, err := s.CreateRecette(r)
	if err != nil {
		t.Fatal(err)
	}
	re.Nom = fmt.Sprintf("Toma tes f a rcies2 %d", time.Now().UnixNano())
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

	groupe, err := s.CreateGroupe(r, sej.Id)
	if err != nil {
		t.Fatal(err)
	}

	rep, err := s.CreateRepas(r, sej.Id, models.NullableId(m.Id))
	if err != nil {
		t.Fatal(err)
	}
	rep.OffsetPersonnes = 55
	rep.Horaire = models.Midi
	if err = s.UpdateManyRepas(r, []RepasWithGroupe{{Repas: rep, Groupes: []models.RepasGroupe{
		{IdRepas: rep.Id, IdGroupe: groupe.Id}},
	}}); err != nil {
		t.Fatal(err)
	}

	a, err := s.LoadSejoursUtilisateur(RequeteContext{idProprietaire: 2})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
	if err = s.DeleteRepas(r, rep.Id); err != nil {
		t.Fatal(err)
	}
	if _, err = s.DeleteGroupe(r, groupe.Id); err != nil {
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
