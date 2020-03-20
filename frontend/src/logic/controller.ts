import {
  Menu,
  Recette,
  Sejour,
  RepasWithGroupe,
  Groupe,
  RepasGroupe,
  Produit
} from "./types";
import { Notifications } from "./notifications";
import { Calculs } from "./calculs";
import { Data, devMode } from "./data";
import { IngredientOptions, New, DetailsRepas } from "./types2";
import { Formatter } from "./formatter";
import { Loggin as Logger } from "./loggin";
import { State } from "./state";
import { searchFunction } from "@/components/utils/utils";

export class Controller {
  readonly data: Data;
  readonly notifications: Notifications;
  readonly calculs: Calculs;
  readonly formatter: Formatter;
  readonly logger: Logger;
  readonly state: State;

  token: string = "";
  idUtilisateur: number | null = devMode ? 2 : null;

  constructor() {
    this.data = new Data(this);
    this.notifications = new Notifications();
    this.calculs = new Calculs(this);
    this.formatter = new Formatter(this);
    this.logger = new Logger(this);
    this.state = new State(this);

    const o = this.logger.checkCookies();
    if (o != null) {
      this.token = o.token;
      this.idUtilisateur = o.idUtilisateur;
      this.state.isLoggedIn = true;
    }
  }

  auth() {
    return {
      username: String(this.idUtilisateur || ""),
      password: this.token
    };
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

  getMenuRecettes(menu: New<Menu>) {
    return (menu.recettes || []).map(rec => this.data.recettes[rec.id_recette]);
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
    if (!item.id_proprietaire.Valid) return null;
    return this.data.utilisateurs[item.id_proprietaire.Int64];
  }

  getRecetteIngredients(idRecette: number): IngredientOptions[] {
    return (this.getRecette(idRecette).ingredients || []).map(ing => {
      return {
        ingredient: (this.data.ingredients || {})[ing.id_ingredient],
        options: ing
      };
    });
  }

  getFournisseur(produit: Produit) {
    return (this.data.fournisseurs || {})[produit.id_fournisseur];
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
    return Math.round(
      (date.valueOf() - dateDebut.valueOf()) / (24 * 60 * 60 * 1000)
    );
  }

  iterateAllRepas(fn: (sejour: Sejour, rep: RepasWithGroupe) => void) {
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
}

// Object principal de stockage des données
// sur le client.
// Une instance de cet objet est créé au chargement,
// puis partagée entre les différents composants.
// Le système de réactivité de vuejs permet de propager
// facilement les changements effectués aux données.
// Ce composant est responsable de la comunication avec le serveur.
export const C = new Controller();
