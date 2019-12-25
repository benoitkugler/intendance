import {
  Ingredient,
  Menu,
  Recette,
  Sejour,
  Repas,
  RecetteIngredient,
  MenuIngredient,
  Utilisateur
} from "./types";

export type New<T extends { id: number }> = Omit<T, "id"> &
  Partial<Pick<T, "id">>;

export type Ingredients = { [key: number]: Ingredient };
export type Recettes = { [key: number]: Recette };
export type Menus = { [key: number]: Menu };
export type Utilisateurs = { [key: number]: Utilisateur };

export type DetailsSejour = Pick<Sejour, "nom" | "date_debut">;
export type DetailsRepas = Pick<
  Repas,
  "horaire" | "id_menu" | "nb_personnes" | "jour_offset"
>;

export interface IngredientOptions {
  ingredient: Ingredient;
  options?: RecetteIngredient | MenuIngredient;
}

export interface PreferencesAgenda {
  restrictSejourCourant: boolean;
  startPremierJour: boolean;
}

export type EditMode = "new" | "edit";
