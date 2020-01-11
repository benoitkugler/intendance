import { Menu, Recette, Sejour, Repas } from "./types";
import { Notifications } from "./notifications";
import { Calculs } from "./calculs";
import { Data, devMode } from "./data";
import { IngredientOptions } from "./types2";
import { Formatter } from "./formatter";
import { Loggin as Logger } from "./loggin";
import { State } from "./state";

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
  }

  auth() {
    return {
      username: String(this.idUtilisateur || ""),
      password: this.token
    };
  }

  getAllIngredients(): IngredientOptions[] {
    return Object.values(this.data.ingredients).map(ing => {
      return { ingredient: ing };
    });
  }

  getMenuRecettes(menu: Menu) {
    return (menu.recettes || []).map(rec => this.data.recettes[rec.id_recette]);
  }

  getMenuIngredients(menu: Menu): IngredientOptions[] {
    return (menu.ingredients || []).map(ing => {
      return {
        ingredient: this.data.ingredients[ing.id_ingredient],
        options: ing
      };
    });
  }

  getMenuOrRecetteProprietaire(item: Menu | Recette) {
    if (!item.id_proprietaire.Valid) return null;
    return this.data.utilisateurs[item.id_proprietaire.Int64];
  }

  getRecetteIngredients(rec: Recette): IngredientOptions[] {
    return (rec.ingredients || []).map(ing => {
      return {
        ingredient: this.data.ingredients[ing.id_ingredient],
        options: ing
      };
    });
  }

  iterateAllRepas(fn: (sejour: Sejour, rep: Repas) => void) {
    Object.values(this.data.sejours.sejours).forEach(sejour => {
      if (!sejour.repass) return;
      sejour.repass.forEach(repas => {
        fn(sejour, repas);
      });
    });
  }

  getRepasGroupes(repas: Repas) {}
}

// Object principal de stockage des données
// sur le client.
// Une instance de cet objet est créé au chargement,
// puis partagée entre les différents composants.
// Le système de réactivité de vuejs permet de propager
// facilement les changements effectués aux données.
// Ce composant est responsable de la comunication avec le serveur.
export const C = new Controller();
