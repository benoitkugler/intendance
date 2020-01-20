import { Controller } from "./controller";
import { Menu, Recette, RepasWithGroupe } from "./types";

const Months = [
  "Janvier",
  "Février",
  "Mars",
  "Avril",
  "Mai",
  "Juin",
  "Juillet",
  "Août",
  "Septembre",
  "Octobre",
  "Novembre",
  "Décembre"
];

export class Formatter {
  private controller: Controller;

  constructor(controller: Controller) {
    this.controller = controller;
  }

  formatRepasName = (r: RepasWithGroupe) => {
    let menuName = "";
    if (!r.id_menu.Valid) {
      menuName = "-";
    } else {
      const menu = this.controller.data.menus[r.id_menu.Int64];
      menuName = menu ? this.formatMenuName(menu) : "";
    }
    return menuName;
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

  static formatDate(dateString: Time) {
    dateString = dateString || "";
    if (dateString.length < 10 || dateString.substr(0, 10) == "0001-01-01") {
      return null;
    }
    const year = dateString.substr(0, 4);
    const month = Number(dateString.substr(5, 2));
    const day = dateString.substr(8, 2);
    return `${day} ${Months[month - 1]} ${year}`;
  }

  static formatQuantite(v: number) {
    return v.toFixed(2);
  }
}
