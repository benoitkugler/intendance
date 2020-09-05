package controller

import (
	"fmt"
	"sort"

	"github.com/benoitkugler/intendance/server/models"
)

// vérifie que tous les ingrédients liés au produit
// appartiennent à l'utilisateur courant
func (ct RequeteContext) checkProprioAllIngredient(produit models.Produit) error {
	rows, err := ct.DB.Query("SELECT id_utilisateur FROM ingredient_produits WHERE id_produit = $1", produit.Id)
	if err != nil {
		return ErrorSQL(err)
	}
	ids, err := models.ScanIds(rows)
	if err != nil {
		return ErrorSQL(err)
	}
	set := models.NewSetFromSlice(ids)
	if len(set) > 1 { // le produit est partagé
		return fmt.Errorf(`Le produit %s est partagé entre plusieurs utilisateurs. 
			Sa suppression est donc désactivée.`, produit.Nom)
	}
	return nil
}

// vérification de la présence dans les commandes
func (ct RequeteContext) checkCommandes(produit models.Produit) error {
	cps, err := models.SelectCommandeProduitsByIdProduits(ct.DB, produit.Id)
	if err != nil {
		return ErrorSQL(err)
	}
	if l := len(cps); l > 0 {
		return fmt.Errorf("Le produit <b>%s</b> est déjà présent dans %d commande(s). Sa modification est désactivé.",
			produit.Nom, l)
	}
	return nil
}

func (ct RequeteContext) GetIngredientProduits(idIngredient int64) (IngredientProduits, error) {
	// sélection des fournisseurs autorisés
	fourns, err := ct.loadFournisseurs()
	if err != nil {
		return IngredientProduits{}, err
	}

	ing := models.Ingredient{Id: idIngredient}
	produits, err := ing.GetProduits(ct.DB, fourns)
	if err != nil {
		return IngredientProduits{}, ErrorSQL(err)
	}

	var out IngredientProduits
	for _, produit := range produits {
		out.Produits = append(out.Produits, produit)
	}
	sort.Slice(out.Produits, func(i, j int) bool {
		return out.Produits[i].Nom < out.Produits[j].Nom
	})
	sort.SliceStable(out.Produits, func(i, j int) bool {
		return out.Produits[i].IdLivraison < out.Produits[j].IdLivraison
	})

	rows, err := ct.DB.Query("SELECT id_produit FROM defaut_produits WHERE id_utilisateur = $1 AND id_ingredient = $2",
		ct.IdProprietaire, idIngredient)
	if err != nil {
		return IngredientProduits{}, ErrorSQL(err)
	}
	defaults, err := models.ScanIds(rows)
	if err != nil {
		return IngredientProduits{}, ErrorSQL(err)
	}
	out.Defaults = models.NewSetFromSlice(defaults)
	return out, nil
}

// AjouteIngredientProduit crée le produit donné et l'associe à l'ingrédient.
func (ct RequeteContext) AjouteIngredientProduit(idIngredient int64, produit models.Produit) (models.Produit, error) {
	contrainte := ContrainteProduit{produit: produit}
	if err := contrainte.Check(); err != nil {
		return produit, err
	}

	ing, err := models.SelectIngredient(ct.DB, idIngredient)
	if err != nil {
		return produit, ErrorSQL(err)
	}

	// L'unité Piece est particulière car elle laisse la responsabilité au produit
	// pour définir l'unité utilisée.
	// Pour les autres, les unités du produit et de l'ingrédient doivent être identiques.
	contrainte2 := ContrainteIngredientProduit{ingredient: ing, produit: produit}
	if err := contrainte2.Check(); err != nil {
		return produit, err
	}

	if _, err := ct.checkFournisseurs(produit); err != nil {
		return produit, err
	}

	tx, err := ct.beginTx()
	if err != nil {
		return produit, err
	}
	produit, err = produit.Insert(tx)
	if err != nil {
		return produit, tx.rollback(err)
	}
	err = models.InsertManyIngredientProduits(tx.Tx,
		models.IngredientProduit{IdIngredient: idIngredient, IdProduit: produit.Id, IdUtilisateur: ct.IdProprietaire})
	if err != nil {
		return produit, tx.rollback(err)
	}
	return produit, tx.Commit()
}

