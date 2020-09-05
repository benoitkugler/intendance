package controller

import (
	"fmt"

	"github.com/benoitkugler/intendance/server/models"
)

func (ct RequeteContext) loadFournisseurs() (models.Fournisseurs, error) {
	rows, err := ct.DB.Query(`SELECT fournisseurs.* FROM fournisseurs 
		JOIN utilisateur_fournisseurs ON utilisateur_fournisseurs.id_fournisseur = fournisseurs.id 
		WHERE utilisateur_fournisseurs.id_utilisateur = $1`, ct.IdProprietaire)
	if err != nil {
		return nil, ErrorSQL(err)
	}
	out, err := models.ScanFournisseurs(rows)
	return out, ErrorSQL(err)
}

// renvoie les livraisons des fournisseurs donnés
func (ct RequeteContext) loadLivraisons(fournisseurs models.Fournisseurs) (models.Livraisons, error) {
	livraisons, err := models.SelectLivraisonsByIdFournisseurs(ct.DB, fournisseurs.Ids()...)
	return livraisons, ErrorSQL(err)
}

func (ct RequeteContext) hasFournisseur(idFournisseur int64) (bool, error) {
	fourns, err := ct.loadFournisseurs()
	if err != nil {
		return false, fmt.Errorf("Impossible de vérifier les fournisseurs : %s", err)
	}
	idsFournisseurs := fourns.Ids().AsSet()
	return idsFournisseurs.Has(idFournisseur), nil
}

// vérifie que le fournisseur du produit fait partie
// des fournisseurs associés à l'utilisateur courant
// renvoie la livraison associée
func (ct RequeteContext) checkFournisseurs(produit models.Produit) (models.Livraison, error) {
	livraison, err := models.SelectLivraison(ct.DB, produit.IdLivraison)
	if err != nil {
		return livraison, ErrorSQL(fmt.Errorf("can't find livraison : %s", err))
	}
	hasFournisseur, err := ct.hasFournisseur(livraison.IdFournisseur)
	if err != nil {
		return livraison, err
	}
	if !hasFournisseur {
		return livraison, fmt.Errorf("Le fournisseur du produit %s ne fait pas partie de vos fournisseurs.", produit.Nom)
	}
	return livraison, nil
}

func (ct RequeteContext) checkLivraisonFournisseur(livraison models.Livraison) error {
	ok, err := ct.hasFournisseur(livraison.IdFournisseur)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("Le fournisseur (%d) ne fait pas partie de vos fournisseurs.", livraison.IdFournisseur)
	}
	return nil
}

// LoadFournisseurs renvoie les fournisseurs associés à l'utilisateur courant,
// ainsi que les contraints de livraisons pertinentes.
func (ct RequeteContext) LoadFournisseurs() (models.Fournisseurs, models.Livraisons, error) {
	fournisseurs, err := ct.loadFournisseurs()
	if err != nil {
		return nil, nil, err
	}

	livraisons, err := ct.loadLivraisons(fournisseurs)
	if err != nil {
		return nil, nil, err
	}
	return fournisseurs, livraisons, nil
}

// CreateFournisseur crée un fournisseur et le lie à l'utilisateur courant
// Une contrainte de livraison "standard" est automatiquement créée
func (ct RequeteContext) CreateFournisseur(fournisseur models.Fournisseur) (out models.Fournisseur, err error) {
	tx, err := ct.beginTx()
	if err != nil {
		return out, err
	}

	out, err = fournisseur.Insert(tx)
	if err != nil {
		return out, tx.rollback(err)
	}

	err = models.InsertManyUtilisateurFournisseurs(tx.Tx,
		models.UtilisateurFournisseur{IdFournisseur: out.Id, IdUtilisateur: ct.IdProprietaire})
	if err != nil {
		return out, tx.rollback(err)
	}

	// ajout d'une contraint de livraison par défaut
	livraison := models.Livraison{
		IdFournisseur:  out.Id,
		Nom:            "",
		JoursLivraison: models.JoursLivraison{true, true, true, true, true, false, false},
		DelaiCommande:  2,
		Anticipation:   1,
	}
	livraison, err = livraison.Insert(tx)
	if err != nil {
		return out, tx.rollback(err)
	}

	return out, tx.Commit()
}

func (ct RequeteContext) UpdateFournisseur(fournisseur models.Fournisseur) (models.Fournisseur, error) {
	hasF, err := ct.hasFournisseur(fournisseur.Id)
	if err != nil {
		return models.Fournisseur{}, err
	}
	if !hasF {
		return models.Fournisseur{}, fmt.Errorf("Le fournisseur %s ne fait pas partie de vos fournisseurs.", fournisseur.Nom)
	}
	fournisseur, err = fournisseur.Update(ct.DB)
	return fournisseur, ErrorSQL(err)
}

