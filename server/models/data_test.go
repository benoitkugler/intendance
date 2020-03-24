package models

import (
	"database/sql"
	"math/rand"
	"time"
)

func randint64() int64 {
	return int64(rand.Intn(1000000))
}

func randtTime() time.Time {
	return time.Unix(int64(rand.Int31()), 5)
}

var letterRunes2 = []rune("azertyuiopqsdfghjklmwxcvbn123456789é@!?&èïab ")

func randstring() string {
	b := make([]rune, 50)
	maxLength := len(letterRunes2)
	for i := range b {
		b[i] = letterRunes2[rand.Intn(maxLength)]
	}
	return string(b)
}

func randCommande() Commande {
	return Commande{
		Id:             randint64(),
		IdProprietaire: randint64(),
		DateEmission:   randtTime(),
		Tag:            randstring(),
	}
}

func randCommandeProduit() CommandeProduit {
	return CommandeProduit{
		IdCommande: randint64(),
		IdProduit:  randint64(),
		Quantite:   randint64(),
	}
}

func randbool() bool {
	i := rand.Int31n(2)
	return i == 1
}

func randArray7bool() [7]bool {
	var out [7]bool
	for i := range out {
		out[i] = randbool()
	}
	return out
}

func randJoursLivraison() JoursLivraison {
	return JoursLivraison(randArray7bool())
}

func randFournisseur() Fournisseur {
	return Fournisseur{
		Id:             randint64(),
		Nom:            randstring(),
		DelaiCommande:  randint64(),
		JoursLivraison: randJoursLivraison(),
	}
}

func randGroupe() Groupe {
	return Groupe{
		Id:          randint64(),
		IdSejour:    randint64(),
		Nom:         randstring(),
		NbPersonnes: randint64(),
		Couleur:     randstring(),
	}
}

func randUnite() Unite {
	return Unite(randstring())
}

func randCategorie() Categorie {
	return Categorie(randstring())
}

func randCallories() Callories {
	return Callories{}
}

func randfloat64() float64 {
	return rand.Float64() * float64(rand.Int31())
}

func randConditionnement() Conditionnement {
	return Conditionnement{
		Quantite: randfloat64(),
		Unite:    randUnite(),
	}
}

func randIngredient() Ingredient {
	return Ingredient{
		Id:              randint64(),
		Nom:             randstring(),
		Unite:           randUnite(),
		Categorie:       randCategorie(),
		Callories:       randCallories(),
		Conditionnement: randConditionnement(),
	}
}

func randIngredientProduit() IngredientProduit {
	return IngredientProduit{
		IdIngredient: randint64(),
		IdProduit:    randint64(),
		IdAjouteur:   randint64(),
	}
}

func randsqlNullInt64() sql.NullInt64 {
	return sql.NullInt64{
		Int64: randint64(),
		Valid: randbool(),
	}
}

func randMenu() Menu {
	return Menu{
		Id:             randint64(),
		IdProprietaire: randsqlNullInt64(),
		Commentaire:    randstring(),
	}
}

func randMenuIngredient() MenuIngredient {
	return MenuIngredient{
		IdMenu:       randint64(),
		IdIngredient: randint64(),
		Quantite:     randfloat64(),
		Cuisson:      randstring(),
	}
}

func randMenuRecette() MenuRecette {
	return MenuRecette{
		IdMenu:    randint64(),
		IdRecette: randint64(),
	}
}

func randProduit() Produit {
	return Produit{
		Id:                   randint64(),
		IdFournisseur:        randint64(),
		Nom:                  randstring(),
		Conditionnement:      randConditionnement(),
		Prix:                 randfloat64(),
		ReferenceFournisseur: randstring(),
		Colisage:             randint64(),
	}
}

func randRecette() Recette {
	return Recette{
		Id:             randint64(),
		IdProprietaire: randsqlNullInt64(),
		Nom:            randstring(),
		ModeEmploi:     randstring(),
	}
}

func randRecetteIngredient() RecetteIngredient {
	return RecetteIngredient{
		IdRecette:    randint64(),
		IdIngredient: randint64(),
		Quantite:     randfloat64(),
		Cuisson:      randstring(),
	}
}

func randHoraire() Horaire {
	return Horaire(randstring())
}

func randRepas() Repas {
	return Repas{
		Id:              randint64(),
		IdSejour:        randint64(),
		IdMenu:          randsqlNullInt64(),
		OffsetPersonnes: randint64(),
		JourOffset:      randint64(),
		Horaire:         randHoraire(),
	}
}

func randRepasGroupe() RepasGroupe {
	return RepasGroupe{
		IdRepas:  randint64(),
		IdGroupe: randint64(),
	}
}

func randSejour() Sejour {
	return Sejour{
		Id:             randint64(),
		IdProprietaire: randint64(),
		DateDebut:      randtTime(),
		Nom:            randstring(),
	}
}

func randUtilisateur() Utilisateur {
	return Utilisateur{
		Id:        randint64(),
		Password:  randstring(),
		Mail:      randstring(),
		PrenomNom: randstring(),
	}
}
