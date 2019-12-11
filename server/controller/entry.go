package controller

import "database/sql"

type Server struct {
	db *sql.DB
}

func (s Server) LoadAgendaUtilisateur() (AgendaUtilisateur, error) {
	tx, err := s.db.Begin()
	
}
