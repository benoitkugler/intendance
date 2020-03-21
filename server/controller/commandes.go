package controller

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/benoitkugler/intendance/server/models"
	"github.com/lib/pq"
)

// Ce fichier implémente la génération d'une commande à
// partir d'une liste d'ingrédients.
// La commande peut prendre en compte plusieurs contraintes (ingrédients, coût, etc...)

// const (
// 	MeilleurPrix         CommandeOptimisation = "prix"
// 	PlusFaibleGaspillage CommandeOptimisation = "gaspillage"
// )

// type CommandeOptimisation string

// type ContrainteProduit struct {
// 	Date         time.Time `json:"date"`
// 	IdIngredient int64     `json:"id_ingredient"`
// 	IdProduit    int64     `json:"id_produit"`
// }

// TODO: ajouter un paramètre de regroupement des jours. Ex :
// un même produit nécessaire sur 2 jour serait commandé groupé pour le premier

const jourDuration = 24 * time.Hour

type CommandeContraintes struct {
	// Force l'utilisation du produit pour l'ingrédient (idIngredient -> idProduit)
	ContrainteProduits map[int64]int64 `json:"contrainte_produits"`

	// Si `true`, regroupe toutes les commandes
	// à la date courante (prototype)
	Regroupe bool `json:"regroupe"`
}

// TimedIngredientQuantite ajoute la date de demande de l'ingrédient
type TimedIngredientQuantite struct {
	IngredientQuantite
	Date time.Time `json:"date"`
}

// CommandeItem représente la commande d'un produit.
type CommandeItem struct {
	Produit models.Produit `json:"produit"`

	// jour conseillé de commande, prenant en compte les délais de livraison
	JourCommande time.Time `json:"jour_commande"`

	Quantite int64 `json:"quantite"`

	// ingrédients liés à ce produit
	Origines []TimedIngredientQuantite `json:"origines"`
}

// renvoie la quantité équivalente à la somme
// des ingrédients contenus dans `l`
// les ingrédients doivent avoir tous la même unité
func aggregeIngredients(l []TimedIngredientQuantite) (float64, error) {
	contrainte := ContrainteListeIngredients{ingredients: l}
	if err := contrainte.Check(); err != nil {
		return 0, nil
	}
	total := 0.
	for _, iq := range l {
		total += iq.Quantite
	}
	return total, nil
}

type timedProduit struct {
	idProduit int64
	date      time.Time
}

type timedProduits struct {
	data map[timedProduit][]TimedIngredientQuantite
	ids  models.Set
}

func newTimedProduits() timedProduits {
	return timedProduits{data: make(map[timedProduit][]TimedIngredientQuantite), ids: models.NewSet()}
}

func (ts timedProduits) addIngredient(idProduit int64, jour time.Time, ing TimedIngredientQuantite) {
	// on normalise les dates
	jour = jour.Truncate(jourDuration)
	key := timedProduit{idProduit: idProduit, date: jour}
	ts.data[key] = append(ts.data[key], ing)
	ts.ids.Add(idProduit)
}

// EtablitCommande calcule pour chaque ingrédient le jour de commande du produit
// et le nombre d'exemplaire.
func (s Server) EtablitCommande(ct RequeteContext, ingredients []DateIngredientQuantites, contraintes CommandeContraintes) ([]CommandeItem, error) {
	// TODO: vérifier les associations ing -> produit,
	// où au moins les contraintes d'unité, etc..
	if err := ct.beginTx(s); err != nil {
		return nil, err
	}
	defer ct.commitTx() // pas de modifications sur les données

	fourns, err := ct.loadFournisseurs()
	if err != nil {
		return nil, err
	}

	// plusieurs ingrédients peuvent donner la même commande
	// au sens (produit, date)
	accu := newTimedProduits()

	for _, demande := range ingredients {
		// la commande doit toujours arriver un jour avant son utilisation
		dateArrivee := demande.Date.Add(-jourDuration)

		for _, ing := range demande.Ingredients {
			targetIdProduit, hasContrainte := contraintes.ContrainteProduits[ing.Ingredient.Id]
			if !hasContrainte {
				// TODO: optimiser
				prods, err := ing.Ingredient.GetProduits(ct.tx, fourns)
				if err != nil {
					return nil, err
				}
				if len(prods) == 0 {
					return nil, fmt.Errorf("L'ingrédient %s n'est associé à aucun produit !", ing.Ingredient.Nom)
				}
				for idProd := range prods {
					targetIdProduit = idProd
					break // on prend arbitrairement le premier
				}
			}

			// TODO: prendre en compte le délai du fournisseur
			// pour l'instant on commande le même jour
			dateCommande := dateArrivee
			if contraintes.Regroupe {
				dateCommande = time.Now()
			}

			// on ajoute au timed-produit l'ingrédient et sa quantité
			// avec la date de demande
			ting := TimedIngredientQuantite{IngredientQuantite: ing, Date: demande.Date}
			accu.addIngredient(targetIdProduit, dateCommande, ting)
		}
	}

	// on récupére les données des produits
	rows, err := ct.tx.Query("SELECT * FROM produits WHERE id = ANY($1)", pq.Int64Array(accu.ids.Keys()))
	if err != nil {
		return nil, ErrorSQL(err)
	}
	produits, err := models.ScanProduits(rows)
	if err != nil {
		return nil, ErrorSQL(err)
	}

	var out []CommandeItem
	// toutes les demandes ont étés regroupées en produit,
	// on peut maitenant calculer le nombre de produit nécessaire
	for key, value := range accu.data {
		total, err := aggregeIngredients(value)
		if err != nil {
			return nil, err
		}
		prod := produits[key.idProduit]
		if prod.Conditionnement.Quantite <= 0 {
			var chunks []string
			for _, ing := range value {
				chunks = append(chunks, "<b>"+ing.Ingredient.Nom+"</b>")
			}
			return nil, fmt.Errorf(`Le conditionnement du produit <b>%s</b> est invalide : <i>%0.3f</i> <br/>
			Ingrédients liés : %s
			`, prod.Nom, prod.Conditionnement.Quantite, strings.Join(chunks, ", "))
		}
		colisage := prod.ColisageNeeded(total)
		out = append(out, CommandeItem{Produit: prod, JourCommande: key.date, Quantite: colisage, Origines: value})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Quantite > out[j].Quantite
	})

	return out, nil
}
