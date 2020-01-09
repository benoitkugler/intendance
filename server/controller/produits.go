package controller

import (
	"database/sql"

	"github.com/benoitkugler/intendance/server/models"
)

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
	rows, err := ct.tx.Query("SELECT id_fournisseur FROM utilisateur_fournisseurs WHERE id_utilisateur = $1",
		ct.idProprietaire)
	if err != nil {
		return IngredientProduits{}, ErrorSQL(err)
	}
	_ids, err := models.ScanInts(rows)
	if err != nil {
		return IngredientProduits{}, ErrorSQL(err)
	}
	idsFournisseurs := NewSetFromSlice(_ids)
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
	ct.commit()

	out.IdDefault = idDefault
	return out, nil
}
