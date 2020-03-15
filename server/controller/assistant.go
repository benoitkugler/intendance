package controller

import (
	"github.com/benoitkugler/intendance/server/models"
)

type OptionsAssistantCreateRepass struct {
	Duree          int  `json:"duree"`
	WithCinquieme  bool `json:"with_cinquieme"`
	WithGouter     bool `json:"with_gouter"`
	DeleteExisting bool `json:"delete_existing"`
}

func (o OptionsAssistantCreateRepass) resoudHoraires() []models.Horaire {
	horaires := []models.Horaire{models.PetitDejeuner, models.Midi, models.Diner}
	if o.WithGouter {
		horaires = append(horaires, models.Gouter)
	}
	if o.WithCinquieme {
		horaires = append(horaires, models.Cinquieme)
	}
	return horaires
}

type InAssistantCreateRepass struct {
	IdSejour       int64                        `json:"id_sejour"`
	Options        OptionsAssistantCreateRepass `json:"options"`
	GroupesSorties map[int][]int64              `json:"groupes_sorties"` // offset -> ids_groupes
}

// crée un repas et y ajoute les groupes donnés
func creeRepasWithGroupe(ct RequeteContext, params InAssistantCreateRepass,
	horaire models.Horaire, jourOffset int, idsGroupes Set) error {
	repas := models.Repas{
		IdSejour:   params.IdSejour,
		Horaire:    horaire,
		JourOffset: int64(jourOffset),
	}
	repas, err := repas.Insert(ct.tx)
	if err != nil {
		return ErrorSQL(err)
	}
	var rg []models.RepasGroupe
	for idGroupe := range idsGroupes {
		rg = append(rg, models.RepasGroupe{IdGroupe: idGroupe, IdRepas: repas.Id})
	}
	if err = models.InsertManyRepasGroupes(ct.tx, rg); err != nil {
		return ErrorSQL(err)
	}
	return nil
}

func (s Server) InitiateRepas(ct RequeteContext, params InAssistantCreateRepass) error {
	if err := ct.beginTx(s); err != nil {
		return err
	}
	if err := s.proprioSejour(ct, models.Sejour{Id: params.IdSejour}, false); err != nil {
		return err
	}

	if params.Options.DeleteExisting {
		// on supprime tous les repas liés au séjour
		// lien repas-groupes
		_, err := s.db.Exec(`DELETE FROM repas_groupes WHERE id_repas = 
			ANY(SELECT id_repas FROM sejours WHERE id_sejour = $1)`, params.IdSejour)
		if err != nil {
			return ct.rollbackTx(err)
		}
		// repas
		_, err = s.db.Exec(`DELETE FROM repas WHERE id_sejour = $1`, params.IdSejour)
		if err != nil {
			return ct.rollbackTx(err)
		}
	}

	// on récupère les groupes du séjour
	rows, err := s.db.Query(`SELECT * FROM groupes WHERE id_sejour = $1`, params.IdSejour)
	if err != nil {
		return ct.rollbackTx(err)
	}
	groupes, err := models.ScanGroupes(rows)
	if err != nil {
		return ct.rollbackTx(err)
	}

	// on calcule les horaires à ajouter
	horaires := params.Options.resoudHoraires()

	for jourOffset := 0; jourOffset < params.Options.Duree; jourOffset++ {
		// calcule les deux listes ('basique' et 'sorties')
		sorties := NewSetFromSlice(params.GroupesSorties[jourOffset])
		basique := NewSet()
		for idGroupe := range groupes {
			if !sorties.Has(idGroupe) {
				basique.Add(idGroupe)
			}
		}

		for _, horaire := range horaires {
			if len(sorties) != 0 {
				// on crée le repas 'sorties'
				err = creeRepasWithGroupe(ct, params, horaire, jourOffset, sorties)
				if err != nil {
					return ct.rollbackTx(err)
				}
			}

			if len(sorties) == 0 || len(basique) > 0 {
				// 'sorties' vide ou 'basique' plein -> repas 'basique'
				err = creeRepasWithGroupe(ct, params, horaire, jourOffset, basique)
				if err != nil {
					return ct.rollbackTx(err)
				}
			}
		}
	}

	if err = ct.commitTx(); err != nil {
		return ct.rollbackTx(err)
	}
	return nil
}
