// this file was automatically generated, DO NOT EDIT
// structs

// types

// struct2ts:github.com/benoitkugler/intendance/server/models.Horaire
export interface Horaire {
  heure: number;
  minute: number;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Repas
export interface Repas {
  id_menu: number;
  nb_personnes: number;
  horaire: Horaire;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Journee
export interface Journee {
  jour_offset: number;
  menus: Repas[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Sejour
export interface Sejour {
  id: number;
  id_proprietaire: number;
  date_debut: Date;
  nom: string;
  journees: { [key: number]: Journee };
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.AgendaUtilisateur
export interface AgendaUtilisateur {
  sejours: Sejour[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutAgenda
export interface OutAgenda {
  token: string;
  agenda: AgendaUtilisateur;
}

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
