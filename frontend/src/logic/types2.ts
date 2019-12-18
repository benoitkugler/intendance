import { Recette, Ingredient, Menu } from "./types";

export type New<T extends { id: number }> = Omit<T, "id"> &
  Partial<Pick<T, "id">>;

export type Ingredients = { [key: number]: Ingredient };
export type Recettes = { [key: number]: Recette };
export type Menus = { [key: number]: Menu };
