import { Ingredient, Menu, Recette, Sejour } from "./types";

export type New<T extends { id: number }> = Omit<T, "id"> &
  Partial<Pick<T, "id">>;

export type Ingredients = { [key: number]: Ingredient | undefined };
export type Recettes = { [key: number]: Recette | undefined };
export type Menus = { [key: number]: Menu | undefined };

export type DetailsSejour = Pick<Sejour, "nom" | "date_debut">;
