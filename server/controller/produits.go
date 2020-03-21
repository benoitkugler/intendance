package controller

import (
	"database/sql"
	"fmt"

	"github.com/benoitkugler/intendance/server/models"
)

func (ct RequeteContext) loadFournisseurs() (models.Fournisseurs, error) {
	rows, err := ct.tx.Query(`SELECT fournisseurs.* FROM fournisseurs 
		JOIN utilisateur_fournisseurs ON utilisateur_fournisseurs.id_fournisseur = fournisseurs.id 
		WHERE utilisateur_fournisseurs.id_utilisateur = $1`, ct.idProprietaire)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	out, err := models.ScanFournisseurs(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	return out, nil
}

func (ct RequeteContext) checkFournisseurs(produit models.Produit) error {
	fourns, err := ct.loadFournisseurs()
	if err != nil {
		return err
	}
	idsFournisseurs := fourns.Ids().AsSet()
	if !idsFournisseurs.Has(produit.IdFournisseur) {
		return fmt.Errorf("Le fournisseur du produit %s ne fait pas partie de vos fournisseurs.", produit.Nom)
	}
	return nil
}

// vérifie que tous les ingrédients liés au produit
// appartiennent à l'utilisateur courant
func (ct RequeteContext) checkProprioAllIngredient(produit models.Produit) error {
	rows, err := ct.tx.Query("SELECT id_ajouteur FROM ingredient_produits WHERE id_produit = $1", produit.Id)
	if err != nil {
		return ErrorSQL(err)
	}
	ids, err := models.ScanInts(rows)
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
	rows, err := ct.tx.Query("SELECT  * FROM commande_produits WHERE id_produit = $1", produit.Id)
	if err != nil {
		return ErrorSQL(err)
	}
	cps, err := models.ScanCommandeProduits(rows)
	if err != nil {
		return ErrorSQL(err)
	}
	if l := len(cps); l > 0 {
		return fmt.Errorf("Le produit <b>%s</b> est déjà présent dans %d commande(s). Sa modification est désactivé.",
			produit.Nom, l)
	}
	return nil
}

// LoadFournisseurs renvoie les fournisseurs associés à l'utilisateur courant
func (s Server) LoadFournisseurs(ct RequeteContext) (models.Fournisseurs, error) {
	if err := ct.beginTx(s); err != nil {
		return nil, err
	}
	defer ct.rollbackTx(nil) // pas de modifications
	return ct.loadFournisseurs()
}

func (s Server) GetIngredientProduits(ct RequeteContext, idIngredient int64) (IngredientProduits, error) {
	if err := ct.beginTx(s); err != nil {
		return IngredientProduits{}, err
	}
	// sélection des fournisseurs autorisés
	fourns, err := ct.loadFournisseurs()
	if err != nil {
		return IngredientProduits{}, err
	}

	ing := models.Ingredient{Id: idIngredient}
	produits, err := ing.GetProduits(ct.tx, fourns)
	if err != nil {
		return IngredientProduits{}, ErrorSQL(err)
	}

	var out IngredientProduits
	for _, produit := range produits {
		out.Produits = append(out.Produits, produit)
	}

	row := ct.tx.QueryRow("SELECT id_produit FROM produits_par_defaut WHERE id_utilisateur = $1 AND id_ingredient = $2",
		ct.idProprietaire, idIngredient)
	var idDefault sql.NullInt64
	err = row.Scan(&idDefault.Int64)
	if err == sql.ErrNoRows {
		// pas de valeur par défaut, idDefault.Valid reste à false
	} else if err != nil { // "vraie" erreur
		return IngredientProduits{}, ErrorSQL(err)
	} else { // on a trouvé une valeur par défault
		idDefault.Valid = true
	}
	if err = ct.commitTx(); err != nil {
		return out, err
	}

	out.IdDefault = idDefault
	return out, nil
}

// AjouteIngredientProduit crée le produit donné et l'associe à l'ingrédient.
func (s Server) AjouteIngredientProduit(ct RequeteContext, idIngredient int64, produit models.Produit) error {
	contrainte := ContrainteProduit{produit: produit}
	if err := contrainte.Check(); err != nil {
		return err
	}

	if err := ct.beginTx(s); err != nil {
		return err
	}
	row := ct.tx.QueryRow("SELECT * FROM ingredients WHERE id = $1", idIngredient)
	ing, err := models.ScanIngredient(row)
	if err != nil {
		return ErrorSQL(err)
	}

	// L'unité Piece est particulière car elle laisse la responsabilité au produit
	// pour définir l'unité utilisée.
	// Pour les autres, les unités du produit et de l'ingrédient doivent être identiques.
	contrainte2 := ContrainteIngredientProduit{ingredient: ing, produit: produit}
	if err := contrainte2.Check(); err != nil {
		return err
	}

	if err := ct.checkFournisseurs(produit); err != nil {
		return err
	}

	produit, err = produit.Insert(ct.tx)
	if err != nil {
		return ErrorSQL(err)
	}
	err = models.InsertManyIngredientProduits(ct.tx, []models.IngredientProduit{
		{IdIngredient: idIngredient, IdProduit: produit.Id, IdAjouteur: ct.idProprietaire},
	})
	if err != nil {
		return ct.rollbackTx(err)
	}
	return ct.commitTx()
}

// UpdateProduit modifie un produit, sous réserve qu'il ne soit pas encore utilisé dans une commande.
func (s Server) UpdateProduit(ct RequeteContext, produit models.Produit) (out models.Produit, err error) {
	contrainte := ContrainteProduit{produit: produit}
	if err := contrainte.Check(); err != nil {
		return out, err
	}

	if err := ct.beginTx(s); err != nil {
		return out, err
	}

	if err := ct.checkFournisseurs(produit); err != nil {
		return out, err
	}

	// vérification de la présence dans les commandes
	if err := ct.checkCommandes(produit); err != nil {
		return out, err
	}

	// vérification des contraintes ingrédient/produit
	rows, err := ct.tx.Query(`SELECT ingredients.* FROM ingredients
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
	out, err = produit.Update(ct.tx)
	if err != nil {
		return out, ct.rollbackTx(err)
	}
	err = ct.commitTx()
	return out, err
}

func (s Server) DeleteProduit(ct RequeteContext, idProduit int64) error {
	if err := ct.beginTx(s); err != nil {
		return err
	}

	row := ct.tx.QueryRow("SELECT * FROM produits WHERE id = $1", idProduit)
	produit, err := models.ScanProduit(row)
	if err != nil {
		return ErrorSQL(err)
	}

	if err := ct.checkFournisseurs(produit); err != nil {
		return err
	}

	if err := ct.checkProprioAllIngredient(produit); err != nil {
		return err
	}

	if err := ct.checkCommandes(produit); err != nil {
		return err
	}

	// suppression de la liaison ingrédient
	_, err = ct.tx.Exec("DELETE FROM ingredient_produits WHERE id_produit = $1", idProduit)
	if err != nil {
		return ct.rollbackTx(err)
	}

	// suppression du produit
	_, err = produit.Delete(ct.tx)
	if err != nil {
		return ct.rollbackTx(err)
	}
	return ct.commitTx()
}
