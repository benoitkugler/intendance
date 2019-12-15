package controller

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

// Server est le controller principal, partagé par toutes les requêtes.
type Server struct {
	db *sql.DB
}

// RequeteContext est créé pour chaque requête.
type RequeteContext struct {
	idProprietaire int64
	tx             *sql.Tx // a créer
}

// rollback the current transaction, caused by `err`, and
// handles the possible error from tx.Rollback()
func (r RequeteContext) rollback(origin error) error {
	if err := r.tx.Rollback(); err != nil {
		origin = fmt.Errorf("Rollback impossible. Erreur originale : %s", origin)
	}
	if _, ok := origin.(errorSQL); ok { // pas besoin de wrapper
		return origin
	}
	return ErrorSQL(origin)
}

// commit the transaction and try to rollback on error
func (r RequeteContext) commit() error {
	if err := r.tx.Commit(); err != nil {
		return r.rollback(err)
	}
	return nil
}

func (ct *RequeteContext) setup(s Server) (err error) {
	ct.tx, err = s.db.Begin()
	if err != nil {
		return ErrorSQL(err)
	}
	return nil
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
	JOIN repass ON repass.id_menu = menus.id 
	WHERE repass.id_sejour = ANY($1)`, sejours.Ids())
	if err != nil {
		err = ErrorSQL(err)
		return
	}
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

	resolvedMenus := make(map[int64]*Repas, len(menus))
	for k, v := range menus {
		resolvedMenus[k] = &Repas{Menu: v}
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
	rows, err = s.db.Query(`SELECT * FROM repass WHERE id_sejour = ANY($1)`, sejours.Ids())
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	sms, err := models.ScanRepass(rows)
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
// ------------------------------------------------------------------------

func (s Server) createIngredient(ct RequeteContext) (out models.Ingredient, err error) {
	if err = ct.setup(s); err != nil {
		return
	}
	out.Nom = fmt.Sprintf("I%d", time.Now().UnixNano())
	out, err = out.Insert(ct.tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = ct.commit()
	return
}

func (s Server) updateIngredient(ct RequeteContext, ig models.Ingredient) error {
	if err := ct.setup(s); err != nil {
		return err
	}
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
	return ct.commit()
}

func (s Server) deleteIngredient(ct RequeteContext, id int64, removeLiensProduits bool) error {
	if err := ct.setup(s); err != nil {
		return err
	}
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
		return ct.rollback(err)
	}
	return ct.commit()
}

// ------------------------------------------------------------------------
// ------------------------ Recettes --------------------------------------
// ------------------------------------------------------------------------

func (s Server) createRecette(ct RequeteContext) (out models.Recette, err error) {
	if err = ct.setup(s); err != nil {
		return
	}
	tx := ct.tx
	out.Nom = fmt.Sprintf("R%d", time.Now().UnixNano())
	out.IdProprietaire = models.NullId(ct.idProprietaire)
	out, err = out.Insert(tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = ct.commit()
	return
}

func (s Server) updateRecette(ct RequeteContext, in models.Recette, ings []models.RecetteIngredient) error {
	for _, r := range ings {
		if r.IdRecette != in.Id {
			return fmt.Errorf("L'ingrédient %d n'est pas associé à la recette fournie !", r.IdIngredient)
		}
	}
	if err := ct.setup(s); err != nil {
		return err
	}
	if err := s.proprioRecette(ct, in, true); err != nil {
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
		return ct.rollback(err)
	}
	err = models.InsertManyRecetteIngredients(tx, ings)
	if err != nil {
		return ct.rollback(err)
	}
	return ct.commit()
}

func (s Server) deleteRecette(ct RequeteContext, id int64) error {
	if err := ct.setup(s); err != nil {
		return err
	}
	if err := s.proprioRecette(ct, models.Recette{Id: id}, false); err != nil {
		return err
	}
	rows, err := ct.tx.Query(`SELECT menus.id FROM menus 
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
	_, err = ct.tx.Exec("DELETE FROM recette_ingredients WHERE id_recette = $1", id)
	if err != nil {
		return ErrorSQL(err)
	}
	_, err = models.Recette{Id: id}.Delete(ct.tx)
	if err != nil {
		return ct.rollback(err)
	}
	return ct.commit()
}

// ------------------------------------------------------------------------
// ----------------------------- Menus ------------------------------------
// ------------------------------------------------------------------------

