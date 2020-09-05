import {
  Menu,
  Recette,
  Sejour,
  RepasComplet,
  Groupe,
  RepasGroupe,
  Produit,
  MenuComplet,
  Livraison,
  Sejours,
  Ingredients,
  RecetteComplet,
  Utilisateur,
  Fournisseurs,
  Livraisons
} from "./types";
import { Notifications } from "./notifications";
import { IngredientOptions, New, DetailsRepas } from "./types2";
import { Formatter } from "./formatter";
import { searchFunction } from "@/components/utils/utils";
import { Calculs } from "./calculs";
import { API, Meta } from "./server";
import { Data } from "./data";

/**  Object principal de stockage des données
 * sur le client.
 * Une instance de cet objet est créé au chargement,
 * puis partagée entre les différents composants.
 * Le système de réactivité de vuejs permet de propager
 * facilement les changements effectués aux données.
 */
export class Controller {
  readonly data: Data;
  readonly calculs: Calculs;
  readonly formatter: Formatter;

  private readonly api: API;

  readonly state: {
    idUtilisateur: number;
    idSejour: number | null;
  };

  constructor(meta: Meta, public readonly notifications: Notifications) {
    this.calculs = new Calculs(this);
    this.formatter = new Formatter(this);

    this.api = new API(this.notifications, meta.token);
    this.state = { idUtilisateur: meta.idUtilisateur, idSejour: null };

    this.data = new Data(this.api);
  }

  getAllIngredients(): IngredientOptions[] {
    return Object.values(this.data.ingredients || {}).map(ing => {
      return { ingredient: ing };
    });
  }

  getMenu(idMenu: number) {
    return (this.data.menus || {})[idMenu];
  }

  getRecette(idRecette: number) {
    return (this.data.recettes || {})[idRecette];
  }

  getIngredient(idIngredient: number) {
    return (this.data.ingredients || {})[idIngredient];
  }

  getMenuRecettes(menu: New<MenuComplet>) {
    return (menu.recettes || []).map(id => this.data.recettes[id]);
  }

  getMenuIngredients(idMenu: number): IngredientOptions[] {
    return (this.getMenu(idMenu).ingredients || []).map(ing => {
      return {
        ingredient: (this.data.ingredients || {})[ing.id_ingredient],
        options: ing
      };
    });
  }

  getMenuOrRecetteProprietaire(item: Menu | Recette) {
    if (!item.id_utilisateur.Valid) return null;
    return this.data.utilisateurs[item.id_utilisateur.Int64];
  }

  getRecetteIngredients(idRecette: number): IngredientOptions[] {
    return (this.getRecette(idRecette).ingredients || []).map(ing => {
      return {
        ingredient: (this.data.ingredients || {})[ing.id_ingredient],
        options: ing
      };
    });
  }

  getFournisseur(idFournisseur: number) {
    return (this.data.fournisseurs || {})[idFournisseur];
  }

  getLivraison(produit: Produit): Livraison {
    return (this.data.livraisons || {})[produit.id_livraison];
  }

  searchMenu(search: string) {
    const menus = Object.values(this.data.menus);
    const predicat = searchFunction(search);
    // on cherche dans le nom du menu, composé des recettes
    return menus.filter(menu => predicat(this.formatter.formatMenuName(menu)));
  }

  offsetToDate(idSejour: number, offset: number) {
    const sejour = (this.data.sejours.sejours || {})[idSejour];
    const dateDebut = new Date(sejour.date_debut);
    dateDebut.setDate(dateDebut.getDate() + offset);
    return dateDebut;
  }

  dateToOffset(idSejour: number, date: Date) {
    const sejour = (this.data.sejours.sejours || {})[idSejour];
    const dateDebut = new Date(sejour.date_debut);
    const dayMs = 24 * 60 * 60 * 1000;
    const truncateDebut = Math.floor(dateDebut.valueOf() / dayMs);
    const truncateFin = Math.floor(date.valueOf() / dayMs);
    return truncateFin - truncateDebut;
  }

  iterateAllRepas(fn: (sejour: Sejour, rep: RepasComplet) => void) {
    Object.values(this.data.sejours.sejours || {}).forEach(sejour => {
      if (!sejour.repass) return;
      sejour.repass.forEach(repas => {
        fn(sejour, repas);
      });
    });
  }

  getRepasGroupes(repas: { groupes: RepasGroupe[] | null }): Groupe[] {
    return (repas.groupes || []).map(
      rg => (this.data.sejours.groupes || {})[rg.id_groupe]
    );
  }

  getRepasNbPersonnes(repas: DetailsRepas) {
    const nb = this.getRepasGroupes(repas)
      .map(g => g.nb_personnes)
      .reduce((a, b) => a + b, repas.offset_personnes);
    return nb >= 0 ? nb : 0;
  }

  getSejour() {
    if (this.state.idSejour == null) return null;
    return (this.data.sejours.sejours || {})[this.state.idSejour] || null;
  }

  // renvoie les groupes du séjour courant
  getGroupes() {
    const idS = this.state.idSejour;
    if (idS == null) return [];
    return Object.values(this.data.sejours.groupes || {}).filter(
      groupe => groupe.id_sejour == idS
    );
  }
}
