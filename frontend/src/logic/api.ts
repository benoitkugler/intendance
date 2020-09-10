// Code generated by apigen. DO NOT EDIT

import Axios, { AxiosResponse } from "axios";

// github.com/benoitkugler/intendance/server/controller.InLoggin
export interface InLoggin {
  mail: string;
  password: string;
}
// github.com/benoitkugler/intendance/server/controller.Utilisateur
export interface Utilisateur {
  id: number;
  prenom_nom: string;
}
// github.com/benoitkugler/intendance/server/controller.OutLoggin
export interface OutLoggin {
  erreur: string;
  token: string;
  utilisateur: Utilisateur;
  expires: number;
}
// github.com/benoitkugler/intendance/server/models.Unite
export enum Unite {
  Kilos = "Kg",
  Litres = "L",
  Piece = "P",
  Zero = ""
}

export const UniteLabels: { [key in Unite]: string } = {
  [Unite.Kilos]: "Kilos",
  [Unite.Litres]: "Litres",
  [Unite.Piece]: "Pièce(s)",
  [Unite.Zero]: "Unité invalide"
};

// github.com/benoitkugler/intendance/server/models.Categorie
export type Categorie = string;
// github.com/benoitkugler/intendance/server/models.Callories
export interface Callories {}
// github.com/benoitkugler/intendance/server/models.Conditionnement
export interface Conditionnement {
  quantite: number;
  unite: Unite;
}
// github.com/benoitkugler/intendance/server/models.Ingredient
export interface Ingredient {
  id: number;
  nom: string;
  unite: Unite;
  categorie: Categorie;
  callories: Callories;
  conditionnement: Conditionnement;
}
// github.com/benoitkugler/intendance/server/models.Ingredients
export type Ingredients = { [key: number]: Ingredient } | null;
export type New<T extends { id: number }> = Omit<T, "id"> &
  Partial<Pick<T, "id">>;
// database/sql.NullInt64
export interface NullInt64 {
  Int64: number;
  Valid: boolean;
}
// github.com/benoitkugler/intendance/server/models.Recette
export interface Recette {
  id: number;
  id_utilisateur: NullInt64;
  nom: string;
  mode_emploi: string;
}
// github.com/benoitkugler/intendance/server/models.LienIngredient
export interface LienIngredient {
  id_ingredient: number;
  quantite: number;
  cuisson: string;
}
// github.com/benoitkugler/intendance/server/models.LienIngredients
export type LienIngredients = LienIngredient[] | null;
// github.com/benoitkugler/intendance/server/controller.RecetteComplet
export type RecetteComplet = {
  ingredients: LienIngredients;
} & Recette;
// github.com/benoitkugler/intendance/server/models.Menu
export interface Menu {
  id: number;
  id_utilisateur: NullInt64;
  commentaire: string;
}
// github.com/benoitkugler/intendance/server/models.Ids
export type Ids = number[] | null;
// github.com/benoitkugler/intendance/server/controller.MenuComplet
export type MenuComplet = {
  recettes: Ids;
  ingredients: LienIngredients;
} & Menu;

class DateTag {
  private _: "D" = "D";
}

class TimeTag {
  private _: "T" = "T";
}

// AAAA-MM-YY date format
export type Date_ = string & DateTag;

// ISO date-time string
export type Time = string & TimeTag;

// github.com/benoitkugler/intendance/server/models.Sejour
export interface Sejour {
  id: number;
  id_utilisateur: number;
  date_debut: Time;
  nom: string;
}
// github.com/benoitkugler/intendance/server/models.SejourFournisseur
export interface SejourFournisseur {
  id_utilisateur: number;
  id_sejour: number;
  id_fournisseur: number;
}
// github.com/benoitkugler/intendance/server/models.Horaire
export enum Horaire {
  Cinquieme = 4,
  Diner = 3,
  Gouter = 2,
  Midi = 1,
  PetitDejeuner = 0
}

