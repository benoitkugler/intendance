package controller

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

// Server est le controller principal, partagé par toutes les requêtes.
type Server struct {
	db      *sql.DB
	devMode bool // contourne l'authentification
}

func NewServer(db *sql.DB, devMode bool) Server {
	return Server{db: db, devMode: devMode}
}

func (s Server) PingDB() error {
	return s.db.Ping()
}

// RequeteContext est créé pour chaque requête.
type RequeteContext struct {
	idProprietaire int64
	tx             *sql.Tx // a créer
	Token          string  // à remplir pendant la phase d'authentification
}

// rollbackTx the current transaction, caused by `err`, and
// handles the possible error from tx.Rollback()
func (ct RequeteContext) rollbackTx(origin error) error {
	if err := ct.tx.Rollback(); err != nil {
		origin = fmt.Errorf("Rollback impossible. Erreur originale : %s", origin)
	}
	if _, ok := origin.(errorSQL); ok { // pas besoin de wrapper
		return origin
	}
	return ErrorSQL(origin)
}

func (ct *RequeteContext) beginTx(s Server) (err error) {
	ct.tx, err = s.db.Begin()
	if err != nil {
		return ErrorSQL(err)
	}
	return nil
}

// commitTx the transaction and try to rollback on error
func (r RequeteContext) commitTx() error {
	if err := r.tx.Commit(); err != nil {
		return r.rollbackTx(err)
	}
	return nil
}

// ---------------------------- Identification ----------------------------

func (s Server) Loggin(mail, password string) (out OutLoggin, err error) {
	r := s.db.QueryRow("SELECT * FROM utilisateurs WHERE mail = $1", mail)
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
	token, err := creeToken(u.Id)
	out = OutLoggin{
		Utilisateur: Utilisateur{Id: u.Id, PrenomNom: u.PrenomNom},
		Token:       token,
	}
	return out, err
}

