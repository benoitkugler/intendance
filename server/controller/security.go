package controller

import (
	"database/sql"
	"fmt"

	"github.com/benoitkugler/intendance/server/models"
)

// Plusieurs items sont liées à un propriétaire.
// Comme les ids sont transmis en clair,
// un utilisateur mal intentionné pourrait chercher
// à accéder à des ressources qui ne lui appartiennent pas.
// Les routines ci-dessous renvoient `nil` si et seulement si
// l'accès est légitimite.

// ct doit déjà être setup
func (ct RequeteContext) proprioRecette(recette models.Recette, checkProprioField bool) error {
	row := ct.DB.QueryRow("SELECT id_utilisateur FROM recettes WHERE id = $1", recette.Id)
	var trueProp sql.NullInt64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(err)
	}
	if trueProp.Valid && trueProp.Int64 != ct.IdProprietaire {
		return fmt.Errorf(`Votre requête est impossible car la <b>recette</b> 
		concernée ne vous <b>appartient pas</b> !`)
	}
	if checkProprioField && trueProp.Valid && ct.IdProprietaire != recette.IdUtilisateur.Int64 {
		return fmt.Errorf(`Votre requête est impossible car le propriétaire indiqué
		est <b>invalide</b> !`)
	}
	return nil
}

func (ct RequeteContext) proprioMenu(menu models.Menu, checkProprioField bool) error {
	row := ct.DB.QueryRow("SELECT id_utilisateur FROM menus WHERE id = $1", menu.Id)
	var trueProp sql.NullInt64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(err)
	}
	if trueProp.Valid && trueProp.Int64 != ct.IdProprietaire {
		return fmt.Errorf(`Votre requête est impossible car le <b>menu</b> 
		concerné ne vous <b>appartient pas</b> !`)
	}
	if checkProprioField && trueProp.Valid && ct.IdProprietaire != menu.IdUtilisateur.Int64 {
		return fmt.Errorf(`Votre requête est impossible car le propriétaire indiqué
		est <b>invalide</b> !`)
	}
	return nil
}

// Vérifie que le séjour donné appartient au propriétaire courant
// Si `checkProprioField`, vérifie aussi que le champ IdUtilisateur est cohérent.
func (ct RequeteContext) proprioSejour(sejour models.Sejour, checkProprioField bool) error {
	row := ct.DB.QueryRow("SELECT id_utilisateur FROM sejours WHERE id = $1", sejour.Id)
	var trueProp int64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(fmt.Errorf("can't load sejour utilisateur: %s", err))
	}
	if trueProp != ct.IdProprietaire {
		return fmt.Errorf(`Votre requête est impossible car le <b>séjour</b> 
		concerné ne vous <b>appartient pas</b> !`)
	}
	if checkProprioField && ct.IdProprietaire != sejour.IdUtilisateur {
		return fmt.Errorf(`Votre requête est impossible car le propriétaire indiqué
		est <b>invalide</b> !`)
	}
	return nil
}

func (ct RequeteContext) proprioGroupe(idGroupe int64) error {
	row := ct.DB.QueryRow(`SELECT sejours.id_utilisateur FROM sejours 
	JOIN groupes ON groupes.id_sejour = sejours.id
	WHERE groupes.id = $1`, idGroupe)
	var trueProp int64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(err)
	}
	if trueProp != ct.IdProprietaire {
		return fmt.Errorf(`Votre requête est impossible car le <b>séjour</b> 
		concerné ne vous <b>appartient pas</b> !`)
	}
	return nil
}

func (ct RequeteContext) proprioRepas(idRepas int64) error {
	row := ct.DB.QueryRow(`SELECT sejours.id_utilisateur FROM sejours 
	JOIN repass ON repass.id_sejour = sejours.id
	WHERE repass.id = $1`, idRepas)
	var trueProp int64
	if err := row.Scan(&trueProp); err != nil {
		return ErrorSQL(err)
	}
	if trueProp != ct.IdProprietaire {
		return fmt.Errorf(`Votre requête est impossible car le <b>séjour</b> 
		concerné ne vous <b>appartient pas</b> !`)
	}
	return nil
}
