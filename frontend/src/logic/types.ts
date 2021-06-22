import {
  Sejour,
  RepasComplet,
  Ingredient,
  LienIngredient,
  NullInt64,
  Livraison,
  New,
} from "./api";

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

export const NullId = (): NullInt64 => {
  return { Valid: false, Int64: 0 };
};
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

export function defaultLivraison(): New<Livraison> {
  return {
    nom: "",
    id_fournisseur: -1,
    jours_livraison: [true, true, true, true, true, false, false],
    delai_commande: 2,
    anticipation: 1,
  };
}

export interface EnumItem<T = string> {
  value: T;
  text: string;
}

export function enumStringToOptions<T extends string>(
  enums: { [key in T]: string }
) {
  const out: EnumItem<T>[] = [];

  for (const value in enums) {
    const text = enums[value];
    out.push({ value, text });
  }
  return out.sort((a, b) => {
    if (a.value == "") {
      return -1;
    }
    return a.text.localeCompare(b.text);
  });
}

export function enumIntToOptions<T extends number>(
  enums: { [key in T]: string }
) {
  const out: EnumItem<T>[] = [];

  for (const value in enums) {
    const text = enums[value];
    out.push({ value: Number(value) as T, text });
  }
  return out.sort((a, b) => {
    if (a.value == 0) {
      return -1;
    }
    return a.value - b.value;
  });
}
