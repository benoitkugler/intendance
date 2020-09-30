package controller

import (
	"time"

	"github.com/benoitkugler/intendance/server/models"
)

type InLoggin struct {
	Mail     string `json:"mail,omitempty"`
	Password string `json:"password,omitempty"`
}

type OutDeleteGroupe struct {
	Id      int64 `json:"id"`
	NbRepas int64 `json:"nb_repas"`
}

type InResoudIngredients struct {
	Mode string `json:"mode"` // "repas" ou "journees"

	IdRepas int64 `json:"id_repas"` // pour Mode == "repas"
	// pour Mode == "repas" seulement
	// donner -1 pour utiliser le nombre de personnes actuel
	NbPersonnes int64 `json:"nb_personnes"`

	IdSejour    int64   `json:"id_sejour"`   // pour Mode == "journees"
	JourOffsets []int64 `json:"jour_offset"` // pour Mode == "journees". Passer nil pour tout le sejour
}

// Fournisseurs et produits

type InSejourFournisseurs struct {
	IdSejour        int64   `json:"id_sejour"`
	IdsFournisseurs []int64 `json:"ids_fournisseurs"`
}

type OutFournisseurs struct {
	Fournisseurs models.Fournisseurs `json:"fournisseurs"`
	Livraisons   models.Livraisons   `json:"livraisons"`
}

type InLieIngredientProduit struct {
	IdIngredient int64 `json:"id_ingredient,omitempty"`
	IdProduit    int64 `json:"id_produit,omitempty"`
}

type InAjouteIngredientProduit struct {
	IdIngredient int64          `json:"id_ingredient,omitempty"`
	Produit      models.Produit `json:"produit,omitempty"`
}

type InSetDefautProduit struct {
	IdIngredient int64 `json:"id_ingredient"`
	IdProduit    int64 `json:"id_produit"`
	On           bool  `json:"on"`
}

type InCommandeComplete struct {
	Ingredients []DateIngredientQuantites `json:"ingredients"`
	Contraintes CommandeContraintes       `json:"contraintes"`
}

type OutCommandeComplete struct {
	Commande []CommandeCompleteItem `json:"commande"`
}

type InCommandeSimple struct {
	Ingredients []DateIngredientQuantites `json:"ingredients"`
	Contraintes CommandeContraintes       `json:"contraintes"`
}

type OutCommandeSimple struct {
	Commande []CommandeSimpleItem `json:"commande"`
}

type RecetteComplet struct {
	models.Recette
	Ingredients models.LienIngredients `json:"ingredients"`
}

type MenuComplet struct {
	models.Menu

	Recettes    models.Ids             `json:"recettes"`
	Ingredients models.LienIngredients `json:"ingredients"`
}

type RepasComplet struct {
	models.Repas

	Groupes     []models.RepasGroupe   `json:"groupes"`
	Recettes    models.Ids             `json:"recettes"`
	Ingredients models.LienIngredients `json:"ingredients"`
}
type SejourRepas struct {
	models.Sejour
	Fournisseurs []models.SejourFournisseur `json:"fournisseurs"`
	Repass       []RepasComplet             `json:"repass"`
}

// Sejours contient les séjours, ainsi que les groupes et repas associés.
type Sejours struct {
	Sejours map[int64]SejourRepas `json:"sejours"`
	Groupes models.Groupes        `json:"groupes"`
}

type Utilisateur struct {
	Id        int64  `json:"id"`
	PrenomNom string `json:"prenom_nom"`
}

type OutLoggin struct {
	Erreur      string      `json:"erreur"`
	Token       string      `json:"token"`
	Utilisateur Utilisateur `json:"utilisateur"`
	Expires     int         `json:"expires"` // en jours
}

type IngredientQuantite struct {
	Ingredient models.Ingredient `json:"ingredient"`
	Quantite   float64           `json:"quantite"`
}

type DateIngredientQuantites struct {
	Date        time.Time            `json:"date"`
	Ingredients []IngredientQuantite `json:"ingredients"`
}

type IngredientProduits struct {
	Produits []models.Produit `json:"produits"`
	Defaults models.Set       `json:"defaults"` // id_produit -> is default
}

type PreviewCommande struct {
	Produits []CommandeCompleteItem
}

type Ingredients = models.Ingredients

type InAssistantCreateRepass struct {
	IdSejour       int64                        `json:"id_sejour"`
	Options        OptionsAssistantCreateRepass `json:"options"`
	GroupesSorties map[int][]int64              `json:"groupes_sorties"` // offset -> ids_groupes
}
