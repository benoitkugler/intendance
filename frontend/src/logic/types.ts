// github.com/benoitkugler/intendance/server/models.Unite
export type Unite = string

// github.com/benoitkugler/intendance/server/models.Conditionnement
export interface Conditionnement {
	quantite: number,
	unite: Unite,
}

// github.com/benoitkugler/intendance/server/models.Produit
export interface Produit {
	id: number,
	id_fournisseur: number,
	nom: string,
	conditionnement: Conditionnement,
	prix: number,
	reference_fournisseur: string,
	colisage: number,
}

// server/views.InAjouteIngredientProduit
export interface InAjouteIngredientProduit {
	id_ingredient: number,
	produit: Produit,
}

// github.com/benoitkugler/intendance/server/controller.OptionsAssistantCreateRepass
export interface OptionsAssistantCreateRepass {
	duree: number,
	with_cinquieme: boolean,
	with_gouter: boolean,
	delete_existing: boolean,
}

// github.com/benoitkugler/intendance/server/controller.InAssistantCreateRepass
export interface InAssistantCreateRepass {
	id_sejour: number,
	options: OptionsAssistantCreateRepass,
	groupes_sorties: { [key: number]: number[] | null } | null,
}

// server/views.InLieIngredientProduit
export interface InLieIngredientProduit {
	id_ingredient: number,
	id_produit: number,
}

// server/views.InLoggin
export interface InLoggin {
	mail: string,
	password: string,
}

// server/views.InResoudIngredients
export interface InResoudIngredients {
	mode: string,
	id_repas: number,
	nb_personnes: number,
	id_sejour: number,
	jour_offset: number[] | null,
}

// server/views.OutDeleteGroupe
export interface OutDeleteGroupe {
	token: string,
	nb_repas: number,
}

// github.com/benoitkugler/intendance/server/models.Groupe
export interface Groupe {
	id: number,
	id_sejour: number,
	nom: string,
	nb_personnes: number,
	couleur: string,
}

// server/views.OutGroupe
export interface OutGroupe {
	token: string,
	groupe: Groupe,
}

// github.com/benoitkugler/intendance/server/models.Categorie
export type Categorie = string

// github.com/benoitkugler/intendance/server/models.Callories
export interface Callories {
}

// github.com/benoitkugler/intendance/server/models.Ingredient
export interface Ingredient {
	id: number,
	nom: string,
	unite: Unite,
	categorie: Categorie,
	callories: Callories,
	conditionnement: Conditionnement,
}

// server/views.OutIngredient
export interface OutIngredient {
	token: string,
	ingredient: Ingredient,
}

// database/sql.NullInt64
export interface NullInt64 {
	Int64: number,
	Valid: boolean,
}

// github.com/benoitkugler/intendance/server/controller.IngredientProduits
export interface IngredientProduits {
	produits: Produit[] | null,
	id_default: NullInt64,
}

// server/views.OutIngredientProduits
export interface OutIngredientProduits {
	token: string,
	produits: IngredientProduits,
}

// github.com/benoitkugler/intendance/server/models.Ingredients
export type Ingredients = { [key: number]: Ingredient } | null

// server/views.OutIngredients
export interface OutIngredients {
	token: string,
	ingredients: Ingredients,
}

// github.com/benoitkugler/intendance/server/controller.Utilisateur
export interface Utilisateur {
	id: number,
	prenom_nom: string,
}

// github.com/benoitkugler/intendance/server/controller.OutLoggin
export interface OutLoggin {
	erreur: string,
	token: string,
	utilisateur: Utilisateur,
}

// github.com/benoitkugler/intendance/server/models.MenuRecette
export interface MenuRecette {
	id_menu: number,
	id_recette: number,
}

// github.com/benoitkugler/intendance/server/models.MenuIngredient
export interface MenuIngredient {
	id_menu: number,
	id_ingredient: number,
	quantite: number,
	cuisson: string,
}

// github.com/benoitkugler/intendance/server/controller.Menu
export interface Menu {
	id: number,
	id_proprietaire: NullInt64,
	commentaire: string,
	recettes: MenuRecette[] | null,
	ingredients: MenuIngredient[] | null,
}

// server/views.OutMenu
export interface OutMenu {
	token: string,
	menu: Menu,
}

// server/views.OutMenus
export interface OutMenus {
	token: string,
	menus: { [key: number]: Menu } | null,
}

// github.com/benoitkugler/intendance/server/models.RecetteIngredient
export interface RecetteIngredient {
	id_recette: number,
	id_ingredient: number,
	quantite: number,
	cuisson: string,
}

// github.com/benoitkugler/intendance/server/controller.Recette
export interface Recette {
	id: number,
	id_proprietaire: NullInt64,
	nom: string,
	mode_emploi: string,
	ingredients: RecetteIngredient[] | null,
}

// server/views.OutRecette
export interface OutRecette {
	token: string,
	recette: Recette,
}

// server/views.OutRecettes
export interface OutRecettes {
	token: string,
	recettes: { [key: number]: Recette } | null,
}

// 
export type Time = string

// github.com/benoitkugler/intendance/server/controller.IngredientQuantite
export interface IngredientQuantite {
	ingredient: Ingredient,
	quantite: number,
}

// github.com/benoitkugler/intendance/server/controller.DateIngredientQuantites
export interface DateIngredientQuantites {
	date: Time,
	ingredients: IngredientQuantite[] | null,
}

// server/views.OutResoudIngredients
export interface OutResoudIngredients {
	token: string,
	date_ingredients: DateIngredientQuantites[] | null,
}

// github.com/benoitkugler/intendance/server/models.Sejour
export interface Sejour {
	id: number,
	id_proprietaire: number,
	date_debut: Time,
	nom: string,
}

// server/views.OutSejour
export interface OutSejour {
	token: string,
	sejour: Sejour,
}

// github.com/benoitkugler/intendance/server/models.Horaire
export type Horaire = string

// github.com/benoitkugler/intendance/server/models.RepasGroupe
export interface RepasGroupe {
	id_repas: number,
	id_groupe: number,
}

// github.com/benoitkugler/intendance/server/controller.RepasWithGroupe
export interface RepasWithGroupe {
	id: number,
	id_sejour: number,
	id_menu: NullInt64,
	offset_personnes: number,
	jour_offset: number,
	horaire: Horaire,
	groupes: RepasGroupe[] | null,
}

// github.com/benoitkugler/intendance/server/controller.SejourRepas
export interface SejourRepas {
	id: number,
	id_proprietaire: number,
	date_debut: Time,
	nom: string,
	repass: RepasWithGroupe[] | null,
}

// github.com/benoitkugler/intendance/server/models.Groupes
export type Groupes = { [key: number]: Groupe } | null

// github.com/benoitkugler/intendance/server/controller.Sejours
export interface Sejours {
	sejours: { [key: number]: SejourRepas } | null,
	groupes: Groupes,
}

// server/views.OutSejours
export interface OutSejours {
	token: string,
	sejours: Sejours,
}

// server/views.OutUtilisateurs
export interface OutUtilisateurs {
	token: string,
	utilisateurs: { [key: number]: Utilisateur } | null,
}

