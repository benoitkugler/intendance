package models

import (
	"database/sql"
	"math/rand"
	"time"
)

// utilitaires pour données aléatoires

var (
	letterRunes  = []rune("azertyuiopqsdfghjklmwxcvbn123456789")
	specialRunes = []rune("é@!?&èïab ")
)

func randstringExt(n int, specialChars bool) string {
	b := make([]rune, n)
	props, maxLength := letterRunes, len(letterRunes)
	if specialChars {
		props = append(props, specialRunes...)
		maxLength += len(specialRunes)
	}
	for i := range b {
		b[i] = props[rand.Intn(maxLength)]
	}
	return string(b)
}

func randstring() string {
	return randstringExt(50, true)
}

func randBool() bool {
	i := rand.Int31n(2)
	return i == 1
}

func randfloat64() float64 {
	return rand.Float64() * float64(rand.Int31())
}

func randNullInt64() sql.NullInt64 {
	return sql.NullInt64{
		Int64: rand.Int63(),
		Valid: randBool(),
	}
}

func randNullString() sql.NullString {
	return sql.NullString{
		String: randstring(),
		Valid:  randBool(),
	}
}

// TODO: update
func randCallories() Callories {
	return Callories{}
}

// TODO: update
func randCategorie() Categorie {
	return ""
}

func randTime() time.Time {
	return time.Unix(int64(rand.Int31()), 5)
}

func randUnite() Unite {
	i := rand.Intn(3)
	return [3]Unite{Litres, Kilos, Piece}[i]
}

func randJoursLivraison() JoursLivraison {
	var out JoursLivraison
	for i := range out {
		out[i] = randBool()
	}
	return out
}

func randConditionnement() Conditionnement {
	return Conditionnement{
		Quantite: randfloat64(),
		Unite:    randUnite(),
	}
}

func randHoraire() Horaire {
	t := randTime()
	return Horaire{
		Heure:  t.Hour(),
		Minute: t.Minute(),
	}
}
