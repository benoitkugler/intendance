import { Ingredient, Menu, Recette, Sejour, Repas } from "./types";

export type New<T extends { id: number }> = Omit<T, "id"> &
  Partial<Pick<T, "id">>;

export type Ingredients = { [key: number]: Ingredient };
export type Recettes = { [key: number]: Recette };
export type Menus = { [key: number]: Menu };

export type DetailsSejour = Pick<Sejour, "nom" | "date_debut">;
export type DetailsRepas = Pick<
  Repas,
  "horaire" | "id_menu" | "nb_personnes" | "jour_offset"
>;

export interface PreferencesAgenda {
  restrictSejourCourant: boolean;
  startPremierJour: boolean;
}
