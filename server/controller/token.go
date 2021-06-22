package controller

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	DeltaTokenJours = 3
	deltaToken      = DeltaTokenJours * 24 * time.Hour
)

// UserMeta are custom claims extending default ones.
type UserMeta struct {
	IdProprietaire int64
	jwt.StandardClaims
}

func (s Server) JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{SigningKey: logs.PASSPHRASE, Claims: &UserMeta{}}
	if s.Dev {
		config.SigningKey = logs.PASSPHRASE_DEV
	}
	return middleware.JWTWithConfig(config)
}

func (s Server) newToken(id int64) (string, error) {
	// Set custom claims
	claims := &UserMeta{
		IdProprietaire: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(deltaToken).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	pass := logs.PASSPHRASE
	if s.Dev {
		pass = logs.PASSPHRASE_DEV
	}
	return token.SignedString(pass)
}

// NewRequeteContext attend une requête authentifiée par JWT
func (s Server) NewRequeteContext(c echo.Context) RequeteContext {
	meta := c.Get("user").(*jwt.Token).Claims.(*UserMeta)
	return RequeteContext{IdProprietaire: meta.IdProprietaire, DB: s.DB}
}

// GetDevToken choisit un utilisateur au hasard et renvoie
// un token de connection
func (s Server) GetDevToken() (string, error) {
	users, err := models.SelectAllUtilisateurs(s.DB)
	if err != nil {
		return "", err
	}
	if len(users) == 0 {
		return "", errors.New("Aucun utilisateur n'est présent dans la base de données.")
	}
	id := users.Ids()[0]
	token, err := s.newToken(id)
	if err != nil {
		return "", err
	}
	type meta struct {
		User  int64  `json:"idUtilisateur"`
		Token string `json:"token"`
	}
	out, err := json.Marshal(meta{User: id, Token: token})
	return string(out), err
}
