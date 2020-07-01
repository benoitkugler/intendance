import {
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
  Sejour,
  OutSejour,
  OutIngredientProduits,
  OutSejours,
  Sejours,
  Groupe,
  OutGroupe,
  OutDeleteGroupe,
  RepasComplet,
  OptionsAssistantCreateRepass,
  InAssistantCreateRepass,
  InAjouteIngredientProduit,
  OutFournisseurs,
  Fournisseurs,
  Ingredients,
  InSejourFournisseurs,
  RecetteComplet,
  MenuComplet,
  Utilisateur,
  Produit,
  InSetDefautProduit,
  Livraisons,
  SejourRepas,
  Fournisseur,
  OutFournisseur,
  Livraison,
  OutLivraison,
  Horaire
} from "./types";
import axios, { AxiosResponse } from "axios";

import { Controller } from "./controller";
import Vue from "vue";
import { New } from "./types2";

export const devMode = process.env.NODE_ENV != "production";
const host = devMode ? "http://localhost:1323" : window.location.origin;
export const ServerURL = host + "/api";

export class Data {
  sejours: Sejours = {
    sejours: {},
    groupes: {}
  };
  ingredients: Ingredients = {};
  recettes: { [key: number]: RecetteComplet } = {};
  menus: { [key: number]: MenuComplet } = {};
  utilisateurs: { [key: number]: Utilisateur } = {};
  fournisseurs: Fournisseurs = {};
  livraisons: Livraisons = {};

  private controller: Controller;

  constructor(controller: Controller) {
    this.controller = controller;
  }

