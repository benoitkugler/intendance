import axios, { AxiosResponse } from "axios";
import {
  AgendaUtilisateur,
  Ingredient,
  OutAgenda,
  OutIngredients,
  OutIngredient,
  OutRecette,
  OutMenu,
  Recette,
  OutRecettes,
  OutMenus,
  Menu
} from "./types";
import { New, Ingredients, Recettes, Menus } from "./types2";

const devMode = process.env.NODE_ENV != "production";
const host = devMode ? "http://localhost:1323" : window.location.origin;
export const ServerURL = host + "/api";

interface Error {
  code: number | null;
  kind: string;
  messageHtml: string;
}

function arrayBufferToString(buffer: ArrayBuffer) {
  const uintArray = new Uint8Array(buffer);
  const encodedString = String.fromCharCode.apply(null, Array.from(uintArray));
  return decodeURIComponent(escape(encodedString));
}

function formateError(error: any): Error {
  let kind: string,
    messageHtml: string,
    code = null;
  if (error.response) {
    // The request was made and the server responded with a status code
    // that falls out of the range of 2xx
    kind = `Erreur côté serveur`;
    code = error.response.status;

    messageHtml = error.response.data.message;
    if (!messageHtml) {
      try {
        const json = arrayBufferToString(error.response.data);
        messageHtml = JSON.parse(json).message;
      } catch (error) {
        messageHtml = `Impossible de décoder la réponse du serveur. <br/>
        Détails : <i>${error}</i>`;
      }
    }
  } else if (error.request) {
    // The request was made but no response was received
    // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
    // http.ClientRequest in node.js
    kind = "Aucune réponse du serveur";
    messageHtml =
      "La requête a bien été envoyée, mais le serveur n'a donné aucune réponse...";
  } else {
    // Something happened in setting up the request that triggered an Error
    kind = "Erreur du client";
    messageHtml = `La requête n'a pu être mise en place. <br/>
                  Détails :  ${error.message} `;
  }
  return { kind, messageHtml, code };
}

class Data {
  agenda: AgendaUtilisateur;
  ingredients: Ingredients;
  recettes: Recettes;
  menus: Menus;

  error: Error | null;
  private token: string;
  private idUtilisateur: number | "*" | null;

  constructor() {
    this.agenda = { sejours: [] };
    this.ingredients = {};
    this.recettes = {};
    this.menus = {};

    this.token = "";
    this.idUtilisateur = devMode ? "*" : null;

    this.error = null;
  }

  private auth() {
    return {
      username: String(this.idUtilisateur || ""),
      password: this.token
    };
  }

  loadAgenda = async () => {
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
      this.error = formateError(error);
    }
  };

  loadIngredients = async () => {
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
      this.error = formateError(error);
    }
  };

  private createOrUpdateIngredient = async (
    ing: New<Ingredient>,
    method: "put" | "post"
  ) => {
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutIngredient> = await f(
        ServerURL + "/ingredients",
        ing,
        {
          auth: this.auth()
        }
      );
      this.ingredients[response.data.ingredient.id] = response.data.ingredient;
      return response.data.ingredient;
    } catch (error) {
      this.error = formateError(error);
    }
  };

  createIngredient = async (ing: New<Ingredient>) => {
    return this.createOrUpdateIngredient(ing, "put");
  };

  updateIngredient = async (ing: Ingredient) => {
    return this.createOrUpdateIngredient(ing, "post");
  };

  deleteIngredient = async (ing: Ingredient, checkProduits: boolean) => {
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
      this.error = formateError(error);
    }
  };

  loadRecettes = async () => {
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
      this.error = formateError(error);
    }
  };

  private createOrUpdateRecette = async (
    recette: New<Recette>,
    method: "put" | "post"
  ) => {
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutRecette> = await f(
        ServerURL + "/recettes",
        recette,
        {
          auth: this.auth()
        }
      );
      this.recettes[response.data.recette.id] = response.data.recette;
      return response.data.recette;
    } catch (error) {
      this.error = formateError(error);
    }
  };

  createRecette = async (recette: New<Recette>) => {
    return this.createOrUpdateRecette(recette, "put");
  };

  updateRecette = async (recette: Recette) => {
    return this.createOrUpdateRecette(recette, "post");
  };

  deleteRecette = async (recette: Recette) => {
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
      this.error = formateError(error);
    }
  };

  loadMenus = async () => {
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
      this.error = formateError(error);
    }
  };

  private createOrUpdateMenu = async (
    menu: New<Menu>,
    method: "put" | "post"
  ) => {
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutMenu> = await f(
        ServerURL + "/menus",
        menu,
        {
          auth: this.auth()
        }
      );
      this.menus[response.data.menu.id] = response.data.menu;
      return response.data.menu;
    } catch (error) {
      this.error = formateError(error);
    }
  };

  createMenu = async (menu: New<Menu>) => {
    return this.createOrUpdateMenu(menu, "put");
  };

  updateMenu = async (menu: Menu) => {
    return this.createOrUpdateMenu(menu, "post");
  };

  deleteMenu = async (menu: Menu) => {
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
      this.error = formateError(error);
    }
  };
}

// Object principal de stockage des données
// sur le client.
// Une instance de cet objet est créé au chargement,
// puis partagée entre les différents composants.
// Le système de réactivité de vuejs permet de propager
// facilement les changements effectués aux données.
// Ce composant est responsable de la comunication avec le serveur.
export const D = new Data();
