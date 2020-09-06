package controller

import (
	"database/sql"
	"fmt"
)

// Server est le controller principal, partagé par toutes les requêtes.
type Server struct {
	DB  *sql.DB
	Dev bool // contourne l'authentification
}

// RequeteContext est créé pour chaque requête.
type RequeteContext struct {
	IdProprietaire int64
	DB             *sql.DB
}

type Tx struct {
	*sql.Tx
	IdProprietaire int64
}

func (ct RequeteContext) beginTx() (Tx, error) {
	tx, err := ct.DB.Begin()
	if err != nil {
		return Tx{}, ErrorSQL(err)
	}
	return Tx{Tx: tx, IdProprietaire: ct.IdProprietaire}, nil
}

// rollbackTx the current transaction, caused by `err` (wrapping it), and
// handles the possible error from tx.Rollback()
func (tx Tx) rollback(origin error) error {
	if err := tx.Rollback(); err != nil {
		origin = fmt.Errorf("Rollback impossible. Erreur originale : %s", origin)
	}
	if _, ok := origin.(errorSQL); ok { // pas besoin de wrapper
		return origin
	}
	return ErrorSQL(origin)
}
