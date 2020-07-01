package models

const (
	PetitDejeuner Horaire = iota // Petit déjeuner
	Midi                         // Midi
	Gouter                       // Goûter
	Diner                        // Dîner
	Cinquieme                    // Cinquième
)

const (
	Zero   Unite = ""   // Unité invalide
	Litres Unite = "L"  // Litres
	Kilos  Unite = "Kg" // Kilos
	Piece  Unite = "P"  // Pièce(s)
)
