import {
  AgendaUtilisateur,
  OutUtilisateurs,
  OutIngredients,
  Ingredient,
  OutIngredient,
  OutRecettes,
  Recette,
  OutRecette,
  OutMenus,
  Menu,
  OutMenu,
  OutAgenda,
  Sejour,
  OutSejour,
  Repas,
  Horaire,
  OutIngredientProduits
} from "./types";
import axios, { AxiosResponse } from "axios";
import { Ingredients, Recettes, Menus, Utilisateurs, New } from "./types2";

import { Controller } from "./controller";
import Vue from "vue";

export const devMode = process.env.NODE_ENV != "production";
const host = devMode ? "http://localhost:1323" : window.location.origin;
export const ServerURL = host + "/api";

export class Data {
  agenda: AgendaUtilisateur = { sejours: {} };
  ingredients: Ingredients = {};
  recettes: Recettes = {};
  menus: Menus = {};
  utilisateurs: Utilisateurs = {};

  private controller: Controller;

  constructor(controller: Controller) {
    this.controller = controller;
  }

  async loadAllMenus() {
    await Promise.all([this.loadIngredients(), this.loadUtilisateurs()]);
    await this.loadRecettes(); // recettes dépend des ingrédients
    await this.loadMenus(); // menus dépends des recettes, ingrédients et utilisateurs
  }

  loadUtilisateurs = async () => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutUtilisateurs> = await axios.get(
        ServerURL + "/utilisateurs",
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      this.utilisateurs = response.data.utilisateurs;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  loadIngredients = async () => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutIngredients> = await axios.get(
        ServerURL + "/ingredients",
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      this.ingredients = response.data.ingredients;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  private createOrUpdateIngredient = async (
    ing: New<Ingredient>,
    method: "put" | "post"
  ) => {
    this.controller.notifications.startSpin();
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutIngredient> = await f(
        ServerURL + "/ingredients",
        ing,
        {
          auth: this.controller.auth()
        }
      );
      Vue.set(
        this.ingredients,
        response.data.ingredient.id,
        response.data.ingredient
      ); // VRC
      this.controller.token = response.data.token;

      return response.data.ingredient;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  createIngredient = async (ing: New<Ingredient>) => {
    return this.createOrUpdateIngredient(ing, "put");
  };

  updateIngredient = async (ing: Ingredient) => {
    return this.createOrUpdateIngredient(ing, "post");
  };

  deleteIngredient = async (ing: Ingredient, checkProduits: boolean) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutIngredients> = await axios.delete(
        ServerURL + "/ingredients",
        {
          params: { id: ing.id, check_produits: checkProduits ? "check" : "" },
          auth: this.controller.auth()
        }
      );
      this.ingredients = response.data.ingredients;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  getIngredientProduits = async (idIngredient: number) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutIngredientProduits> = await axios.get(
        ServerURL + "/recettes",
        {
          auth: this.controller.auth(),
          params: {
            id: idIngredient
          }
        }
      );
      this.controller.token = response.data.token;
      return response.data.produits;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  loadRecettes = async () => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutRecettes> = await axios.get(
        ServerURL + "/recettes",
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      this.recettes = response.data.recettes;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  private createOrUpdateRecette = async (
    recette: New<Recette>,
    method: "put" | "post"
  ) => {
    this.controller.notifications.startSpin();
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutRecette> = await f(
        ServerURL + "/recettes",
        recette,
        {
          auth: this.controller.auth()
        }
      );
      Vue.set(this.recettes, response.data.recette.id, response.data.recette); // VRC
      this.controller.token = response.data.token;

      return response.data.recette;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  createRecette = async (recette: New<Recette>) => {
    return this.createOrUpdateRecette(recette, "put");
  };

  updateRecette = async (recette: Recette) => {
    return this.createOrUpdateRecette(recette, "post");
  };

  deleteRecette = async (recette: Recette) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutRecettes> = await axios.delete(
        ServerURL + "/recettes",
        {
          params: { id: recette.id },
          auth: this.controller.auth()
        }
      );
      this.recettes = response.data.recettes;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  loadMenus = async () => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutMenus> = await axios.get(
        ServerURL + "/menus",
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      this.menus = response.data.menus;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  private createOrUpdateMenu = async (
    menu: New<Menu>,
    method: "put" | "post"
  ) => {
    this.controller.notifications.startSpin();
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutMenu> = await f(
        ServerURL + "/menus",
        menu,
        {
          auth: this.controller.auth()
        }
      );
      Vue.set(this.menus, response.data.menu.id, response.data.menu); // VRC
      this.controller.token = response.data.token;

      return response.data.menu;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  createMenu = async (menu: New<Menu>) => {
    return this.createOrUpdateMenu(menu, "put");
  };

  updateMenu = async (menu: Menu) => {
    return this.createOrUpdateMenu(menu, "post");
  };

  deleteMenu = async (menu: Menu) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutMenus> = await axios.delete(
        ServerURL + "/menus",
        {
          params: { id: menu.id },
          auth: this.controller.auth()
        }
      );
      this.menus = response.data.menus;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  loadAgenda = async () => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutAgenda> = await axios.get(
        ServerURL + "/agenda",
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      this.agenda = response.data.agenda;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  private createOrUpdateSejour = async (
    sejour: New<Sejour>,
    method: "put" | "post"
  ) => {
    this.controller.notifications.startSpin();
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutSejour> = await f(
        ServerURL + "/sejours",
        sejour,
        {
          auth: this.controller.auth()
        }
      );
      const entry = this.agenda.sejours[response.data.sejour.id] || {
        journees: [],
        sejour: response.data.sejour
      };
      entry.sejour = response.data.sejour;
      Vue.set(this.agenda.sejours, response.data.sejour.id, entry); // VRC
      this.controller.token = response.data.token;
      return response.data.sejour;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  createSejour = async (sejour: New<Sejour>) => {
    return this.createOrUpdateSejour(sejour, "put");
  };

  updateSejour = async (sejour: Sejour) => {
    return this.createOrUpdateSejour(sejour, "post");
  };

  deleteSejour = async (sejour: Sejour) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutAgenda> = await axios.delete(
        ServerURL + "/sejours",
        {
          params: { id: sejour.id },
          auth: this.controller.auth()
        }
      );
      this.agenda = response.data.agenda;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  createRepas = async (repas: New<Repas>) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutAgenda> = await axios.put(
        ServerURL + "/sejours/repas",
        repas,
        {
          auth: this.controller.auth()
        }
      );
      this.agenda = response.data.agenda;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  updateManyRepas = async (repass: Repas[]) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutAgenda> = await axios.post(
        ServerURL + "/sejours/repas",
        repass,
        {
          auth: this.controller.auth()
        }
      );
      this.agenda = response.data.agenda;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  deleteRepas = async (repas: Repas) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutAgenda> = await axios.delete(
        ServerURL + "/sejours/repas",
        {
          params: { id: repas.id },
          auth: this.controller.auth()
        }
      );
      this.agenda = response.data.agenda;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  // gère l'erreur d'un séjour introuvable
  private getSejour(idSejour: number) {
    const sejour = this.agenda.sejours[idSejour];
    if (!sejour) {
      this.controller.notifications.setError({
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
      this.controller.notifications.setError({
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
    if (this.controller.notifications.getError() === null) {
      this.controller.notifications.setMessage(
        "Les journées ont étés échangées avec succès."
      );
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
    if (this.controller.notifications.getError() === null) {
      this.controller.notifications.setMessage(
        "Le repas a été déplacé avec succès."
      );
    }
  }
}
