package controller

import (
	"database/sql"

	"github.com/benoitkugler/intendance/server/datamodel"
)

type Server struct {
	db *sql.DB
}

func (s Server) loadAgendaUtilisateur(idUtilisateur int64) (out AgendaUtilisateur, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	rows, err := tx.Query("SELECT * FROM sejours WHERE id_proprietaire = $1", idUtilisateur)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	sejours, err := datamodel.ScanSejours(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	rows, err = tx.Query(`SELECT menus.* FROM menus 
	JOIN sejours_menus ON sejours_menus.id_menu = menus.id 
	WHERE sejours_menus.id_sejour = ANY($1)`, sejours.Ids())
	menus, err := datamodel.ScanMenus(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	rows, err = tx.Query(`SELECT recettes.* FROM recettes 
	JOIN menus_recettes ON menus_recettes.id_recette = recettes.id 
	WHERE menus_recettes.id_menu = ANY($1)`, menus.Ids())
	recettes, err := datamodel.ScanRecettes(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	rows, err = tx.Query(`SELECT ingredients.* FROM ingredients 
		JOIN menus_ingredients ON menus_ingredients.id_ingredient = ingredients.id 
		WHERE menus_ingredients.id_menu = ANY($1)
		UNION
		SELECT ingredients.* FROM ingredients 
		JOIN recettes_ingredients ON recettes_ingredients.id_ingredient = ingredients.id 
		WHERE recettes_ingredients.id_recette = ANY($2)`, menus.Ids(), recettes.Ids())
	ingredients, err := datamodel.ScanIngredients(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	resolvedRecettes := make(map[int64]*Recette, len(recettes))
	for k, v := range recettes {
		resolvedRecettes[k] = &Recette{Recette: v}
	}
	rows, err = tx.Query(`SELECT * FROM recettes_ingredients WHERE id_recette = ANY($1)`, recettes.Ids())
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	ris, err := datamodel.ScanRecetteIngredients(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	for _, l := range ris {
		resolvedRecettes[l.IdRecette].Ingredients = append(resolvedRecettes[l.IdRecette].Ingredients, IngredientRecette{
			Ingredient: ingredients[l.IdIngredient],
			Quantite:   l.Quantite,
			Cuisson:    l.Cuisson,
		})
	}

	resolvedMenus := make(map[int64]*Menu, len(menus))
	for k, v := range menus {
		resolvedMenus[k] = &Menu{Menu: v}
	}
	rows, err = tx.Query(`SELECT * FROM menus_recettes WHERE id_menu = ANY($1)`, menus.Ids())
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	mrs, err := datamodel.ScanMenuRecettes(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	for _, l := range mrs {
		resolvedMenus[l.IdMenu].Recettes = append(resolvedMenus[l.IdMenu].Recettes, *resolvedRecettes[l.IdRecette])
	}

	rows, err = tx.Query(`SELECT * FROM menus_ingredients WHERE id_menu = ANY($1)`, menus.Ids())
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	mis, err := datamodel.ScanMenuIngredients(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	for _, l := range mis {
		resolvedMenus[l.IdMenu].Ingredients = append(resolvedMenus[l.IdMenu].Ingredients, IngredientRecette{
			Ingredient: ingredients[l.IdIngredient],
			Quantite:   l.Quantite,
			Cuisson:    l.Cuisson,
		})
	}

	resolvedSejours := make(map[int64]*Sejour, len(sejours))
	for k, v := range sejours {
		resolvedSejours[k] = &Sejour{Sejour: v, Journees: map[int64]Journee{}}
	}
	rows, err = tx.Query(`SELECT * FROM sejours_menus WHERE id_sejour = ANY($1)`, sejours.Ids())
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	sms, err := datamodel.ScanSejourMenus(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	for _, l := range sms {
		m := resolvedMenus[l.IdMenu]
		m.NbPersonnes = l.NbPersonnes
		m.Horaire = l.Horaire

		journee := resolvedSejours[l.IdSejour].Journees[l.JourOffset]
		journee.JourOffset = l.JourOffset
		journee.Menus = append(journee.Menus, *m)
		resolvedSejours[l.IdSejour].Journees[l.JourOffset] = journee
	}

	for _, v := range resolvedSejours {
		out.Sejours = append(out.Sejours, *v)
	}
	return out, nil
}

func (s Server) createIngredient() (out datamodel.Ingredient, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	out, err = out.Insert(tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = tx.Commit()
	return
}

func (s Server) updateIngredient(ig datamodel.Ingredient) error {
	tx, err := s.db.Begin()
	if err != nil {
		return ErrorSQL(err)
	}
	// vérification de la compatilibité des unités et des contionnements
	rows, err := tx.Query(`SELECT produits.* FROM produits 
	JOIN ingredients_produits ON ingredients_produits.id_produit = produits.id 
	WHERE ingredients_produits = $1`, ig.Id)
	if err != nil {
		return ErrorSQL(err)
	}
	produits, err := datamodel.ScanProduits(rows)
	if err != nil {
		return ErrorSQL(err)
	}
	for _, prod := range produits {
		if ig.Unite != datamodel.Piece && ig.Unite != prod.Conditionnement.Unite {
			return ErrorIngredientProduitUnite{ingredient: ig, produit: prod}
		}
		if !ig.Conditionnement.IsNull() && ig.Conditionnement != prod.Conditionnement {
			return ErrorIngredientProduitConditionnement{ingredient: ig, produit: prod}
		}
	}

	// modification
	ig, err = ig.Update(tx)
	if err != nil {
		return ErrorSQL(err)
	}
	if err = tx.Commit(); err != nil {
		return ErrorSQL(err)
	}
	return nil
}

func (s Server) deleteIngredient(id int64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return ErrorSQL(err)
	}

	rows, err := tx.Query("SELECT recettes.* JOIN ")
}
