package controller

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/benoitkugler/intendance/server/models"
	"github.com/lib/pq"
)

// ---------------------------- Identification ----------------------------

func (s Server) Loggin(mail, password string) (out OutLoggin, err error) {
	r := s.DB.QueryRow("SELECT * FROM utilisateurs WHERE mail = $1", mail)
	u, err := models.ScanUtilisateur(r)
	if err == sql.ErrNoRows {
		out.Erreur = fmt.Sprintf(`L'adresse mail <i>%s</i> ne figure pas dans nos <b>utilisateurs</b>. <br/>
					Peut-être souhaitez-vous créer un compte ?`, mail)
		return out, nil
	}
	if err != nil {
		return out, ErrorSQL(err)
	}
	if u.Password != password {
		out.Erreur = "Le mot de passe est invalide."
		return out, nil
	}
	token, err := s.newToken(u.Id)
	out = OutLoggin{
		Utilisateur: Utilisateur{Id: u.Id, PrenomNom: u.PrenomNom},
		Token:       token,
		Expires:     DeltaTokenJours,
	}
	return out, err
}

func (ct RequeteContext) LoadSejoursUtilisateur() (out Sejours, err error) {
	rows, err := ct.DB.Query(`SELECT groupes.* FROM groupes 
	JOIN sejours ON sejours.id = groupes.id_sejour
	WHERE sejours.id_utilisateur = $1`, ct.IdProprietaire)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	out.Groupes, err = models.ScanGroupes(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	sejours, err := models.SelectSejoursByIdUtilisateurs(ct.DB, ct.IdProprietaire)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	// on charge les fournisseurs des séjours
	sejoursFournisseurs, err := models.SelectSejourFournisseursByIdSejours(ct.DB, sejours.Ids()...)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	// on résoud les repas
	repas, err := models.SelectRepassByIdSejours(ct.DB, sejours.Ids()...)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	// on charge les groupes, recettes et ingrédients pour chaque repas
	idsRepas := repas.Ids()
	repasGroupes, err := models.SelectRepasGroupesByIdRepass(ct.DB, idsRepas...)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	repasRecettes, err := models.SelectRepasRecettesByIdRepass(ct.DB, idsRepas...)
	if err != nil {
		return out, ErrorSQL(err)
	}
	repasIngredients, err := models.SelectRepasIngredientsByIdRepass(ct.DB, idsRepas...)
	if err != nil {
		return out, ErrorSQL(err)
	}

	// on commence par associer les groupes, recettes, ingredients aux repas
	tmpGroupes := map[int64][]models.RepasGroupe{} // idRepas -> groupes
	tmpRecettes := map[int64]models.Ids{}
	tmpIngredients := map[int64][]models.LienIngredient{}
	for _, lien := range repasGroupes {
		tmpGroupes[lien.IdRepas] = append(tmpGroupes[lien.IdRepas], lien)
	}
	for _, lien := range repasRecettes {
		tmpRecettes[lien.IdRepas] = append(tmpRecettes[lien.IdRepas], lien.IdRecette)
	}
	for _, lien := range repasIngredients {
		tmpIngredients[lien.IdRepas] = append(tmpIngredients[lien.IdRepas], lien.LienIngredient)
	}

	// puis on associe à chaque séjour ses repas
	tmpRepas := map[int64][]RepasComplet{} // idSejour -> repas
	for _, rep := range repas {            // on ajoute les repas aux séjours
		repG := RepasComplet{Repas: rep, Groupes: tmpGroupes[rep.Id], Recettes: tmpRecettes[rep.Id], Ingredients: tmpIngredients[rep.Id]}
		tmpRepas[repG.IdSejour] = append(tmpRepas[repG.IdSejour], repG)
	}

	// et à chaque séjour ses fournisseurs
	tmpFournisseurs := map[int64][]models.SejourFournisseur{}
	for _, rep := range sejoursFournisseurs {
		tmpFournisseurs[rep.IdSejour] = append(tmpFournisseurs[rep.IdSejour], rep)
	}

	// finalement on assemble tout
	out.Sejours = make(map[int64]SejourRepas, len(sejours))
	for k, v := range sejours { // informations basiques
		out.Sejours[k] = SejourRepas{Sejour: v, Repass: tmpRepas[k], Fournisseurs: tmpFournisseurs[k]}
	}

	return out, nil
}

// LoadUtilisateurs renvois les données publiques des utilisateurs enregistrés.
func (s Server) LoadUtilisateurs() (map[int64]Utilisateur, error) {
	users, err := models.SelectAllUtilisateurs(s.DB)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	out := make(map[int64]Utilisateur, len(users))
	for k, v := range users {
		out[k] = Utilisateur{Id: v.Id, PrenomNom: v.PrenomNom}
	}
	return out, nil
}

// ------------------------------------------------------------------------
// ---------------------------- Ingrédients -------------------------------
// ------------------------------------------------------------------------

func (s Server) LoadIngredients() (models.Ingredients, error) {
	out, err := models.SelectAllIngredients(s.DB)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	return out, nil
}

func (ct RequeteContext) CreateIngredient() (out models.Ingredient, err error) {
	out.Nom = fmt.Sprintf("I%d", time.Now().UnixNano())
	out, err = out.Insert(ct.DB)
	return out, ErrorSQL(err)
}

func (ct RequeteContext) UpdateIngredient(ig models.Ingredient) (models.Ingredient, error) {
	// contrainte interne à l'ingrédient
	contrainte := ContrainteIngredient{ingredient: ig}
	if err := contrainte.Check(); err != nil {
		return ig, err
	}

	tx, err := ct.beginTx()
	if err != nil {
		return ig, err
	}

	// vérification de la compatilibité des unités et des contionnements
	produits, err := ig.GetProduits(tx.Tx, nil)
	if err != nil {
		return ig, ErrorSQL(err)
	}
	for _, prod := range produits {
		contrainte := ContrainteIngredientProduit{ingredient: ig, produit: prod}
		if err := contrainte.Check(); err != nil {
			return ig, err
		}
	}

	// modification
	ig, err = ig.Update(tx.Tx)
	if err != nil {
		return ig, ErrorSQL(err)
	}
	return ig, tx.Commit()
}

func (ct RequeteContext) DeleteIngredient(id int64, checkProduits bool) error {
	tx, err := ct.beginTx()
	if err != nil {
		return err
	}

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
	check.produits, err = ing.GetProduits(tx, nil)
	if err != nil {
		return ErrorSQL(err)
	}

	if !checkProduits { // on regarde uniquement les recettes et menus
		if len(check.recettes)+len(check.menus) > 0 {
			check.produits = nil
			return check
		}

		err = models.DeleteIngredientProduitsByIdIngredients(tx, id)
		if err != nil {
			return ErrorSQL(err)
		}
	} else { // on regarde aussi les produits
		if len(check.recettes)+len(check.menus)+len(check.produits) > 0 {
			return check
		}
	}
	// tout bon, on peut supprimer
	if _, err = models.DeleteIngredientById(tx, ing.Id); err != nil {
		return tx.rollback(err)
	}
	return tx.Commit()
}

// ------------------------------------------------------------------------
// ------------------------ Recettes --------------------------------------
// ------------------------------------------------------------------------
func (s Server) LoadRecettes() (out map[int64]*RecetteComplet, err error) {
	recettes, err := models.SelectAllRecettes(s.DB)
	if err != nil {
		return out, ErrorSQL(err)
	}
	ris, err := models.SelectAllRecetteIngredients(s.DB)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	resolvedRecettes := make(map[int64]*RecetteComplet, len(recettes))
	for id, r := range recettes {
		resolvedRecettes[id] = &RecetteComplet{Recette: r}
	}
	for _, lien := range ris {
		resolvedRecettes[lien.IdRecette].Ingredients = append(resolvedRecettes[lien.IdRecette].Ingredients, lien.LienIngredient)
	}
	return resolvedRecettes, nil
}

func (ct RequeteContext) CreateRecette() (out models.Recette, err error) {
	out.Nom = fmt.Sprintf("R%d", time.Now().UnixNano())
	out.IdUtilisateur = models.NullableId(ct.IdProprietaire)
	out, err = out.Insert(ct.DB)
	return out, ErrorSQL(err)
}

func (ct RequeteContext) UpdateRecette(in RecetteComplet) (RecetteComplet, error) {
	if err := ct.proprioRecette(in.Recette, true); err != nil {
		return in, err
	}

	tx, err := ct.beginTx()
	if err != nil {
		return in, err
	}
	//TODO: notification aux utilisateurs avec possibilité de copie
	in.Recette, err = in.Recette.Update(tx)
	if err != nil {
		return in, tx.rollback(err)
	}

	err = models.DeleteRecetteIngredientsByIdRecettes(tx, in.Recette.Id)
	if err != nil {
		return in, tx.rollback(err)
	}

	ings := in.Ingredients.AsRecetteIngredients(in.Recette.Id)
	err = models.InsertManyRecetteIngredients(tx.Tx, ings...)
	if err != nil {
		return in, tx.rollback(err)
	}
	return in, tx.Commit()
}

func (ct RequeteContext) DeleteRecette(id int64) error {
	if err := ct.proprioRecette(models.Recette{Id: id}, false); err != nil {
		return err
	}
	tx, err := ct.beginTx()
	if err != nil {
		return err
	}
	rows, err := tx.Tx.Query(`SELECT menus.id FROM menus 
	JOIN menu_recettes ON menu_recettes.id_menu = menus.id
	WHERE menu_recettes.id_recette = $1`, id)
	if err != nil {
		return tx.rollback(err)
	}
	ids, err := models.ScanIds(rows)
	if err != nil {
		return tx.rollback(err)
	}
	//TODO: notification aux utilisateurs avec possibilité de copie
	// nécessite de rassembler les données nécessaires à la re-création
	if L := len(ids); L > 0 {
		_ = tx.rollback(err)
		return fmt.Errorf(`Cette recette est présente dans <b>%d menu(s)</b>.
		Si vous souhaitez vraiment la supprimer, il faudra d'abord l'en retirer.`, L)
	}
	err = models.DeleteRecetteIngredientsByIdRecettes(tx, id)
	if err != nil {
		return tx.rollback(err)
	}
	_, err = models.DeleteRecetteById(tx.Tx, id)
	if err != nil {
		return tx.rollback(err)
	}
	return tx.Commit()
}

// ------------------------------------------------------------------------
// ----------------------------- Menus ------------------------------------
// ------------------------------------------------------------------------

func (s Server) LoadMenus() (out map[int64]*MenuComplet, err error) {
	menus, err := models.SelectAllMenus(s.DB)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	out = make(map[int64]*MenuComplet, len(menus))
	for k, v := range menus { // base
		out[k] = &MenuComplet{Menu: v}
	}
	mrs, err := models.SelectAllMenuRecettes(s.DB)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	for _, l := range mrs {
		out[l.IdMenu].Recettes = append(out[l.IdMenu].Recettes, l.IdRecette)
	}

	mis, err := models.SelectAllMenuIngredients(s.DB)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	for _, l := range mis {
		out[l.IdMenu].Ingredients = append(out[l.IdMenu].Ingredients, l.LienIngredient)
	}
	return out, nil
}

func (ct RequeteContext) CreateMenu() (out models.Menu, err error) {
	out.IdUtilisateur = models.NullableId(ct.IdProprietaire)
	out, err = out.Insert(ct.DB)
	return out, ErrorSQL(err)
}

func (ct RequeteContext) UpdateMenu(in MenuComplet) (MenuComplet, error) {
	if err := ct.proprioMenu(in.Menu, true); err != nil {
		return in, err
	}
	tx, err := ct.beginTx()
	if err != nil {
		return in, err
	}
	//TODO: notification aux utilisateurs avec possibilité de copie
	in.Menu, err = in.Menu.Update(tx)
	if err != nil {
		return in, tx.rollback(err)
	}

	err = models.DeleteMenuRecettesByIdMenus(tx, in.Id)
	if err != nil {
		return in, tx.rollback(err)
	}
	recettes := in.Recettes.AsMenuRecettes(in.Menu.Id)
	err = models.InsertManyMenuRecettes(tx.Tx, recettes...)
	if err != nil {
		return in, tx.rollback(err)
	}

	err = models.DeleteMenuIngredientsByIdMenus(tx, in.Id)
	if err != nil {
		return in, tx.rollback(err)
	}
	ings := in.Ingredients.AsMenuIngredients(in.Menu.Id)
	err = models.InsertManyMenuIngredients(tx.Tx, ings...)
	if err != nil {
		return in, tx.rollback(err)
	}
	return in, tx.Commit()
}

func (ct RequeteContext) DeleteMenu(id int64) error {
	if err := ct.proprioMenu(models.Menu{Id: id}, false); err != nil {
		return err
	}
	tx, err := ct.beginTx()
	if err != nil {
		return err
	}
	// supression des liens
	err = models.DeleteMenuRecettesByIdMenus(tx, id)
	if err != nil {
		return tx.rollback(err)
	}
	err = models.DeleteMenuIngredientsByIdMenus(tx, id)
	if err != nil {
		return tx.rollback(err)
	}
	_, err = models.DeleteMenuById(tx.Tx, id)
	if err != nil {
		return tx.rollback(err)
	}
	return tx.Commit()
}

// ------------------------------------------------------------------------
// ----------------------- Séjour et repas --------------------------------
// ------------------------------------------------------------------------

func (ct RequeteContext) CreateSejour() (out models.Sejour, err error) {
	out.IdUtilisateur = ct.IdProprietaire
	out, err = out.Insert(ct.DB)
	return out, ErrorSQL(err)
}

func (ct RequeteContext) UpdateSejour(in models.Sejour) (models.Sejour, error) {
	if err := ct.proprioSejour(in, true); err != nil {
		return in, err
	}
	in, err := in.Update(ct.DB)
	return in, ErrorSQL(err)
}

func (ct RequeteContext) DeleteSejour(id int64) error {
	if err := ct.proprioSejour(models.Sejour{Id: id}, false); err != nil {
		return err
	}
	tx, err := ct.beginTx()
	if err != nil {
		return err
	}

	// table de lien
	_, err = tx.Exec(`DELETE FROM repas_groupes 
	USING repass WHERE repass.id = repas_groupes.id_repas 
	AND repass.id_sejour = $1`, id)
	if err != nil {
		return tx.rollback(err)
	}
	_, err = models.DeleteRepassByIdSejours(tx, id)
	if err != nil {
		return tx.rollback(err)
	}
	_, err = models.DeleteGroupesByIdSejours(tx, id)
	if err != nil {
		return tx.rollback(err)
	}
	_, err = models.DeleteSejourById(tx.Tx, id)
	if err != nil {
		return tx.rollback(err)
	}
	return tx.Commit()
}

// Groupes

func (ct RequeteContext) CreateGroupe(idSejour int64) (out models.Groupe, err error) {
	if err = ct.proprioSejour(models.Sejour{Id: idSejour}, false); err != nil {
		return
	}

	out.IdSejour = idSejour
	out.Nom = fmt.Sprintf("G%d", time.Now().UnixNano())
	out, err = out.Insert(ct.DB)
	return out, ErrorSQL(err)
}

func (ct RequeteContext) UpdateGroupe(in models.Groupe) (models.Groupe, error) {
	if err := ct.proprioGroupe(in.Id); err != nil {
		return in, err
	}
	in, err := in.Update(ct.DB)
	return in, ErrorSQL(err)
}

// DeleteGroupe supprime le groupe et renvoie le nombre de repas touchés.
func (ct RequeteContext) DeleteGroupe(id int64) (OutDeleteGroupe, error) {
	if err := ct.proprioGroupe(id); err != nil {
		return OutDeleteGroupe{}, err
	}
	tx, err := ct.beginTx()
	if err != nil {
		return OutDeleteGroupe{}, err
	}

	// on enlève le groupe des repas
	res, err := tx.Exec("DELETE FROM repas_groupes WHERE id_groupe = $1", id)
	if err != nil {
		return OutDeleteGroupe{}, tx.rollback(err)
	}
	nbDelete, err := res.RowsAffected()
	if err != nil {
		return OutDeleteGroupe{}, tx.rollback(err)
	}
	_, err = models.DeleteGroupeById(tx.Tx, id)
	if err != nil {
		return OutDeleteGroupe{}, tx.rollback(err)
	}
	return OutDeleteGroupe{NbRepas: nbDelete, Id: id}, tx.Commit()
}

// Repas

func (ct RequeteContext) CreateRepas(idSejour int64) (out models.Repas, err error) {
	if err = ct.proprioSejour(models.Sejour{Id: idSejour}, false); err != nil {
		return
	}

	out.IdSejour = idSejour
	out, err = out.Insert(ct.DB)
	return out, ErrorSQL(err)
}

func (ct RequeteContext) UpdateManyRepas(repass []RepasComplet) error {
	tx, err := ct.beginTx()
	if err != nil {
		return err
	}
	var repasIds pq.Int64Array
	cribleRepasGroupes := map[models.RepasGroupe]bool{}         // pour respecter l'unicité
	cribleRepasRecettes := map[models.RepasRecette]bool{}       // pour respecter l'unicité
	cribleRepasIngredients := map[models.RepasIngredient]bool{} // pour respecter l'unicité
	for _, repas := range repass {
		if err := ct.proprioRepas(repas.Id); err != nil {
			return tx.rollback(err)
		}
		if _, err := repas.Repas.Update(tx.Tx); err != nil {
			return tx.rollback(err)
		}
		repasIds = append(repasIds, repas.Id)
		for _, rg := range repas.Groupes {
			cribleRepasGroupes[rg] = true
		}
		for _, rr := range repas.Recettes {
			cribleRepasRecettes[models.RepasRecette{IdRepas: repas.Id, IdRecette: rr}] = true
		}
		for _, ri := range repas.Ingredients {
			cribleRepasIngredients[models.RepasIngredient{IdRepas: repas.Id, LienIngredient: ri}] = true
		}
	}

	// mise à jour des groupes
	err = models.DeleteRepasGroupesByIdRepass(tx, repasIds...)
	if err != nil {
		return tx.rollback(err)
	}
	batchRepasGroupes := make([]models.RepasGroupe, 0, len(cribleRepasGroupes))
	for rg := range cribleRepasGroupes {
		batchRepasGroupes = append(batchRepasGroupes, rg)
	}
	if err = models.InsertManyRepasGroupes(tx.Tx, batchRepasGroupes...); err != nil {
		return tx.rollback(err)
	}

	// mise à jour des recettes
	err = models.DeleteRepasRecettesByIdRepass(tx, repasIds...)
	if err != nil {
		return tx.rollback(err)
	}
	batchRepasRecettes := make([]models.RepasRecette, 0, len(cribleRepasRecettes))
	for rg := range cribleRepasRecettes {
		batchRepasRecettes = append(batchRepasRecettes, rg)
	}
	if err = models.InsertManyRepasRecettes(tx.Tx, batchRepasRecettes...); err != nil {
		return tx.rollback(err)
	}

	// mise à jour des ingredients
	err = models.DeleteRepasIngredientsByIdRepass(tx, repasIds...)
	if err != nil {
		return tx.rollback(err)
	}
	batchRepasIngredients := make([]models.RepasIngredient, 0, len(cribleRepasIngredients))
	for rg := range cribleRepasIngredients {
		batchRepasIngredients = append(batchRepasIngredients, rg)
	}
	if err = models.InsertManyRepasIngredients(tx.Tx, batchRepasIngredients...); err != nil {
		return tx.rollback(err)
	}

	return tx.Commit()
}

func (ct RequeteContext) DeleteRepas(id int64) error {
	if err := ct.proprioRepas(id); err != nil {
		return err
	}
	tx, err := ct.beginTx()
	if err != nil {
		return err
	}

	// suppression des liens par contrainte

	_, err = models.DeleteRepasById(tx.Tx, id)
	if err != nil {
		return tx.rollback(err)
	}
	return tx.Commit()
}
