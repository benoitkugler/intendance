import { Repas, Menu } from "./types";
import { D } from "./controller";
import { G } from "./getters";

export function formatRepasName(r: Repas) {
  const menu = D.menus[r.id_menu];
  if (!menu) return "";
  return formatMenuName(menu);
}

export function formatMenuName(menu: Menu) {
  const recs = G.getMenuRecettes(menu);
  if (recs.length == 0) return `(${menu.id})`;
  return recs.map(rec => rec.nom || "").join(", ");
}
