import { Notifications } from "./notifications";
import { IngredientOptions, DetailsRepas } from "./types";
import { Formatter } from "./formatter";
import { searchFunction } from "@/components/utils/utils";
import { API, Meta } from "./server";
import {
  MenuComplet,
  Menu,
  Recette,
  Produit,
  Livraison,
  Sejour,
  RepasComplet,
  RepasGroupe,
  Groupe,
  InResoudIngredients,
  Horaire,
  New
} from "./api";

/**  Object principal de stockage des données
 * sur le client.
 * Une instance de cet objet est créé au chargement,
 * puis partagée entre les différents composants.
 * Le système de réactivité de vuejs permet de propager
 * facilement les changements effectués aux données.
 */
export class Controller {
  readonly formatter: Formatter;

  readonly api: API;

  readonly state: {
    idUtilisateur: number;
    idSejour: number | null;
  };

  constructor(meta: Meta, public readonly notifications: Notifications) {
    this.formatter = new Formatter(this);

    this.api = new API(this.notifications, meta.token);
    this.state = { idUtilisateur: meta.idUtilisateur, idSejour: null };
  }

  getAllIngredients(): IngredientOptions[] {
    return Object.values(this.api.ingredients || {}).map(ing => {
      return { ingredient: ing };
    });
  }

  getMenu(idMenu: number) {
    return (this.api.menus || {})[idMenu];
  }

  getRecette(idRecette: number) {
    return (this.api.recettes || {})[idRecette];
  }

  getIngredient(idIngredient: number) {
    return (this.api.ingredients || {})[idIngredient];
  }

  getMenuRecettes(menu: New<MenuComplet>) {
    return (menu.recettes || []).map(id => this.api.recettes[id]);
  }

  getMenuIngredients(idMenu: number): IngredientOptions[] {
    return (this.getMenu(idMenu).ingredients || []).map(ing => {
      return {
        ingredient: (this.api.ingredients || {})[ing.id_ingredient],
        options: ing
      };
    });
  }

  getMenuOrRecetteProprietaire(item: Menu | Recette) {
    if (!item.id_utilisateur.Valid) return null;
    return this.api.utilisateurs[item.id_utilisateur.Int64];
  }

  getRecetteIngredients(idRecette: number): IngredientOptions[] {
    return (this.getRecette(idRecette).ingredients || []).map(ing => {
      return {
        ingredient: (this.api.ingredients || {})[ing.id_ingredient],
        options: ing
      };
    });
  }

  getFournisseur(idFournisseur: number) {
    return (this.api.fournisseurs || {})[idFournisseur];
  }

  getLivraison(produit: Produit): Livraison {
    return (this.api.livraisons || {})[produit.id_livraison];
  }

  searchMenu(search: string) {
    const menus = Object.values(this.api.menus);
    const predicat = searchFunction(search);
    // on cherche dans le nom du menu, composé des recettes
    return menus.filter(menu => predicat(this.formatter.formatMenuName(menu)));
  }

  offsetToDate(idSejour: number, offset: number) {
    const sejour = (this.api.sejours.sejours || {})[idSejour];
    const dateDebut = new Date(sejour.date_debut);
    dateDebut.setDate(dateDebut.getDate() + offset);
    return dateDebut;
  }

  dateToOffset(idSejour: number, date: Date) {
    const sejour = (this.api.sejours.sejours || {})[idSejour];
    const dateDebut = new Date(sejour.date_debut);
    const dayMs = 24 * 60 * 60 * 1000;
    const truncateDebut = Math.floor(dateDebut.valueOf() / dayMs);
    const truncateFin = Math.floor(date.valueOf() / dayMs);
    return truncateFin - truncateDebut;
  }

  iterateAllRepas(fn: (sejour: Sejour, rep: RepasComplet) => void) {
    Object.values(this.api.sejours.sejours || {}).forEach(sejour => {
      if (!sejour.repass) return;
      sejour.repass.forEach(repas => {
        fn(sejour, repas);
      });
    });
  }

  getRepasGroupes(repas: { groupes: RepasGroupe[] | null }): Groupe[] {
    return (repas.groupes || []).map(
      rg => (this.api.sejours.groupes || {})[rg.id_groupe]
    );
  }

  getRepasNbPersonnes(repas: DetailsRepas) {
    const nb = this.getRepasGroupes(repas)
      .map(g => g.nb_personnes)
      .reduce((a, b) => a + b, repas.offset_personnes);
    return nb >= 0 ? nb : 0;
  }

  getSejour() {
    if (this.state.idSejour == null) return null;
    return (this.api.sejours.sejours || {})[this.state.idSejour] || null;
  }

  // renvoie les groupes du séjour courant
  getGroupes() {
    const idS = this.state.idSejour;
    if (idS == null) return [];
    return Object.values(this.api.sejours.groupes).filter(
      groupe => groupe.id_sejour == idS
    );
  }

  resoudIngredientsRepas(idRepas: number, nbPersonnes?: number) {
    const params: InResoudIngredients = {
      mode: "repas",
      id_repas: idRepas,
      nb_personnes: nbPersonnes == undefined ? -1 : nbPersonnes,
      id_sejour: -1, // ignoré
      jour_offset: [] // ignoré
    };
    return this.api.ResoudIngredients(params);
  }

  resoudIngredientsJournees(idSejour: number, jourOffsets: number[]) {
    const params: InResoudIngredients = {
      mode: "journees",
      id_repas: -1, // ignoré
      nb_personnes: -1, // ignoré
      id_sejour: idSejour,
      jour_offset: jourOffsets // journées données
    };
    return this.api.ResoudIngredients(params);
  }

  /** vérifie que la date est après le début du séjour et gère l'erreur
   * si la date est valide, renvoie l'offset correspondant */
  private checkOffset(sejour: Sejour, jour: Date) {
    const dateDebut = new Date(sejour.date_debut);
    if (jour < dateDebut) {
      // invalide
      this.notifications.setError({
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
    const sejour = this.api.sejours.sejours[idSejour];
    if (!sejour) return;
    const offsetTo = this.checkOffset(sejour, to);
    if (offsetTo === undefined) return;
    const offsetFrom = this.checkOffset(sejour, from);
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
    const out = await this.api.UpdateManyRepas(modifs);
    if (out === undefined) return;
    this.notifications.setMessage(
      "Les journées ont étés échangées avec succès."
    );
  }

  // modifie le moment du repas
  async deplaceRepas(repas: RepasComplet, jour: Date, horaire: Horaire) {
    const sejour = this.api.sejours.sejours[repas.id_sejour];
    if (!sejour) return;
    const offset = this.checkOffset(sejour, jour);
    if (offset === undefined) return;
    repas.jour_offset = offset;
    repas.horaire = horaire;
    const out = await this.api.UpdateManyRepas([repas]);
    if (out === undefined) return;
    this.notifications.setMessage("Le repas a été déplacé avec succès.");
  }
}
