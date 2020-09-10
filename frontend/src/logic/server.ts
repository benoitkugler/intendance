import Cookie from "js-cookie";
import Vue from "vue";
import {
  AbstractAPI,
  Sejours,
  Sejour,
  Groupe,
  OutDeleteGroupe,
  OutFournisseurs,
  Fournisseur,
  Livraison,
  IngredientProduits,
  Produit,
  OutCommande,
  DateIngredientQuantites,
  InResoudIngredients
} from "./api";
import { Notifications } from "./notifications";
import { Controller } from "./controller";
import {
  InLoggin,
  SejourRepas,
  Groupes,
  Ingredient,
  RecetteComplet,
  MenuComplet,
  Utilisateur,
  Fournisseurs,
  Livraisons,
  OutLoggin,
  Ingredients
} from "./api";

export const devMode = process.env.NODE_ENV != "production";
const host = devMode ? "http://localhost:1323" : window.location.origin;

export class API extends AbstractAPI {
  sejours: {
    sejours: { [key: number]: SejourRepas };
    groupes: NonNullable<Groupes>;
  } = { sejours: {}, groupes: {} };
  ingredients: { [key: number]: Ingredient } = {};
  recettes: { [key: number]: RecetteComplet } = {};
  menus: { [key: number]: MenuComplet } = {};
  utilisateurs: { [key: number]: Utilisateur } = {};
  fournisseurs: NonNullable<Fournisseurs> = {};
  livraisons: NonNullable<Livraisons> = {};

  constructor(private notifications: Notifications, token: string) {
    super(host, token, {});
  }

  handleError(err: any) {
    this.notifications.setAxiosError(err);
  }

  startRequest() {
    this.notifications.spin = true;
  }

  static Loggin(notifications: Notifications, params: InLoggin) {
    // Le serveur ignore le token sur cette requête
    return new API(notifications, "").Loggin(params);
  }

  protected onSuccessLoggin(data: OutLoggin) {}

  protected onSuccessGetUtilisateurs(
    data: { [key: number]: Utilisateur } | null
  ): void {
    this.utilisateurs = data || {};
    this.notifications.setMessage("Utilisateurs chargés.");
  }

  protected onSuccessGetIngredients(data: Ingredients): void {
    this.ingredients = data || {};
    this.notifications.setMessage("Ingrédients chargés.");
  }

