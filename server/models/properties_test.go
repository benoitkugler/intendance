package models

import (
	"fmt"
	"testing"
	"time"
)

type testColisage struct {
	dem  float64
	cond float64
	col  int64
	exp  int64
}

// demandée, produit, colisage
var testsColisage = []testColisage{
	{dem: 2, cond: 2, col: 1, exp: 1},
	{dem: 2.2, cond: 2, col: 1, exp: 2},
	{dem: 1.8, cond: 2, col: 1, exp: 1},
	{dem: 2, cond: 2, col: 2, exp: 2},
	{dem: 2.5, cond: 1, col: 2, exp: 4},
}

func TestColisage(t *testing.T) {
	for _, test := range testsColisage {
		p := Produit{Conditionnement: Conditionnement{Quantite: test.cond}, Colisage: test.col}
		got := p.ColisageNeeded(test.dem)
		if got != test.exp {
			t.Errorf("expected %d, got %d", test.exp, got)
		}
	}
}

func TestDateLivraison(t *testing.T) {
	l := Livraison{Anticipation: 1, DelaiCommande: 4, JoursLivraison: JoursLivraison{true, true, false, false, false, false, false}}
	ti := time.Date(2020, time.March, 26, 0, 0, 0, 0, time.UTC)
	cd, lv := l.DateCommande(ti)
	fmt.Println("commande : ", cd)
	fmt.Println("livraison : ", lv)
}
