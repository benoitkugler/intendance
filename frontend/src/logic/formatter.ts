import { Controller } from "./controller";
import {
  RepasComplet,
  MenuComplet,
  Menu,
  Recette,
  Time,
  Conditionnement,
  UniteLabels,
  Produit,
  Livraison,
} from "./api";

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
  "Décembre",
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
    let out = recs.map((rec) => rec.nom || "").join(", ");
    if (nbIngs > 0) {
      out += ` - ${nbIngs} ing.`;
    }
    return out;
  };

  formatMenuOrRecetteProprietaire = (item: Menu | Recette) => {
    const prop = this.controller.getMenuOrRecetteProprietaire(item);
    if (prop == null) {
      return "public";
    } else if (prop.id == this.controller.state.idUtilisateur) {
      return "vous appartient";
    }
    return `appartient à ${prop.prenom_nom}`;
  };

  static formatDate(dateString: Time) {
    const date = new Date(dateString);
    if (isNaN(date.valueOf())) return "";
    return date.toLocaleDateString("fr-FR", {
      weekday: "short",
      day: "numeric",
      month: "short",
    });
  }

  static formatQuantite(v: number) {
    return v.toFixed(2);
  }

  static formatConditionnement(c: Conditionnement) {
    return Formatter.formatQuantite(c.quantite) + " " + UniteLabels[c.unite];
  }

  formatFournisseur(produit: Produit) {
    const livraison = this.controller.getLivraison(produit);
    return this.formatLivraison(livraison);
  }

  formatLivraison(livraison: Livraison) {
    const fournisseur = this.controller.getFournisseur(
      livraison.id_fournisseur
    );
    let fourstring = fournisseur.nom;
    if (livraison.nom) {
      fourstring += ` <i>(${livraison.nom})</i>`;
    }
    return fourstring;
  }

  formatProduit = (produit: Produit) => {
    const fourstring = this.formatFournisseur(produit);
    return `${fourstring} - ${produit.nom}`;
  };
}
