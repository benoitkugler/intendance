package controller

import (
	"fmt"
)

// Plusieurs items sont liées à un propriétaire.
// Comme les ids sont transmis, en clair,
// un utilisateur mal intentionné pourrait chercher
// à accéder à des ressources qui ne lui appartiennent pas.
// Les routines ci-dessous renvoient `nil` si et seulement si
// l'accès est légitimite.

func (s Server) proprioRecette(ct Requete, idRecette int64) error {
	row := ct.tx.QueryRow("SELECT id_proprietaire FROM recettes WHERE id = $1", idRecette)
	var trueProp int64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(err)
	}
	if trueProp != ct.idProprietaire {
		return fmt.Errorf(`Votre requête est impossible car la <b>recette</b> 
		concernée ne vous <b>appartient pas</b> !`)
	}
	return nil
}
