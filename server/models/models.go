// Définit la structure de la base de données.
package models

import (
	"database/sql"
	"time"
)

// Le logiciel est disponible sous la forme d'une **application web** accessible par mail/password.
// sql: ADD UNIQUE(mail)
type Utilisateur struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
	Mail     string `json:"mail"`

	PrenomNom string `json:"prenom_nom"`
}

// Les ingrédients sont _partagés_ : n'importe quel intendant peut utiliser un ingrédient déjà défini, et ajouter un produit lié. En revanche, les recettes et menus ne sont modifiables que par son _propriétaire_ (mais copiables libremement). En cas de modification par le propriétaire, les intendants utilisant la ressource sont notifiés par mail et peuvent choisir d'accepter la modification ou de s'approprier la ressource en la copiant.
// Les recettes et menus peuvent n'être liés à aucun propriétaire, et sont alors éditable par tout le monde.
//
// sql: ADD UNIQUE(nom)
type Ingredient struct {
	Id    int64  `json:"id"`
	Nom   string `json:"nom"`
	Unite Unite  `json:"unite"`

	Categorie Categorie `json:"categorie"`
	Callories Callories `json:"callories"`
	// optionnel, zero signifie pas de contrainte
	Conditionnement Conditionnement `json:"conditionnement,omitempty"`
}

type LienIngredient struct {
	IdIngredient int64   `json:"id_ingredient"`
	Quantite     float64 `json:"quantite"`
	Cuisson      string  `json:"cuisson"`
}

type Recette struct {
	Id            int64         `json:"id"`
	IdUtilisateur sql.NullInt64 `json:"id_utilisateur"`
	Nom           string        `json:"nom"`

	ModeEmploi string `json:"mode_emploi"`
}

// sql: ADD UNIQUE(id_recette, id_ingredient)
type RecetteIngredient struct {
	IdRecette int64 `json:"id_recette"`
	LienIngredient
}

// Menu définit un raccourci pour organiser recettes et ingrédients
// Voir `Repas` pour un repas effectif.
type Menu struct {
	Id            int64         `json:"id"`
	IdUtilisateur sql.NullInt64 `json:"id_utilisateur"`

	Commentaire string `json:"commentaire"`
}

// sql: ADD UNIQUE(id_menu, id_ingredient)
type MenuIngredient struct {
	IdMenu int64 `json:"id_menu"`
	LienIngredient
}

// sql: ADD UNIQUE(id_menu, id_recette)
type MenuRecette struct {
	IdMenu    int64 `json:"id_menu"`
	IdRecette int64 `json:"id_recette"`
}

// Les séjours sont _privés_, mais les journées formées peuvent être copiées.
// sql: ADD UNIQUE(id, id_utilisateur)
type Sejour struct {
	Id            int64 `json:"id"`
	IdUtilisateur int64 `json:"id_utilisateur"`

	// Fixe l'origine du séjour.
	// Une journée est déterminé par un "offset"
	// relatif à cette date.
	DateDebut time.Time `json:"date_debut"`
	Nom       string    `json:"nom"`
}

// Groupe est un groupe de personnes lié à un séjour
type Groupe struct {
	Id          int64  `json:"id"`
	IdSejour    int64  `json:"id_sejour"`
	Nom         string `json:"nom"`
	NbPersonnes int64  `json:"nb_personnes"`
	Couleur     string `json:"couleur"`
}

// Le concept de journée nécessite d'être lié à la donnée du nombre de personnes pour chaque menu. Cela ne colle pas bien avec un schéma SQL classique. De plus, une journée n'a pas vraiment d'intérêt à être partagée : la modification sur une journée entrainerait celle sur une autre, ce qui est serait plutôt déroutant.
// On propose donc de ne pas utiliser de table "journée", mais de construire (dynamiquement) les journées à partir de la table _repas_ (voir ci dessous). En revanche, le concept de journée sera bien présent pour l'utilisateur, pour organiser son emploi du temps, ou pour copier des journées déjà existantes.
//
// Repas représente un repas effectif, lié à un séjour.
// Il est constitué de recettes et d'ingrédients (de la même manière qu'un menu)
type Repas struct {
	Id              int64   `json:"id"`
	IdSejour        int64   `json:"id_sejour"`
	OffsetPersonnes int64   `json:"offset_personnes"`
	JourOffset      int64   `json:"jour_offset"`
	Horaire         Horaire `json:"horaire"`
	Anticipation    int64   `json:"anticipation"` // commande les ingrédients en avance (en jours)
}

// sql: ADD UNIQUE(id_repas, id_ingredient)
type RepasIngredient struct {
	IdRepas int64 `json:"id_repas" sql_foreign:"CASCADE"`
	LienIngredient
}

