import {
  RepasComplet,
  MenuRecette,
  RepasRecette,
  MenuIngredient,
  RepasIngredient,
  Menu
} from "@/logic/types";
import { compareArrays } from "@/components/utils/utils";

export function toDateVuetify(d: Date) {
  return d.toISOString().substr(0, 10);
}

export function formatNbOffset(repas: RepasComplet) {
  const n = repas.offset_personnes;
  if (n != 0) {
    return `${n > 0 ? "+" : ""}${n} pers.`;
  }
  return "";
}

export function asRepasRecette(
  recette: MenuRecette,
  idRepas: number
): RepasRecette {
  return { id_repas: idRepas, id_recette: recette.id_recette };
}

type ingredient = Pick<
  MenuIngredient,
  "id_ingredient" | "quantite" | "cuisson"
>;

function extractIngredient(ingredient: ingredient): ingredient {
  return {
    id_ingredient: ingredient.id_ingredient,
    quantite: ingredient.quantite,
    cuisson: ingredient.cuisson
  };
}

export function asRepasIngredient(
  ingredient: MenuIngredient,
  idRepas: number
): RepasIngredient {
  return {
    id_repas: idRepas,
    ...extractIngredient(ingredient)
  };
}

// renvoie `true` si le menu et le repas ont le même contenu
export function compareRecettesIngredient(menu: Menu, repas: RepasComplet) {
  const mrs = (menu.recettes || []).map(r => r.id_recette);
  const rrs = (repas.recettes || []).map(r => r.id_recette);
  const mis = (menu.ingredients || []).map(extractIngredient);
  const ris = (repas.ingredients || []).map(extractIngredient);
  return compareArrays(mrs, rrs) && compareArrays(mis, ris);
}
