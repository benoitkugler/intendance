package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/benoitkugler/intendance/server/models"
)

func main() {
	path := "tarifs.csv"
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	allProduits := map[string]models.Ingredient{}
	for _, line := range lines {
		description := line[2]
		// on ne conserve que le premier mot de la description
		ingredient := strings.Title(strings.Split(description, " ")[0])
		unite := parseUnite(line[12])
		categorie := models.Categorie(strings.Title(strings.Split(line[1], " ")[0]))
		if ingredient != "" {
			allProduits[ingredient] = models.Ingredient{Nom: ingredient, Unite: unite, Categorie: categorie}
		}

	}
	query := generateSQL(allProduits)

	err = ioutil.WriteFile("ingredients.sql", []byte(query), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func parseUnite(unite string) models.Unite {
	switch strings.ToLower(strings.TrimSpace(unite)) {
	case "l":
		return models.Litres
	case "kg":
		return models.Kilos
	default:
		return models.Piece
	}
}

func generateSQL(ingredients map[string]models.Ingredient) string {
	var values []string
	for _, ing := range ingredients {
		nom := strings.ReplaceAll(ing.Nom, "'", "''")
		value := fmt.Sprintf("('%s', '%s', '%s', '{}', '{}')", nom, string(ing.Unite), ing.Categorie)
		values = append(values, value)
	}
	query := "INSERT INTO ingredients (nom, unite, categorie, callories, conditionnement) VALUES " + strings.Join(values, ", ") + ";"
	return query
}
