// DON'T EDIT - automatically generated by structgen //

package models

import "database/sql"

func queriesCommande(tx *sql.Tx, item Commande) (Commande, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM commandes")
	if err != nil {
		return item, err
	}
	items, err := ScanCommandes(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectCommande(tx, item.Id)

	return item, err
}

func queriesCommandeProduit(tx *sql.Tx, item CommandeProduit) (CommandeProduit, error) {
	err := InsertManyCommandeProduits(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM commande_produits")
	if err != nil {
		return item, err
	}
	items, err := ScanCommandeProduits(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM commande_produits WHERE 
		id_commande = $1 AND id_produit = $2;`, item.IdCommande, item.IdProduit)
	_, err = ScanCommandeProduit(row)

	return item, err
}

func queriesDefautProduit(tx *sql.Tx, item DefautProduit) (DefautProduit, error) {
	err := InsertManyDefautProduits(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM defaut_produits")
	if err != nil {
		return item, err
	}
	items, err := ScanDefautProduits(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM defaut_produits WHERE 
		id_utilisateur = $1 AND id_ingredient = $2 AND id_fournisseur = $3 AND id_produit = $4;`, item.IdUtilisateur, item.IdIngredient, item.IdFournisseur, item.IdProduit)
	_, err = ScanDefautProduit(row)

	return item, err
}

func queriesFournisseur(tx *sql.Tx, item Fournisseur) (Fournisseur, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM fournisseurs")
	if err != nil {
		return item, err
	}
	items, err := ScanFournisseurs(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectFournisseur(tx, item.Id)

	return item, err
}

func queriesGroupe(tx *sql.Tx, item Groupe) (Groupe, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM groupes")
	if err != nil {
		return item, err
	}
	items, err := ScanGroupes(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectGroupe(tx, item.Id)

	return item, err
}

func queriesIngredient(tx *sql.Tx, item Ingredient) (Ingredient, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM ingredients")
	if err != nil {
		return item, err
	}
	items, err := ScanIngredients(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectIngredient(tx, item.Id)

	return item, err
}

func queriesIngredientProduit(tx *sql.Tx, item IngredientProduit) (IngredientProduit, error) {
	err := InsertManyIngredientProduits(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM ingredient_produits")
	if err != nil {
		return item, err
	}
	items, err := ScanIngredientProduits(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM ingredient_produits WHERE 
		id_ingredient = $1 AND id_produit = $2 AND id_utilisateur = $3;`, item.IdIngredient, item.IdProduit, item.IdUtilisateur)
	_, err = ScanIngredientProduit(row)

	return item, err
}

func queriesLienIngredient(tx *sql.Tx, item LienIngredient) (LienIngredient, error) {
	err := InsertManyLienIngredients(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM lien_ingredients")
	if err != nil {
		return item, err
	}
	items, err := ScanLienIngredients(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM lien_ingredients WHERE 
		id_ingredient = $1;`, item.IdIngredient)
	_, err = ScanLienIngredient(row)

	return item, err
}

func queriesLivraison(tx *sql.Tx, item Livraison) (Livraison, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM livraisons")
	if err != nil {
		return item, err
	}
	items, err := ScanLivraisons(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectLivraison(tx, item.Id)

	return item, err
}

func queriesMenu(tx *sql.Tx, item Menu) (Menu, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM menus")
	if err != nil {
		return item, err
	}
	items, err := ScanMenus(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectMenu(tx, item.Id)

	return item, err
}

func queriesMenuIngredient(tx *sql.Tx, item MenuIngredient) (MenuIngredient, error) {
	err := InsertManyMenuIngredients(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM menu_ingredients")
	if err != nil {
		return item, err
	}
	items, err := ScanMenuIngredients(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM menu_ingredients WHERE 
		id_menu = $1 AND id_ingredient = $2;`, item.IdMenu, item.IdIngredient)
	_, err = ScanMenuIngredient(row)

	return item, err
}

func queriesMenuRecette(tx *sql.Tx, item MenuRecette) (MenuRecette, error) {
	err := InsertManyMenuRecettes(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM menu_recettes")
	if err != nil {
		return item, err
	}
	items, err := ScanMenuRecettes(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM menu_recettes WHERE 
		id_menu = $1 AND id_recette = $2;`, item.IdMenu, item.IdRecette)
	_, err = ScanMenuRecette(row)

	return item, err
}

func queriesProduit(tx *sql.Tx, item Produit) (Produit, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM produits")
	if err != nil {
		return item, err
	}
	items, err := ScanProduits(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectProduit(tx, item.Id)

	return item, err
}

func queriesRecette(tx *sql.Tx, item Recette) (Recette, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM recettes")
	if err != nil {
		return item, err
	}
	items, err := ScanRecettes(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectRecette(tx, item.Id)

	return item, err
}

func queriesRecetteIngredient(tx *sql.Tx, item RecetteIngredient) (RecetteIngredient, error) {
	err := InsertManyRecetteIngredients(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM recette_ingredients")
	if err != nil {
		return item, err
	}
	items, err := ScanRecetteIngredients(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM recette_ingredients WHERE 
		id_recette = $1 AND id_ingredient = $2;`, item.IdRecette, item.IdIngredient)
	_, err = ScanRecetteIngredient(row)

	return item, err
}

func queriesRepas(tx *sql.Tx, item Repas) (Repas, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM repass")
	if err != nil {
		return item, err
	}
	items, err := ScanRepass(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectRepas(tx, item.Id)

	return item, err
}

func queriesRepasGroupe(tx *sql.Tx, item RepasGroupe) (RepasGroupe, error) {
	err := InsertManyRepasGroupes(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM repas_groupes")
	if err != nil {
		return item, err
	}
	items, err := ScanRepasGroupes(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM repas_groupes WHERE 
		id_repas = $1 AND id_groupe = $2;`, item.IdRepas, item.IdGroupe)
	_, err = ScanRepasGroupe(row)

	return item, err
}

func queriesRepasIngredient(tx *sql.Tx, item RepasIngredient) (RepasIngredient, error) {
	err := InsertManyRepasIngredients(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM repas_ingredients")
	if err != nil {
		return item, err
	}
	items, err := ScanRepasIngredients(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM repas_ingredients WHERE 
		id_repas = $1 AND id_ingredient = $2;`, item.IdRepas, item.IdIngredient)
	_, err = ScanRepasIngredient(row)

	return item, err
}

func queriesRepasRecette(tx *sql.Tx, item RepasRecette) (RepasRecette, error) {
	err := InsertManyRepasRecettes(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM repas_recettes")
	if err != nil {
		return item, err
	}
	items, err := ScanRepasRecettes(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM repas_recettes WHERE 
		id_repas = $1 AND id_recette = $2;`, item.IdRepas, item.IdRecette)
	_, err = ScanRepasRecette(row)

	return item, err
}

func queriesSejour(tx *sql.Tx, item Sejour) (Sejour, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM sejours")
	if err != nil {
		return item, err
	}
	items, err := ScanSejours(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectSejour(tx, item.Id)

	return item, err
}

func queriesSejourFournisseur(tx *sql.Tx, item SejourFournisseur) (SejourFournisseur, error) {
	err := InsertManySejourFournisseurs(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM sejour_fournisseurs")
	if err != nil {
		return item, err
	}
	items, err := ScanSejourFournisseurs(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM sejour_fournisseurs WHERE 
		id_utilisateur = $1 AND id_sejour = $2 AND id_fournisseur = $3;`, item.IdUtilisateur, item.IdSejour, item.IdFournisseur)
	_, err = ScanSejourFournisseur(row)

	return item, err
}

func queriesUtilisateur(tx *sql.Tx, item Utilisateur) (Utilisateur, error) {
	item, err := item.Insert(tx)

	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM utilisateurs")
	if err != nil {
		return item, err
	}
	items, err := ScanUtilisateurs(rows)
	if err != nil {
		return item, err
	}

	_ = items.Ids()

	item, err = item.Update(tx)
	if err != nil {
		return item, err
	}
	_, err = SelectUtilisateur(tx, item.Id)

	return item, err
}

func queriesUtilisateurFournisseur(tx *sql.Tx, item UtilisateurFournisseur) (UtilisateurFournisseur, error) {
	err := InsertManyUtilisateurFournisseurs(tx, item)
	if err != nil {
		return item, err
	}
	rows, err := tx.Query("SELECT * FROM utilisateur_fournisseurs")
	if err != nil {
		return item, err
	}
	items, err := ScanUtilisateurFournisseurs(rows)
	if err != nil {
		return item, err
	}

	_ = len(items)

	row := tx.QueryRow(`SELECT * FROM utilisateur_fournisseurs WHERE 
		id_utilisateur = $1 AND id_fournisseur = $2;`, item.IdUtilisateur, item.IdFournisseur)
	_, err = ScanUtilisateurFournisseur(row)

	return item, err
}
