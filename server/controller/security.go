package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	mathRand "math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
)

const DeltaToken = 24 * time.Hour

func encrypt(data []byte) (string, error) {
	block, _ := aes.NewCipher(logs.PASSPHRASE)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

func decrypt(dataStr string) ([]byte, error) {
	data, err := base64.RawURLEncoding.DecodeString(dataStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(logs.PASSPHRASE)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(data) <= nonceSize {
		return nil, errors.New("data too short")
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

type tokenData struct {
	Salt          int32
	Time          time.Time
	IdUtilisateur int64
}

func creeToken(idUtilisateur int64) (string, error) {
	t := tokenData{
		Time:          time.Now(),
		IdUtilisateur: idUtilisateur,
		Salt:          mathRand.Int31(),
	}
	b, err := json.Marshal(t)
	if err != nil {
		return "", ErrorAuth(err)
	}
	return encrypt(b)
}

func refreshToken(token string, idUtilisateur int64) (newToken string, err error) {
	b, err := decrypt(token)
	if err != nil {
		return "", err
	}
	var data tokenData
	if err = json.Unmarshal(b, &data); err != nil {
		return "", err
	}
	if idUtilisateur != data.IdUtilisateur {
		return "", errors.New("Token corrompu : utilisateur invalide.")
	}
	diff := time.Since(data.Time)
	if diff > DeltaToken {
		diff = diff.Truncate(time.Second)
		return "", fmt.Errorf("Session écoulée (dernière action il y a %s). Veuillez vous reconnecter.", diff)
	}
	return creeToken(idUtilisateur)
}

// Authentifie vérifie le champ `BasicAuth` de `r`.
// Renvoie l'ID du propriétaire et un token mis à jour.
func (s Server) Authentifie(r *http.Request) (ct RequeteContext, err error) {
	idString, token, _ := r.BasicAuth()
	id0, err := strconv.Atoi(idString)
	if err != nil {
		return ct, ErrorAuth(err)
	}
	id := int64(id0)
	if s.devMode { // on autorise toutes les requêtes
		return RequeteContext{idProprietaire: id}, nil
	}
	token, err = refreshToken(token, id)
	if err != nil {
		return ct, ErrorAuth(err)
	}
	return RequeteContext{idProprietaire: id, Token: token}, nil
}

// Plusieurs items sont liées à un propriétaire.
// Comme les ids sont transmis, en clair,
// un utilisateur mal intentionné pourrait chercher
// à accéder à des ressources qui ne lui appartiennent pas.
// Les routines ci-dessous renvoient `nil` si et seulement si
// l'accès est légitimite.

// ct doit déjà être setup
func (ct RequeteContext) proprioRecette(recette models.Recette, checkProprioField bool) error {
	row := ct.tx.QueryRow("SELECT id_proprietaire FROM recettes WHERE id = $1", recette.Id)
	var trueProp sql.NullInt64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(err)
	}
	if trueProp.Valid && trueProp.Int64 != ct.idProprietaire {
		return fmt.Errorf(`Votre requête est impossible car la <b>recette</b> 
		concernée ne vous <b>appartient pas</b> !`)
	}
	if checkProprioField && trueProp.Valid && ct.idProprietaire != recette.IdProprietaire.Int64 {
		return fmt.Errorf(`Votre requête est impossible car le propriétaire indiqué
		est <b>invalide</b> !`)
	}
	return nil
}

func (ct RequeteContext) proprioMenu(menu models.Menu, checkProprioField bool) error {
	row := ct.tx.QueryRow("SELECT id_proprietaire FROM menus WHERE id = $1", menu.Id)
	var trueProp sql.NullInt64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(err)
	}
	if trueProp.Valid && trueProp.Int64 != ct.idProprietaire {
		return fmt.Errorf(`Votre requête est impossible car le <b>menu</b> 
		concerné ne vous <b>appartient pas</b> !`)
	}
	if checkProprioField && trueProp.Valid && ct.idProprietaire != menu.IdProprietaire.Int64 {
		return fmt.Errorf(`Votre requête est impossible car le propriétaire indiqué
		est <b>invalide</b> !`)
	}
	return nil
}

// Vérifie que le séjour donné appartient au propriétaire courant
// Si `checkProprioField`, vérifie aussi que le champ IdProprietaire est cohérent.
func (ct RequeteContext) proprioSejour(sejour models.Sejour, checkProprioField bool) error {
	row := ct.tx.QueryRow("SELECT id_proprietaire FROM sejours WHERE id = $1", sejour.Id)
	var trueProp int64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(err)
	}
	if trueProp != ct.idProprietaire {
		return fmt.Errorf(`Votre requête est impossible car le <b>séjour</b> 
		concerné ne vous <b>appartient pas</b> !`)
	}
	if checkProprioField && ct.idProprietaire != sejour.IdProprietaire {
		return fmt.Errorf(`Votre requête est impossible car le propriétaire indiqué
		est <b>invalide</b> !`)
	}
	return nil
}

func (ct RequeteContext) proprioGroupe(idGroupe int64) error {
	row := ct.tx.QueryRow(`SELECT sejours.id_proprietaire FROM sejours 
	JOIN groupes ON groupes.id_sejour = sejours.id
	WHERE groupes.id = $1`, idGroupe)
	var trueProp int64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(err)
	}
	if trueProp != ct.idProprietaire {
		return fmt.Errorf(`Votre requête est impossible car le <b>séjour</b> 
		concerné ne vous <b>appartient pas</b> !`)
	}
	return nil
}

func (ct RequeteContext) proprioRepas(idRepas int64) error {
	row := ct.tx.QueryRow(`SELECT sejours.id_proprietaire FROM sejours 
	JOIN repass ON repass.id_sejour = sejours.id
	WHERE repass.id = $1`, idRepas)
	var trueProp int64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(err)
	}
	if trueProp != ct.idProprietaire {
		return fmt.Errorf(`Votre requête est impossible car le <b>séjour</b> 
		concerné ne vous <b>appartient pas</b> !`)
	}
	return nil
}
