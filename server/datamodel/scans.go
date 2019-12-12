// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package datamodel

import (
	"database/sql"

	"github.com/lib/pq"
)

func ScanUtilisateur(r *sql.Row) (Utilisateur, error) {
	var s Utilisateur
	if err := r.Scan(
		&s.Id,
		&s.Password,
		&s.Mail,
		&s.PrenomNom,
	); err != nil {
		return Utilisateur{}, err
	}
	return s, nil
}

type Utilisateurs map[int64]Utilisateur

func (m Utilisateurs) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanUtilisateurs(rs *sql.Rows) (Utilisateurs, error) {
	structs := make(Utilisateurs, 16)
	var err error
	for rs.Next() {
		var s Utilisateur
		if err = rs.Scan(
			&s.Id,
			&s.Password,
			&s.Mail,
			&s.PrenomNom,
		); err != nil {
			return nil, err
		}
		structs[s.Id] = s
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert Utilisateur in the database and returns the item with id filled.
func (item Utilisateur) Insert(tx *sql.Tx) (out Utilisateur, err error) {
	r := tx.QueryRow(`INSERT INTO utilisateurs (
		password,mail,prenom_nom
		) VALUES (
		$1,$2,$3
		) RETURNING 
		id,password,mail,prenom_nom;
		`, item.Password, item.Mail, item.PrenomNom)
	return ScanUtilisateur(r)
}

// Update Utilisateur in the database and returns the new version.
func (item Utilisateur) Update(tx *sql.Tx) (out Utilisateur, err error) {
	r := tx.QueryRow(`UPDATE utilisateurs SET (
		password,mail,prenom_nom
		) = (
		$2,$3,$4
		) WHERE id = $1 RETURNING 
		id,password,mail,prenom_nom;
		`, item.Id, item.Password, item.Mail, item.PrenomNom)
	return ScanUtilisateur(r)
}

// Delete Utilisateur in the database and the return the id.
// Only the field 'Id' is used.
func (item Utilisateur) Delete(tx *sql.Tx) (int64, error) {
	var deleted_id int64
	r := tx.QueryRow("DELETE FROM utilisateurs WHERE id = $1 RETURNING id;", item.Id)
	err := r.Scan(&deleted_id)
	return deleted_id, err
}

func ScanIngredient(r *sql.Row) (Ingredient, error) {
	var s Ingredient
	if err := r.Scan(
		&s.Id,
		&s.Nom,
		&s.Unite,
		&s.Categorie,
		&s.Callories,
		&s.Conditionnement,
	); err != nil {
		return Ingredient{}, err
	}
	return s, nil
}

type Ingredients map[int64]Ingredient

func (m Ingredients) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanIngredients(rs *sql.Rows) (Ingredients, error) {
	structs := make(Ingredients, 16)
	var err error
	for rs.Next() {
		var s Ingredient
		if err = rs.Scan(
			&s.Id,
			&s.Nom,
			&s.Unite,
			&s.Categorie,
			&s.Callories,
			&s.Conditionnement,
		); err != nil {
			return nil, err
		}
		structs[s.Id] = s
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert Ingredient in the database and returns the item with id filled.
func (item Ingredient) Insert(tx *sql.Tx) (out Ingredient, err error) {
	r := tx.QueryRow(`INSERT INTO ingredients (
		nom,unite,categorie,callories,conditionnement
		) VALUES (
		$1,$2,$3,$4,$5
		) RETURNING 
		id,nom,unite,categorie,callories,conditionnement;
		`, item.Nom, item.Unite, item.Categorie, item.Callories, item.Conditionnement)
	return ScanIngredient(r)
}

// Update Ingredient in the database and returns the new version.
func (item Ingredient) Update(tx *sql.Tx) (out Ingredient, err error) {
	r := tx.QueryRow(`UPDATE ingredients SET (
		nom,unite,categorie,callories,conditionnement
		) = (
		$2,$3,$4,$5,$6
		) WHERE id = $1 RETURNING 
		id,nom,unite,categorie,callories,conditionnement;
		`, item.Id, item.Nom, item.Unite, item.Categorie, item.Callories, item.Conditionnement)
	return ScanIngredient(r)
}

// Delete Ingredient in the database and the return the id.
// Only the field 'Id' is used.
func (item Ingredient) Delete(tx *sql.Tx) (int64, error) {
	var deleted_id int64
	r := tx.QueryRow("DELETE FROM ingredients WHERE id = $1 RETURNING id;", item.Id)
	err := r.Scan(&deleted_id)
	return deleted_id, err
}

func ScanRecette(r *sql.Row) (Recette, error) {
	var s Recette
	if err := r.Scan(
		&s.Id,
		&s.IdProprietaire,
		&s.Nom,
		&s.ModeEmploi,
	); err != nil {
		return Recette{}, err
	}
	return s, nil
}

type Recettes map[int64]Recette

func (m Recettes) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanRecettes(rs *sql.Rows) (Recettes, error) {
	structs := make(Recettes, 16)
	var err error
	for rs.Next() {
		var s Recette
		if err = rs.Scan(
			&s.Id,
			&s.IdProprietaire,
			&s.Nom,
			&s.ModeEmploi,
		); err != nil {
			return nil, err
		}
		structs[s.Id] = s
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert Recette in the database and returns the item with id filled.
func (item Recette) Insert(tx *sql.Tx) (out Recette, err error) {
	r := tx.QueryRow(`INSERT INTO recettes (
		id_proprietaire,nom,mode_emploi
		) VALUES (
		$1,$2,$3
		) RETURNING 
		id,id_proprietaire,nom,mode_emploi;
		`, item.IdProprietaire, item.Nom, item.ModeEmploi)
	return ScanRecette(r)
}

// Update Recette in the database and returns the new version.
func (item Recette) Update(tx *sql.Tx) (out Recette, err error) {
	r := tx.QueryRow(`UPDATE recettes SET (
		id_proprietaire,nom,mode_emploi
		) = (
		$2,$3,$4
		) WHERE id = $1 RETURNING 
		id,id_proprietaire,nom,mode_emploi;
		`, item.Id, item.IdProprietaire, item.Nom, item.ModeEmploi)
	return ScanRecette(r)
}

// Delete Recette in the database and the return the id.
// Only the field 'Id' is used.
func (item Recette) Delete(tx *sql.Tx) (int64, error) {
	var deleted_id int64
	r := tx.QueryRow("DELETE FROM recettes WHERE id = $1 RETURNING id;", item.Id)
	err := r.Scan(&deleted_id)
	return deleted_id, err
}

func ScanRecetteIngredient(r *sql.Row) (RecetteIngredient, error) {
	var s RecetteIngredient
	if err := r.Scan(
		&s.IdRecette,
		&s.IdIngredient,
		&s.Quantite,
		&s.Cuisson,
	); err != nil {
		return RecetteIngredient{}, err
	}
	return s, nil
}

type RecetteIngredients map[int64]RecetteIngredient

func (m RecetteIngredients) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanRecetteIngredients(rs *sql.Rows) ([]RecetteIngredient, error) {
	structs := make([]RecetteIngredient, 0, 16)
	var err error
	for rs.Next() {
		var s RecetteIngredient
		if err = rs.Scan(
			&s.IdRecette,
			&s.IdIngredient,
			&s.Quantite,
			&s.Cuisson,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert the link RecetteIngredient in the database.
func (item RecetteIngredient) Insert(tx *sql.Tx) error {
	_, err := tx.Exec(`INSERT INTO recette_ingredients (
		id_recette,id_ingredient,quantite,cuisson
		) VALUES (
		$1,$2,$3,$4
		);
		`, item.IdRecette, item.IdIngredient, item.Quantite, item.Cuisson)
	return err
}

// Delete the link RecetteIngredient in the database.
// Only the 'IdRecette' 'IdIngredient' fields are used.
func (item RecetteIngredient) Delete(tx *sql.Tx) error {
	_, err := tx.Exec(`DELETE FROM recette_ingredients WHERE 
	id_recette = $1 AND id_ingredient = $2;`, item.IdRecette, item.IdIngredient)
	return err
}

func ScanMenu(r *sql.Row) (Menu, error) {
	var s Menu
	if err := r.Scan(
		&s.Id,
		&s.IdProprietaire,
		&s.Commentaire,
	); err != nil {
		return Menu{}, err
	}
	return s, nil
}

type Menus map[int64]Menu

func (m Menus) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanMenus(rs *sql.Rows) (Menus, error) {
	structs := make(Menus, 16)
	var err error
	for rs.Next() {
		var s Menu
		if err = rs.Scan(
			&s.Id,
			&s.IdProprietaire,
			&s.Commentaire,
		); err != nil {
			return nil, err
		}
		structs[s.Id] = s
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert Menu in the database and returns the item with id filled.
func (item Menu) Insert(tx *sql.Tx) (out Menu, err error) {
	r := tx.QueryRow(`INSERT INTO menus (
		id_proprietaire,commentaire
		) VALUES (
		$1,$2
		) RETURNING 
		id,id_proprietaire,commentaire;
		`, item.IdProprietaire, item.Commentaire)
	return ScanMenu(r)
}

// Update Menu in the database and returns the new version.
func (item Menu) Update(tx *sql.Tx) (out Menu, err error) {
	r := tx.QueryRow(`UPDATE menus SET (
		id_proprietaire,commentaire
		) = (
		$2,$3
		) WHERE id = $1 RETURNING 
		id,id_proprietaire,commentaire;
		`, item.Id, item.IdProprietaire, item.Commentaire)
	return ScanMenu(r)
}

// Delete Menu in the database and the return the id.
// Only the field 'Id' is used.
func (item Menu) Delete(tx *sql.Tx) (int64, error) {
	var deleted_id int64
	r := tx.QueryRow("DELETE FROM menus WHERE id = $1 RETURNING id;", item.Id)
	err := r.Scan(&deleted_id)
	return deleted_id, err
}

func ScanMenuIngredient(r *sql.Row) (MenuIngredient, error) {
	var s MenuIngredient
	if err := r.Scan(
		&s.IdMenu,
		&s.IdIngredient,
		&s.Quantite,
		&s.Cuisson,
	); err != nil {
		return MenuIngredient{}, err
	}
	return s, nil
}

type MenuIngredients map[int64]MenuIngredient

func (m MenuIngredients) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanMenuIngredients(rs *sql.Rows) ([]MenuIngredient, error) {
	structs := make([]MenuIngredient, 0, 16)
	var err error
	for rs.Next() {
		var s MenuIngredient
		if err = rs.Scan(
			&s.IdMenu,
			&s.IdIngredient,
			&s.Quantite,
			&s.Cuisson,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert the link MenuIngredient in the database.
func (item MenuIngredient) Insert(tx *sql.Tx) error {
	_, err := tx.Exec(`INSERT INTO menu_ingredients (
		id_menu,id_ingredient,quantite,cuisson
		) VALUES (
		$1,$2,$3,$4
		);
		`, item.IdMenu, item.IdIngredient, item.Quantite, item.Cuisson)
	return err
}

// Delete the link MenuIngredient in the database.
// Only the 'IdMenu' 'IdIngredient' fields are used.
func (item MenuIngredient) Delete(tx *sql.Tx) error {
	_, err := tx.Exec(`DELETE FROM menu_ingredients WHERE 
	id_menu = $1 AND id_ingredient = $2;`, item.IdMenu, item.IdIngredient)
	return err
}

func ScanMenuRecette(r *sql.Row) (MenuRecette, error) {
	var s MenuRecette
	if err := r.Scan(
		&s.IdMenu,
		&s.IdRecette,
	); err != nil {
		return MenuRecette{}, err
	}
	return s, nil
}

type MenuRecettes map[int64]MenuRecette

func (m MenuRecettes) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanMenuRecettes(rs *sql.Rows) ([]MenuRecette, error) {
	structs := make([]MenuRecette, 0, 16)
	var err error
	for rs.Next() {
		var s MenuRecette
		if err = rs.Scan(
			&s.IdMenu,
			&s.IdRecette,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert the link MenuRecette in the database.
func (item MenuRecette) Insert(tx *sql.Tx) error {
	_, err := tx.Exec(`INSERT INTO menu_recettes (
		id_menu,id_recette
		) VALUES (
		$1,$2
		);
		`, item.IdMenu, item.IdRecette)
	return err
}

// Delete the link MenuRecette in the database.
// Only the 'IdMenu' 'IdRecette' fields are used.
func (item MenuRecette) Delete(tx *sql.Tx) error {
	_, err := tx.Exec(`DELETE FROM menu_recettes WHERE 
	id_menu = $1 AND id_recette = $2;`, item.IdMenu, item.IdRecette)
	return err
}

func ScanSejour(r *sql.Row) (Sejour, error) {
	var s Sejour
	if err := r.Scan(
		&s.Id,
		&s.IdProprietaire,
		&s.DateDebut,
		&s.Nom,
	); err != nil {
		return Sejour{}, err
	}
	return s, nil
}

type Sejours map[int64]Sejour

func (m Sejours) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanSejours(rs *sql.Rows) (Sejours, error) {
	structs := make(Sejours, 16)
	var err error
	for rs.Next() {
		var s Sejour
		if err = rs.Scan(
			&s.Id,
			&s.IdProprietaire,
			&s.DateDebut,
			&s.Nom,
		); err != nil {
			return nil, err
		}
		structs[s.Id] = s
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert Sejour in the database and returns the item with id filled.
func (item Sejour) Insert(tx *sql.Tx) (out Sejour, err error) {
	r := tx.QueryRow(`INSERT INTO sejours (
		id_proprietaire,date_debut,nom
		) VALUES (
		$1,$2,$3
		) RETURNING 
		id,id_proprietaire,date_debut,nom;
		`, item.IdProprietaire, item.DateDebut, item.Nom)
	return ScanSejour(r)
}

// Update Sejour in the database and returns the new version.
func (item Sejour) Update(tx *sql.Tx) (out Sejour, err error) {
	r := tx.QueryRow(`UPDATE sejours SET (
		id_proprietaire,date_debut,nom
		) = (
		$2,$3,$4
		) WHERE id = $1 RETURNING 
		id,id_proprietaire,date_debut,nom;
		`, item.Id, item.IdProprietaire, item.DateDebut, item.Nom)
	return ScanSejour(r)
}

// Delete Sejour in the database and the return the id.
// Only the field 'Id' is used.
func (item Sejour) Delete(tx *sql.Tx) (int64, error) {
	var deleted_id int64
	r := tx.QueryRow("DELETE FROM sejours WHERE id = $1 RETURNING id;", item.Id)
	err := r.Scan(&deleted_id)
	return deleted_id, err
}

func ScanSejourMenu(r *sql.Row) (SejourMenu, error) {
	var s SejourMenu
	if err := r.Scan(
		&s.IdSejour,
		&s.IdMenu,
		&s.NbPersonnes,
		&s.JourOffset,
		&s.Horaire,
	); err != nil {
		return SejourMenu{}, err
	}
	return s, nil
}

type SejourMenus map[int64]SejourMenu

func (m SejourMenus) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanSejourMenus(rs *sql.Rows) ([]SejourMenu, error) {
	structs := make([]SejourMenu, 0, 16)
	var err error
	for rs.Next() {
		var s SejourMenu
		if err = rs.Scan(
			&s.IdSejour,
			&s.IdMenu,
			&s.NbPersonnes,
			&s.JourOffset,
			&s.Horaire,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert the link SejourMenu in the database.
func (item SejourMenu) Insert(tx *sql.Tx) error {
	_, err := tx.Exec(`INSERT INTO sejour_menus (
		id_sejour,id_menu,nb_personnes,jour_offset,horaire
		) VALUES (
		$1,$2,$3,$4,$5
		);
		`, item.IdSejour, item.IdMenu, item.NbPersonnes, item.JourOffset, item.Horaire)
	return err
}

// Delete the link SejourMenu in the database.
// Only the 'IdSejour' 'IdMenu' fields are used.
func (item SejourMenu) Delete(tx *sql.Tx) error {
	_, err := tx.Exec(`DELETE FROM sejour_menus WHERE 
	id_sejour = $1 AND id_menu = $2;`, item.IdSejour, item.IdMenu)
	return err
}

func ScanFournisseur(r *sql.Row) (Fournisseur, error) {
	var s Fournisseur
	if err := r.Scan(
		&s.Id,
		&s.Nom,
		&s.DelaiCommande,
		&s.JoursLivraison,
	); err != nil {
		return Fournisseur{}, err
	}
	return s, nil
}

type Fournisseurs map[int64]Fournisseur

func (m Fournisseurs) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanFournisseurs(rs *sql.Rows) (Fournisseurs, error) {
	structs := make(Fournisseurs, 16)
	var err error
	for rs.Next() {
		var s Fournisseur
		if err = rs.Scan(
			&s.Id,
			&s.Nom,
			&s.DelaiCommande,
			&s.JoursLivraison,
		); err != nil {
			return nil, err
		}
		structs[s.Id] = s
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert Fournisseur in the database and returns the item with id filled.
func (item Fournisseur) Insert(tx *sql.Tx) (out Fournisseur, err error) {
	r := tx.QueryRow(`INSERT INTO fournisseurs (
		nom,delai_commande,jours_livraison
		) VALUES (
		$1,$2,$3
		) RETURNING 
		id,nom,delai_commande,jours_livraison;
		`, item.Nom, item.DelaiCommande, item.JoursLivraison)
	return ScanFournisseur(r)
}

// Update Fournisseur in the database and returns the new version.
func (item Fournisseur) Update(tx *sql.Tx) (out Fournisseur, err error) {
	r := tx.QueryRow(`UPDATE fournisseurs SET (
		nom,delai_commande,jours_livraison
		) = (
		$2,$3,$4
		) WHERE id = $1 RETURNING 
		id,nom,delai_commande,jours_livraison;
		`, item.Id, item.Nom, item.DelaiCommande, item.JoursLivraison)
	return ScanFournisseur(r)
}

// Delete Fournisseur in the database and the return the id.
// Only the field 'Id' is used.
func (item Fournisseur) Delete(tx *sql.Tx) (int64, error) {
	var deleted_id int64
	r := tx.QueryRow("DELETE FROM fournisseurs WHERE id = $1 RETURNING id;", item.Id)
	err := r.Scan(&deleted_id)
	return deleted_id, err
}

func ScanProduit(r *sql.Row) (Produit, error) {
	var s Produit
	if err := r.Scan(
		&s.Id,
		&s.IdFournisseur,
		&s.Nom,
		&s.Conditionnement,
		&s.Prix,
		&s.ReferenceFournisseur,
		&s.Colisage,
	); err != nil {
		return Produit{}, err
	}
	return s, nil
}

type Produits map[int64]Produit

func (m Produits) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanProduits(rs *sql.Rows) (Produits, error) {
	structs := make(Produits, 16)
	var err error
	for rs.Next() {
		var s Produit
		if err = rs.Scan(
			&s.Id,
			&s.IdFournisseur,
			&s.Nom,
			&s.Conditionnement,
			&s.Prix,
			&s.ReferenceFournisseur,
			&s.Colisage,
		); err != nil {
			return nil, err
		}
		structs[s.Id] = s
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert Produit in the database and returns the item with id filled.
func (item Produit) Insert(tx *sql.Tx) (out Produit, err error) {
	r := tx.QueryRow(`INSERT INTO produits (
		id_fournisseur,nom,conditionnement,prix,reference_fournisseur,colisage
		) VALUES (
		$1,$2,$3,$4,$5,$6
		) RETURNING 
		id,id_fournisseur,nom,conditionnement,prix,reference_fournisseur,colisage;
		`, item.IdFournisseur, item.Nom, item.Conditionnement, item.Prix, item.ReferenceFournisseur, item.Colisage)
	return ScanProduit(r)
}

// Update Produit in the database and returns the new version.
func (item Produit) Update(tx *sql.Tx) (out Produit, err error) {
	r := tx.QueryRow(`UPDATE produits SET (
		id_fournisseur,nom,conditionnement,prix,reference_fournisseur,colisage
		) = (
		$2,$3,$4,$5,$6,$7
		) WHERE id = $1 RETURNING 
		id,id_fournisseur,nom,conditionnement,prix,reference_fournisseur,colisage;
		`, item.Id, item.IdFournisseur, item.Nom, item.Conditionnement, item.Prix, item.ReferenceFournisseur, item.Colisage)
	return ScanProduit(r)
}

// Delete Produit in the database and the return the id.
// Only the field 'Id' is used.
func (item Produit) Delete(tx *sql.Tx) (int64, error) {
	var deleted_id int64
	r := tx.QueryRow("DELETE FROM produits WHERE id = $1 RETURNING id;", item.Id)
	err := r.Scan(&deleted_id)
	return deleted_id, err
}

func ScanCommande(r *sql.Row) (Commande, error) {
	var s Commande
	if err := r.Scan(
		&s.Id,
		&s.IdProprietaire,
		&s.DateLivraison,
	); err != nil {
		return Commande{}, err
	}
	return s, nil
}

type Commandes map[int64]Commande

func (m Commandes) Ids() pq.Int64Array {
	out := make(pq.Int64Array, 0, len(m))
	for i := range m {
		out = append(out, i)
	}
	return out
}

func ScanCommandes(rs *sql.Rows) (Commandes, error) {
	structs := make(Commandes, 16)
	var err error
	for rs.Next() {
		var s Commande
		if err = rs.Scan(
			&s.Id,
			&s.IdProprietaire,
			&s.DateLivraison,
		); err != nil {
			return nil, err
		}
		structs[s.Id] = s
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

// Insert Commande in the database and returns the item with id filled.
func (item Commande) Insert(tx *sql.Tx) (out Commande, err error) {
	r := tx.QueryRow(`INSERT INTO commandes (
		id_proprietaire,date_livraison
		) VALUES (
		$1,$2
		) RETURNING 
		id,id_proprietaire,date_livraison;
		`, item.IdProprietaire, item.DateLivraison)
	return ScanCommande(r)
}

// Update Commande in the database and returns the new version.
func (item Commande) Update(tx *sql.Tx) (out Commande, err error) {
	r := tx.QueryRow(`UPDATE commandes SET (
		id_proprietaire,date_livraison
		) = (
		$2,$3
		) WHERE id = $1 RETURNING 
		id,id_proprietaire,date_livraison;
		`, item.Id, item.IdProprietaire, item.DateLivraison)
	return ScanCommande(r)
}

// Delete Commande in the database and the return the id.
// Only the field 'Id' is used.
func (item Commande) Delete(tx *sql.Tx) (int64, error) {
	var deleted_id int64
	r := tx.QueryRow("DELETE FROM commandes WHERE id = $1 RETURNING id;", item.Id)
	err := r.Scan(&deleted_id)
	return deleted_id, err
}