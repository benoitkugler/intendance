// this file was automatically generated, DO NOT EDIT
// structs
// struct2ts:github.com/benoitkugler/intendance/server/models.Callories
export interface Callories {}

// struct2ts:github.com/benoitkugler/intendance/server/models.Conditionnement
export interface Conditionnement {
  quantite: number;
  unite: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/models.Ingredient
export interface Ingredient {
  id: number;
  nom: string;
  unite: string;
  categorie: string;
  callories: Callories;
  conditionnement: Conditionnement;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutIngredient
export interface OutIngredient {
  token: string;
  ingredient: Ingredient;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutIngredients
export interface OutIngredients {
  token: string;
  ingredients: { [key: number]: Ingredient };
}

// struct2ts:database/sql.NullInt64
export interface NullInt64 {
  Int64: number;
  Valid: boolean;
}

// struct2ts:github.com/benoitkugler/intendance/server/models.RecetteIngredient
export interface RecetteIngredient {
  id_recette: number;
  id_ingredient: number;
  quantite: number;
  cuisson: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Recette
export interface Recette {
  id: number;
  id_proprietaire: NullInt64;
  nom: string;
  mode_emploi: string;
  ingredients: RecetteIngredient[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutRecette
export interface OutRecette {
  token: string;
  recette: Recette;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutRecettes
export interface OutRecettes {
  token: string;
  recettes: { [key: number]: Recette };
}

// struct2ts:github.com/benoitkugler/intendance/server/models.MenuRecette
export interface MenuRecette {
  id_menu: number;
  id_recette: number;
}

// struct2ts:github.com/benoitkugler/intendance/server/models.MenuIngredient
export interface MenuIngredient {
  id_menu: number;
  id_ingredient: number;
  quantite: number;
  cuisson: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Menu
export interface Menu {
  id: number;
  id_proprietaire: NullInt64;
  commentaire: string;
  recettes: MenuRecette[] | null;
  ingredients: MenuIngredient[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutMenu
export interface OutMenu {
  token: string;
  menu: Menu;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutMenus
export interface OutMenus {
  token: string;
  menus: { [key: number]: Menu };
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Utilisateur
export interface Utilisateur {
  id: number;
  prenom_nom: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutUtilisateurs
export interface OutUtilisateurs {
  token: string;
  utilisateurs: { [key: number]: Utilisateur };
}

// struct2ts:github.com/benoitkugler/intendance/server/models.Sejour
export interface Sejour {
  id: number;
  id_proprietaire: number;
  date_debut: Time;
  nom: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutSejour
export interface OutSejour {
  token: string;
  sejour: Sejour;
}

// struct2ts:github.com/benoitkugler/intendance/server/models.Groupe
export interface Groupe {
  id: number;
  id_sejour: number;
  nom: string;
  nb_personnes: number;
  couleur: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutGroupe
export interface OutGroupe {
  token: string;
  groupe: Groupe;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutDeleteGroupe
export interface OutDeleteGroupe {
  token: string;
  nb_repas: number;
}

// struct2ts:github.com/benoitkugler/intendance/server/models.Horaire
export interface Horaire {
  heure: number;
  minute: number;
}

// struct2ts:github.com/benoitkugler/intendance/server/models.RepasGroupe
export interface RepasGroupe {
  id_repas: number;
  id_groupe: number;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.RepasWithGroupe
export interface RepasWithGroupe {
  id: number;
  id_sejour: number;
  id_menu: number;
  offset_personnes: number;
  jour_offset: number;
  horaire: Horaire;
  groupes: RepasGroupe[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.SejourRepas
export interface SejourRepas {
  id: number;
  id_proprietaire: number;
  date_debut: Time;
  nom: string;
  repass: RepasWithGroupe[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Sejours
export interface Sejours {
  sejours: { [key: number]: SejourRepas };
  groupes: { [key: number]: Groupe };
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutSejours
export interface OutSejours {
  token: string;
  sejours: Sejours;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.InResoudIngredients
export interface InResoudIngredients {
  mode: string;
  id_repas: number;
  nb_personnes: number;
  id_sejour: number;
  jour_offset: number[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.IngredientQuantite
export interface IngredientQuantite {
  ingredient: Ingredient;
  quantite: number;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.DateIngredientQuantites
export interface DateIngredientQuantites {
  date: Time;
  ingredients: IngredientQuantite[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutResoudIngredients
export interface OutResoudIngredients {
  token: string;
  date_ingredients: DateIngredientQuantites[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.OutLoggin
export interface OutLoggin {
  erreur: string;
  token: string;
  utilisateur: Utilisateur;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.InLoggin
export interface InLoggin {
  mail: string;
  password: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/models.Produit
export interface Produit {
  id: number;
  id_fournisseur: number;
  nom: string;
  conditionnement: Conditionnement;
  prix: number;
  reference_fournisseur: string;
  colisage: number;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.IngredientProduits
export interface IngredientProduits {
  produits: Produit[] | null;
  id_default: NullInt64;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutIngredientProduits
export interface OutIngredientProduits {
  token: string;
  produits: IngredientProduits;
}

// types