  protected onSuccessCreateIngredient(data: Ingredient): void {
    Vue.set(this.ingredients || {}, data.id, data); // VRC
    this.notifications.setMessage("Ingrédient ajouté.");
  }
  protected onSuccessUpdateIngredient(data: Ingredient): void {
    Vue.set(this.ingredients || {}, data.id, data); // VRC
    this.notifications.setMessage("Ingrédient mis à jour.");
  }
  protected onSuccessDeleteIngredient(data: Ingredients): void {
    this.ingredients = data || {};
    this.notifications.setMessage("Ingrédient bien supprimé.");
  }
  protected onSuccessGetRecettes(
    data: { [key: number]: RecetteComplet } | null
  ): void {
    this.recettes = data || {};
    this.notifications.setMessage("Recettes chargées.");
  }
  protected onSuccessCreateRecette(data: RecetteComplet): void {
    Vue.set(this.recettes, data.id, data); // VRC
    this.notifications.setMessage("Recette ajoutée.");
  }
  protected onSuccessUpdateRecette(data: RecetteComplet): void {
    Vue.set(this.recettes, data.id, data); // VRC
    this.notifications.setMessage("Recette mise à jour.");
  }
  protected onSuccessDeleteRecette(
    data: { [key: number]: RecetteComplet } | null
  ): void {
    this.recettes = data || {};
    this.notifications.setMessage("Recette supprimée.");
  }
  protected onSuccessGetMenus(
    data: { [key: number]: MenuComplet } | null
  ): void {
    this.menus = data || {};
    this.notifications.setMessage("Menus chargés.");
  }
  protected onSuccessCreateMenu(data: MenuComplet): void {
    Vue.set(this.menus, data.id, data); // VRC
    this.notifications.setMessage("Menu créé.");
  }
  protected onSuccessUpdateMenu(data: MenuComplet): void {
    Vue.set(this.menus, data.id, data); // VRC
    this.notifications.setMessage("Menu mis à jour.");
  }
  protected onSuccessDeleteMenu(
    data: { [key: number]: MenuComplet } | null
  ): void {
    this.menus = data || {};
    this.notifications.setMessage("Menu supprimé.");
  }
  protected onSuccessGetSejours(data: Sejours): void {
    this.sejours.sejours = data.sejours || {};
    this.sejours.groupes = data.groupes || {};
    this.notifications.setMessage("Séjours chargés.");
  }
  protected onSuccessCreateSejour(data: Sejour): void {
    const newSejour: SejourRepas = { ...data, repass: [], fournisseurs: [] };
    Vue.set(this.sejours.sejours, newSejour.id, newSejour); // VRC
    this.notifications.setMessage("Sejour créé");
  }
  protected onSuccessUpdateSejour(data: Sejour): void {
    const oldSejour = this.sejours.sejours[data.id];
    const updatedSejour: SejourRepas = {
      ...data,
      repass: oldSejour.repass,
      fournisseurs: oldSejour.fournisseurs
    };
    this.sejours.sejours[data.id] = updatedSejour;
    this.notifications.setMessage("Sejour mis à jour.");
  }
  protected onSuccessDeleteSejour(data: Sejours): void {
    this.sejours.sejours = data.sejours || {};
    this.sejours.groupes = data.groupes || {};
    this.notifications.setMessage("Sejour supprimé.");
  }
  protected onSuccessCreateGroupe(data: Groupe): void {
    Vue.set(this.sejours.groupes, data.id, data); // VRC
    this.notifications.setMessage("Groupe créé.");
  }
  protected onSuccessUpdateGroupe(data: Groupe): void {
    Vue.set(this.sejours.groupes, data.id, data); // VRC
    this.notifications.setMessage("Groupe mis à jour.");
  }
  protected onSuccessDeleteGroupe(data: OutDeleteGroupe): void {
    Vue.delete(this.sejours.groupes || {}, data.id); // VRC
    this.notifications.setMessage(
      `Groupe supprimé, ${data.nb_repas} repas supprimé(s)`
    );
  }
  protected onSuccessUpdateSejourFournisseurs(data: Sejours): void {
    this.sejours.sejours = data.sejours || {};
    this.sejours.groupes = data.groupes || {};
    this.notifications.setMessage("Fournisseurs mis à jour.");
  }
  protected onSuccessCreateRepas(data: Sejours): void {
    this.sejours.sejours = data.sejours || {};
    this.sejours.groupes = data.groupes || {};
    this.notifications.setMessage("Repas créé.");
  }
  protected onSuccessUpdateManyRepas(data: Sejours): void {
    this.sejours.sejours = data.sejours || {};
    this.sejours.groupes = data.groupes || {};
    this.notifications.setMessage("Repas mis à jour.");
  }
  protected onSuccessDeleteRepas(data: Sejours): void {
    this.sejours.sejours = data.sejours || {};
    this.sejours.groupes = data.groupes || {};
    this.notifications.setMessage("Repas supprimé.");
  }
  protected onSuccessAssistantCreateRepas(data: Sejours): void {
    this.sejours.sejours = data.sejours || {};
    this.sejours.groupes = data.groupes || {};
    this.notifications.setMessage("Repas ajoutés.");
  }
  protected onSuccessResoudIngredients(
    data: DateIngredientQuantites[] | null
  ): void {
    this.notifications.spin = false;
  }

