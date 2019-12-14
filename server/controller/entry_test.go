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
	a, err := s.loadAgendaUtilisateur(2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

type agendaFormatter struct {
	s      strings.Builder
	indent int
}

func (s *agendaFormatter) Indent()  { s.indent++ }
func (s *agendaFormatter) DeIdent() { s.indent-- }

func (s *agendaFormatter) Printf(format string, args ...interface{}) {
	c := fmt.Sprintf(format, args...)
	s.s.WriteString(strings.Repeat("\t", s.indent) + c + "\n")
}

func (a AgendaUtilisateur) String() string {
	var out agendaFormatter
	out.Printf("Séjours :")
	out.Indent()
	for _, s := range a.Sejours {
		out.Printf("Séjour %s, début : %s", s.Nom, s.DateDebut)
		out.Indent()
		for _, j := range s.Journees {
			out.Printf("Journée %d", j.JourOffset)
		}
		out.DeIdent()
	}
	out.DeIdent()
	return out.s.String()
}
