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
	i3.IdProprietaire = i1.Id
	i3, err = queriesRecette(tx, i3)
	if err != nil {
		t.Fatal(err)
	}

	i4 := randMenu()
	i4.IdProprietaire = i1.Id
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
