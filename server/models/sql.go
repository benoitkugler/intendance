package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq" // side effect

	"github.com/benoitkugler/intendance/logs"
)

func ConnectDB(credences logs.SQL) (*sql.DB, error) {
	port := credences.Port
	if port == 0 {
		port = 5432
	}
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		credences.Host, port, credences.User, credences.Password, credences.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("connexion DB : %s", err)
	}
	return db, nil
}

// ------------------- Helpers --------------------------------

func ScanInts(rs *sql.Rows) ([]int64, error) {
	ints := make([]int64, 0, 16)
	var err error
	for rs.Next() {
		var s int64
		if err = rs.Scan(&s); err != nil {
			return nil, err
		}
		ints = append(ints, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return ints, nil
}

// GetProduits renvoie les produits associé à l'ingrédient.
// Seul le champ 'Id' est utilisé
func (ig Ingredient) GetProduits(tx *sql.Tx) (Produits, error) {
	rows, err := tx.Query(`SELECT produits.* FROM produits 
		JOIN ingredient_produits ON ingredient_produits.id_produit = produits.id 
		WHERE ingredient_produits.id_ingredient = $1`, ig.Id)
	if err != nil {
		return nil, err
	}
	return ScanProduits(rows)
}

// ------------------- Json encoding of custom types --------------------------

func loadJSON(out interface{}, src interface{}) error {
	if src == nil {
		return nil //zero value out
	}
	bs, ok := src.([]byte)
	if !ok {
		return errors.New("not a []byte")
	}
	return json.Unmarshal(bs, out)
}

func dumpJSON(s interface{}) (driver.Value, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return driver.Value(b), nil
}

func (s *Callories) Scan(src interface{}) error {
	return loadJSON(s, src)
}

func (s Callories) Value() (driver.Value, error) {
	return dumpJSON(s)
}

func (s *Conditionnement) Scan(src interface{}) error {
	return loadJSON(s, src)
}

func (s Conditionnement) Value() (driver.Value, error) {
	return dumpJSON(s)
}

func (s *Horaire) Scan(src interface{}) error {
	return loadJSON(s, src)
}

func (s Horaire) Value() (driver.Value, error) {
	return dumpJSON(s)
}

func (s *JoursLivraison) Scan(src interface{}) error {
	tmp := pq.BoolArray{}
	if err := tmp.Scan(src); err != nil {
		return err
	}
	if L := len(tmp); L != 7 {
		return fmt.Errorf("wrong length for JoursLivraison : expected 7, got %d", L)
	}
	for i, v := range tmp {
		(*s)[i] = v
	}
	return nil
}

func (s JoursLivraison) Value() (driver.Value, error) {
	tmp := make(pq.BoolArray, 7)
	for i, v := range s {
		tmp[i] = v
	}
	return tmp.Value()
}
