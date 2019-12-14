package models

import (
	"testing"

	"github.com/benoitkugler/intendance/logs"
)

func TestSql(t *testing.T) {
	db, err := ConnectDB(logs.DB_DEV)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	// tables primaires
	i1 := randUtilisateur()
	i1, err = queriesUtilisateur(tx, i1)
	if err != nil {
		t.Fatal(err)
	}

	i2 := randIngredient()
	i2, err = queriesIngredient(tx, i2)
	if err != nil {
		t.Fatal(err)
	}

	i3 := randRecette()
	i3.IdProprietaire = NullId(i1.Id)
	i3, err = queriesRecette(tx, i3)
	if err != nil {
		t.Fatal(err)
	}

	i4 := randMenu()
	i4.IdProprietaire = NullId(i1.Id)
	i4, err = queriesMenu(tx, i4)
	if err != nil {
		t.Fatal(err)
	}

	i5 := randSejour()
	i5.IdProprietaire = i1.Id
	i5, err = queriesSejour(tx, i5)
	if err != nil {
		t.Fatal(err)
	}

	i6 := randFournisseur()
	i6, err = queriesFournisseur(tx, i6)
	if err != nil {
		t.Fatal(err)
	}

	i7 := randProduit()
	i7.IdFournisseur = i6.Id
	i7, err = queriesProduit(tx, i7)
	if err != nil {
		t.Fatal(err)
	}

	i8 := randCommande()
	i8.IdProprietaire = i1.Id
	i8, err = queriesCommande(tx, i8)
	if err != nil {
		t.Fatal(err)
	}

	// tables de lien
	l1 := randRecetteIngredient()
	l1.IdIngredient = i2.Id
	l1.IdRecette = i3.Id
	l1, err = queriesRecetteIngredient(tx, l1)
	if err != nil {
		t.Fatal(err)
	}

	l2 := randMenuIngredient()
	l2.IdIngredient = i2.Id
	l2.IdMenu = i4.Id
	l2, err = queriesMenuIngredient(tx, l2)
	if err != nil {
		t.Fatal(err)
	}

	l3 := randMenuRecette()
	l3.IdRecette = i3.Id
	l3.IdMenu = i4.Id
	l3, err = queriesMenuRecette(tx, l3)
	if err != nil {
		t.Fatal(err)
	}

	l4 := randSejourMenu()
	l4.IdSejour = i5.Id
	l4.IdMenu = i4.Id
	l4, err = queriesSejourMenu(tx, l4)
	if err != nil {
		t.Fatal(err)
	}

	l5 := randIngredientProduit()
	l5.IdIngredient = i2.Id
	l5.IdProduit = i7.Id
	l5, err = queriesIngredientProduit(tx, l5)
	if err != nil {
		t.Fatal(err)
	}

	// suppressions des liens
	if err := l1.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if err := l2.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if err := l3.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if err := l4.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if err := l5.Delete(tx); err != nil {
		t.Fatal(err)
	}

	//suppressions des objets
	if _, err := i8.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if _, err := i7.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if _, err := i6.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if _, err := i5.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if _, err := i4.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if _, err := i3.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if _, err := i2.Delete(tx); err != nil {
		t.Fatal(err)
	}
	if _, err := i1.Delete(tx); err != nil {
		t.Fatal(err)
	}
}

func TestProduits(t *testing.T) {
	db, err := ConnectDB(logs.DB_DEV)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	fr := Fournisseur{Nom: "SuperU"}
	fr, err = fr.Insert(tx)
	if err != nil {
		t.Fatal(err)
	}

	ig := Ingredient{Nom: "Test"}
	ig, err = ig.Insert(tx)
	if err != nil {
		t.Fatal(err)
	}

	pr1 := Produit{Nom: "Prod1", IdFournisseur: fr.Id}
	pr1, err = pr1.Insert(tx)
	if err != nil {
		t.Fatal(err)
	}
	pr2 := Produit{Nom: "Prod2", IdFournisseur: fr.Id}
	pr2, err = pr2.Insert(tx)
	if err != nil {
		t.Fatal(err)
	}

	rows, err := tx.Query("SELECT id FROM produits")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = ScanInts(rows); err != nil {
		t.Fatal(err)
	}

	err = InsertManyIngredientProduits(tx, []IngredientProduit{
		{IdIngredient: ig.Id, IdProduit: pr1.Id},
		{IdIngredient: ig.Id, IdProduit: pr2.Id},
	})
	if err != nil {
		t.Fatal(err)
	}

	prods, err := ig.GetProduits(tx)
	if L := len(prods); L != 2 {
		t.Errorf("expected 2 produits, got %d", L)
	}
}