export const HoraireLabels: { [key in Horaire]: string } = {
  [Horaire.Cinquieme]: "Cinquième",
  [Horaire.Diner]: "Dîner",
  [Horaire.Gouter]: "Goûter",
  [Horaire.Midi]: "Midi",
  [Horaire.PetitDejeuner]: "Petit déjeuner"
};

// github.com/benoitkugler/intendance/server/models.Repas
export interface Repas {
  id: number;
  id_sejour: number;
  offset_personnes: number;
  jour_offset: number;
  horaire: Horaire;
  anticipation: number;
}
// github.com/benoitkugler/intendance/server/models.RepasGroupe
export interface RepasGroupe {
  id_repas: number;
  id_groupe: number;
}
// github.com/benoitkugler/intendance/server/controller.RepasComplet
export type RepasComplet = {
  groupes: RepasGroupe[] | null;
  recettes: Ids;
  ingredients: LienIngredients;
} & Repas;
// github.com/benoitkugler/intendance/server/controller.SejourRepas
export type SejourRepas = {
  fournisseurs: SejourFournisseur[] | null;
  repass: RepasComplet[] | null;
} & Sejour;
// github.com/benoitkugler/intendance/server/models.Groupe
export interface Groupe {
  id: number;
  id_sejour: number;
  nom: string;
  nb_personnes: number;
  couleur: string;
}
// github.com/benoitkugler/intendance/server/models.Groupes
export type Groupes = { [key: number]: Groupe } | null;
// github.com/benoitkugler/intendance/server/controller.Sejours
export interface Sejours {
  sejours: { [key: number]: SejourRepas } | null;
  groupes: Groupes;
}
// github.com/benoitkugler/intendance/server/controller.OutDeleteGroupe
export interface OutDeleteGroupe {
  id: number;
  nb_repas: number;
}
// github.com/benoitkugler/intendance/server/controller.InSejourFournisseurs
export interface InSejourFournisseurs {
  id_sejour: number;
  ids_fournisseurs: number[] | null;
}
// github.com/benoitkugler/intendance/server/controller.OptionsAssistantCreateRepass
export interface OptionsAssistantCreateRepass {
  duree: number;
  with_gouter: boolean;
  cinquieme: Ids;
  delete_existing: boolean;
}
// github.com/benoitkugler/intendance/server/controller.InAssistantCreateRepass
export interface InAssistantCreateRepass {
  id_sejour: number;
  options: OptionsAssistantCreateRepass;
  groupes_sorties: { [key: number]: number[] | null } | null;
}
// github.com/benoitkugler/intendance/server/controller.InResoudIngredients
export interface InResoudIngredients {
  mode: string;
  id_repas: number;
  nb_personnes: number;
  id_sejour: number;
  jour_offset: number[] | null;
}
// github.com/benoitkugler/intendance/server/controller.IngredientQuantite
export interface IngredientQuantite {
  ingredient: Ingredient;
  quantite: number;
}
// github.com/benoitkugler/intendance/server/controller.DateIngredientQuantites
export interface DateIngredientQuantites {
  date: Time;
  ingredients: IngredientQuantite[] | null;
}
// github.com/benoitkugler/intendance/server/models.Fournisseur
export interface Fournisseur {
  id: number;
  nom: string;
  lieu: string;
}
// github.com/benoitkugler/intendance/server/models.Fournisseurs
export type Fournisseurs = { [key: number]: Fournisseur } | null;
// github.com/benoitkugler/intendance/server/models.JoursLivraison
export type JoursLivraison = boolean[];
// github.com/benoitkugler/intendance/server/models.Livraison
export interface Livraison {
  id: number;
  id_fournisseur: number;
  nom: string;
  jours_livraison: JoursLivraison;
  delai_commande: number;
  anticipation: number;
}
// github.com/benoitkugler/intendance/server/models.Livraisons
export type Livraisons = { [key: number]: Livraison } | null;
// github.com/benoitkugler/intendance/server/controller.OutFournisseurs
export interface OutFournisseurs {
  fournisseurs: Fournisseurs;
  livraisons: Livraisons;
}
// github.com/benoitkugler/intendance/server/models.Produit
export interface Produit {
  id: number;
  id_livraison: number;
  nom: string;
  conditionnement: Conditionnement;
  prix: number;
  reference_fournisseur: string;
  colisage: number;
}
// github.com/benoitkugler/intendance/server/models.Set
export type Set = { [key: number]: boolean } | null;
// github.com/benoitkugler/intendance/server/controller.IngredientProduits
export interface IngredientProduits {
  produits: Produit[] | null;
  defaults: Set;
}
// github.com/benoitkugler/intendance/server/controller.InAjouteIngredientProduit
export interface InAjouteIngredientProduit {
  id_ingredient: number;
  produit: Produit;
}
// github.com/benoitkugler/intendance/server/controller.InSetDefautProduit
export interface InSetDefautProduit {
  id_ingredient: number;
  id_produit: number;
  on: boolean;
}
// github.com/benoitkugler/intendance/server/controller.CommandeContraintes
export interface CommandeContraintes {
  contrainte_produits: { [key: number]: number } | null;
  regroupe: boolean;
}
// github.com/benoitkugler/intendance/server/controller.InCommande
export interface InCommande {
  ingredients: DateIngredientQuantites[] | null;
  contraintes: CommandeContraintes;
}
// github.com/benoitkugler/intendance/server/controller.TimedIngredientQuantite
export type TimedIngredientQuantite = {
  date: Time;
} & IngredientQuantite;
// github.com/benoitkugler/intendance/server/controller.CommandeItem
export interface CommandeItem {
  produit: Produit;
  jour_commande: Time;
  quantite: number;
  origines: TimedIngredientQuantite[] | null;
}
// github.com/benoitkugler/intendance/server/controller.Ambiguites
export type Ambiguites = { [key: number]: Produit[] | null } | null;
// github.com/benoitkugler/intendance/server/controller.OutCommande
export interface OutCommande {
  commande: CommandeItem[] | null;
  ambiguites: Ambiguites;
}

