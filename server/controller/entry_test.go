package controller

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
)

// s.db should be close after use
func setupTest(t *testing.T) (Server, RequeteContext) {
	db, err := models.ConnectDB(logs.DB_DEV)
	if err != nil {
		t.Fatal(err)
	}
	s := Server{DB: db}
	ct := RequeteContext{IdProprietaire: 2, DB: db}
	return s, ct
}

func TestLoggin(t *testing.T) {
	db, err := models.ConnectDB(logs.DB_DEV)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	s := Server{DB: db}

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
	s, ct := setupTest(t)
	defer s.DB.Close()
	a, err := ct.LoadSejoursUtilisateur()
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
	s, ct := setupTest(t)
	defer s.DB.Close()

	ig, err := ct.CreateIngredient()
	if err != nil {
		t.Fatal(err)
	}
	ig.Nom = fmt.Sprintf("Tom acs tesd sdl sddddds ddsd35 %d", time.Now().UnixNano())
	ig.Unite = models.Kilos
	ig, err = ct.UpdateIngredient(ig)
	if err != nil {
		t.Fatal(err)
	}
	re, err := ct.CreateRecette()
	if err != nil {
		t.Fatal(err)
	}
	re.Nom = fmt.Sprintf("Toma tes f a rcies2 %d", time.Now().UnixNano())
	if err != nil {
		t.Fatal(err)
	}
	_, err = ct.UpdateRecette(RecetteComplet{Recette: re, Ingredients: models.LienIngredients{
		{IdIngredient: ig.Id, Quantite: 4},
	}})
	if err != nil {
		t.Fatal(err)
	}
	m, err := ct.CreateMenu()
	if err != nil {
		t.Fatal(err)
	}
	m.Commentaire = "Un menu bien équilibré"
	_, err = ct.UpdateMenu(MenuComplet{Menu: m, Recettes: models.Ids{
		re.Id,
	}, Ingredients: models.LienIngredients{
		{IdIngredient: ig.Id},
	}})
	if err != nil {
		t.Fatal(err)
	}
	sej, err := ct.CreateSejour()
	if err != nil {
		t.Fatal(err)
	}
	sej.Nom = "C2"
	if _, err = ct.UpdateSejour(sej); err != nil {
		t.Fatal(err)
	}

	groupe, err := ct.CreateGroupe(sej.Id)
	if err != nil {
		t.Fatal(err)
	}

	rep, err := ct.CreateRepas(sej.Id)
	if err != nil {
		t.Fatal(err)
	}
	rep.OffsetPersonnes = 55
	rep.Horaire = models.Midi
	if err = ct.UpdateManyRepas([]RepasComplet{{Repas: rep, Groupes: []models.RepasGroupe{
		{IdRepas: rep.Id, IdGroupe: groupe.Id}},
	}}); err != nil {
		t.Fatal(err)
	}

	a, err := ct.LoadSejoursUtilisateur()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
	if err = ct.DeleteRepas(rep.Id); err != nil {
		t.Fatal(err)
	}
	if _, err = ct.DeleteGroupe(groupe.Id); err != nil {
		t.Fatal(err)
	}
	if err = ct.DeleteSejour(sej.Id); err != nil {
		t.Fatal(err)
	}
	if err = ct.DeleteMenu(m.Id); err != nil {
		t.Fatal(err)
	}
	if err = ct.DeleteRecette(re.Id); err != nil {
		t.Fatal(err)
	}
	if err = ct.DeleteIngredient(ig.Id, true); err != nil {
		t.Fatal(err)
	}
}
