import { Controller } from "./controller";
import { Repas, Menu, Recette, Horaire } from "./types";

export class Formatter {
  private controller: Controller;

  constructor(controller: Controller) {
    this.controller = controller;
  }

  formatRepasName = (r: Repas) => {
    const menu = this.controller.data.menus[r.id_menu];
    const menuName = menu ? this.formatMenuName(menu) : "";
    return ` (${r.nb_personnes}) - ${menuName}`;
  };

  formatMenuName = (menu: Menu) => {
    const recs = this.controller.getMenuRecettes(menu);
    if (recs.length == 0) return `(${menu.id})`;
    return recs.map(rec => rec.nom || "").join(", ");
  };

  formatMenuOrRecetteProprietaire = (item: Menu | Recette) => {
    const prop = this.controller.getMenuOrRecetteProprietaire(item);
    if (prop == null) {
      return "public";
    } else if (prop.id == this.controller.idUtilisateur) {
      return "vous appartient";
    }
    return `appartient à ${prop.prenom_nom}`;
  };

  static horaireToTime(horaire: Horaire) {
    return (
      ("00" + horaire.heure).substr(-2, 2) +
      ":" +
      ("00" + horaire.minute).substr(-2, 2)
    );
  }
}
