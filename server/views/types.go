package views

import (
	"github.com/benoitkugler/intendance/server/controller"
	"github.com/benoitkugler/intendance/server/models"
)

type InLoggin struct {
	Mail     string `json:"mail,omitempty"`
	Password string `json:"password,omitempty"`
}

type OutSejours struct {
	Token   string             `json:"token"`
	Sejours controller.Sejours `json:"sejours"`
}

type OutIngredient struct {
	Token      string            `json:"token"`
	Ingredient models.Ingredient `json:"ingredient"`
}

type OutIngredients struct {
	Token       string             `json:"token"`
	Ingredients models.Ingredients `json:"ingredients"`
}

type OutRecette struct {
	Token   string             `json:"token"`
	Recette controller.Recette `json:"recette"`
}

type OutRecettes struct {
	Token    string                        `json:"token"`
	Recettes map[int64]*controller.Recette `json:"recettes"`
}

type OutMenu struct {
	Token string          `json:"token"`
	Menu  controller.Menu `json:"menu"`
}

type OutMenus struct {
	Token string                     `json:"token"`
	Menus map[int64]*controller.Menu `json:"menus"`
}

type OutSejour struct {
	Token  string        `json:"token"`
	Sejour models.Sejour `json:"sejour"`
}

type OutGroupe struct {
	Token  string        `json:"token"`
	Groupe models.Groupe `json:"groupe"`
}

type OutDeleteGroupe struct {
	Token   string `json:"token"`
	NbRepas int    `json:"nb_repas"`
}

// to be converted by struct2ts
type InAssistantCreateRepass = controller.InAssistantCreateRepass

type OutUtilisateurs struct {
	Token        string                           `json:"token"`
	Utilisateurs map[int64]controller.Utilisateur `json:"utilisateurs"`
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

type OutResoudIngredients struct {
	Token           string                               `json:"token"`
	DateIngredients []controller.DateIngredientQuantites `json:"date_ingredients"`
}

type InLieIngredientProduit struct {
	IdIngredient int64 `json:"id_ingredient,omitempty"`
	IdProduit    int64 `json:"id_produit,omitempty"`
}

type InAjouteIngredientProduit struct {
	IdIngredient int64          `json:"id_ingredient,omitempty"`
	Produit      models.Produit `json:"produit,omitempty"`
}

type OutIngredientProduits struct {
	Token    string                        `json:"token"`
	Produits controller.IngredientProduits `json:"produits"`
}
