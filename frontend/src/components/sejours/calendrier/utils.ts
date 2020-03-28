import { RepasComplet, MenuComplet, LienIngredient } from "@/logic/types";
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

function extractIngredient(ingredient: LienIngredient): LienIngredient {
  return {
    id_ingredient: ingredient.id_ingredient,
    quantite: ingredient.quantite,
    cuisson: ingredient.cuisson
  };
}

// renvoie `true` si le menu et le repas ont le même contenu
export function compareRecettesIngredient(
  menu: MenuComplet,
  repas: RepasComplet
) {
  const mis = (menu.ingredients || []).map(extractIngredient);
  const ris = (repas.ingredients || []).map(extractIngredient);
  return (
    compareArrays(menu.recettes || [], repas.recettes || []) &&
    compareArrays(mis, ris)
  );
}
