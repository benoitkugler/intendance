package controller

import (
	"database/sql"
	"errors"

	"github.com/benoitkugler/intendance/server/models"
)

// renvoie les fournisseurs lié à l'utilisateur
func (ct RequeteContext) loadIdsFournisseurs() (Set, error) {
	rows, err := ct.tx.Query("SELECT id_fournisseur FROM utilisateur_fournisseurs WHERE id_utilisateur = $1",
		ct.idProprietaire)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	_ids, err := models.ScanInts(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	return NewSetFromSlice(_ids), nil
}

func (s Server) GetIngredientProduits(ct RequeteContext, idIngredient int64) (IngredientProduits, error) {
	if err := ct.beginTx(s); err != nil {
		return IngredientProduits{}, err
	}
	ing := models.Ingredient{Id: idIngredient}
	produits, err := ing.GetProduits(ct.tx)
	if err != nil {
		return IngredientProduits{}, ErrorSQL(err)
	}

	// sélection des fournisseurs autorisés
	idsFournisseurs, err := ct.loadIdsFournisseurs()
	if err != nil {
		return IngredientProduits{}, err
	}
	var out IngredientProduits
	for _, produit := range produits {
		if idsFournisseurs.Has(produit.IdFournisseur) {
			out.Produits = append(out.Produits, produit)
		}
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
	validUnite := ing.Unite == models.Piece || ing.Unite == produit.Conditionnement.Unite
	validConditionnement := ing.Conditionnement.IsNull() || ing.Conditionnement == produit.Conditionnement
	if !(validUnite && validConditionnement) {
		return ErrorLieIngredientProduit(validUnite, validConditionnement)
	}

	idsFournisseurs, err := ct.loadIdsFournisseurs()
	if err != nil {
		return err
	}
	if !idsFournisseurs.Has(produit.IdFournisseur) {
		return errors.New("Le fournisseur du produit ne fait pas partie de vos fournisseurs.")
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
