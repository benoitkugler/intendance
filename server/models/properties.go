package models

import (
	"fmt"
	"math"
)

// ColisageNeeded renvoie le nombre nécessaire d'exemplaire
// du produit pour obtenir (au mieux) `quantite`
// (exprimée dans l'unité du produit).
// Le colisage est pris en compte (c'est à dire que le résultat
// est un multiple du colisage du produit).
func (p Produit) ColisageNeeded(quantite float64) (int64, error) {
	if p.Conditionnement.Quantite <= 0 {
		return 0, fmt.Errorf("Le conditionnement du produit <b>%s</b> est invalide : <i>%0.3f</i>", p.Nom, p.Conditionnement.Quantite)
	}
	// arrondi au supérieur pour ne pas manquer
	nb := int64(math.Ceil(quantite / p.Conditionnement.Quantite))
	colisage := p.Colisage
	if colisage == 0 { // la valeur par défaut pour le colisage est de 1
		colisage = 1
	}
	if reste := nb % colisage; reste != 0 {
		// on ajoute pour être un multiple
		nb = nb - reste + colisage
	}
	return nb, nil
}
