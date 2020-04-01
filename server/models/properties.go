package models

import (
	"math"
	"time"
)

const jourDuration = 24 * time.Hour

// ColisageNeeded renvoie le nombre nécessaire d'exemplaire
// du produit pour obtenir (au mieux) `quantite`
// (exprimée dans l'unité du produit).
// Le colisage est pris en compte (c'est à dire que le résultat
// est un multiple du colisage du produit).
func (p Produit) ColisageNeeded(quantite float64) int64 {
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
	return nb
}

// BestJour remonte jusqu'au premier jour de livraison possible
// (incluant `date`).
func (js JoursLivraison) BestJour(dateLivraison time.Time) time.Time {
	wd := int(dateLivraison.Weekday())
	for i := 0; i < 7; i++ {
		index := (wd - i) % 7 // on remonte le temps
		if js[index] {        // jour ouvré
			return dateLivraison.Add(-time.Duration(index) * jourDuration)
		}
	}
	// `js` n'est jamais ouvert : cela ne devrait pas arriver
	return dateLivraison
}

// DateCommande calcule la date conseillée de commande la plus proche possible de `dateDemande`
// en tenant compte des contraintes.
func (l Livraison) DateCommande(dateDemande time.Time) time.Time {
	// on soustrait l'anticipation
	dateLivraison := dateDemande.Add(-time.Duration(l.Anticipation) * jourDuration)

	// on remonte jusqu'au premier jour de livraison possible
	dateLivraison = l.JoursLivraison.BestJour(dateLivraison)

	// on soustrait le délai
	dateLivraison = dateLivraison.Add(-time.Duration(l.DelaiCommande) * jourDuration)
	return dateLivraison
}
