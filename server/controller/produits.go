package controller

import (
	"fmt"
	"sort"

	"github.com/benoitkugler/intendance/server/models"
)

// vérifie que tous les ingrédients liés au produit
// appartiennent à l'utilisateur courant
func (ct RequeteContext) checkProprioAllIngredient(produit models.Produit) error {
	rows, err := ct.tx.Query("SELECT id_utilisateur FROM ingredient_produits WHERE id_produit = $1", produit.Id)
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

// vérifie que la livraison éventuelle est lié au fournisseur du produit
// cette contrainte n'est pas assuré par le modèle SQL
func (ct RequeteContext) checkLivraison(produit models.Produit) error {
	if !produit.IdLivraison.Valid { // rien à vérifier
		return nil
	}
	row := ct.tx.QueryRow("SELECT * FROM livraisons WHERE id = $1", produit.IdLivraison.Int64)
	livraison, err := models.ScanLivraison(row)
	if err != nil {
		return ErrorSQL(err)
	}
	if !livraison.IdFournisseur.Valid { // la contrainte est universelle -> OK
		return nil
	}
	if livraison.IdFournisseur.Int64 != produit.IdFournisseur {
		return fmt.Errorf("La contrainte de livraison %s est liée à un fournisseur différent de celui du produit %s.",
			livraison.Nom, produit.Nom)
	}
	return nil
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
	sort.Slice(out.Produits, func(i, j int) bool {
		return out.Produits[i].Nom < out.Produits[j].Nom
	})
	sort.SliceStable(out.Produits, func(i, j int) bool {
		return out.Produits[i].IdFournisseur < out.Produits[j].IdFournisseur
	})

	rows, err := ct.tx.Query("SELECT id_produit FROM defaut_produits WHERE id_utilisateur = $1 AND id_ingredient = $2",
		ct.idProprietaire, idIngredient)
	if err != nil {
		return IngredientProduits{}, ErrorSQL(err)
	}
	defaults, err := models.ScanInts(rows)
	if err != nil {
		return IngredientProduits{}, ErrorSQL(err)
	}
	out.Defaults = models.NewSetFromSlice(defaults)
	if err = ct.commitTx(); err != nil {
		return out, err
	}

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
		{IdIngredient: idIngredient, IdProduit: produit.Id, IdUtilisateur: ct.idProprietaire},
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

	if err := ct.checkLivraison(produit); err != nil {
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

	// suppression des préférences
	_, err = ct.tx.Exec("DELETE FROM defaut_produits WHERE id_produit = $1", idProduit)
	if err != nil {
		return ct.rollbackTx(err)
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

// SetDefautProduit met à jour le produit par défaut pour l'ingrédient donné
// `GetIngredientProduits` devrait être appelé ensuite.
func (s Server) SetDefautProduit(ct RequeteContext, idIngredient int64, idProduit int64, on bool) error {
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

	// on enlève un éventuel produit par défaut du même fournisseur
	// par la même occasion, on supprime la valeur par défaut (cas `off` == false )
	_, err = ct.tx.Exec("DELETE FROM defaut_produits WHERE id_utilisateur = $1 AND id_ingredient = $2 AND id_fournisseur = $3",
		ct.idProprietaire, idIngredient, produit.IdFournisseur)
	if err != nil {
		return ct.rollbackTx(err)
	}
	if on { // on ajoute le nouveau défaut
		dp := models.DefautProduit{
			IdFournisseur: produit.IdFournisseur, // déduit du produit
			IdIngredient:  idIngredient,
			IdProduit:     idProduit,
			IdUtilisateur: ct.idProprietaire}
		if err := models.InsertManyDefautProduits(ct.tx, []models.DefautProduit{dp}); err != nil {
			return ct.rollbackTx(err)
		}
	}
	return ct.commitTx()
}
