package controller

import (
	"github.com/benoitkugler/intendance/server/models"
)

type OptionsAssistantCreateRepass struct {
	Duree          int        `json:"duree"`
	WithGouter     bool       `json:"with_gouter"`
	Cinquieme      models.Ids `json:"cinquieme"` // ids groupes
	DeleteExisting bool       `json:"delete_existing"`
}

// le cas du cinquième est à  part
func (o OptionsAssistantCreateRepass) resoudHoraires() []models.Horaire {
	horaires := []models.Horaire{models.PetitDejeuner, models.Midi, models.Diner}
	if o.WithGouter {
		horaires = append(horaires, models.Gouter)
	}
	return horaires
}

// crée un repas et y ajoute les groupes donnés
// ne commit pas; ne rollback pas
func (tx Tx) creeRepasComplet(params InAssistantCreateRepass,
	horaire models.Horaire, jourOffset int, idsGroupes models.Set) error {
	repas := models.Repas{
		IdSejour:   params.IdSejour,
		Horaire:    horaire,
		JourOffset: int64(jourOffset),
	}
	repas, err := repas.Insert(tx)
	if err != nil {
		return ErrorSQL(err)
	}
	var rg []models.RepasGroupe
	for idGroupe := range idsGroupes {
		rg = append(rg, models.RepasGroupe{IdGroupe: idGroupe, IdRepas: repas.Id})
	}
	if err = models.InsertManyRepasGroupes(tx.Tx, rg...); err != nil {
		return ErrorSQL(err)
	}
	return nil
}

func (ct RequeteContext) InitiateRepas(params InAssistantCreateRepass) error {
	if err := ct.proprioSejour(models.Sejour{Id: params.IdSejour}, false); err != nil {
		return err
	}

	tx, err := ct.beginTx()
	if err != nil {
		return err
	}
	if params.Options.DeleteExisting {
		// on supprime tous les repas liés au séjour
		// lien repas-groupes
		_, err := tx.Exec(`DELETE FROM repas_groupes WHERE id_repas = 
			ANY(SELECT id_repas FROM sejours WHERE id = $1)`, params.IdSejour)
		if err != nil {
			return tx.rollback(err)
		}

		// lien repas-recettes
		_, err = tx.Exec(`DELETE FROM repas_recettes WHERE id_repas = 
		ANY(SELECT id_repas FROM sejours WHERE id = $1)`, params.IdSejour)
		if err != nil {
			return tx.rollback(err)
		}

		// lien repas-ingredients
		_, err = tx.Exec(`DELETE FROM repas_ingredients WHERE id_repas = 
		ANY(SELECT id_repas FROM sejours WHERE id = $1)`, params.IdSejour)
		if err != nil {
			return tx.rollback(err)
		}

		// repas
		_, err = tx.Exec(`DELETE FROM repass WHERE id_sejour = $1`, params.IdSejour)
		if err != nil {
			return tx.rollback(err)
		}
	}

	// on récupère les groupes du séjour
	groupes, err := models.SelectGroupesByIdSejours(tx, params.IdSejour)
	if err != nil {
		return tx.rollback(err)
	}

	// on calcule les horaires à ajouter
	horaires := params.Options.resoudHoraires()

	for jourOffset := 0; jourOffset < params.Options.Duree; jourOffset++ {
		// calcule les deux listes ('basique' et 'sorties')
		sorties := models.NewSetFromSlice(params.GroupesSorties[jourOffset])
		basique := models.NewSet()
		for idGroupe := range groupes {
			if !sorties.Has(idGroupe) {
				basique.Add(idGroupe)
			}
		}

		for _, horaire := range horaires {
			if len(sorties) != 0 {
				// on crée le repas 'sorties'
				err = tx.creeRepasComplet(params, horaire, jourOffset, sorties)
				if err != nil {
					return tx.rollback(err)
				}
			}

			if len(sorties) == 0 || len(basique) > 0 {
				// 'sorties' vide ou 'basique' plein -> repas 'basique'
				err = tx.creeRepasComplet(params, horaire, jourOffset, basique)
				if err != nil {
					return tx.rollback(err)
				}
			}
		}

		// cas du cinquième
		if groupes5 := params.Options.Cinquieme; len(groupes5) > 0 {
			err = tx.creeRepasComplet(params, models.Cinquieme, jourOffset, groupes5.AsSet())
			if err != nil {
				return tx.rollback(err)
			}
		}
	}

	err = tx.Commit()
	return err
}