export abstract class AbstractAPI {
  constructor(
    protected baseUrl: string,
    protected authToken: string,
    protected urlParams: {}
  ) {}

  abstract handleError(error: any): void;

  abstract startRequest(): void;

  getHeaders() {
    return { Authorization: "Bearer " + this.authToken };
  }

  protected async rawLoggin(params: InLoggin) {
    const fullUrl = this.baseUrl + "/api/loggin";
    const rep: AxiosResponse<OutLoggin> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawLoggin and handles the error
  async Loggin(params: InLoggin) {
    this.startRequest();
    try {
      const out = await this.rawLoggin(params);
      this.onSuccessLoggin(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessLoggin(data: OutLoggin): void;

  protected async rawGetUtilisateurs() {
    const fullUrl = this.baseUrl + "/api/utilisateurs";
    const rep: AxiosResponse<{
      [key: number]: Utilisateur;
    } | null> = await Axios.get(fullUrl, { headers: this.getHeaders() });
    return rep.data;
  }

  // wraps rawGetUtilisateurs and handles the error
  async GetUtilisateurs() {
    this.startRequest();
    try {
      const out = await this.rawGetUtilisateurs();
      this.onSuccessGetUtilisateurs(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessGetUtilisateurs(
    data: { [key: number]: Utilisateur } | null
  ): void;

  protected async rawGetIngredients() {
    const fullUrl = this.baseUrl + "/api/ingredients";
    const rep: AxiosResponse<Ingredients> = await Axios.get(fullUrl, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawGetIngredients and handles the error
  async GetIngredients() {
    this.startRequest();
    try {
      const out = await this.rawGetIngredients();
      this.onSuccessGetIngredients(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessGetIngredients(data: Ingredients): void;

  protected async rawCreateIngredient(params: New<Ingredient>) {
    const fullUrl = this.baseUrl + "/api/ingredients";
    const rep: AxiosResponse<Ingredient> = await Axios.put(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawCreateIngredient and handles the error
  async CreateIngredient(params: New<Ingredient>) {
    this.startRequest();
    try {
      const out = await this.rawCreateIngredient(params);
      this.onSuccessCreateIngredient(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessCreateIngredient(data: Ingredient): void;

  protected async rawUpdateIngredient(params: Ingredient) {
    const fullUrl = this.baseUrl + "/api/ingredients";
    const rep: AxiosResponse<Ingredient> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawUpdateIngredient and handles the error
  async UpdateIngredient(params: Ingredient) {
    this.startRequest();
    try {
      const out = await this.rawUpdateIngredient(params);
      this.onSuccessUpdateIngredient(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessUpdateIngredient(data: Ingredient): void;

  protected async rawDeleteIngredient(params: {
    id: number;
    check_produits: boolean;
  }) {
    const fullUrl = this.baseUrl + "/api/ingredients";
    const rep: AxiosResponse<Ingredients> = await Axios.delete(fullUrl, {
      params: {
        id: String(params["id"]),
        check_produits: params["check_produits"] ? "ok" : ""
      },
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawDeleteIngredient and handles the error
  async DeleteIngredient(params: { id: number; check_produits: boolean }) {
    this.startRequest();
    try {
      const out = await this.rawDeleteIngredient(params);
      this.onSuccessDeleteIngredient(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessDeleteIngredient(data: Ingredients): void;

  protected async rawGetRecettes() {
    const fullUrl = this.baseUrl + "/api/recettes";
    const rep: AxiosResponse<{
      [key: number]: RecetteComplet;
    } | null> = await Axios.get(fullUrl, { headers: this.getHeaders() });
    return rep.data;
  }

  // wraps rawGetRecettes and handles the error
  async GetRecettes() {
    this.startRequest();
    try {
      const out = await this.rawGetRecettes();
      this.onSuccessGetRecettes(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessGetRecettes(
    data: { [key: number]: RecetteComplet } | null
  ): void;

  protected async rawCreateRecette(params: New<RecetteComplet>) {
    const fullUrl = this.baseUrl + "/api/recettes";
    const rep: AxiosResponse<RecetteComplet> = await Axios.put(
      fullUrl,
      params,
      { headers: this.getHeaders() }
    );
    return rep.data;
  }

  // wraps rawCreateRecette and handles the error
  async CreateRecette(params: New<RecetteComplet>) {
    this.startRequest();
    try {
      const out = await this.rawCreateRecette(params);
      this.onSuccessCreateRecette(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessCreateRecette(data: RecetteComplet): void;

  protected async rawUpdateRecette(params: RecetteComplet) {
    const fullUrl = this.baseUrl + "/api/recettes";
    const rep: AxiosResponse<RecetteComplet> = await Axios.post(
      fullUrl,
      params,
      { headers: this.getHeaders() }
    );
    return rep.data;
  }

  // wraps rawUpdateRecette and handles the error
  async UpdateRecette(params: RecetteComplet) {
    this.startRequest();
    try {
      const out = await this.rawUpdateRecette(params);
      this.onSuccessUpdateRecette(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessUpdateRecette(data: RecetteComplet): void;

  protected async rawDeleteRecette(params: { id: number }) {
    const fullUrl = this.baseUrl + "/api/recettes";
    const rep: AxiosResponse<{
      [key: number]: RecetteComplet;
    } | null> = await Axios.delete(fullUrl, {
      params: { id: String(params["id"]) },
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawDeleteRecette and handles the error
  async DeleteRecette(params: { id: number }) {
    this.startRequest();
    try {
      const out = await this.rawDeleteRecette(params);
      this.onSuccessDeleteRecette(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessDeleteRecette(
    data: { [key: number]: RecetteComplet } | null
  ): void;

  protected async rawGetMenus() {
    const fullUrl = this.baseUrl + "/api/menus";
    const rep: AxiosResponse<{
      [key: number]: MenuComplet;
    } | null> = await Axios.get(fullUrl, { headers: this.getHeaders() });
    return rep.data;
  }

  // wraps rawGetMenus and handles the error
  async GetMenus() {
    this.startRequest();
    try {
      const out = await this.rawGetMenus();
      this.onSuccessGetMenus(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessGetMenus(
    data: { [key: number]: MenuComplet } | null
  ): void;

  protected async rawCreateMenu(params: New<MenuComplet>) {
    const fullUrl = this.baseUrl + "/api/menus";
    const rep: AxiosResponse<MenuComplet> = await Axios.put(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawCreateMenu and handles the error
  async CreateMenu(params: New<MenuComplet>) {
    this.startRequest();
    try {
      const out = await this.rawCreateMenu(params);
      this.onSuccessCreateMenu(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessCreateMenu(data: MenuComplet): void;

  protected async rawUpdateMenu(params: MenuComplet) {
    const fullUrl = this.baseUrl + "/api/menus";
    const rep: AxiosResponse<MenuComplet> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawUpdateMenu and handles the error
  async UpdateMenu(params: MenuComplet) {
    this.startRequest();
    try {
      const out = await this.rawUpdateMenu(params);
      this.onSuccessUpdateMenu(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessUpdateMenu(data: MenuComplet): void;

  protected async rawDeleteMenu(params: { id: number }) {
    const fullUrl = this.baseUrl + "/api/menus";
    const rep: AxiosResponse<{
      [key: number]: MenuComplet;
    } | null> = await Axios.delete(fullUrl, {
      params: { id: String(params["id"]) },
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawDeleteMenu and handles the error
  async DeleteMenu(params: { id: number }) {
    this.startRequest();
    try {
      const out = await this.rawDeleteMenu(params);
      this.onSuccessDeleteMenu(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessDeleteMenu(
    data: { [key: number]: MenuComplet } | null
  ): void;

  protected async rawGetSejours() {
    const fullUrl = this.baseUrl + "/api/sejours";
    const rep: AxiosResponse<Sejours> = await Axios.get(fullUrl, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawGetSejours and handles the error
  async GetSejours() {
    this.startRequest();
    try {
      const out = await this.rawGetSejours();
      this.onSuccessGetSejours(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessGetSejours(data: Sejours): void;

  protected async rawCreateSejour(params: New<Sejour>) {
    const fullUrl = this.baseUrl + "/api/sejours";
    const rep: AxiosResponse<Sejour> = await Axios.put(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawCreateSejour and handles the error
  async CreateSejour(params: New<Sejour>) {
    this.startRequest();
    try {
      const out = await this.rawCreateSejour(params);
      this.onSuccessCreateSejour(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessCreateSejour(data: Sejour): void;

  protected async rawUpdateSejour(params: Sejour) {
    const fullUrl = this.baseUrl + "/api/sejours";
    const rep: AxiosResponse<Sejour> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawUpdateSejour and handles the error
  async UpdateSejour(params: Sejour) {
    this.startRequest();
    try {
      const out = await this.rawUpdateSejour(params);
      this.onSuccessUpdateSejour(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessUpdateSejour(data: Sejour): void;

  protected async rawDeleteSejour(params: { id: number }) {
    const fullUrl = this.baseUrl + "/api/sejours";
    const rep: AxiosResponse<Sejours> = await Axios.delete(fullUrl, {
      params: { id: String(params["id"]) },
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawDeleteSejour and handles the error
  async DeleteSejour(params: { id: number }) {
    this.startRequest();
    try {
      const out = await this.rawDeleteSejour(params);
      this.onSuccessDeleteSejour(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessDeleteSejour(data: Sejours): void;

  protected async rawCreateGroupe(params: New<Groupe>) {
    const fullUrl = this.baseUrl + "/api/groupes";
    const rep: AxiosResponse<Groupe> = await Axios.put(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawCreateGroupe and handles the error
  async CreateGroupe(params: New<Groupe>) {
    this.startRequest();
    try {
      const out = await this.rawCreateGroupe(params);
      this.onSuccessCreateGroupe(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessCreateGroupe(data: Groupe): void;

  protected async rawUpdateGroupe(params: Groupe) {
    const fullUrl = this.baseUrl + "/api/groupes";
    const rep: AxiosResponse<Groupe> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawUpdateGroupe and handles the error
  async UpdateGroupe(params: Groupe) {
    this.startRequest();
    try {
      const out = await this.rawUpdateGroupe(params);
      this.onSuccessUpdateGroupe(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessUpdateGroupe(data: Groupe): void;

  protected async rawDeleteGroupe(params: { id: number }) {
    const fullUrl = this.baseUrl + "/api/groupes";
    const rep: AxiosResponse<OutDeleteGroupe> = await Axios.delete(fullUrl, {
      params: { id: String(params["id"]) },
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawDeleteGroupe and handles the error
  async DeleteGroupe(params: { id: number }) {
    this.startRequest();
    try {
      const out = await this.rawDeleteGroupe(params);
      this.onSuccessDeleteGroupe(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessDeleteGroupe(data: OutDeleteGroupe): void;

  protected async rawUpdateSejourFournisseurs(params: InSejourFournisseurs) {
    const fullUrl = this.baseUrl + "/api/sejours/fournisseurs";
    const rep: AxiosResponse<Sejours> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawUpdateSejourFournisseurs and handles the error
  async UpdateSejourFournisseurs(params: InSejourFournisseurs) {
    this.startRequest();
    try {
      const out = await this.rawUpdateSejourFournisseurs(params);
      this.onSuccessUpdateSejourFournisseurs(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessUpdateSejourFournisseurs(data: Sejours): void;

  protected async rawCreateRepas(params: New<RepasComplet>) {
    const fullUrl = this.baseUrl + "/api/sejours/repas";
    const rep: AxiosResponse<Sejours> = await Axios.put(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawCreateRepas and handles the error
  async CreateRepas(params: New<RepasComplet>) {
    this.startRequest();
    try {
      const out = await this.rawCreateRepas(params);
      this.onSuccessCreateRepas(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessCreateRepas(data: Sejours): void;

  protected async rawUpdateManyRepas(params: RepasComplet[] | null) {
    const fullUrl = this.baseUrl + "/api/sejours/repas";
    const rep: AxiosResponse<Sejours> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawUpdateManyRepas and handles the error
  async UpdateManyRepas(params: RepasComplet[] | null) {
    this.startRequest();
    try {
      const out = await this.rawUpdateManyRepas(params);
      this.onSuccessUpdateManyRepas(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessUpdateManyRepas(data: Sejours): void;

  protected async rawDeleteRepas(params: { id: number }) {
    const fullUrl = this.baseUrl + "/api/sejours/repas";
    const rep: AxiosResponse<Sejours> = await Axios.delete(fullUrl, {
      params: { id: String(params["id"]) },
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawDeleteRepas and handles the error
  async DeleteRepas(params: { id: number }) {
    this.startRequest();
    try {
      const out = await this.rawDeleteRepas(params);
      this.onSuccessDeleteRepas(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessDeleteRepas(data: Sejours): void;

  protected async rawAssistantCreateRepas(params: InAssistantCreateRepass) {
    const fullUrl = this.baseUrl + "/api/sejours/assistant";
    const rep: AxiosResponse<Sejours> = await Axios.put(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawAssistantCreateRepas and handles the error
  async AssistantCreateRepas(params: InAssistantCreateRepass) {
    this.startRequest();
    try {
      const out = await this.rawAssistantCreateRepas(params);
      this.onSuccessAssistantCreateRepas(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessAssistantCreateRepas(data: Sejours): void;

  protected async rawResoudIngredients(params: InResoudIngredients) {
    const fullUrl = this.baseUrl + "/api/resolution";
    const rep: AxiosResponse<
      DateIngredientQuantites[] | null
    > = await Axios.post(fullUrl, params, { headers: this.getHeaders() });
    return rep.data;
  }

  // wraps rawResoudIngredients and handles the error
  async ResoudIngredients(params: InResoudIngredients) {
    this.startRequest();
    try {
      const out = await this.rawResoudIngredients(params);
      this.onSuccessResoudIngredients(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessResoudIngredients(
    data: DateIngredientQuantites[] | null
  ): void;

  protected async rawGetFournisseurs() {
    const fullUrl = this.baseUrl + "/api/fournisseurs";
    const rep: AxiosResponse<OutFournisseurs> = await Axios.get(fullUrl, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawGetFournisseurs and handles the error
  async GetFournisseurs() {
    this.startRequest();
    try {
      const out = await this.rawGetFournisseurs();
      this.onSuccessGetFournisseurs(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessGetFournisseurs(data: OutFournisseurs): void;

  protected async rawCreateFournisseur(params: New<Fournisseur>) {
    const fullUrl = this.baseUrl + "/api/fournisseurs";
    const rep: AxiosResponse<OutFournisseurs> = await Axios.put(
      fullUrl,
      params,
      { headers: this.getHeaders() }
    );
    return rep.data;
  }

  // wraps rawCreateFournisseur and handles the error
  async CreateFournisseur(params: New<Fournisseur>) {
    this.startRequest();
    try {
      const out = await this.rawCreateFournisseur(params);
      this.onSuccessCreateFournisseur(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessCreateFournisseur(data: OutFournisseurs): void;

  protected async rawUpdateFournisseur(params: Fournisseur) {
    const fullUrl = this.baseUrl + "/api/fournisseurs";
    const rep: AxiosResponse<Fournisseur> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawUpdateFournisseur and handles the error
  async UpdateFournisseur(params: Fournisseur) {
    this.startRequest();
    try {
      const out = await this.rawUpdateFournisseur(params);
      this.onSuccessUpdateFournisseur(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessUpdateFournisseur(data: Fournisseur): void;

  protected async rawDeleteFournisseur(params: { id: number }) {
    const fullUrl = this.baseUrl + "/api/fournisseurs";
    const rep: AxiosResponse<OutFournisseurs> = await Axios.delete(fullUrl, {
      params: { id: String(params["id"]) },
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawDeleteFournisseur and handles the error
  async DeleteFournisseur(params: { id: number }) {
    this.startRequest();
    try {
      const out = await this.rawDeleteFournisseur(params);
      this.onSuccessDeleteFournisseur(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessDeleteFournisseur(data: OutFournisseurs): void;

  protected async rawCreateLivraison(params: New<Livraison>) {
    const fullUrl = this.baseUrl + "/api/livraisons";
    const rep: AxiosResponse<Livraison> = await Axios.put(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawCreateLivraison and handles the error
  async CreateLivraison(params: New<Livraison>) {
    this.startRequest();
    try {
      const out = await this.rawCreateLivraison(params);
      this.onSuccessCreateLivraison(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessCreateLivraison(data: Livraison): void;

  protected async rawUpdateLivraison(params: Livraison) {
    const fullUrl = this.baseUrl + "/api/livraisons";
    const rep: AxiosResponse<Livraison> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawUpdateLivraison and handles the error
  async UpdateLivraison(params: Livraison) {
    this.startRequest();
    try {
      const out = await this.rawUpdateLivraison(params);
      this.onSuccessUpdateLivraison(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessUpdateLivraison(data: Livraison): void;

  protected async rawDeleteLivraison(params: { id: number }) {
    const fullUrl = this.baseUrl + "/api/livraisons";
    const rep: AxiosResponse<OutFournisseurs> = await Axios.delete(fullUrl, {
      params: { id: String(params["id"]) },
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawDeleteLivraison and handles the error
  async DeleteLivraison(params: { id: number }) {
    this.startRequest();
    try {
      const out = await this.rawDeleteLivraison(params);
      this.onSuccessDeleteLivraison(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessDeleteLivraison(data: OutFournisseurs): void;

  protected async rawGetIngredientProduits(params: { id: number }) {
    const fullUrl = this.baseUrl + "/api/ingredient-produit";
    const rep: AxiosResponse<IngredientProduits> = await Axios.get(fullUrl, {
      params: { id: String(params["id"]) },
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawGetIngredientProduits and handles the error
  async GetIngredientProduits(params: { id: number }) {
    this.startRequest();
    try {
      const out = await this.rawGetIngredientProduits(params);
      this.onSuccessGetIngredientProduits(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessGetIngredientProduits(
    data: IngredientProduits
  ): void;

  protected async rawAjouteIngredientProduit(
    params: InAjouteIngredientProduit
  ) {
    const fullUrl = this.baseUrl + "/api/ingredient-produit";
    const rep: AxiosResponse<IngredientProduits> = await Axios.post(
      fullUrl,
      params,
      { headers: this.getHeaders() }
    );
    return rep.data;
  }

  // wraps rawAjouteIngredientProduit and handles the error
  async AjouteIngredientProduit(params: InAjouteIngredientProduit) {
    this.startRequest();
    try {
      const out = await this.rawAjouteIngredientProduit(params);
      this.onSuccessAjouteIngredientProduit(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessAjouteIngredientProduit(
    data: IngredientProduits
  ): void;

  protected async rawSetDefautProduit(params: InSetDefautProduit) {
    const fullUrl = this.baseUrl + "/api/ingredient-produit-defaut";
    const rep: AxiosResponse<IngredientProduits> = await Axios.post(
      fullUrl,
      params,
      { headers: this.getHeaders() }
    );
    return rep.data;
  }

  // wraps rawSetDefautProduit and handles the error
  async SetDefautProduit(params: InSetDefautProduit) {
    this.startRequest();
    try {
      const out = await this.rawSetDefautProduit(params);
      this.onSuccessSetDefautProduit(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessSetDefautProduit(data: IngredientProduits): void;

  protected async rawUpdateProduit(params: Produit) {
    const fullUrl = this.baseUrl + "/api/produits";
    const rep: AxiosResponse<Produit> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawUpdateProduit and handles the error
  async UpdateProduit(params: Produit) {
    this.startRequest();
    try {
      const out = await this.rawUpdateProduit(params);
      this.onSuccessUpdateProduit(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessUpdateProduit(data: Produit): void;

  protected async rawDeleteProduit(params: { id: number }) {
    const fullUrl = this.baseUrl + "/api/produits";
    const rep: AxiosResponse<any> = await Axios.delete(fullUrl, {
      params: { id: String(params["id"]) },
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawDeleteProduit and handles the error
  async DeleteProduit(params: { id: number }) {
    this.startRequest();
    try {
      const out = await this.rawDeleteProduit(params);
      this.onSuccessDeleteProduit(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessDeleteProduit(data: any): void;

  protected async rawEtablitCommande(params: InCommande) {
    const fullUrl = this.baseUrl + "/api/commande";
    const rep: AxiosResponse<OutCommande> = await Axios.post(fullUrl, params, {
      headers: this.getHeaders()
    });
    return rep.data;
  }

  // wraps rawEtablitCommande and handles the error
  async EtablitCommande(params: InCommande) {
    this.startRequest();
    try {
      const out = await this.rawEtablitCommande(params);
      this.onSuccessEtablitCommande(out);
      return out;
    } catch (error) {
      this.handleError(error);
    }
  }

  protected abstract onSuccessEtablitCommande(data: OutCommande): void;
}