  protected onSuccessGetFournisseurs(data: OutFournisseurs): void {
    this.fournisseurs = data.fournisseurs || {};
    this.livraisons = data.livraisons || {};
    this.notifications.setMessage("Fournisseurs chargés.");
  }
  protected onSuccessCreateFournisseur(data: OutFournisseurs): void {
    this.fournisseurs = data.fournisseurs || {};
    this.livraisons = data.livraisons || {};
    this.notifications.setMessage("Fournisseur créé.");
  }
  protected onSuccessUpdateFournisseur(data: Fournisseur): void {
    Vue.set(this.fournisseurs, data.id, data); // VRC
    this.notifications.setMessage("Fournisseur mis à jour.");
  }
  protected onSuccessDeleteFournisseur(data: OutFournisseurs): void {
    this.fournisseurs = data.fournisseurs || {};
    this.livraisons = data.livraisons || {};
    this.notifications.setMessage("Fournisseur supprimé.");
  }
  protected onSuccessCreateLivraison(data: Livraison): void {
    Vue.set(this.livraisons, data.id, data); // VRC
    this.notifications.setMessage("Contrainte de livraison créé");
  }
  protected onSuccessUpdateLivraison(data: Livraison): void {
    Vue.set(this.livraisons, data.id, data); // VRC
    this.notifications.setMessage("Contrainte de livraison mise à jour.");
  }
  protected onSuccessDeleteLivraison(data: OutFournisseurs): void {
    this.fournisseurs = data.fournisseurs || {};
    this.livraisons = data.livraisons || {};
    this.notifications.setMessage("Contrainte de livraison supprimée.");
  }
  protected onSuccessGetIngredientProduits(data: IngredientProduits): void {
    this.notifications.setMessage("Produits chargés.");
  }
  protected onSuccessAjouteIngredientProduit(data: IngredientProduits): void {
    this.notifications.setMessage("Produit ajouté.");
  }
  protected onSuccessSetDefautProduit(data: IngredientProduits): void {
    this.notifications.setMessage("Produit choisi par défaut.");
  }
  protected onSuccessUpdateProduit(data: Produit): void {
    this.notifications.setMessage("Produit mis à jour.");
  }
  protected onSuccessDeleteProduit(data: any): void {
    this.notifications.setMessage("Produit supprimé.");
  }
  protected onSuccessEtablitCommande(data: OutCommande): void {
    this.notifications.setMessage("Commande établie.");
  }

  // charge en parallèle les données nécessaires aux menus
  async loadAllMenus() {
    this.notifications.spin = true;
    try {
      const datas = await Promise.all([
        this.rawGetIngredients(),
        this.rawGetUtilisateurs(),
        this.rawGetFournisseurs(),
        this.rawGetRecettes(), // recettes dépend des ingrédients
        this.rawGetMenus() // menus dépends des recettes, ingrédients et utilisateurs
      ]);
      this.ingredients = datas[0] || {};
      this.utilisateurs = datas[1] || {};
      this.fournisseurs = datas[2].fournisseurs || {};
      this.livraisons = datas[2].livraisons || {};
      this.recettes = datas[3] || {};
      this.menus = datas[4] || {};
      this.notifications.setMessage("Menus chargés.");
    } catch (error) {
      this.handleError(error);
    }
  }
}

export interface Meta {
  idUtilisateur: number;
  token: string;
}

export const metaDev: Meta = {
  idUtilisateur: 1,
  token:
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZFByb3ByaWV0YWlyZSI6MSwiZXhwIjoxNjAwMDE1NjcxfQ.4IiKMIVsoWNBuTq3DWMf_PfiVQPG2IMqQrJLL7hnN28"
};

/** Exposes une série d'utilitaires gérant la connection (par mot de passe ou par cookies) */
export class LogginController {
  /** Vérifie les cookies */
  static checkCookies(): Meta | null {
    const token = Cookie.get("token");
    const idUtilisateur = Cookie.get("id_utilisateur");
    if (token == undefined || idUtilisateur == undefined) {
      return null;
    }
    return { idUtilisateur: Number(idUtilisateur), token };
  }

  // renvoie un message d'erreur ou la chaine vide
  // si le mot de passe est correct.
  static async loggin(notifications: Notifications, params: InLoggin) {
    const out = await API.Loggin(notifications, params); // token is ignored here
    if (out === undefined) return; // erreur inatendue

    // save for future connections
    Cookie.set("token", out.token, { expires: out.expires });
    Cookie.set("id_utilisateur", out.utilisateur, { expires: out.expires });

    notifications.setMessage(
      `Connecté sous le nom de <b>${out.utilisateur.prenom_nom}</b>.`
    );
    return out;
  }

  static logout() {
    Cookie.remove("token");
    Cookie.remove("id_utilisateur");
  }
}

export const N = new Notifications();
