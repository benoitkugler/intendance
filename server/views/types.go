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

type OutDeleteGroupe struct {
	NbRepas int `json:"nb_repas"`
}

// to be converted by structgen
type InAssistantCreateRepass = controller.InAssistantCreateRepass

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

type InCommande struct {
	Ingredients []controller.DateIngredientQuantites `json:"ingredients"`
	Contraintes controller.CommandeContraintes       `json:"contraintes"`
}

type OutCommande struct {
	Commande   []controller.CommandeItem `json:"commande"`
	Ambiguites controller.Ambiguites     `json:"ambiguites"`
}
