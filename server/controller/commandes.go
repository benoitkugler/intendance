package controller

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/benoitkugler/intendance/server/models"
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

// Ambiguites indique les ingrédients pour lesquelles plusieurs produits sont disponibles
type Ambiguites map[int64][]models.Produit // id_ingredient -> produits

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

type cacheIngredientProduits struct {
	produits     models.Produits
	associations map[int64][]models.Produit // id ingredient -> produits
	defauts      map[int64][]models.Produit // id ingredient -> produits par défaut
}

// charge en (seulement) 3 requêtes les infos nécessaires
// à l'établissement d'une commande
// se restreint aux produits associés aux livraisons données
func (ct RequeteContext) resoudProduits(idsIngredients models.Ids, livraisons models.Livraisons) (cacheIngredientProduits, error) {
	out := cacheIngredientProduits{ // initialization des map
		produits:     make(models.Produits),
		associations: make(map[int64][]models.Produit),
		defauts:      make(map[int64][]models.Produit),
	}

	rows, err := ct.tx.Query(`SELECT ingredient_produits.* FROM ingredient_produits 
	JOIN produits ON ingredient_produits.id_produit = produits.id 
	WHERE ingredient_produits.id_ingredient = ANY($1)
	AND produits.id_livraison = ANY($2)`,
		idsIngredients.AsSQL(), livraisons.Ids().AsSQL())
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}
	ingredientsProduits, err := models.ScanIngredientProduits(rows)
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}

	rows, err = ct.tx.Query(`SELECT produits.* FROM produits 
	JOIN ingredient_produits ON ingredient_produits.id_produit = produits.id 
	WHERE ingredient_produits.id_ingredient = ANY($1)
	AND produits.id_livraison = ANY($2)`,
		idsIngredients.AsSQL(), livraisons.Ids().AsSQL())
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}
	out.produits, err = models.ScanProduits(rows)
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}

	// on regroupe les produits par ingrédients
	for _, ingProd := range ingredientsProduits {
		out.associations[ingProd.IdIngredient] = append(out.associations[ingProd.IdIngredient], out.produits[ingProd.IdProduit])
	}

	rows, err = ct.tx.Query("SELECT * FROM defaut_produits WHERE id_produit = ANY($1)", out.produits.Ids().AsSQL())
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}
	defauts, err := models.ScanDefautProduits(rows)
	if err != nil {
		return cacheIngredientProduits{}, ErrorSQL(err)
	}
	for _, def := range defauts {
		out.defauts[def.IdIngredient] = append(out.defauts[def.IdIngredient], out.produits[def.IdProduit])
	}
	return out, nil
}

type timedProduit struct {
	idProduit int64
	date      time.Time
}

type timedProduits map[timedProduit][]TimedIngredientQuantite

func newTimedProduits() timedProduits {
	return timedProduits(make(map[timedProduit][]TimedIngredientQuantite))
}

func (ts timedProduits) addIngredient(idProduit int64, jour time.Time, ing TimedIngredientQuantite) {
	// on normalise les dates
	jour = jour.Truncate(jourDuration)
	key := timedProduit{idProduit: idProduit, date: jour}
	ts[key] = append(ts[key], ing)
}

// ajuste la date de tous les produits au premier jour de commande
func (ts timedProduits) groupe() timedProduits {
	var first time.Time
	for key := range ts {
		if first.IsZero() || key.date.Before(first) {
			first = key.date
		}
	}
	// on regroupe toutes les demandes sur le même jour
	out := newTimedProduits()
	for key, v := range ts {
		out[timedProduit{idProduit: key.idProduit, date: first}] = append(out[timedProduit{idProduit: key.idProduit, date: first}], v...)
	}
	return out
}

