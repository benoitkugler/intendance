package datamodel

const (
	Litres Unite = "L"
	Kilos  Unite = "Kg"
	Piece  Unite = "P"
)

type Unite string

type Callories struct{}

// Lundi = 0, Mardi = 1 , etc...
type JoursLivraison [7]bool

type Conditionnement struct {
	Quantite float64
	Unite    Unite
}

func (c Conditionnement) IsNull() bool {
	return c == Conditionnement{}
}

// Horaire définie l'horaire d'un repas.
// Le frontend peut définir certains horaires classiques.
type Horaire struct {
	Heure  int
	Minute int
}
