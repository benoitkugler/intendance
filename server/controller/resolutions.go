package controller

import (
	"database/sql"
	"sort"

	"github.com/benoitkugler/intendance/server/models"
	"github.com/lib/pq"
)

// Ce fichier implémente le calcul des ingrédients pour un repas, une journée,
// ou un séjour, ainsi que l'association avec les produits correspondants.

// données nécessaire à la résolution des ingrédients
// d'un (ou plusieurs) repas
type dataRepas struct {
	repass             models.Repass
	groupes            models.Groupes
	menuIngredients    []models.MenuIngredient
	recetteIngredients []models.RecetteIngredient
	ingredients        models.Ingredients

	menuRecettes map[int64]models.Set // id menu -> ids recettes
	repasGroupes map[int64]models.Set // id repas -> ids groupes
}

// idIngredient -> quantité
type quantites = map[int64]float64

// ajoute id2 au crible de id1
func addToCribles(crible map[int64]models.Set, id1, id2 int64) {
	s := crible[id1]
	if s == nil {
		s = models.NewSet()
	}
	s.Add(id2)
	crible[id1] = s
}

// prend en entrée le résultat d'une requête renvoyant des repas
func (s Server) loadDataRepas(rowsRepas *sql.Rows) (out dataRepas, err error) {
	out.repass, err = models.ScanRepass(rowsRepas)
	if err != nil {
		return out, err
	}
	idRepass := make(pq.Int64Array, 0, len(out.repass))
	for idRepas := range out.repass {
		idRepass = append(idRepass, idRepas)
	}

	rows, err := s.db.Query(`SELECT * FROM repas_groupes WHERE id_repas = ANY($1)`, idRepass)
	if err != nil {
		return out, err
	}
	tmp, err := models.ScanRepasGroupes(rows)
	if err != nil {
		return out, err
	}
	out.repasGroupes = make(map[int64]models.Set)
	for _, rg := range tmp {
		addToCribles(out.repasGroupes, rg.IdRepas, rg.IdGroupe)
	}

	rows, err = s.db.Query(`SELECT groupes.* FROM groupes 
		JOIN repas_groupes ON repas_groupes.id_groupe = groupes.id
		WHERE repas_groupes.id_repas = ANY($1)`, idRepass)
	if err != nil {
		return out, err
	}
	out.groupes, err = models.ScanGroupes(rows)
	if err != nil {
		return out, err
	}

	rows, err = s.db.Query(`SELECT menu_ingredients.* FROM menu_ingredients
	JOIN repass ON repass.id_menu = menu_ingredients.id_menu 
	WHERE repass.id = ANY($1)`, idRepass)
	if err != nil {
		return out, err
	}
	out.menuIngredients, err = models.ScanMenuIngredients(rows)
	if err != nil {
		return out, err
	}

	rows, err = s.db.Query(`SELECT menu_recettes.* FROM menu_recettes
	JOIN repass ON repass.id_menu = menu_recettes.id_menu 
	WHERE repass.id = ANY($1)`, idRepass)
	if err != nil {
		return out, err
	}
	menuRecettes, err := models.ScanMenuRecettes(rows)
	if err != nil {
		return out, err
	}

	out.menuRecettes = make(map[int64]models.Set)
	for _, menuRecette := range menuRecettes {
		addToCribles(out.menuRecettes, menuRecette.IdMenu, menuRecette.IdRecette)
	}

	rows, err = s.db.Query(`SELECT recette_ingredients.* FROM recette_ingredients 
	JOIN menu_recettes ON menu_recettes.id_recette = recette_ingredients.id_recette
	JOIN repass ON repass.id_menu = menu_recettes.id_menu 
	WHERE repass.id = ANY($1)`, idRepass)
	if err != nil {
		return out, err
	}
	out.recetteIngredients, err = models.ScanRecetteIngredients(rows)
	if err != nil {
		return out, err
	}

	idsIngredients := models.NewSet()
	for _, ing := range out.menuIngredients {
		idsIngredients.Add(ing.IdIngredient)
	}
	for _, ing := range out.recetteIngredients {
		idsIngredients.Add(ing.IdIngredient)
	}
	rows, err = s.db.Query("SELECT * FROM ingredients WHERE id = ANY($1)", pq.Int64Array(idsIngredients.Keys()))
	if err != nil {
		return out, err
	}
	out.ingredients, err = models.ScanIngredients(rows)
	if err != nil {
		return out, err
	}
	return out, nil
}

