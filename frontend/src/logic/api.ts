import * as types from "./types";

import Axios, { AxiosResponse } from "axios";

export abstract class API {
  constructor(
    protected baseUrl: string,
    protected authToken: string,
    protected urlParams: {}
  ) {}

  abstract handleError(error: any): void;

  getHeaders() {
    return { Authorization: "Bearer " + this.authToken };
  }

  async Loggin(params: types.InLoggin) {
    try {
      const fullUrl = this.baseUrl + "/api/loggin";
      const rep: AxiosResponse<types.OutLoggin> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async GetUtilisateurs() {
    try {
      const fullUrl = this.baseUrl + "/api/utilisateurs";
      const rep: AxiosResponse<{
        [key: number]: types.Utilisateur;
      } | null> = await Axios.get(fullUrl);
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async GetIngredients() {
    try {
      const fullUrl = this.baseUrl + "/api/ingredients";
      const rep: AxiosResponse<types.Ingredients> = await Axios.get(fullUrl);
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async CreateIngredient(params: types.Ingredient) {
    try {
      const fullUrl = this.baseUrl + "/api/ingredients";
      const rep: AxiosResponse<types.Ingredient> = await Axios.put(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async UpdateIngredient(params: types.Ingredient) {
    try {
      const fullUrl = this.baseUrl + "/api/ingredients";
      const rep: AxiosResponse<types.Ingredient> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async DeleteIngredient(params: { id: string; check_produits: string }) {
    try {
      const fullUrl = this.baseUrl + "/api/ingredients";
      const rep: AxiosResponse<types.Ingredients> = await Axios.delete(
        fullUrl,
        { params: params, headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async GetRecettes() {
    try {
      const fullUrl = this.baseUrl + "/api/recettes";
      const rep: AxiosResponse<{
        [key: number]: types.RecetteComplet;
      } | null> = await Axios.get(fullUrl);
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async CreateRecette(params: types.RecetteComplet) {
    try {
      const fullUrl = this.baseUrl + "/api/recettes";
      const rep: AxiosResponse<types.RecetteComplet> = await Axios.put(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async UpdateRecette(params: types.RecetteComplet) {
    try {
      const fullUrl = this.baseUrl + "/api/recettes";
      const rep: AxiosResponse<types.RecetteComplet> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async DeleteRecette(params: { id: string }) {
    try {
      const fullUrl = this.baseUrl + "/api/recettes";
      const rep: AxiosResponse<{
        [key: number]: types.RecetteComplet;
      } | null> = await Axios.delete(fullUrl, {
        params: params,
        headers: this.getHeaders()
      });
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async GetMenus() {
    try {
      const fullUrl = this.baseUrl + "/api/menus";
      const rep: AxiosResponse<{
        [key: number]: types.MenuComplet;
      } | null> = await Axios.get(fullUrl);
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async CreateMenu(params: types.MenuComplet) {
    try {
      const fullUrl = this.baseUrl + "/api/menus";
      const rep: AxiosResponse<types.MenuComplet> = await Axios.put(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async UpdateMenu(params: types.MenuComplet) {
    try {
      const fullUrl = this.baseUrl + "/api/menus";
      const rep: AxiosResponse<types.MenuComplet> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async DeleteMenu(params: { id: string }) {
    try {
      const fullUrl = this.baseUrl + "/api/menus";
      const rep: AxiosResponse<{
        [key: number]: types.MenuComplet;
      } | null> = await Axios.delete(fullUrl, {
        params: params,
        headers: this.getHeaders()
      });
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async GetSejours() {
    try {
      const fullUrl = this.baseUrl + "/api/sejours";
      const rep: AxiosResponse<types.Sejours> = await Axios.get(fullUrl);
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async CreateSejour(params: types.Sejour) {
    try {
      const fullUrl = this.baseUrl + "/api/sejours";
      const rep: AxiosResponse<types.Sejour> = await Axios.put(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async UpdateSejour(params: types.Sejour) {
    try {
      const fullUrl = this.baseUrl + "/api/sejours";
      const rep: AxiosResponse<types.Sejour> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async DeleteSejour(params: { id: string }) {
    try {
      const fullUrl = this.baseUrl + "/api/sejours";
      const rep: AxiosResponse<types.Sejours> = await Axios.delete(fullUrl, {
        params: params,
        headers: this.getHeaders()
      });
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async CreateGroupe(params: types.Groupe) {
    try {
      const fullUrl = this.baseUrl + "/api/groupes";
      const rep: AxiosResponse<types.Groupe> = await Axios.put(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async UpdateGroupe(params: types.Groupe) {
    try {
      const fullUrl = this.baseUrl + "/api/groupes";
      const rep: AxiosResponse<types.Groupe> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async DeleteGroupe(params: { id: string }) {
    try {
      const fullUrl = this.baseUrl + "/api/groupes";
      const rep: AxiosResponse<types.OutDeleteGroupe> = await Axios.delete(
        fullUrl,
        { params: params, headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async UpdateSejourFournisseurs(params: types.InSejourFournisseurs) {
    try {
      const fullUrl = this.baseUrl + "/api/sejours/fournisseurs";
      const rep: AxiosResponse<types.Sejours> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async CreateRepas(params: types.RepasComplet) {
    try {
      const fullUrl = this.baseUrl + "/api/sejours/repas";
      const rep: AxiosResponse<types.Sejours> = await Axios.put(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async UpdateRepas(params: types.RepasComplet[] | null) {
    try {
      const fullUrl = this.baseUrl + "/api/sejours/repas";
      const rep: AxiosResponse<types.Sejours> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async DeleteRepas(params: { id: string }) {
    try {
      const fullUrl = this.baseUrl + "/api/sejours/repas";
      const rep: AxiosResponse<types.Sejours> = await Axios.delete(fullUrl, {
        params: params,
        headers: this.getHeaders()
      });
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async AssistantCreateRepas(params: types.InAssistantCreateRepass) {
    try {
      const fullUrl = this.baseUrl + "/api/sejours/assistant";
      const rep: AxiosResponse<types.Sejours> = await Axios.put(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async ResoudIngredients(params: types.InResoudIngredients) {
    try {
      const fullUrl = this.baseUrl + "/api/resolution";
      const rep: AxiosResponse<
        types.DateIngredientQuantites[] | null
      > = await Axios.post(fullUrl, params, { headers: this.getHeaders() });
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async GetFournisseurs() {
    try {
      const fullUrl = this.baseUrl + "/api/fournisseurs";
      const rep: AxiosResponse<any> = await Axios.get(fullUrl);
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async CreateFournisseur(params: types.Fournisseur) {
    try {
      const fullUrl = this.baseUrl + "/api/fournisseurs";
      const rep: AxiosResponse<any> = await Axios.put(fullUrl, params, {
        headers: this.getHeaders()
      });
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async UpdateFournisseur(params: types.Fournisseur) {
    try {
      const fullUrl = this.baseUrl + "/api/fournisseurs";
      const rep: AxiosResponse<types.Fournisseur> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async DeleteFournisseur(params: { id: string }) {
    try {
      const fullUrl = this.baseUrl + "/api/fournisseurs";
      const rep: AxiosResponse<any> = await Axios.delete(fullUrl, {
        params: params,
        headers: this.getHeaders()
      });
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async CreateLivraison(params: types.Livraison) {
    try {
      const fullUrl = this.baseUrl + "/api/livraisons";
      const rep: AxiosResponse<types.Livraison> = await Axios.put(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async UpdateLivraison(params: types.Livraison) {
    try {
      const fullUrl = this.baseUrl + "/api/livraisons";
      const rep: AxiosResponse<types.Livraison> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async DeleteLivraison(params: { id: string }) {
    try {
      const fullUrl = this.baseUrl + "/api/livraisons";
      const rep: AxiosResponse<any> = await Axios.delete(fullUrl, {
        params: params,
        headers: this.getHeaders()
      });
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async GetIngredientProduits(params: { id: string }) {
    try {
      const fullUrl = this.baseUrl + "/api/ingredient-produit";
      const rep: AxiosResponse<types.IngredientProduits> = await Axios.get(
        fullUrl,
        { params: params, headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async AjouteIngredientProduit(params: types.InAjouteIngredientProduit) {
    try {
      const fullUrl = this.baseUrl + "/api/ingredient-produit";
      const rep: AxiosResponse<types.IngredientProduits> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async SetDefautProduit(params: types.InSetDefautProduit) {
    try {
      const fullUrl = this.baseUrl + "/api/ingredient-produit-defaut";
      const rep: AxiosResponse<types.IngredientProduits> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async UpdateProduit(params: types.Produit) {
    try {
      const fullUrl = this.baseUrl + "/api/produits";
      const rep: AxiosResponse<types.Produit> = await Axios.post(
        fullUrl,
        params,
        { headers: this.getHeaders() }
      );
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async DeleteProduit(params: { id: string }) {
    try {
      const fullUrl = this.baseUrl + "/api/produits";
      const rep: AxiosResponse<any> = await Axios.delete(fullUrl, {
        params: params,
        headers: this.getHeaders()
      });
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  async EtablitCommande(params: types.InCommande) {
    try {
      const fullUrl = this.baseUrl + "/api/commande";
      const rep: AxiosResponse<any> = await Axios.post(fullUrl, params, {
        headers: this.getHeaders()
      });
      return rep.data;
    } catch (error) {
      this.handleError(error);
    }
  }
}
