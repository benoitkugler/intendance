import {
  Menu,
  Recette,
  Ingredient,
  RecetteComplet,
  MenuComplet
} from "@/logic/types";

import { IngredientOptions, New } from "@/logic/types2";

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
  unite: "",
  categorie: "",
  conditionnement: { unite: "", quantite: 0 },
  callories: {}
};

export const DefautRecette: New<RecetteComplet> = {
  id_utilisateur: { Valid: true, Int64: -1 },
  ingredients: [],
  mode_emploi: "",
  nom: ""
};

export const DefautMenu: New<MenuComplet> = {
  commentaire: "",
  id_utilisateur: { Valid: true, Int64: -1 },
  ingredients: [],
  recettes: []
};
