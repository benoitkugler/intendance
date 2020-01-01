package controller

import (
	"sort"

	"github.com/benoitkugler/intendance/server/models"
	"github.com/lib/pq"
)

// Ce fichier implémente le calcul des ingrédients pour un repas, une journée,
// ou un séjour, ainsi que l'association avec les produits correspondants.

func (s Server) ResoudIngredients(idRepas int64) ([]IngredientQuantite, error) {
	r := s.db.QueryRow("SELECT * FROM repass WHERE id = $1", idRepas)
	repas, err := models.ScanRepas(r)
	if err != nil {
		return nil, ErrorSQL(err)
	}

	rows, err := s.db.Query(`SELECT recette_ingredients.* FROM recette_ingredients 
	JOIN menu_recettes ON menu_recettes.id_recette = recette_ingredients.id_recette
	JOIN repass ON repass.id_menu = menu_recettes.id_menu 
	WHERE repass.id = $1`, idRepas)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	recetteIngredients, err := models.ScanRecetteIngredients(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	rows, err = s.db.Query(`SELECT menu_ingredients.* FROM menu_ingredients
	JOIN repass ON repass.id_menu = menu_ingredients.id_menu 
	WHERE repass.id = $1`, idRepas)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	menuIngredients, err := models.ScanMenuIngredients(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}

	total := map[int64]float64{} // idIngredient -> quantité pour le nombre de personnes souhaité
	ids := make(pq.Int64Array, 0, len(recetteIngredients)+len(menuIngredients))
	nbPersonnes := float64(repas.NbPersonnes)
	for _, rIng := range recetteIngredients {
		total[rIng.IdIngredient] += nbPersonnes * rIng.Quantite
		ids = append(ids, rIng.IdIngredient)
	}
	for _, rIng := range menuIngredients {
		total[rIng.IdIngredient] += nbPersonnes * rIng.Quantite
		ids = append(ids, rIng.IdIngredient)
	}
	rows, err = s.db.Query("SELECT * FROM ingredients WHERE id = ANY($1)", ids)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	ingredients, err := models.ScanIngredients(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	out := make([]IngredientQuantite, 0, len(ingredients))
	for _, ing := range ingredients {
		if quantite := total[ing.Id]; quantite > 0 {
			out = append(out, IngredientQuantite{Ingredient: ing, Quantite: quantite})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Quantite < out[j].Quantite })
	return out, nil
}
