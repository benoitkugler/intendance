import {
  Ingredient,
  Menu,
  Recette,
  Sejour,
  RecetteIngredient,
  MenuIngredient,
  Utilisateur,
  Groupe,
  RepasComplet,
  NullInt64
} from "./types";

export type New<T extends { id: number }> = Omit<T, "id"> &
  Partial<Pick<T, "id">>;

export type Recettes = { [key: number]: Recette };
export type Menus = { [key: number]: Menu };
export type Utilisateurs = { [key: number]: Utilisateur };
export type Groupes = { [key: number]: Groupe };

export type DetailsSejour = Pick<Sejour, "nom" | "date_debut">;
export type DetailsRepas = Pick<
  RepasComplet,
  | "horaire"
  | "offset_personnes"
  | "jour_offset"
  | "groupes"
  | "recettes"
  | "ingredients"
>;

export interface IngredientOptions {
  ingredient: Ingredient;
  options?: RecetteIngredient | MenuIngredient;
}

export interface PreferencesAgenda {
  startPremierJour: boolean;
}

export type EditMode = "new" | "edit";

export type CalendarMode = "groupes" | "menus";

export const NullId: NullInt64 = { Valid: false, Int64: 0 };
export function toNullableId(id: number): NullInt64 {
  return { Valid: true, Int64: id };
}

export function deepcopy<T>(v: T): T {
  return JSON.parse(JSON.stringify(v));
}

// compare object as JSON
export function deepequal<T>(v1: T, v2: T): boolean {
  return JSON.stringify(v1) === JSON.stringify(v2);
}
