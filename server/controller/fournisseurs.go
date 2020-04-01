package controller

import (
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
	livraison, err := models.SelectLivraison(ct.tx, produit.IdLivraison)
	if err != nil {
		return livraison, err
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
func (s Server) LoadFournisseurs(ct RequeteContext) (models.Fournisseurs, models.Livraisons, error) {
	if err := ct.beginTx(s); err != nil {
		return nil, nil, err
	}
	defer ct.rollbackTx(nil) // pas de modifications
	fournisseurs, err := ct.loadFournisseurs()
	if err != nil {
		return nil, nil, err
	}

	// on sélectionne les livraisons liées aux fournisseurs et les livraisons universelles
	rows, err := ct.tx.Query("SELECT * FROM livraisons WHERE id_fournisseur = ANY($1) OR id_fournisseur IS null",
		fournisseurs.Ids().AsSQL())
	if err != nil {
		return nil, nil, ErrorSQL(err)
	}
	livraisons, err := models.ScanLivraisons(rows)
	if err != nil {
		return nil, nil, ErrorSQL(err)
	}
	return fournisseurs, livraisons, nil
}

// CreateFournisseur crée un fournisseur et le lie à l'utilisateur courant
// Une contrainte de livraison "standard" est automatiquement créée
func (s Server) CreateFournisseur(ct RequeteContext, fournisseur models.Fournisseur) (out models.Fournisseur, err error) {
	if err := ct.beginTx(s); err != nil {
		return out, err
	}

	out, err = fournisseur.Insert(ct.tx)
	if err != nil {
		return out, ct.rollbackTx(err)
	}

	err = models.InsertManyUtilisateurFournisseurs(ct.tx, []models.UtilisateurFournisseur{
		{IdFournisseur: out.Id, IdUtilisateur: ct.idProprietaire},
	})
	if err != nil {
		return out, ct.rollbackTx(err)
	}

	// ajout d'une contraint de livraison par défaut
	livraison := models.Livraison{
		IdFournisseur:  out.Id,
		Nom:            "",
		JoursLivraison: models.JoursLivraison{true, true, true, true, true, false, false},
		DelaiCommande:  2,
		Anticipation:   1,
	}
	livraison, err = livraison.Insert(ct.tx)
	if err != nil {
		return out, ct.rollbackTx(err)
	}

	return out, ct.commitTx()
}

func (s Server) UpdateFournisseur(ct RequeteContext, fournisseur models.Fournisseur) (models.Fournisseur, error) {
	if err := ct.beginTx(s); err != nil {
		return models.Fournisseur{}, err
	}
	hasF, err := ct.hasFournisseur(fournisseur.Id)
	if err != nil {
		return models.Fournisseur{}, ct.rollbackTx(err)
	}
	if !hasF {
		ct.rollbackTx(err)
		return models.Fournisseur{}, fmt.Errorf("Le fournisseur %s ne fait pas partie de vos fournisseurs.", fournisseur.Nom)
	}
	fournisseur, err = fournisseur.Update(ct.tx)
	if err != nil {
		return models.Fournisseur{}, ErrorSQL(err)
	}
	return fournisseur, ct.commitTx()
}

func (s Server) DeleteFournisseur(ct RequeteContext, idFournisseur int64) error {
	if err := ct.beginTx(s); err != nil {
		return err
	}
	hasF, err := ct.hasFournisseur(idFournisseur)
	if err != nil {
		return ct.rollbackTx(err)
	}
	if !hasF {
		ct.rollbackTx(nil)
		return fmt.Errorf("Le fournisseur (%d) ne fait pas partie de vos fournisseurs.", idFournisseur)
	}

	rows, err := ct.tx.Query(`SELECT  * FROM commande_produits 
		JOIN produits ON produits.id = commande_produits.id_produit 
		JOIN livraisons ON produits.id_livraison = livraisons.id
		WHERE livraisons.id_fournisseur = $1`, idFournisseur)
	if err != nil {
		return ErrorSQL(err)
	}
	cps, err := models.ScanCommandeProduits(rows)
	if err != nil {
		return ErrorSQL(err)
	}
	if L := len(cps); L > 0 {
		ct.rollbackTx(nil)
		return fmt.Errorf("%d produit(s) lié(s) au fournisseur sont déjà utilisés dans une commande.", L)
	}

	// sejours
	_, err = ct.tx.Exec("DELETE FROM sejour_fournisseurs WHERE id_fournisseur = $1", idFournisseur)
	if err != nil {
		return ct.rollbackTx(err)
	}

	// ingredients
	_, err = ct.tx.Exec(`DELETE FROM ingredient_produits USING produits, livraisons
		WHERE ingredient_produits.id_produit = produits.id 
		AND produits.id_livraison = livraisons.id
		AND livraisons.id_fournisseur = $1`, idFournisseur)
	if err != nil {
		return ct.rollbackTx(err)
	}

	// defaut
	_, err = ct.tx.Exec(`DELETE FROM defaut_produits USING produits, livraisons
		WHERE defaut_produits.id_produit = produits.id 
		AND produits.id_livraison = livraisons.id
		AND livraisons.id_fournisseur = $1`, idFournisseur)
	if err != nil {
		return ct.rollbackTx(err)
	}

	// produits
	_, err = ct.tx.Exec(`DELETE FROM produits USING livraisons
		WHERE produits.id_livraison = livraisons.id AND livraisons.id_fournisseur = $1`, idFournisseur)
	if err != nil {
		return ct.rollbackTx(err)
	}

	// livraisons
	_, err = ct.tx.Exec("DELETE FROM livraisons WHERE id_fournisseur = $1", idFournisseur)
	if err != nil {
		return ct.rollbackTx(err)
	}
	// utilisateurs
	_, err = ct.tx.Exec("DELETE FROM utilisateur_fournisseurs WHERE id_fournisseur = $1", idFournisseur)
	if err != nil {
		return ct.rollbackTx(err)
	}

	_, err = models.Fournisseur{Id: idFournisseur}.Delete(ct.tx)
	if err != nil {
		return ct.rollbackTx(err)
	}
	return ct.commitTx()
}

func (s Server) UpdateSejourFournisseurs(ct RequeteContext, idSejour int64, idsFournisseurs []int64) error {
	if err := ct.beginTx(s); err != nil {
		return err
	}
	if err := ct.proprioSejour(models.Sejour{Id: idSejour}, false); err != nil {
		return err
	}

	// reset les fournisseurs du séjour ...
	_, err := ct.tx.Exec("DELETE FROM sejour_fournisseurs WHERE id_sejour = $1", idSejour)
	if err != nil {
		return ct.rollbackTx(err)
	}
	sf := make([]models.SejourFournisseur, len(idsFournisseurs))
	for i, id := range idsFournisseurs {
		sf[i] = models.SejourFournisseur{IdSejour: idSejour, IdFournisseur: id}
	}
	// ... et rajoute les nouveaux
	if err := models.InsertManySejourFournisseurs(ct.tx, sf); err != nil {
		return ct.rollbackTx(err)
	}
	return ct.commitTx()
}

func (s Server) CreateLivraison(ct RequeteContext, livraison models.Livraison) (models.Livraison, error) {
	if err := ct.beginTx(s); err != nil {
		return models.Livraison{}, err
	}

	if err := ct.checkLivraisonFournisseur(livraison); err != nil {
		return models.Livraison{}, err
	}

	livraison, err := livraison.Insert(ct.tx)
	if err != nil {
		return models.Livraison{}, ErrorSQL(err)
	}
	return livraison, ct.commitTx()
}

func (s Server) UpdateLivraison(ct RequeteContext, livraison models.Livraison) (models.Livraison, error) {
	if err := ct.beginTx(s); err != nil {
		return models.Livraison{}, err
	}

	if err := ct.checkLivraisonFournisseur(livraison); err != nil {
		return models.Livraison{}, err
	}

	// on vérifie que les produits ayant cette contrainte sont tous du founisseur

	livraison, err := livraison.Update(ct.tx)
	if err != nil {
		return models.Livraison{}, ErrorSQL(err)
	}
	return livraison, ct.commitTx()
}

// DeleteLivraison supprime la livraison
func (s Server) DeleteLivraison(ct RequeteContext, idLivraison int64) error {
	if err := ct.beginTx(s); err != nil {
		return err
	}

	livraison, err := models.SelectLivraison(ct.tx, idLivraison)
	if err != nil {
		return ct.rollbackTx(err)
	}

	if err := ct.checkLivraisonFournisseur(livraison); err != nil {
		ct.rollbackTx(nil)
		return err
	}

	// on modifie les produits concernés
	_, err = ct.tx.Exec("UPDATE produits SET id_livraison = null WHERE id_livraison = $1", idLivraison)
	if err != nil {
		return ct.rollbackTx(err)
	}

	_, err = livraison.Delete(ct.tx)
	if err != nil {
		return ct.rollbackTx(err)
	}
	return ct.commitTx()
}
