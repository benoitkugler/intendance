import axios, { AxiosResponse } from "axios";
import {
  AgendaUtilisateur,
  Ingredient,
  Menu,
  OutAgenda,
  OutIngredient,
  OutIngredients,
  OutMenu,
  OutMenus,
  OutRecette,
  OutRecettes,
  OutSejour,
  Recette,
  Repas,
  Sejour,
  Horaire,
  OutUtilisateurs
} from "./types";
import { Ingredients, Menus, New, Recettes, Utilisateurs } from "./types2";
import { NS } from "./notifications";
import Vue from "vue";

export const devMode = process.env.NODE_ENV != "production";
const host = devMode ? "http://localhost:1323" : window.location.origin;
export const ServerURL = host + "/api";

class Data {
  agenda: AgendaUtilisateur = { sejours: {} };
  ingredients: Ingredients = {};
  recettes: Recettes = {};
  menus: Menus = {};
  utilisateurs: Utilisateurs = {};

  private token: string = "";
  idUtilisateur: number | null = devMode ? 2 : null;

  private auth() {
    return {
      username: String(this.idUtilisateur || ""),
      password: this.token
    };
  }

  loadUtilisateurs = async () => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutUtilisateurs> = await axios.get(
        ServerURL + "/utilisateurs",
        {
          auth: this.auth()
        }
      );
      this.token = response.data.token;
      this.utilisateurs = response.data.utilisateurs;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  loadIngredients = async () => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutIngredients> = await axios.get(
        ServerURL + "/ingredients",
        {
          auth: this.auth()
        }
      );
      this.token = response.data.token;
      this.ingredients = response.data.ingredients;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  private createOrUpdateIngredient = async (
    ing: New<Ingredient>,
    method: "put" | "post"
  ) => {
    NS.startSpin();
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutIngredient> = await f(
        ServerURL + "/ingredients",
        ing,
        {
          auth: this.auth()
        }
      );
      Vue.set(
        this.ingredients,
        response.data.ingredient.id,
        response.data.ingredient
      ); // VRC
      this.token = response.data.token;

      return response.data.ingredient;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  createIngredient = async (ing: New<Ingredient>) => {
    return this.createOrUpdateIngredient(ing, "put");
  };

  updateIngredient = async (ing: Ingredient) => {
    return this.createOrUpdateIngredient(ing, "post");
  };

  deleteIngredient = async (ing: Ingredient, checkProduits: boolean) => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutIngredients> = await axios.delete(
        ServerURL + "/ingredients",
        {
          params: { id: ing.id, check_produits: checkProduits ? "check" : "" },
          auth: this.auth()
        }
      );
      this.ingredients = response.data.ingredients;
      this.token = response.data.token;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  loadRecettes = async () => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutRecettes> = await axios.get(
        ServerURL + "/recettes",
        {
          auth: this.auth()
        }
      );
      this.token = response.data.token;
      this.recettes = response.data.recettes;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  private createOrUpdateRecette = async (
    recette: New<Recette>,
    method: "put" | "post"
  ) => {
    NS.startSpin();
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutRecette> = await f(
        ServerURL + "/recettes",
        recette,
        {
          auth: this.auth()
        }
      );
      Vue.set(this.recettes, response.data.recette.id, response.data.recette); // VRC
      this.token = response.data.token;

      return response.data.recette;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  createRecette = async (recette: New<Recette>) => {
    return this.createOrUpdateRecette(recette, "put");
  };

  updateRecette = async (recette: Recette) => {
    return this.createOrUpdateRecette(recette, "post");
  };

  deleteRecette = async (recette: Recette) => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutRecettes> = await axios.delete(
        ServerURL + "/recettes",
        {
          params: { id: recette.id },
          auth: this.auth()
        }
      );
      this.recettes = response.data.recettes;
      this.token = response.data.token;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  loadMenus = async () => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutMenus> = await axios.get(
        ServerURL + "/menus",
        {
          auth: this.auth()
        }
      );
      this.token = response.data.token;
      this.menus = response.data.menus;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  private createOrUpdateMenu = async (
    menu: New<Menu>,
    method: "put" | "post"
  ) => {
    NS.startSpin();
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutMenu> = await f(
        ServerURL + "/menus",
        menu,
        {
          auth: this.auth()
        }
      );
      Vue.set(this.menus, response.data.menu.id, response.data.menu); // VRC
      this.token = response.data.token;

      return response.data.menu;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  createMenu = async (menu: New<Menu>) => {
    return this.createOrUpdateMenu(menu, "put");
  };

  updateMenu = async (menu: Menu) => {
    return this.createOrUpdateMenu(menu, "post");
  };

  deleteMenu = async (menu: Menu) => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutMenus> = await axios.delete(
        ServerURL + "/menus",
        {
          params: { id: menu.id },
          auth: this.auth()
        }
      );
      this.menus = response.data.menus;
      this.token = response.data.token;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  loadAgenda = async () => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutAgenda> = await axios.get(
        ServerURL + "/agenda",
        {
          auth: this.auth()
        }
      );
      this.token = response.data.token;
      this.agenda = response.data.agenda;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  private createOrUpdateSejour = async (
    sejour: New<Sejour>,
    method: "put" | "post"
  ) => {
    NS.startSpin();
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutSejour> = await f(
        ServerURL + "/sejours",
        sejour,
        {
          auth: this.auth()
        }
      );
      const entry = this.agenda.sejours[response.data.sejour.id] || {
        journees: [],
        sejour: response.data.sejour
      };
      entry.sejour = response.data.sejour;
      Vue.set(this.agenda.sejours, response.data.sejour.id, entry); // VRC
      this.token = response.data.token;
      return response.data.sejour;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  createSejour = async (sejour: New<Sejour>) => {
    return this.createOrUpdateSejour(sejour, "put");
  };

  updateSejour = async (sejour: Sejour) => {
    return this.createOrUpdateSejour(sejour, "post");
  };

  deleteSejour = async (sejour: Sejour) => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutAgenda> = await axios.delete(
        ServerURL + "/sejours",
        {
          params: { id: sejour.id },
          auth: this.auth()
        }
      );
      this.agenda = response.data.agenda;
      this.token = response.data.token;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  createRepas = async (repas: New<Repas>) => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutAgenda> = await axios.put(
        ServerURL + "/sejours/repas",
        repas,
        {
          auth: this.auth()
        }
      );
      this.agenda = response.data.agenda;
      this.token = response.data.token;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  updateManyRepas = async (repass: Repas[]) => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutAgenda> = await axios.post(
        ServerURL + "/sejours/repas",
        repass,
        {
          auth: this.auth()
        }
      );
      this.agenda = response.data.agenda;
      this.token = response.data.token;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  deleteRepas = async (repas: Repas) => {
    NS.startSpin();
    try {
      const response: AxiosResponse<OutAgenda> = await axios.delete(
        ServerURL + "/sejours/repas",
        {
          params: { id: repas.id },
          auth: this.auth()
        }
      );
      this.agenda = response.data.agenda;
      this.token = response.data.token;
    } catch (error) {
      NS.setAxiosError(error);
    }
  };

  // gère l'erreur d'un séjour introuvable
  private getSejour(idSejour: number) {
    const sejour = this.agenda.sejours[idSejour];
    if (!sejour) {
      NS.setError({
        code: null,
        kind: "Séjour introuvable",
        messageHtml: `Le séjour concerné (id ${idSejour}) est <i>introuvable<i> <br/>
            Il s'agit probablement d'une erreur de synchronisation avec le serveur. <br/>
            Pour la corriger, merci de <b>recharger la page</b>.`
      });
    }
    return sejour;
  }

  // vérifie que la date est après le début du séjour et gère l'erreur
  // si la date est valide, renvoie l'offset correspondant
  getOffset(sejour: Sejour, jour: Date) {
    const dateDebut = new Date(sejour.date_debut);
    if (jour < dateDebut) {
      // invalide
      NS.setError({
        code: null,
        kind: "Jour invalide",
        messageHtml: `La date ciblée (${jour.toLocaleDateString()}) est <i>antérieure</i> au début du séjour.<br/>
                  Si vous souhaitez déplacer un repas sur cette journée, 
                  veuillez d'abord <b>modifier la date de début</b> du séjour <b>${
                    sejour.nom
                  }</b>`
      });
      return;
    }
    const offset = Math.ceil(
      (jour.valueOf() - dateDebut.valueOf()) / (1000 * 60 * 60 * 24)
    );
    return offset;
  }

  // échange les deux journées, en modifiant les dates
  // des repas concernés pour le séjour donné.
  async switchDays(idSejour: number, from: Date, to: Date) {
    const sejour = this.getSejour(idSejour);
    if (!sejour) return;
    const offsetTo = this.getOffset(sejour.sejour, to);
    if (offsetTo === undefined) return;
    const offsetFrom = this.getOffset(sejour.sejour, from);
    if (offsetFrom === undefined) return;
    if (offsetFrom == offsetTo) return;
    const menusFrom =
      (sejour.journees[offsetFrom] || { menus: [] }).menus || [];
    const menusTo = (sejour.journees[offsetTo] || { menus: [] }).menus || [];
    menusFrom.forEach(m => (m.jour_offset = offsetTo));
    menusTo.forEach(m => (m.jour_offset = offsetFrom));
    const modifs = menusFrom.concat(menusTo);
    if (modifs.length === 0) return;
    await this.updateManyRepas(modifs);
    if (NS.getError() === null) {
      NS.setMessage("Les journées ont étés échangées avec succès.");
    }
  }

  // modifie le moment du repas
  async deplaceRepas(repas: Repas, jour: Date, horaire: Horaire) {
    const sejour = this.getSejour(repas.id_sejour);
    if (!sejour) return;
    const offset = this.getOffset(sejour.sejour, jour);
    if (offset === undefined) return;
    repas.jour_offset = offset;
    repas.horaire = horaire;
    await this.updateManyRepas([repas]);
    if (NS.getError() === null) {
      NS.setMessage("Le repas a été déplacé avec succès.");
    }
  }
}

// Object principal de stockage des données
// sur le client.
// Une instance de cet objet est créé au chargement,
// puis partagée entre les différents composants.
// Le système de réactivité de vuejs permet de propager
// facilement les changements effectués aux données.
// Ce composant est responsable de la comunication avec le serveur.
export const D = new Data();