// sql: ADD UNIQUE(id_repas, id_recette)
type RepasRecette struct {
	IdRepas   int64 `json:"id_repas" sql_foreign:"CASCADE"`
	IdRecette int64 `json:"id_recette"`
}

// sql: ADD UNIQUE(id_repas, id_groupe)
type RepasGroupe struct {
	IdRepas  int64 `json:"id_repas" sql_foreign:"CASCADE"`
	IdGroupe int64 `json:"id_groupe" sql_foreign:"CASCADE"`
}

// Fournisseur définit un fournisseur.
// Chaque fournisseur possède au moins une contrainte de livraison
// (voir `Livraison`), et peut en posséder plusieurs.
// sql: ADD UNIQUE(nom)
type Fournisseur struct {
	Id   int64  `json:"id"`
	Nom  string `json:"nom"`
	Lieu string `json:"lieu"`
}

// Enregistre les fournisseurs associés à chaque utilisateur
// sql: ADD UNIQUE(id_utilisateur,id_fournisseur)
type UtilisateurFournisseur struct {
	IdUtilisateur int64 `json:"id_utilisateur"`
	IdFournisseur int64 `json:"id_fournisseur"`
}

// Enregistre les fournisseurs associés à un séjour.
// Note: Le champ `IdUtilisateur` permet d'assurer par une contrainte
// que les fournisseurs soient associés à l'utilisateur du séjour.
//
// sql: ADD UNIQUE(id_sejour,id_fournisseur)
// sql: ADD FOREIGN KEY (id_utilisateur, id_sejour) REFERENCES sejours (id_utilisateur, id)
// sql: ADD FOREIGN KEY (id_utilisateur, id_fournisseur) REFERENCES utilisateur_fournisseurs (id_utilisateur, id_fournisseur)
type SejourFournisseur struct {
	IdUtilisateur int64 `json:"id_utilisateur,omitempty"`
	IdSejour      int64 `json:"id_sejour,omitempty"`
	IdFournisseur int64 `json:"id_fournisseur,omitempty"`
}

// sql: ADD CHECK(prix >= 0)
// sql: ADD UNIQUE(id_livraison, nom)
type Produit struct {
	Id          int64 `json:"id"`
	IdLivraison int64 `json:"id_livraison"`

	Nom             string          `json:"nom"`
	Conditionnement Conditionnement `json:"conditionnement"`
	Prix            float64         `json:"prix"`

	ReferenceFournisseur string `json:"reference_fournisseur"`
	// zero signifie pas de contrainte
	Colisage int64 `json:"colisage"`
}

// Livraison enregistre les contraintes d'un fournisseur
// quant à la livraison d'une gamme de produit.
//
// sql: ADD CHECK(anticipation >= 0)
// sql: ADD CHECK(delai_commande >= 0)
// sql: ADD UNIQUE(id_fournisseur, nom)
type Livraison struct {
	Id            int64 `json:"id"`
	IdFournisseur int64 `json:"id_fournisseur" sql_foreign:"CASCADE"`

	Nom            string         `json:"nom"`
	JoursLivraison JoursLivraison `json:"jours_livraison"` // jours possibles de livraison
	DelaiCommande  int64          `json:"delai_commande"`  // nombre de jours à anticiper par rapport au jour de livraison
	Anticipation   int64          `json:"anticipation"`    // nombre de jours entre la livraison et l'utilisation (défaut : 1)
}

// sql: ADD UNIQUE(id_ingredient, id_produit)
type IngredientProduit struct {
	IdIngredient int64 `json:"id_ingredient"`
	IdProduit    int64 `json:"id_produit" sql_foreign:"CASCADE"`

	IdUtilisateur int64 `json:"id_utilisateur"` // ajouteur
}

// sql: ADD UNIQUE(id_utilisateur, id_ingredient, id_fournisseur)
// sql: ADD FOREIGN KEY (id_utilisateur, id_fournisseur) REFERENCES utilisateur_fournisseurs (id_utilisateur, id_fournisseur)
// sql: ADD FOREIGN KEY (id_ingredient, id_produit) REFERENCES ingredient_produits (id_ingredient, id_produit)
type DefautProduit struct {
	IdUtilisateur int64 `json:"id_utilisateur"`
	IdIngredient  int64 `json:"id_ingredient"`
	IdFournisseur int64 `json:"id_fournisseur"`
	IdProduit     int64 `json:"id_produit" sql_foreign:"CASCADE"`
}

type Commande struct {
	Id            int64     `json:"id"`
	IdUtilisateur int64     `json:"id_utilisateur"`
	DateEmission  time.Time `json:"date_emission"`
	Tag           string    `json:"tag"`
}

// sql: ADD UNIQUE(id_commande, id_produit)
type CommandeProduit struct {
	IdCommande int64 `json:"id_commande"`
	IdProduit  int64 `json:"id_produit"`
	Quantite   int64 `json:"quantite"`
}
