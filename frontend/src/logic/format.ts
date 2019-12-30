import { Repas, Menu, Recette } from "./types";
import { D } from "./controller";
import { G } from "./getters";

export function formatRepasName(r: Repas) {
  const menu = D.menus[r.id_menu];
  const menuName = menu ? formatMenuName(menu) : "";
  return ` (${r.nb_personnes}) - ${menuName}`;
}

export function formatMenuName(menu: Menu) {
  const recs = G.getMenuRecettes(menu);
  if (recs.length == 0) return `(${menu.id})`;
  return recs.map(rec => rec.nom || "").join(", ");
}

export function formatMenuOrRecetteProprietaire(item: Menu | Recette) {
  const prop = G.getMenuOrRecetteProprietaire(item);
  if (prop == null) {
    return "public";
  } else if (prop.id == D.idUtilisateur) {
    return "vous appartient";
  }
  return `appartient à ${prop.prenom_nom}`;
}
