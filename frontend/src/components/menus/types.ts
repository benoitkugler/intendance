import {
  Ingredient,
  RecetteComplet,
  MenuComplet,
  Unite,
  New,
} from "@/logic/api";

type ModeMenu = "visu" | "editMenu" | "editRecette" | "editIngredient";
export interface SelectionMenu {
  idMenu: number | null;
  idRecette: number | null;
  idIngredient: number | null;
}
export interface StateMenus {
  mode: ModeMenu;
  selection: SelectionMenu;
}

export const DefautIngredient: New<Ingredient> = {
  nom: "",
  unite: Unite.Litres,
  categorie: "",
  conditionnement: { unite: Unite.Litres, quantite: 0 },
  callories: {},
};

export const DefautRecette: New<RecetteComplet> = {
  id_utilisateur: { Valid: true, Int64: -1 },
  ingredients: [],
  mode_emploi: "",
  nom: "",
};

export const DefautMenu: New<MenuComplet> = {
  commentaire: "",
  id_utilisateur: { Valid: true, Int64: -1 },
  ingredients: [],
  recettes: [],
};
