package controller

import (
	"errors"
	"time"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
	"github.com/dgrijalva/jwt-go"
)

const DeltaToken = 72 * time.Hour

// UserMeta are custom claims extending default ones.
type UserMeta struct {
	IdProprietaire int64
	jwt.StandardClaims
}

func newToken(id int64) (string, error) {
	// Set custom claims
	claims := &UserMeta{
		IdProprietaire: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(DeltaToken).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString(logs.PASSPHRASE)
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
	return newToken(users.Ids()[0])
}
