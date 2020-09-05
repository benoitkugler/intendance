import {
  Sejours,
  Ingredients,
  RecetteComplet,
  MenuComplet,
  Utilisateur,
  Fournisseurs,
  Livraisons,
  OutFournisseurs,
  OutUtilisateurs,
  OutIngredients,
  Ingredient,
  OutIngredient,
  OutIngredientProduits,
  InAjouteIngredientProduit,
  InSetDefautProduit,
  OutRecettes,
  OutRecette,
  OutMenus,
  OutMenu,
  OutSejours,
  Sejour,
  OutSejour,
  SejourRepas,
  Fournisseur,
  OutFournisseur,
  InSejourFournisseurs,
  Livraison,
  OutLivraison,
  Groupe,
  OutGroupe,
  OutDeleteGroupe,
  RepasComplet,
  Horaire,
  OptionsAssistantCreateRepass,
  InAssistantCreateRepass,
  Groupes
} from "./types";
import { Controller } from "./controller";
import { AxiosResponse } from "axios";
import { New } from "./types2";
import Vue from "vue";
import { API } from "./server";

/* Ce composant est responsable de la comunication avec le serveur, via une classe API
 * Il stocke et met à jour les données client.
 */
export class Data {
  sejours: {
    sejours: { [key: number]: SejourRepas };
    groupes: NonNullable<Groupes>;
  } = { sejours: {}, groupes: {} };
  ingredients: NonNullable<Ingredients> = {};
  recettes: { [key: number]: RecetteComplet } = {};
  menus: { [key: number]: MenuComplet } = {};
  utilisateurs: { [key: number]: Utilisateur } = {};
  fournisseurs: NonNullable<Fournisseurs> = {};
  livraisons: NonNullable<Livraisons> = {};

  constructor(private api: API) {}

  async loadFournisseurs() {
    const out = await this.api.GetFournisseurs();
    if (out === undefined) return; // erreur
    this.fournisseurs = out.fournisseurs || {};
    this.livraisons = out.livraisons || {};
  }

  // charge transitivement les données nécessaires aux menus
  async loadAllMenus() {
    this.loadFournisseurs();
    await Promise.all([this.loadIngredients(), this.loadUtilisateurs()]);
    await this.loadRecettes(); // recettes dépend des ingrédients
    await this.loadMenus(); // menus dépends des recettes, ingrédients et utilisateurs
  }

  loadUtilisateurs = async () => {
    const out = await this.api.GetUtilisateurs();
    if (out === undefined) return;
    this.utilisateurs = out || {};
  };

  loadIngredients = async () => {
    const out = await this.api.GetIngredients();
    if (out === undefined) return;
    this.ingredients = out || {};
  };

  private createOrUpdateIngredient = async (
    ing: Ingredient,
    method: "create" | "update"
  ) => {
    const f =
      method == "create"
        ? this.api.CreateIngredient
        : this.api.UpdateIngredient;
    const out = await f(ing);
    if (out === undefined) return;
    Vue.set(this.ingredients || {}, out.id, out); // VRC
    return out;
  };

  createIngredient = async (ing: New<Ingredient>) => {
    return this.createOrUpdateIngredient({ id: 0, ...ing }, "create");
  };

  updateIngredient = async (ing: Ingredient) => {
    return this.createOrUpdateIngredient(ing, "update");
  };

  deleteIngredient = async (idIngredient: number, checkProduits: boolean) => {
    const out = await this.api.DeleteIngredient({
      id: String(idIngredient),
      check_produits: checkProduits ? "check" : ""
    });
    if (out === undefined) return;
    this.ingredients = out || {};
  };

  getIngredientProduits = async (idIngredient: number) => {
    return this.api.GetIngredientProduits({ id: String(idIngredient) });
  };

  // renvoie la liste des produits mise à jour
  ajouteIngredientProduit = async (ip: InAjouteIngredientProduit) => {
    return this.api.AjouteIngredientProduit(ip);
  };

  // `getIngredientProduits` devrait être appelé ensuite
  deleteProduit = async (idProduit: number) => {
    await this.api.DeleteProduit({ id: String(idProduit) });
  };

  // renvoie la liste des produits mise à jour
  setDefautProduit = async (params: InSetDefautProduit) => {
    return this.api.SetDefautProduit(params);
  };

  // TODO: WIP
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
