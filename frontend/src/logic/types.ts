// this file was automatically generated, DO NOT EDIT
// structs

// types
export type Ingredients = { [key: number]: Ingredient };

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

// struct2ts:database/sql.NullInt64
export interface NullInt64 {
  Int64: number;
  Valid: boolean;
}

// struct2ts:github.com/benoitkugler/intendance/server/models.Recette
export interface Recette {
  id: number;
  id_proprietaire: NullInt64;
  nom: string;
  mode_emploi: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/models.Menu
export interface Menu {
  id: number;
  id_proprietaire: NullInt64;
  commentaire: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Items
export interface Items {
  ingredients: { [key: number]: Ingredient };
  recettes: { [key: number]: Recette };
  menus: { [key: number]: Menu };
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutAgenda
export interface OutAgenda {
  token: string;
  agenda: AgendaUtilisateur;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutCreateIngredient
export interface OutCreateIngredient {
  token: string;
  ingredient: Ingredient;
}

// struct2ts:github.com/benoitkugler/intendance/server/views.OutUpdateIngredient
export interface OutUpdateIngredient {
  token: string;
  ingredients: { [key: number]: Ingredient };
}
