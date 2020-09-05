package controller

import (
	"database/sql"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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

// NewRequeteContext attend une requête authentifiée par JWT
func (s Server) NewRequeteContext(c echo.Context) RequeteContext {
	meta := c.Get("user").(*jwt.Token).Claims.(*UserMeta)
	return RequeteContext{IdProprietaire: meta.IdProprietaire, DB: s.DB}
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