func (s Server) createMenu(ct RequeteContext) (out models.Menu, err error) {
	if err = ct.setup(s); err != nil {
		return
	}
	tx := ct.tx
	out.IdProprietaire = models.NullId(ct.idProprietaire)
	out, err = out.Insert(tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = ct.commit()
	return
}

func (s Server) updateMenu(ct RequeteContext, in models.Menu,
	recettes []models.MenuRecette, ings []models.MenuIngredient) error {
	for _, r := range recettes {
		if r.IdMenu != in.Id {
			return fmt.Errorf("La recette %d n'est pas associée au menu fourni !", r.IdRecette)
		}
	}
	for _, r := range ings {
		if r.IdMenu != in.Id {
			return fmt.Errorf("L'ingrédient %d n'est pas associé au menu fourni !", r.IdIngredient)
		}
	}
	if err := ct.setup(s); err != nil {
		return err
	}
	if err := s.proprioMenu(ct, in, true); err != nil {
		return err
	}
	tx := ct.tx
	//TODO: notification aux utilisateurs avec possibilité de copie
	in, err := in.Update(tx)
	if err != nil {
		return ErrorSQL(err)
	}
	_, err = tx.Exec("DELETE FROM menu_recettes WHERE id_menu = $1", in.Id)
	if err != nil {
		return ct.rollback(err)
	}
	err = models.InsertManyMenuRecettes(tx, recettes)
	if err != nil {
		return ct.rollback(err)
	}
	_, err = tx.Exec("DELETE FROM menu_ingredients WHERE id_menu = $1", in.Id)
	if err != nil {
		return ct.rollback(err)
	}
	err = models.InsertManyMenuIngredients(tx, ings)
	if err != nil {
		return ct.rollback(err)
	}
	return ct.commit()
}

func (s Server) deleteMenu(ct RequeteContext, id int64) error {
	if err := ct.setup(s); err != nil {
		return err
	}
	if err := s.proprioMenu(ct, models.Menu{Id: id}, false); err != nil {
		return err
	}
	rows, err := ct.tx.Query(`SELECT sejours.id FROM sejours 
	JOIN repass ON repass.id_sejour = sejours.id
	WHERE repass.id_menu = $1`, id)
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
		return fmt.Errorf(`Ce menu est présent dans <b>%d menu(s)</b>.
		Si vous souhaitez vraiment le supprimer, il faudra d'abord l'en retirer.`, L)
	}
	_, err = ct.tx.Exec("DELETE FROM menu_recettes WHERE id_menu = $1", id)
	if err != nil {
		return ErrorSQL(err)
	}
	_, err = ct.tx.Exec("DELETE FROM menu_ingredients WHERE id_menu = $1", id)
	if err != nil {
		return ct.rollback(err)
	}
	_, err = models.Menu{Id: id}.Delete(ct.tx)
	if err != nil {
		return ct.rollback(err)
	}
	return ct.commit()
}

// ------------------------------------------------------------------------
// ----------------------- Séjour et repas --------------------------------
// ------------------------------------------------------------------------

func (s Server) createSejour(ct RequeteContext) (out models.Sejour, err error) {
	if err = ct.setup(s); err != nil {
		return
	}
	tx := ct.tx
	out.IdProprietaire = ct.idProprietaire
	out, err = out.Insert(tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = ct.commit()
	return
}

func (s Server) updateSejour(ct RequeteContext, in models.Sejour) error {
	if err := ct.setup(s); err != nil {
		return err
	}
	if err := s.proprioSejour(ct, in, true); err != nil {
		return err
	}
	tx := ct.tx
	in, err := in.Update(tx)
	if err != nil {
		return ErrorSQL(err)
	}
	return ct.commit()
}

func (s Server) deleteSejour(ct RequeteContext, id int64) error {
	if err := ct.setup(s); err != nil {
		return err
	}
	if err := s.proprioSejour(ct, models.Sejour{Id: id}, false); err != nil {
		return err
	}

	_, err := ct.tx.Exec("DELETE FROM repass WHERE id_sejour = $1", id)
	if err != nil {
		return ErrorSQL(err)
	}
	_, err = models.Sejour{Id: id}.Delete(ct.tx)
	if err != nil {
		return ct.rollback(err)
	}
	return ct.commit()
}

func (s Server) createRepas(ct RequeteContext, idSejour, idMenu int64) (out models.Repas, err error) {
	if err = ct.setup(s); err != nil {
		return
	}
	if err = s.proprioSejour(ct, models.Sejour{Id: idSejour}, false); err != nil {
		return
	}
	tx := ct.tx
	out.IdSejour = idSejour
	out.IdMenu = idMenu
	out, err = out.Insert(tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = ct.commit()
	return
}

func (s Server) updateManyRepas(ct RequeteContext, repass []models.Repas) error {
	if err := ct.setup(s); err != nil {
		return err
	}
	for _, repas := range repass {
		if err := s.proprioRepas(ct, repas.Id); err != nil {
			return ct.rollback(err)
		}
		if _, err := repas.Update(ct.tx); err != nil {
			return ct.rollback(err)
		}
	}
	return ct.commit()
}

func (s Server) deleteRepas(ct RequeteContext, id int64) error {
	if err := ct.setup(s); err != nil {
		return err
	}
	if err := s.proprioRepas(ct, id); err != nil {
		return err
	}
	_, err := models.Repas{Id: id}.Delete(ct.tx)
	if err != nil {
		return ct.rollback(err)
	}
	return ct.commit()
}
