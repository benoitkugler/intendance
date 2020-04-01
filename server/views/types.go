package views

import (
	"github.com/benoitkugler/intendance/server/controller"
	"github.com/benoitkugler/intendance/server/models"
)

type InLoggin struct {
	Mail     string `json:"mail,omitempty"`
	Password string `json:"password,omitempty"`
}

type OutLoggin = controller.OutLoggin

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
	Token   string                    `json:"token"`
	Recette controller.RecetteComplet `json:"recette"`
}

type OutRecettes struct {
	Token    string                               `json:"token"`
	Recettes map[int64]*controller.RecetteComplet `json:"recettes"`
}

type OutMenu struct {
	Token string                 `json:"token"`
	Menu  controller.MenuComplet `json:"menu"`
}

type OutMenus struct {
	Token string                            `json:"token"`
	Menus map[int64]*controller.MenuComplet `json:"menus"`
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

// Fournisseurs et produits

type InSejourFournisseurs struct {
	IdSejour        int64   `json:"id_sejour"`
	IdsFournisseurs []int64 `json:"ids_fournisseurs"`
}

type OutFournisseur struct {
	Token       string             `json:"token"`
	Fournisseur models.Fournisseur `json:"fournisseur"`
}

type OutFournisseurs struct {
	Token        string              `json:"token"`
	Fournisseurs models.Fournisseurs `json:"fournisseurs"`
	Livraisons   models.Livraisons   `json:"livraisons"`
}

type OutLivraison struct {
	Token     string           `json:"token"`
	Livraison models.Livraison `json:"livraison"`
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

type OutProduit struct {
	Token   string         `json:"token"`
	Produit models.Produit `json:"produit"`
}

type InSetDefautProduit struct {
	IdIngredient int64 `json:"id_ingredient"`
	IdProduit    int64 `json:"id_produit"`
	On           bool  `json:"on"`
}

type InCommande struct {
	Ingredients []controller.DateIngredientQuantites `json:"ingredients"`
	Contraintes controller.CommandeContraintes       `json:"contraintes"`
}

type OutCommande struct {
	Token    string                    `json:"token"`
	Commande []controller.CommandeItem `json:"commande"`
}