  async loadFournisseurs() {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutFournisseurs> = await axios.get(
        ServerURL + "/fournisseurs",
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      this.fournisseurs = response.data.fournisseurs || {};
      this.livraisons = response.data.livraisons || {};
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  }

  // charge transitivement les données nécessaires aux menus
  async loadAllMenus() {
    this.loadFournisseurs();
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
      this.utilisateurs = response.data.utilisateurs || {};
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
      this.ingredients = response.data.ingredients || {};
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
        this.ingredients || {},
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

  deleteIngredient = async (idIngredient: number, checkProduits: boolean) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutIngredients> = await axios.delete(
        ServerURL + "/ingredients",
        {
          params: {
            id: idIngredient,
            check_produits: checkProduits ? "check" : ""
          },
          auth: this.controller.auth()
        }
      );
      this.ingredients = response.data.ingredients || {};
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  getIngredientProduits = async (idIngredient: number) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutIngredientProduits> = await axios.get(
        ServerURL + "/ingredient-produit",
        {
          auth: this.controller.auth(),
          params: {
            id: idIngredient
          }
        }
      );
      this.controller.token = response.data.token;
      this.controller.notifications.setMessage(null);
      return response.data.produits;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  // renvoie la liste des produits mise à jour
  ajouteIngredientProduit = async (ip: InAjouteIngredientProduit) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutIngredientProduits> = await axios.post(
        ServerURL + "/ingredient-produit",
        ip,
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      this.controller.notifications.setMessage("Produit associé avec succès.");
      return response.data.produits;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  // `getIngredientProduits` devrait être appelé ensuite
  deleteProduit = async (idProduit: number) => {
    this.controller.notifications.startSpin();
    try {
      await axios.delete(ServerURL + "/produits", {
        params: { id: idProduit },
        auth: this.controller.auth()
      });
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  // renvoie la liste des produits mise à jour
  setDefautProduit = async (params: InSetDefautProduit) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutIngredientProduits> = await axios.post(
        ServerURL + "/ingredient-produit-defaut",
        params,
        {
          auth: this.controller.auth()
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
      this.recettes = response.data.recettes || {};
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  private createOrUpdateRecette = async (
    recette: New<RecetteComplet>,
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

  createRecette = async (recette: New<RecetteComplet>) => {
    return this.createOrUpdateRecette(recette, "put");
  };

  updateRecette = async (recette: RecetteComplet) => {
    return this.createOrUpdateRecette(recette, "post");
  };

  deleteRecette = async (idRecette: number) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutRecettes> = await axios.delete(
        ServerURL + "/recettes",
        {
          params: { id: idRecette },
          auth: this.controller.auth()
        }
      );
      this.recettes = response.data.recettes || {};
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
      this.menus = response.data.menus || {};
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  private createOrUpdateMenu = async (
    menu: New<MenuComplet>,
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

  createMenu = async (menu: New<MenuComplet>) => {
    return this.createOrUpdateMenu(menu, "put");
  };

  updateMenu = async (menu: MenuComplet) => {
    return this.createOrUpdateMenu(menu, "post");
  };

  deleteMenu = async (idMenu: number) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutMenus> = await axios.delete(
        ServerURL + "/menus",
        {
          params: { id: idMenu },
          auth: this.controller.auth()
        }
      );
      this.menus = response.data.menus || {};
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  loadSejours = async () => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutSejours> = await axios.get(
        ServerURL + "/sejours",
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      this.sejours = response.data.sejours;
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
      this.controller.token = response.data.token;
      return response.data.sejour;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  createSejour = async (sejour: New<Sejour>) => {
    const sej = await this.createOrUpdateSejour(sejour, "put");
    if (sej === undefined) return;
    const newSejour: SejourRepas = { ...sej, repass: [], fournisseurs: [] };
    Vue.set(this.sejours.sejours || {}, newSejour.id, newSejour); // VRC
    return newSejour;
  };

  updateSejour = async (sejour: Sejour) => {
    const sej = await this.createOrUpdateSejour(sejour, "post");
    if (sej === undefined) return;
    const oldSejour = (this.sejours.sejours || {})[sej.id];
    const updatedSejour: SejourRepas = {
      ...sej,
      repass: oldSejour.repass,
      fournisseurs: oldSejour.fournisseurs
    };
    (this.sejours.sejours || {})[sej.id] = updatedSejour;
    return updatedSejour;
  };

  deleteSejour = async (sejour: Sejour) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutSejours> = await axios.delete(
        ServerURL + "/sejours",
        {
          params: { id: sejour.id },
          auth: this.controller.auth()
        }
      );
      this.sejours = response.data.sejours;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  // Fournisseurs
  createFournisseur = async (fournisseur: New<Fournisseur>) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutFournisseurs> = await axios.put(
        ServerURL + "/fournisseurs",
        fournisseur,
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      this.fournisseurs = response.data.fournisseurs || {};
      this.livraisons = response.data.livraisons || {};
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  updateFournisseur = async (fournisseur: Fournisseur) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutFournisseur> = await axios.post(
        ServerURL + "/fournisseurs",
        fournisseur,
        {
          auth: this.controller.auth()
        }
      );
      Vue.set(
        this.fournisseurs || {},
        response.data.fournisseur.id,
        response.data.fournisseur
      ); // VRC
      this.controller.token = response.data.token;
      return response.data.fournisseur;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  deleteFournisseur = async (idFournisseur: number) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutFournisseurs> = await axios.delete(
        ServerURL + "/fournisseurs",
        {
          params: { id: idFournisseur },
          auth: this.controller.auth()
        }
      );
      this.fournisseurs = response.data.fournisseurs || {};
      this.livraisons = response.data.livraisons || {};
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  updateSejourFournisseurs = async (
    sejour: Sejour,
    idsFournisseurs: number[]
  ) => {
    this.controller.notifications.startSpin();
    const params: InSejourFournisseurs = {
      id_sejour: sejour.id,
      ids_fournisseurs: idsFournisseurs
    };
    try {
      const response: AxiosResponse<OutSejours> = await axios.post(
        ServerURL + "/sejours/fournisseurs",
        params,
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      this.sejours = response.data.sejours;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  // Livraisons
  private createOrUpdateLivraison = async (
    livraison: New<Livraison>,
    method: "put" | "post"
  ) => {
    this.controller.notifications.startSpin();
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutLivraison> = await f(
        ServerURL + "/livraisons",
        livraison,
        {
          auth: this.controller.auth()
        }
      );
      Vue.set(
        this.livraisons || {},
        response.data.livraison.id,
        response.data.livraison
      ); // VRC
      this.controller.token = response.data.token;

      return response.data.livraison;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  createLivraison = async (livraison: New<Livraison>) => {
    return this.createOrUpdateLivraison(livraison, "put");
  };

  updateLivraison = async (livraison: Livraison) => {
    return this.createOrUpdateLivraison(livraison, "post");
  };

  deleteLivraison = async (idLivraison: number) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutFournisseurs> = await axios.delete(
        ServerURL + "/livraisons",
        {
          params: { id: idLivraison },
          auth: this.controller.auth()
        }
      );
      this.livraisons = response.data.livraisons || {};
      this.fournisseurs = response.data.fournisseurs || {};
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  private createOrUpdateGroupe = async (
    groupe: New<Groupe>,
    method: "put" | "post"
  ) => {
    this.controller.notifications.startSpin();
    const f = method == "put" ? axios.put : axios.post;
    try {
      const response: AxiosResponse<OutGroupe> = await f(
        ServerURL + "/groupes",
        groupe,
        {
          auth: this.controller.auth()
        }
      );
      Vue.set(
        this.sejours.groupes || {},
        response.data.groupe.id,
        response.data.groupe
      ); // VRC
      this.controller.token = response.data.token;
      return response.data.groupe;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  createGroupe = async (groupe: New<Groupe>) => {
    return this.createOrUpdateGroupe(groupe, "put");
  };

  updateGroupe = async (groupe: Groupe) => {
    return this.createOrUpdateGroupe(groupe, "post");
  };

  // Renvoie le nombre de repas touchés par la suppression
  deleteGroupe = async (groupe: Groupe) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutDeleteGroupe> = await axios.delete(
        ServerURL + "/groupes",
        {
          params: { id: groupe.id },
          auth: this.controller.auth()
        }
      );
      Vue.delete(this.sejours.groupes || {}, groupe.id); // VRC
      this.controller.token = response.data.token;
      return response.data.nb_repas;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  createRepas = async (repas: New<RepasComplet>) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutSejours> = await axios.put(
        ServerURL + "/sejours/repas",
        repas,
        {
          auth: this.controller.auth()
        }
      );
      this.sejours = response.data.sejours;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  updateManyRepas = async (repass: RepasComplet[]) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutSejours> = await axios.post(
        ServerURL + "/sejours/repas",
        repass,
        {
          auth: this.controller.auth()
        }
      );
      this.sejours = response.data.sejours;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

  deleteRepas = async (repas: RepasComplet) => {
    this.controller.notifications.startSpin();
    try {
      const response: AxiosResponse<OutSejours> = await axios.delete(
        ServerURL + "/sejours/repas",
        {
          params: { id: repas.id },
          auth: this.controller.auth()
        }
      );
      this.sejours = response.data.sejours;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  };

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

  // gère l'erreur d'un séjour introuvable
  getSejour(idSejour: number) {
    return (this.sejours.sejours || {})[idSejour];
  }

  // échange les deux journées, en modifiant les dates
  // des repas concernés pour le séjour donné.
  async switchDays(idSejour: number, from: Date, to: Date) {
    const sejour = this.getSejour(idSejour);
    if (!sejour) return;
    const offsetTo = this.getOffset(sejour, to);
    if (offsetTo === undefined) return;
    const offsetFrom = this.getOffset(sejour, from);
    if (offsetFrom === undefined) return;
    if (offsetFrom == offsetTo) return;
    const repasFrom = (sejour.repass || []).filter(
      rep => rep.jour_offset == offsetFrom
    );
    const repasTo = (sejour.repass || []).filter(
      rep => rep.jour_offset == offsetTo
    );
    repasFrom.forEach(m => (m.jour_offset = offsetTo));
    repasTo.forEach(m => (m.jour_offset = offsetFrom));
    const modifs = repasFrom.concat(repasTo);
    if (modifs.length === 0) return;
    await this.updateManyRepas(modifs);
    if (this.controller.notifications.getError() === null) {
      this.controller.notifications.setMessage(
        "Les journées ont étés échangées avec succès."
      );
    }
  }

  // modifie le moment du repas
  async deplaceRepas(repas: RepasComplet, jour: Date, horaire: Horaire) {
    const sejour = this.getSejour(repas.id_sejour);
    if (!sejour) return;
    const offset = this.getOffset(sejour, jour);
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

  async assitantCreateRepass(
    idSejour: number,
    options: OptionsAssistantCreateRepass,
    groupesSorties: { [key: number]: number[] }
  ) {
    this.controller.notifications.startSpin();
    const params: InAssistantCreateRepass = {
      id_sejour: idSejour,
      options: options,
      groupes_sorties: groupesSorties
    };
    try {
      const response: AxiosResponse<OutSejours> = await axios.put(
        ServerURL + "/sejours/assistant",
        params,
        {
          auth: this.controller.auth()
        }
      );
      this.sejours = response.data.sejours;
      this.controller.token = response.data.token;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  }
}
