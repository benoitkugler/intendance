package models

import (
	"fmt"
)

type Unite string

func (u Unite) String() string {
	switch u {
	case Zero:
		return "<unité manquante>"
	case Litres:
		return "L"
	case Kilos:
		return "Kg"
	case Piece:
		return "Pièce(s)"
	}
	return "<unité inconnue>"
}

// Categorie permet de classer les ingrédients
type Categorie string

type Callories struct{}

// Lundi = 0, Mardi = 1 , etc...
type JoursLivraison [7]bool

type Conditionnement struct {
	Quantite float64 `json:"quantite"`
	Unite    Unite   `json:"unite"`
}

func (c Conditionnement) String() string {
	return fmt.Sprintf("%.02f %s", c.Quantite, c.Unite)
}

func (c Conditionnement) IsNull() bool {
	return c == Conditionnement{}
}

// Horaire définie l'horaire d'un repas.
type Horaire uint8
