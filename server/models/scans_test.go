// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package models

import (
	"database/sql"
	"fmt"
	"math/rand"
)

func randUtilisateur() Utilisateur {
	return Utilisateur{
		Id:        rand.Int63n(1 << 20),
		Password:  randstring(),
		Mail:      randstring(),
		PrenomNom: randstring(),
	}
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
	row := tx.QueryRow("SELECT * FROM utilisateurs WHERE id = $1", item.Id)

	_, err = ScanUtilisateur(row)
	return item, err
}

func randIngredient() Ingredient {
	return Ingredient{
		Id:              rand.Int63n(1 << 20),
		Nom:             randstring(),
		Unite:           randUnite(),
		Categorie:       randstring(),
		Callories:       randCallories(),
		Conditionnement: randConditionnement(),
	}
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
	row := tx.QueryRow("SELECT * FROM ingredients WHERE id = $1", item.Id)

	_, err = ScanIngredient(row)
	return item, err
}

func randRecette() Recette {
	return Recette{
		Id:             rand.Int63n(1 << 20),
		IdProprietaire: randNullInt64(),
		Nom:            randstring(),
		ModeEmploi:     randstring(),
	}
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
	row := tx.QueryRow("SELECT * FROM recettes WHERE id = $1", item.Id)

	_, err = ScanRecette(row)
	return item, err
}

func randRecetteIngredient() RecetteIngredient {
	return RecetteIngredient{
		IdRecette:    rand.Int63n(1 << 20),
		IdIngredient: rand.Int63n(1 << 20),
		Quantite:     randfloat64(),
		Cuisson:      randstring(),
	}
}

func queriesRecetteIngredient(tx *sql.Tx, item RecetteIngredient) (RecetteIngredient, error) {
	err := InsertManyRecetteIngredients(tx, []RecetteIngredient{item})
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

	fmt.Println(len(items))

	row := tx.QueryRow(`SELECT * FROM recette_ingredients WHERE 
		id_recette = $1 AND id_ingredient = $2;`, item.IdRecette, item.IdIngredient)

	_, err = ScanRecetteIngredient(row)
	return item, err
}

func randMenu() Menu {
	return Menu{
		Id:             rand.Int63n(1 << 20),
		IdProprietaire: randNullInt64(),
		Commentaire:    randstring(),
	}
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
	row := tx.QueryRow("SELECT * FROM menus WHERE id = $1", item.Id)

	_, err = ScanMenu(row)
	return item, err
}

func randMenuIngredient() MenuIngredient {
	return MenuIngredient{
		IdMenu:       rand.Int63n(1 << 20),
		IdIngredient: rand.Int63n(1 << 20),
		Quantite:     randfloat64(),
		Cuisson:      randstring(),
	}
}

func queriesMenuIngredient(tx *sql.Tx, item MenuIngredient) (MenuIngredient, error) {
	err := InsertManyMenuIngredients(tx, []MenuIngredient{item})
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

	fmt.Println(len(items))

	row := tx.QueryRow(`SELECT * FROM menu_ingredients WHERE 
		id_menu = $1 AND id_ingredient = $2;`, item.IdMenu, item.IdIngredient)

	_, err = ScanMenuIngredient(row)
	return item, err
}

func randMenuRecette() MenuRecette {
	return MenuRecette{
		IdMenu:    rand.Int63n(1 << 20),
		IdRecette: rand.Int63n(1 << 20),
	}
}

func queriesMenuRecette(tx *sql.Tx, item MenuRecette) (MenuRecette, error) {
	err := InsertManyMenuRecettes(tx, []MenuRecette{item})
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

	fmt.Println(len(items))

	row := tx.QueryRow(`SELECT * FROM menu_recettes WHERE 
		id_menu = $1 AND id_recette = $2;`, item.IdMenu, item.IdRecette)

	_, err = ScanMenuRecette(row)
	return item, err
}

func randSejour() Sejour {
	return Sejour{
		Id:             rand.Int63n(1 << 20),
		IdProprietaire: rand.Int63n(1 << 20),
		DateDebut:      randTime(),
		Nom:            randstring(),
	}
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
	row := tx.QueryRow("SELECT * FROM sejours WHERE id = $1", item.Id)

	_, err = ScanSejour(row)
	return item, err
}

func randRepas() Repas {
	return Repas{
		Id:          rand.Int63n(1 << 20),
		IdSejour:    rand.Int63n(1 << 20),
		IdMenu:      rand.Int63n(1 << 20),
		NbPersonnes: rand.Int63n(1 << 20),
		JourOffset:  rand.Int63n(1 << 20),
		Horaire:     randHoraire(),
	}
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
	row := tx.QueryRow("SELECT * FROM repass WHERE id = $1", item.Id)

	_, err = ScanRepas(row)
	return item, err
}

func randFournisseur() Fournisseur {
	return Fournisseur{
		Id:             rand.Int63n(1 << 20),
		Nom:            randstring(),
		DelaiCommande:  rand.Int63n(1 << 20),
		JoursLivraison: randJoursLivraison(),
	}
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
	row := tx.QueryRow("SELECT * FROM fournisseurs WHERE id = $1", item.Id)

	_, err = ScanFournisseur(row)
	return item, err
}

func randProduit() Produit {
	return Produit{
		Id:                   rand.Int63n(1 << 20),
		IdFournisseur:        rand.Int63n(1 << 20),
		Nom:                  randstring(),
		Conditionnement:      randConditionnement(),
		Prix:                 randfloat64(),
		ReferenceFournisseur: randstring(),
		Colisage:             rand.Int63n(1 << 20),
	}
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
	row := tx.QueryRow("SELECT * FROM produits WHERE id = $1", item.Id)

	_, err = ScanProduit(row)
	return item, err
}

func randIngredientProduit() IngredientProduit {
	return IngredientProduit{
		IdIngredient: rand.Int63n(1 << 20),
		IdProduit:    rand.Int63n(1 << 20),
	}
}

func queriesIngredientProduit(tx *sql.Tx, item IngredientProduit) (IngredientProduit, error) {
	err := InsertManyIngredientProduits(tx, []IngredientProduit{item})
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

	fmt.Println(len(items))

	row := tx.QueryRow(`SELECT * FROM ingredient_produits WHERE 
		id_ingredient = $1 AND id_produit = $2;`, item.IdIngredient, item.IdProduit)

	_, err = ScanIngredientProduit(row)
	return item, err
}

func randCommande() Commande {
	return Commande{
		Id:             rand.Int63n(1 << 20),
		IdProprietaire: rand.Int63n(1 << 20),
		DateLivraison:  randTime(),
	}
}

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
	row := tx.QueryRow("SELECT * FROM commandes WHERE id = $1", item.Id)

	_, err = ScanCommande(row)
	return item, err
}