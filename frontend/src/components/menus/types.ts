import { Menu, Recette } from "@/logic/types";

import { IngredientOptions } from "@/logic/types2";

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
