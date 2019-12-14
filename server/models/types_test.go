package models

import "testing"

func TestString(t *testing.T) {
	if Litres.String() != "L" {
		t.Error()
	}
	if Unite("").String() != "unité inconnue" {
		t.Error()
	}

	c := Conditionnement{Quantite: 0.4, Unite: Kilos}
	if s := c.String(); s != "0.40 Kg" {
		t.Errorf("got %s", s)
	}
	if c.IsNull() {
		t.Error()
	}
}