func (ct RequeteContext) DeleteFournisseur(idFournisseur int64) error {
	hasF, err := ct.hasFournisseur(idFournisseur)
	if err != nil {
		return err
	}
	if !hasF {
		return fmt.Errorf("Le fournisseur (%d) ne fait pas partie de vos fournisseurs.", idFournisseur)
	}
	tx, err := ct.beginTx()
	if err != nil {
		return err
	}

	rows, err := tx.Query(`SELECT  * FROM commande_produits 
		JOIN produits ON produits.id = commande_produits.id_produit 
		JOIN livraisons ON produits.id_livraison = livraisons.id
		WHERE livraisons.id_fournisseur = $1`, idFournisseur)
	if err != nil {
		return tx.rollback(err)
	}
	cps, err := models.ScanCommandeProduits(rows)
	if err != nil {
		return tx.rollback(err)
	}
	if L := len(cps); L > 0 {
		_ = tx.rollback(nil)
		return fmt.Errorf("%d produit(s) lié(s) au fournisseur sont déjà utilisés dans une commande.", L)
	}

	// sejours
	err = models.DeleteSejourFournisseursByIdFournisseurs(tx, idFournisseur)
	if err != nil {
		return tx.rollback(err)
	}

	// ingredients
	_, err = tx.Exec(`DELETE FROM ingredient_produits USING produits, livraisons
		WHERE ingredient_produits.id_produit = produits.id 
		AND produits.id_livraison = livraisons.id
		AND livraisons.id_fournisseur = $1`, idFournisseur)
	if err != nil {
		return tx.rollback(err)
	}

	// defaut
	_, err = tx.Exec(`DELETE FROM defaut_produits USING produits, livraisons
		WHERE defaut_produits.id_produit = produits.id 
		AND produits.id_livraison = livraisons.id
		AND livraisons.id_fournisseur = $1`, idFournisseur)
	if err != nil {
		return tx.rollback(err)
	}

	// produits
	_, err = tx.Exec(`DELETE FROM produits USING livraisons
		WHERE produits.id_livraison = livraisons.id AND livraisons.id_fournisseur = $1`, idFournisseur)
	if err != nil {
		return tx.rollback(err)
	}

	// livraisons par cascade

	// utilisateurs
	err = models.DeleteUtilisateurFournisseursByIdFournisseurs(tx, idFournisseur)
	if err != nil {
		return tx.rollback(err)
	}

	_, err = models.DeleteFournisseurById(tx, idFournisseur)
	if err != nil {
		return tx.rollback(err)
	}
	return tx.Commit()
}

func (ct RequeteContext) UpdateSejourFournisseurs(idSejour int64, idsFournisseurs []int64) error {
	if err := ct.proprioSejour(models.Sejour{Id: idSejour}, false); err != nil {
		return err
	}
	tx, err := ct.beginTx()
	if err != nil {
		return err
	}

	// reset les fournisseurs du séjour ...
	err = models.DeleteSejourFournisseursByIdSejours(tx, idSejour)
	if err != nil {
		return tx.rollback(err)
	}
	sf := make([]models.SejourFournisseur, len(idsFournisseurs))
	for i, id := range idsFournisseurs {
		sf[i] = models.SejourFournisseur{IdUtilisateur: ct.IdProprietaire, IdSejour: idSejour, IdFournisseur: id}
	}
	// ... et rajoute les nouveaux
	if err := models.InsertManySejourFournisseurs(tx.Tx, sf...); err != nil {
		return tx.rollback(err)
	}
	return tx.Commit()
}

func (ct RequeteContext) CreateLivraison(livraison models.Livraison) (models.Livraison, error) {
	contrainte := ContrainteLivraison{livraison}
	if err := contrainte.Check(); err != nil {
		return models.Livraison{}, err
	}

	if err := ct.checkLivraisonFournisseur(livraison); err != nil {
		return models.Livraison{}, err
	}

	livraison, err := livraison.Insert(ct.DB)
	return livraison, ErrorSQL(err)
}

func (ct RequeteContext) UpdateLivraison(livraison models.Livraison) (models.Livraison, error) {
	contrainte := ContrainteLivraison{livraison}
	if err := contrainte.Check(); err != nil {
		return models.Livraison{}, err
	}

	// on vérifie que les produits ayant cette contrainte sont tous du founisseur
	if err := ct.checkLivraisonFournisseur(livraison); err != nil {
		return models.Livraison{}, err
	}

	livraison, err := livraison.Update(ct.DB)
	return livraison, ErrorSQL(err)
}

// DeleteLivraison supprime la livraison
func (ct RequeteContext) DeleteLivraison(idLivraison int64) error {
	livraison, err := models.SelectLivraison(ct.DB, idLivraison)
	if err != nil {
		return ErrorSQL(err)
	}

	if err := ct.checkLivraisonFournisseur(livraison); err != nil {
		return err
	}

	tx, err := ct.beginTx()
	if err != nil {
		return err
	}

	// on modifie les produits concernés
	_, err = tx.Exec("UPDATE produits SET id_livraison = null WHERE id_livraison = $1", idLivraison)
	if err != nil {
		return tx.rollback(err)
	}

	_, err = models.DeleteLivraisonById(tx, livraison.Id)
	if err != nil {
		return tx.rollback(err)
	}
	return tx.Commit()
}