// si `nbPersonnes` vaut -1, le nombre de personne du repas est utilisé
func (d dataRepas) resoudRepas(idRepas, nbPersonnes int64, quantite quantites) {
	repas := d.repass[idRepas]
	if !repas.IdMenu.Valid { // on ignore ce repas
		return
	}

	if nbPersonnes == -1 { // on résoud le nombre de personnes
		nbPersonnes = repas.OffsetPersonnes
		for idGroupe := range d.repasGroupes[idRepas] { // parcourt des groupes
			nbPersonnes += d.groupes[idGroupe].NbPersonnes
		}
		if nbPersonnes < 0 {
			nbPersonnes = 0
		}
	}
	nbPersonnesF := float64(nbPersonnes)
	cribleRecettes := d.menuRecettes[repas.IdMenu.Int64]
	for _, recetteIngredient := range d.recetteIngredients {
		if cribleRecettes.Has(recetteIngredient.IdRecette) {
			quantite[recetteIngredient.IdIngredient] += nbPersonnesF * recetteIngredient.Quantite
		}
	}
	for _, menuIngredient := range d.menuIngredients {
		if menuIngredient.IdMenu == repas.IdMenu.Int64 {
			quantite[menuIngredient.IdIngredient] += nbPersonnesF * menuIngredient.Quantite
		}
	}
}

func (d dataRepas) formatQuantites(quantites quantites) []IngredientQuantite {
	out := make([]IngredientQuantite, 0, len(quantites))
	// filtre les ingrédients inutiles
	for idIngredient, quantite := range quantites {
		if quantite > 0 {
			out = append(out, IngredientQuantite{Ingredient: d.ingredients[idIngredient], Quantite: quantite})
		}
	}
	// par ordre décroissant
	sort.Slice(out, func(i, j int) bool { return out[i].Quantite > out[j].Quantite })
	return out
}

func (s Server) ResoudIngredientsRepas(idRepas, nbPersonnes int64) ([]IngredientQuantite, error) {
	rows, err := s.db.Query(`SELECT * FROM repass WHERE id = $1`, idRepas)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	data, err := s.loadDataRepas(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	// idIngredient -> quantité pour le nombre de personnes souhaité
	quantites := quantites{}
	data.resoudRepas(idRepas, nbPersonnes, quantites)
	out := data.formatQuantites(quantites)
	return out, nil
}

// ResoudIngredientsJournees renvoies le total des ingrédients.
// Si `journeesOffsets` vaut nil, tout le séjour est utilisé.
func (s Server) ResoudIngredientsJournees(idSejour int64, journeesOffsets []int64) ([]DateIngredientQuantites, error) {
	r := s.db.QueryRow("SELECT * FROM sejours WHERE id = $1", idSejour)
	sejour, err := models.ScanSejour(r)
	if err != nil {
		return nil, ErrorSQL(err)
	}

	rows, err := s.db.Query(`SELECT * FROM repass WHERE repass.id_sejour = $1`, idSejour)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	data, err := s.loadDataRepas(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}

	all := journeesOffsets == nil
	crible := models.NewSetFromSlice(journeesOffsets)
	joursQuantites := map[int64]quantites{}
	for _, repas := range data.repass {
		if all || crible.Has(repas.JourOffset) {
			quantite := joursQuantites[repas.JourOffset]
			if quantite == nil {
				quantite = map[int64]float64{}
			}
			data.resoudRepas(repas.Id, -1, quantite)
			joursQuantites[repas.JourOffset] = quantite
		}
	}

	out := make([]DateIngredientQuantites, 0, len(joursQuantites))
	for jourOffset, quantites := range joursQuantites {
		ings := data.formatQuantites(quantites)
		if len(ings) == 0 {
			continue
		}
		out = append(out, DateIngredientQuantites{
			Ingredients: ings,
			Date:        sejour.DateFromOffset(jourOffset),
		})
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Date.Before(out[j].Date)
	})
	return out, nil
}
