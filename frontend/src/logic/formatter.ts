import { Controller } from "./controller";
import {
  Menu,
  Recette,
  RepasComplet,
  Time,
  Conditionnement,
  MenuComplet
} from "./types";
import { fmtUnite } from "./enums";

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

  formatRepasName = (r: RepasComplet) => {
    return `${(r.recettes || []).length} rec. - ${
      (r.ingredients || []).length
    } ing.`;
  };

  formatMenuName = (menu: MenuComplet) => {
    const recs = this.controller.getMenuRecettes(menu);
    if (recs.length == 0) return `(${menu.id})`;
    const nbIngs = (menu.ingredients || []).length;
    let out = recs.map(rec => rec.nom || "").join(", ");
    if (nbIngs > 0) {
      out += ` - ${nbIngs} ing.`;
    }
    return out;
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

  static formatConditionnement(c: Conditionnement) {
    return Formatter.formatQuantite(c.quantite) + " " + fmtUnite(c.unite);
  }
}
