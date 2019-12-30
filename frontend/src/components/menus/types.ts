import { Menu, Recette, Ingredient } from "@/logic/types";

import { IngredientOptions, New } from "@/logic/types2";

type ModeMenu = "visu" | "editMenu" | "editRecette" | "editIngredient";
interface SelectionMenu {
  menu: Menu | null;
  recette: Recette | null;
  ingredient: IngredientOptions | null;
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

export const DefautRecette: New<Recette> = {
  id_proprietaire: { Valid: true, Int64: -1 },
  ingredients: [],
  mode_emploi: "",
  nom: ""
};

export const DefautMenu: New<Menu> = {
  commentaire: "",
  id_proprietaire: { Valid: true, Int64: -1 },
  ingredients: [],
  recettes: []
};
