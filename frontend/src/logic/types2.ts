import {
  Ingredient,
  Sejour,
  RepasComplet,
  NullInt64,
  LienIngredient
} from "./types";

export type New<T extends { id: number }> = Omit<T, "id"> &
  Partial<Pick<T, "id">>;

export type DetailsSejour = Omit<Sejour, "id" | "id_utilisateur">;
export type DetailsRepas = Omit<RepasComplet, "id" | "id_sejour">;

export interface IngredientOptions {
  ingredient: Ingredient;
  options?: LienIngredient;
}

export interface PreferencesAgenda {
  startPremierJour: boolean;
}

export type EditMode = "new" | "edit";

export type ViewMode = "month" | "day";

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
