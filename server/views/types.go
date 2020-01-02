package views

import (
	"github.com/benoitkugler/intendance/server/controller"
	"github.com/benoitkugler/intendance/server/models"
)

type OutAgenda struct {
	Token  string                       `json:"token"`
	Agenda controller.AgendaUtilisateur `json:"agenda"`
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