func (s Server) LoadSejoursUtilisateur(ct RequeteContext) (out Sejours, err error) {
	rows, err := s.db.Query(`SELECT groupes.* FROM groupes 
	JOIN sejours ON sejours.id = groupes.id_sejour
	WHERE sejours.id_proprietaire = $1`, ct.idProprietaire)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	out.Groupes, err = models.ScanGroupes(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}

	rows, err = s.db.Query("SELECT * FROM sejours WHERE id_proprietaire = $1", ct.idProprietaire)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	sejours, err := models.ScanSejours(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	out.Sejours = make(map[int64]*SejourJournees, len(sejours))
	for k, v := range sejours {
		out.Sejours[k] = &SejourJournees{Sejour: v, Journees: map[int64]Journee{}}
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
		journee := out.Sejours[l.IdSejour].Journees[l.JourOffset]
		journee.JourOffset = l.JourOffset
		journee.Repas = append(journee.Repas, l)
		out.Sejours[l.IdSejour].Journees[l.JourOffset] = journee
	}
	return out, nil
}

// LoadUtilisateurs renvois les données publiques des utilisateurs enregistrés.
func (s Server) LoadUtilisateurs() (map[int64]Utilisateur, error) {
	rows, err := s.db.Query("SELECT * FROM utilisateurs")
	if err != nil {
		return nil, ErrorSQL(err)
	}
	users, err := models.ScanUtilisateurs(rows)
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
	rows, err := s.db.Query("SELECT * FROM ingredients")
	if err != nil {
		return nil, ErrorSQL(err)
	}
	out, err := models.ScanIngredients(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	return out, nil
}

func (s Server) CreateIngredient(ct RequeteContext) (out models.Ingredient, err error) {
	if err = ct.beginTx(s); err != nil {
		return
	}
	out.Nom = fmt.Sprintf("I%d", time.Now().UnixNano())
	out, err = out.Insert(ct.tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = ct.commitTx()
	return
}

func (s Server) UpdateIngredient(ct RequeteContext, ig models.Ingredient) (models.Ingredient, error) {
	if err := ct.beginTx(s); err != nil {
		return ig, err
	}
	tx := ct.tx
	// vérification de la compatilibité des unités et des contionnements
	produits, err := ig.GetProduits(tx)
	if err != nil {
		return ig, ErrorSQL(err)
	}
	for _, prod := range produits {
		if ig.Unite != models.Piece && ig.Unite != prod.Conditionnement.Unite {
			return ig, ErrorIngredientProduitUnite{ingredient: ig, produit: prod}
		}
		if !ig.Conditionnement.IsNull() && ig.Conditionnement != prod.Conditionnement {
			return ig, ErrorIngredientProduitConditionnement{ingredient: ig, produit: prod}
		}
	}

	// modification
	ig, err = ig.Update(tx)
	if err != nil {
		return ig, ErrorSQL(err)
	}
	return ig, ct.commitTx()
}

func (s Server) DeleteIngredient(ct RequeteContext, id int64, checkProduits bool) error {
	if err := ct.beginTx(s); err != nil {
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

	if !checkProduits { // on regarde uniquement les recettes et menus
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
		return ct.rollbackTx(err)
	}
	return ct.commitTx()
}

// ------------------------------------------------------------------------
// ------------------------ Recettes --------------------------------------
// ------------------------------------------------------------------------
func (s Server) LoadRecettes() (out map[int64]*Recette, err error) {
	rows, err := s.db.Query("SELECT * FROM recettes")
	if err != nil {
		return out, ErrorSQL(err)
	}
	recettes, err := models.ScanRecettes(rows)
	if err != nil {
		return out, ErrorSQL(err)
	}
	rows, err = s.db.Query(`SELECT * FROM recette_ingredients`)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	ris, err := models.ScanRecetteIngredients(rows)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	resolvedRecettes := make(map[int64]*Recette, len(recettes))
	for id, r := range recettes {
		resolvedRecettes[id] = &Recette{Recette: r}
	}
	for _, lien := range ris {
		resolvedRecettes[lien.IdRecette].Ingredients = append(resolvedRecettes[lien.IdRecette].Ingredients, lien)
	}
	return resolvedRecettes, nil
}

func (s Server) CreateRecette(ct RequeteContext) (out models.Recette, err error) {
	if err = ct.beginTx(s); err != nil {
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
	err = ct.commitTx()
	return
}

func (s Server) UpdateRecette(ct RequeteContext, in Recette) (Recette, error) {
	for _, r := range in.Ingredients {
		if r.IdRecette != in.Id {
			return in, fmt.Errorf("L'ingrédient %d n'est pas associé à la recette fournie !", r.IdIngredient)
		}
	}
	if err := ct.beginTx(s); err != nil {
		return in, err
	}
	if err := s.proprioRecette(ct, in.Recette, true); err != nil {
		return in, err
	}
	tx := ct.tx
	//TODO: notification aux utilisateurs avec possibilité de copie
	var err error
	in.Recette, err = in.Recette.Update(tx)
	if err != nil {
		return in, ErrorSQL(err)
	}
	_, err = tx.Exec("DELETE FROM recette_ingredients WHERE id_recette = $1", in.Id)
	if err != nil {
		return in, ct.rollbackTx(err)
	}
	err = models.InsertManyRecetteIngredients(tx, in.Ingredients)
	if err != nil {
		return in, ct.rollbackTx(err)
	}
	return in, ct.commitTx()
}

func (s Server) DeleteRecette(ct RequeteContext, id int64) error {
	if err := ct.beginTx(s); err != nil {
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
		return ct.rollbackTx(err)
	}
	return ct.commitTx()
}

// ------------------------------------------------------------------------
// ----------------------------- Menus ------------------------------------
// ------------------------------------------------------------------------

func (s Server) LoadMenus() (out map[int64]*Menu, err error) {
	rows, err := s.db.Query("SELECT * FROM menus")
	if err != nil {
		return nil, ErrorSQL(err)
	}
	menus, err := models.ScanMenus(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	out = make(map[int64]*Menu, len(menus))
	for k, v := range menus { // base
		out[k] = &Menu{Menu: v}
	}
	rows, err = s.db.Query(`SELECT * FROM menu_recettes`)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	mrs, err := models.ScanMenuRecettes(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	for _, l := range mrs {
		out[l.IdMenu].Recettes = append(out[l.IdMenu].Recettes, l)
	}

	rows, err = s.db.Query(`SELECT * FROM menu_ingredients`)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	mis, err := models.ScanMenuIngredients(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	for _, l := range mis {
		out[l.IdMenu].Ingredients = append(out[l.IdMenu].Ingredients, l)
	}
	return out, nil
}

func (s Server) CreateMenu(ct RequeteContext) (out models.Menu, err error) {
	if err = ct.beginTx(s); err != nil {
		return
	}
	tx := ct.tx
	out.IdProprietaire = models.NullId(ct.idProprietaire)
	out, err = out.Insert(tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = ct.commitTx()
	return
}

func (s Server) UpdateMenu(ct RequeteContext, in Menu) (Menu, error) {
	for _, r := range in.Recettes {
		if r.IdMenu != in.Id {
			return in, fmt.Errorf("La recette %d n'est pas associée au menu fourni !", r.IdRecette)
		}
	}
	for _, r := range in.Ingredients {
		if r.IdMenu != in.Id {
			return in, fmt.Errorf("L'ingrédient %d n'est pas associé au menu fourni !", r.IdIngredient)
		}
	}
	if err := ct.beginTx(s); err != nil {
		return in, err
	}
	if err := s.proprioMenu(ct, in.Menu, true); err != nil {
		return in, err
	}
	tx := ct.tx
	//TODO: notification aux utilisateurs avec possibilité de copie
	var err error
	in.Menu, err = in.Menu.Update(tx)
	if err != nil {
		return in, ErrorSQL(err)
	}
	_, err = tx.Exec("DELETE FROM menu_recettes WHERE id_menu = $1", in.Id)
	if err != nil {
		return in, ct.rollbackTx(err)
	}
	err = models.InsertManyMenuRecettes(tx, in.Recettes)
	if err != nil {
		return in, ct.rollbackTx(err)
	}
	_, err = tx.Exec("DELETE FROM menu_ingredients WHERE id_menu = $1", in.Id)
	if err != nil {
		return in, ct.rollbackTx(err)
	}
	err = models.InsertManyMenuIngredients(tx, in.Ingredients)
	if err != nil {
		return in, ct.rollbackTx(err)
	}
	return in, ct.commitTx()
}

func (s Server) DeleteMenu(ct RequeteContext, id int64) error {
	if err := ct.beginTx(s); err != nil {
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
		return fmt.Errorf(`Ce menu est présent dans <b>%d séjours(s)</b>.
		Si vous souhaitez vraiment le supprimer, il faudra d'abord l'en retirer.`, L)
	}
	_, err = ct.tx.Exec("DELETE FROM menu_recettes WHERE id_menu = $1", id)
	if err != nil {
		return ErrorSQL(err)
	}
	_, err = ct.tx.Exec("DELETE FROM menu_ingredients WHERE id_menu = $1", id)
	if err != nil {
		return ct.rollbackTx(err)
	}
	_, err = models.Menu{Id: id}.Delete(ct.tx)
	if err != nil {
		return ct.rollbackTx(err)
	}
	return ct.commitTx()
}

// ------------------------------------------------------------------------
// ----------------------- Séjour et repas --------------------------------
// ------------------------------------------------------------------------

func (s Server) CreateSejour(ct RequeteContext) (out models.Sejour, err error) {
	if err = ct.beginTx(s); err != nil {
		return
	}
	tx := ct.tx
	out.IdProprietaire = ct.idProprietaire
	out, err = out.Insert(tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = ct.commitTx()
	return
}

func (s Server) UpdateSejour(ct RequeteContext, in models.Sejour) (models.Sejour, error) {
	if err := ct.beginTx(s); err != nil {
		return in, err
	}
	if err := s.proprioSejour(ct, in, true); err != nil {
		return in, err
	}
	tx := ct.tx
	in, err := in.Update(tx)
	if err != nil {
		return in, ErrorSQL(err)
	}
	return in, ct.commitTx()
}

func (s Server) DeleteSejour(ct RequeteContext, id int64) error {
	if err := ct.beginTx(s); err != nil {
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
		return ct.rollbackTx(err)
	}
	return ct.commitTx()
}

// Groupes

func (s Server) CreateGroupe(ct RequeteContext, idSejour int64) (out models.Groupe, err error) {
	if err = ct.beginTx(s); err != nil {
		return
	}
	tx := ct.tx
	if err = s.proprioSejour(ct, models.Sejour{Id: idSejour}, false); err != nil {
		return
	}
	out.IdSejour = idSejour
	out.Nom = fmt.Sprintf("G%d", time.Now().UnixNano())
	out, err = out.Insert(tx)
	if err != nil {
		err = ErrorSQL(err)
		return
	}
	err = ct.commitTx()
	return
}

func (s Server) UpdateGroupe(ct RequeteContext, in models.Groupe) (models.Groupe, error) {
	if err := ct.beginTx(s); err != nil {
		return in, err
	}
	if err := s.proprioGroupe(ct, in.Id); err != nil {
		return in, err
	}
	tx := ct.tx
	in, err := in.Update(tx)
	if err != nil {
		return in, ErrorSQL(err)
	}
	return in, ct.commitTx()
}

// DeleteGroupe supprime le groupe et renvoie le nombre de repas touchés.
func (s Server) DeleteGroupe(ct RequeteContext, id int64) (int, error) {
	if err := ct.beginTx(s); err != nil {
		return 0, err
	}
	if err := s.proprioGroupe(ct, id); err != nil {
		return 0, err
	}

	// on enlève le groupe des repas
	res, err := ct.tx.Exec("DELETE FROM repas_groupes WHERE id_groupe = $1", id)
	if err != nil {
		return 0, ErrorSQL(err)
	}
	nbDelete, err := res.RowsAffected()
	if err != nil {
		return 0, ct.rollbackTx(err)
	}
	_, err = models.Groupe{Id: id}.Delete(ct.tx)
	if err != nil {
		return 0, ct.rollbackTx(err)
	}
	return int(nbDelete), ct.commitTx()
}

// Repas

func (s Server) CreateRepas(ct RequeteContext, idSejour, idMenu int64) (out models.Repas, err error) {
	if err = ct.beginTx(s); err != nil {
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
	err = ct.commitTx()
	return
}

func (s Server) UpdateManyRepas(ct RequeteContext, repass []models.Repas) error {
	if err := ct.beginTx(s); err != nil {
		return err
	}
	for _, repas := range repass {
		if err := s.proprioRepas(ct, repas.Id); err != nil {
			return ct.rollbackTx(err)
		}
		if _, err := repas.Update(ct.tx); err != nil {
			return ct.rollbackTx(err)
		}
	}
	return ct.commitTx()
}

func (s Server) DeleteRepas(ct RequeteContext, id int64) error {
	if err := ct.beginTx(s); err != nil {
		return err
	}
	if err := s.proprioRepas(ct, id); err != nil {
		return err
	}
	_, err := models.Repas{Id: id}.Delete(ct.tx)
	if err != nil {
		return ct.rollbackTx(err)
	}
	return ct.commitTx()
}