// EtablitCommande calcule pour chaque ingrédient le jour de commande du produit
// et le nombre d'exemplaire.
func (s Server) EtablitCommande(ct RequeteContext, ingredients []DateIngredientQuantites, contraintes CommandeContraintes) ([]CommandeItem, Ambiguites, error) {
	// TODO: vérifier les associations ing -> produit,
	// où au moins les contraintes d'unité, etc..

	if err := ct.beginTx(s); err != nil {
		return nil, nil, err
	}
	defer ct.rollbackTx(nil) // pas de modifications sur les données

	fourns, err := ct.loadFournisseurs()
	if err != nil {
		return nil, nil, err
	}
	livraisons, err := ct.loadLivraisons(fourns)
	if err != nil {
		return nil, nil, err
	}

	allIngredients := models.Ingredients{}
	for _, iq := range ingredients {
		for _, ing := range iq.Ingredients {
			allIngredients[ing.Ingredient.Id] = ing.Ingredient
		}
	}

	// on récupére les données des produits
	data, err := ct.resoudProduits(allIngredients.Ids(), livraisons)
	if err != nil {
		return nil, nil, err
	}

	// on commence par associer à chaque ingrédient un produit
	// indépendement de la date d'utilisation
	ambiguites := make(Ambiguites)
	targetProduits := make(models.Produits)
	for idIngredient, ingredient := range allIngredients {
		var (
			ambs          []models.Produit
			targetProduit models.Produit
		)
		targetIdProduit, hasContrainte := contraintes.ContrainteProduits[idIngredient]
		if hasContrainte {
			var has bool
			targetProduit, has = data.produits[targetIdProduit]
			if !has { // le produit imposé n'est pas conforme
				return nil, nil, fmt.Errorf("Le produit (%d) n'est pas associé à l'ingrédient <b>%s</b> !",
					targetIdProduit, ingredient.Nom)
			}
		} else {
			// on récupère les produits associés
			prods := data.associations[idIngredient]

			switch {
			case len(prods) == 0: // l'absence de produit est fatale
				return nil, nil, fmt.Errorf("L'ingrédient %s n'est associé à aucun produit !", ingredient.Nom)
			case len(prods) > 1: // on essaye de résoudre les ambiguités
				defauts := data.defauts[idIngredient]
				switch {
				case len(defauts) == 0:
					// pas de valeur par défaut : on fait un choix arbitraire
					ambs = prods
					targetProduit = prods[0]
				case len(defauts) > 1:
					// on restreint l'ambiguité aux défauts et on fait un choix arbitraire
					ambs = defauts
					targetProduit = defauts[0]
				default:
					// OK : on utilise l'unique produit par défaut
					targetProduit = defauts[0]
				}
			default:
				// OK : on utilise le seul produit
				targetProduit = prods[0]
			}
		}
		sort.Slice(ambs, func(i, j int) bool {
			return ambs[i].Nom < ambs[j].Nom
		})
		targetProduits[idIngredient] = targetProduit
		if len(ambs) > 0 {
			ambiguites[idIngredient] = ambs
		}
	}

	// puis on utilise les produits trouvés pour
	// ajuster les dates de commandes

	// plusieurs ingrédients peuvent donner la même commande
	// au sens (produit, date)
	accu := newTimedProduits()

	for _, demande := range ingredients {
		for _, ing := range demande.Ingredients {
			targetProduit := targetProduits[ing.Ingredient.Id]
			livraison := livraisons[targetProduit.IdLivraison]

			dateCommande, _ := livraison.DateCommande(demande.Date)

			// on ajoute au timed-produit l'ingrédient et sa quantité
			// avec la date de demande
			ting := TimedIngredientQuantite{IngredientQuantite: ing, Date: demande.Date}
			accu.addIngredient(targetProduit.Id, dateCommande, ting)
		}
	}

	if contraintes.Regroupe {
		accu = accu.groupe()
	}

	var out []CommandeItem
	// toutes les demandes ont étés regroupées en produit,
	// on peut maitenant calculer le nombre de produit nécessaire
	for key, value := range accu {
		total, err := aggregeIngredients(value)
		if err != nil {
			return nil, nil, err
		}
		prod := data.produits[key.idProduit]
		if prod.Conditionnement.Quantite <= 0 {
			var chunks []string
			for _, ing := range value {
				chunks = append(chunks, "<b>"+ing.Ingredient.Nom+"</b>")
			}
			return nil, nil, fmt.Errorf(`Le conditionnement du produit <b>%s</b> est invalide : <i>%0.3f</i> <br/>
			Ingrédients liés : %s
			`, prod.Nom, prod.Conditionnement.Quantite, strings.Join(chunks, ", "))
		}
		colisage := prod.ColisageNeeded(total)
		out = append(out, CommandeItem{Produit: prod, JourCommande: key.date, Quantite: colisage, Origines: value})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Quantite > out[j].Quantite
	})

	return out, ambiguites, nil
}