// UpdateProduit modifie un produit, sous réserve qu'il ne soit pas encore utilisé dans une commande.
func (ct RequeteContext) UpdateProduit(produit models.Produit) (out models.Produit, err error) {
	contrainte := ContrainteProduit{produit: produit}
	if err := contrainte.Check(); err != nil {
		return out, err
	}

	if _, err := ct.checkFournisseurs(produit); err != nil {
		return out, err
	}

	// vérification de la présence dans les commandes
	if err := ct.checkCommandes(produit); err != nil {
		return out, err
	}

	// vérification des contraintes ingrédient/produit
	rows, err := ct.DB.Query(`SELECT ingredients.* FROM ingredients
		JOIN ingredient_produits ON ingredient_produits.id_ingredient = ingredients.id
		WHERE ingredient_produits.id_produit = $1`, produit.Id)
	if err != nil {
		return out, ErrorSQL(err)
	}
	ings, err := models.ScanIngredients(rows)
	if err != nil {
		return out, ErrorSQL(err)
	}
	for _, ing := range ings {
		contrainte := ContrainteIngredientProduit{ingredient: ing, produit: produit}
		if err := contrainte.Check(); err != nil {
			return out, err
		}
	}

	// finalement, mise à jour
	out, err = produit.Update(ct.DB)
	return out, ErrorSQL(err)
}

func (ct RequeteContext) DeleteProduit(idProduit int64) error {
	produit, err := models.SelectProduit(ct.DB, idProduit)
	if err != nil {
		return ErrorSQL(err)
	}

	if _, err := ct.checkFournisseurs(produit); err != nil {
		return err
	}

	if err := ct.checkProprioAllIngredient(produit); err != nil {
		return err
	}

	if err := ct.checkCommandes(produit); err != nil {
		return err
	}

	// suppression des préférences par cascade

	// suppression de la liaison ingrédient par cascade

	// suppression du produit
	_, err = models.DeleteProduitById(ct.DB, produit.Id)
	return ErrorSQL(err)
}

// SetDefautProduit met à jour le produit par défaut pour l'ingrédient donné
// `GetIngredientProduits` devrait être appelé ensuite.
func (ct RequeteContext) SetDefautProduit(idIngredient int64, idProduit int64, on bool) error {
	produit, err := models.SelectProduit(ct.DB, idProduit)
	if err != nil {
		return ErrorSQL(err)
	}

	livraison, err := ct.checkFournisseurs(produit)
	if err != nil {
		return err
	}

	tx, err := ct.beginTx()
	if err != nil {
		return err
	}
	// on enlève un éventuel produit par défaut du même fournisseur
	// par la même occasion, on supprime la valeur par défaut (cas `on` == false )
	_, err = tx.Exec(`DELETE FROM defaut_produits 
		WHERE defaut_produits.id_fournisseur = $1
			AND defaut_produits.id_utilisateur = $2
			AND defaut_produits.id_ingredient = $3`,
		livraison.IdFournisseur, ct.IdProprietaire, idIngredient)
	if err != nil {
		return tx.rollback(err)
	}
	if on { // on ajoute le nouveau défaut
		dp := models.DefautProduit{
			IdFournisseur: livraison.IdFournisseur, // déduit du produit
			IdIngredient:  idIngredient,
			IdProduit:     idProduit,
			IdUtilisateur: ct.IdProprietaire,
		}
		if err := models.InsertManyDefautProduits(tx.Tx, dp); err != nil {
			return tx.rollback(err)
		}
	}
	return tx.Commit()
}
