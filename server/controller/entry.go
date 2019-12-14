package controller

import (
	"database/sql"
	"fmt"

	"github.com/benoitkugler/intendance/server/models"
)

// Server est le controller principal, partagé par toutes les requêtes.
type Server struct {
	db *sql.DB
}

// Requete est créé pour chaque requête.
type Requete struct {
	tx             *sql.Tx
	idProprietaire int64
}

// rollback the current transaction, caused by `err`, and
// handles the possible error from tx.Rollback()
func rollback(tx *sql.Tx, origin error) error {
	if err := tx.Rollback(); err != nil {
		origin = fmt.Errorf("Rollback impossible. Erreur originale : %s", origin)
	}
	return ErrorSQL(origin)
}

// commit the transaction and try to rollback on error
func commit(tx *sql.Tx) error {
	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}
	return nil
}

func (s Server) NewRequete(idUtilisateur int64) (Requete, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return Requete{}, ErrorSQL(err)
	}
	return Requete{tx: tx, idProprietaire: idUtilisateur}, nil
}

func (s Server) loadAgendaUtilisateur(idProprietaire int64) (out AgendaUtilisateur, err error) {
	rows, err := s.db.Query("SELECT * FROM sejours WHERE id_proprietaire = $1", idProprietaire)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	sejours, err := models.ScanSejours(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	rows, err = s.db.Query(`SELECT menus.* FROM menus 
	JOIN sejour_menus ON sejour_menus.id_menu = menus.id 
	WHERE sejour_menus.id_sejour = ANY($1)`, sejours.Ids())
	menus, err := models.ScanMenus(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	rows, err = s.db.Query(`SELECT recettes.* FROM recettes 
	JOIN menu_recettes ON menu_recettes.id_recette = recettes.id 
	WHERE menu_recettes.id_menu = ANY($1)`, menus.Ids())
	recettes, err := models.ScanRecettes(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	rows, err = s.db.Query(`SELECT ingredients.* FROM ingredients 
		JOIN menu_ingredients ON menu_ingredients.id_ingredient = ingredients.id 
		WHERE menu_ingredients.id_menu = ANY($1)
		UNION
		SELECT ingredients.* FROM ingredients 
		JOIN recette_ingredients ON recette_ingredients.id_ingredient = ingredients.id 
		WHERE recette_ingredients.id_recette = ANY($2)`, menus.Ids(), recettes.Ids())
	ingredients, err := models.ScanIngredients(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	resolvedRecettes := make(map[int64]*Recette, len(recettes))
	for k, v := range recettes {
		resolvedRecettes[k] = &Recette{Recette: v}
	}
	rows, err = s.db.Query(`SELECT * FROM recette_ingredients WHERE id_recette = ANY($1)`, recettes.Ids())
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	ris, err := models.ScanRecetteIngredients(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	for _, l := range ris {
		resolvedRecettes[l.IdRecette].Ingredients = append(resolvedRecettes[l.IdRecette].Ingredients, IngredientRecette{
			Ingredient:        ingredients[l.IdIngredient],
			RecetteIngredient: l,
		})
	}

	resolvedMenus := make(map[int64]*Menu, len(menus))
	for k, v := range menus {
		resolvedMenus[k] = &Menu{Menu: v}
	}
	rows, err = s.db.Query(`SELECT * FROM menu_recettes WHERE id_menu = ANY($1)`, menus.Ids())
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	mrs, err := models.ScanMenuRecettes(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	for _, l := range mrs {
		resolvedMenus[l.IdMenu].Recettes = append(resolvedMenus[l.IdMenu].Recettes, *resolvedRecettes[l.IdRecette])
	}

	rows, err = s.db.Query(`SELECT * FROM menu_ingredients WHERE id_menu = ANY($1)`, menus.Ids())
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	mis, err := models.ScanMenuIngredients(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	for _, l := range mis {
		resolvedMenus[l.IdMenu].Ingredients = append(resolvedMenus[l.IdMenu].Ingredients, IngredientMenu{
			Ingredient:     ingredients[l.IdIngredient],
			MenuIngredient: l,
		})
	}

	resolvedSejours := make(map[int64]*Sejour, len(sejours))
	for k, v := range sejours {
		resolvedSejours[k] = &Sejour{Sejour: v, Journees: map[int64]Journee{}}
	}
	rows, err = s.db.Query(`SELECT * FROM sejour_menus WHERE id_sejour = ANY($1)`, sejours.Ids())
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	sms, err := models.ScanSejourMenus(rows)
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

// ------------------------------------------------------------------------
// --------------------- Ingrédients --------------------------------------

func (s Server) createIngredient(ct Requete) (out models.Ingredient, err error) {
	out, err = out.Insert(ct.tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = commit(ct.tx)
	return
}

func (s Server) updateIngredient(ct Requete, ig models.Ingredient) error {
	tx := ct.tx
	// vérification de la compatilibité des unités et des contionnements
	produits, err := ig.GetProduits(tx)
	if err != nil {
		return ErrorSQL(err)
	}
	for _, prod := range produits {
		if ig.Unite != models.Piece && ig.Unite != prod.Conditionnement.Unite {
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
	return commit(tx)
}

func (s Server) deleteIngredient(ct Requete, id int64, removeLiensProduits bool) error {
	tx := ct.tx

	var check ErrorIngredientUsed
	rows, err := tx.Query(`SELECT recettes.* FROM recettes 
	JOIN recette_ingredients ON recette_ingredients.id_recette = recettes.id 
	WHERE recette_ingredients.id_ingredient = $1`, id)
	if err != nil {
		return ErrorSQL(err)
	}
	check.recettes, err = models.ScanRecettes(rows)
	if err != nil {
		return ErrorSQL(err)
	}

	rows, err = tx.Query(`SELECT menus.* FROM menus 
	JOIN menu_ingredients ON menu_ingredients.id_menu = menus.id 
	WHERE menu_ingredients.id_ingredient = $1`, id)
	if err != nil {
		return ErrorSQL(err)
	}
	check.menus, err = models.ScanMenus(rows)
	if err != nil {
		return ErrorSQL(err)
	}

	ing := models.Ingredient{Id: id}
	check.produits, err = ing.GetProduits(tx)
	if err != nil {
		return ErrorSQL(err)
	}

	if removeLiensProduits { // on regarde uniquement les recettes et menus
		if len(check.recettes)+len(check.menus) > 0 {
			check.produits = nil
			return check
		}

		_, err = tx.Exec(`DELETE FROM ingredient_produits WHERE id_ingredient = $1`, id)
		if err != nil {
			return ErrorSQL(err)
		}
	} else { // on regarde aussi les produits
		if len(check.recettes)+len(check.menus)+len(check.produits) > 0 {
			return check
		}
	}
	// tout bon, on peut supprimer
	if _, err = ing.Delete(tx); err != nil {
		return rollback(tx, err)
	}
	return commit(tx)
}

// ------------------------------------------------------------------------
// ------------------------ Recettes --------------------------------------

func (s Server) createRecette(ct Requete) (out models.Recette, err error) {
	tx := ct.tx
	out.IdProprietaire = models.NullId(ct.idProprietaire)
	out, err = out.Insert(tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = commit(ct.tx)
	return
}

func (s Server) updateRecette(ct Requete, in models.Recette, ings []models.RecetteIngredient) error {
	if err := s.proprioRecette(ct, in.Id); err != nil {
		return err
	}
	tx := ct.tx
	//TODO: notification aux utilisateurs avec possibilité de copie
	in, err := in.Update(tx)
	if err != nil {
		return ErrorSQL(err)
	}
	_, err = tx.Exec("DELETE FROM recette_ingredients WHERE id_recette = $1", in.Id)
	if err != nil {
		return rollback(tx, err)
	}
	err = models.InsertManyRecetteIngredients(tx, ings)
	if err != nil {
		return rollback(tx, err)
	}
	return nil
}

func (s Server) deleteRecette(ct Requete, id int64) error {
	if err := s.proprioRecette(ct, id); err != nil {
		return err
	}
	rows, err := ct.tx.Query(`SELECT menus.id_proprietaire FROM menus 
	JOIN menu_recettes ON menu_recettes.id_menu = menus.id
	WHERE menu_recettes.id_recette = $1`, id)
	if err != nil {
		return ErrorSQL(err)
	}
	ids, err := models.ScanInts(rows)
	if err != nil {
		return ErrorSQL(err)
	}
	//TODO: notification aux utilisateurs avec possibilité de copie
	// nécessite de rassembler les données nécessaires à la re-création

	if L := len(ids); L > 0 {
		return fmt.Errorf(`Cette recette est présente dans <b>%d menu(s)</b>.
		Si vous souhaitez vraiment la supprimer, il faudra d'abord l'en retirer.`, L)
	}
	_, err = models.Recette{Id: id}.Delete(ct.tx)
	if err != nil {
		return ErrorSQL(err)
	}
	return commit(ct.tx)
}
