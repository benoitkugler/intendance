import { Repas, Menu } from "./types";
import { D } from "./controller";

export function formatRepasName(r: Repas) {
  const menu = D.menus[r.id_menu];
  if (!menu) return "";
  return formatMenuName(menu);
}

export function formatMenuName(menu: Menu) {
  if (!menu.recettes) return `(${menu.id})`;
  return menu.recettes
    .map(rec => D.recettes[rec.id_recette]?.nom || "")
    .join(", ");
}
