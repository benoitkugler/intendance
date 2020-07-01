// DO NOT EDIT -- autogenerated by structgen
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

// github.com/benoitkugler/intendance/server/models.Conditionnement
export interface Conditionnement {
  quantite: number;
  unite: Unite;
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

// github.com/benoitkugler/intendance/server/views.InAjouteIngredientProduit
export interface InAjouteIngredientProduit {
  id_ingredient: number;
  produit: Produit;
}

// github.com/benoitkugler/intendance/server/models.Ids
export type Ids = number[] | null;

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

// ISO date string
export type Time = string;

// github.com/benoitkugler/intendance/server/models.Categorie
export type Categorie = string;

// github.com/benoitkugler/intendance/server/models.Callories
export interface Callories {}

// github.com/benoitkugler/intendance/server/models.Ingredient
export interface Ingredient {
  id: number;
  nom: string;
  unite: Unite;
  categorie: Categorie;
  callories: Callories;
  conditionnement: Conditionnement;
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

// github.com/benoitkugler/intendance/server/controller.CommandeContraintes
export interface CommandeContraintes {
  contrainte_produits: { [key: number]: number } | null;
  regroupe: boolean;
}

// github.com/benoitkugler/intendance/server/views.InCommande
export interface InCommande {
  ingredients: DateIngredientQuantites[] | null;
  contraintes: CommandeContraintes;
}

// github.com/benoitkugler/intendance/server/views.InLieIngredientProduit
export interface InLieIngredientProduit {
  id_ingredient: number;
  id_produit: number;
}

// github.com/benoitkugler/intendance/server/views.InLoggin
export interface InLoggin {
  mail: string;
  password: string;
}

// github.com/benoitkugler/intendance/server/views.InResoudIngredients
export interface InResoudIngredients {
  mode: string;
  id_repas: number;
  nb_personnes: number;
  id_sejour: number;
  jour_offset: number[] | null;
}

// github.com/benoitkugler/intendance/server/views.InSejourFournisseurs
export interface InSejourFournisseurs {
  id_sejour: number;
  ids_fournisseurs: number[] | null;
}

// github.com/benoitkugler/intendance/server/views.InSetDefautProduit
export interface InSetDefautProduit {
  id_ingredient: number;
  id_produit: number;
  on: boolean;
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

// github.com/benoitkugler/intendance/server/views.OutCommande
export interface OutCommande {
  token: string;
  commande: CommandeItem[] | null;
  ambiguites: Ambiguites;
}

// github.com/benoitkugler/intendance/server/views.OutDeleteGroupe
export interface OutDeleteGroupe {
  token: string;
  nb_repas: number;
}

// github.com/benoitkugler/intendance/server/models.Fournisseur
export interface Fournisseur {
  id: number;
  nom: string;
  lieu: string;
}

// github.com/benoitkugler/intendance/server/views.OutFournisseur
export interface OutFournisseur {
  token: string;
  fournisseur: Fournisseur;
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

// github.com/benoitkugler/intendance/server/views.OutFournisseurs
export interface OutFournisseurs {
  token: string;
  fournisseurs: Fournisseurs;
  livraisons: Livraisons;
}

// github.com/benoitkugler/intendance/server/models.Groupe
export interface Groupe {
  id: number;
  id_sejour: number;
  nom: string;
  nb_personnes: number;
  couleur: string;
}

// github.com/benoitkugler/intendance/server/views.OutGroupe
export interface OutGroupe {
  token: string;
  groupe: Groupe;
}

// github.com/benoitkugler/intendance/server/views.OutIngredient
export interface OutIngredient {
  token: string;
  ingredient: Ingredient;
}

// github.com/benoitkugler/intendance/server/models.Set
export type Set = { [key: number]: boolean } | null;

// github.com/benoitkugler/intendance/server/controller.IngredientProduits
export interface IngredientProduits {
  produits: Produit[] | null;
  defaults: Set;
}

// github.com/benoitkugler/intendance/server/views.OutIngredientProduits
export interface OutIngredientProduits {
  token: string;
  produits: IngredientProduits;
}

// github.com/benoitkugler/intendance/server/models.Ingredients
export type Ingredients = { [key: number]: Ingredient } | null;

// github.com/benoitkugler/intendance/server/views.OutIngredients
export interface OutIngredients {
  token: string;
  ingredients: Ingredients;
}

// github.com/benoitkugler/intendance/server/views.OutLivraison
export interface OutLivraison {
  token: string;
  livraison: Livraison;
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
}

// database/sql.NullInt64
export interface NullInt64 {
  Int64: number;
  Valid: boolean;
}

// github.com/benoitkugler/intendance/server/models.Menu
export interface Menu {
  id: number;
  id_utilisateur: NullInt64;
  commentaire: string;
}

// github.com/benoitkugler/intendance/server/models.LienIngredient
export interface LienIngredient {
  id_ingredient: number;
  quantite: number;
  cuisson: string;
}

// github.com/benoitkugler/intendance/server/models.LienIngredients
export type LienIngredients = LienIngredient[] | null;

// github.com/benoitkugler/intendance/server/controller.MenuComplet
export type MenuComplet = {
  recettes: Ids;
  ingredients: LienIngredients;
} & Menu;

// github.com/benoitkugler/intendance/server/views.OutMenu
export interface OutMenu {
  token: string;
  menu: MenuComplet;
}

// github.com/benoitkugler/intendance/server/views.OutMenus
export interface OutMenus {
  token: string;
  menus: { [key: number]: MenuComplet } | null;
}

// github.com/benoitkugler/intendance/server/views.OutProduit
export interface OutProduit {
  token: string;
  produit: Produit;
}

// github.com/benoitkugler/intendance/server/models.Recette
export interface Recette {
  id: number;
  id_utilisateur: NullInt64;
  nom: string;
  mode_emploi: string;
}

// github.com/benoitkugler/intendance/server/controller.RecetteComplet
export type RecetteComplet = {
  ingredients: LienIngredients;
} & Recette;

// github.com/benoitkugler/intendance/server/views.OutRecette
export interface OutRecette {
  token: string;
  recette: RecetteComplet;
}

// github.com/benoitkugler/intendance/server/views.OutRecettes
export interface OutRecettes {
  token: string;
  recettes: { [key: number]: RecetteComplet } | null;
}

// github.com/benoitkugler/intendance/server/views.OutResoudIngredients
export interface OutResoudIngredients {
  token: string;
  date_ingredients: DateIngredientQuantites[] | null;
}

// github.com/benoitkugler/intendance/server/models.Sejour
export interface Sejour {
  id: number;
  id_utilisateur: number;
  date_debut: Time;
  nom: string;
}

// github.com/benoitkugler/intendance/server/views.OutSejour
export interface OutSejour {
  token: string;
  sejour: Sejour;
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

// github.com/benoitkugler/intendance/server/models.Groupes
export type Groupes = { [key: number]: Groupe } | null;

// github.com/benoitkugler/intendance/server/controller.Sejours
export interface Sejours {
  sejours: { [key: number]: SejourRepas } | null;
  groupes: Groupes;
}

// github.com/benoitkugler/intendance/server/views.OutSejours
export interface OutSejours {
  token: string;
  sejours: Sejours;
}

// github.com/benoitkugler/intendance/server/views.OutUtilisateurs
export interface OutUtilisateurs {
  token: string;
  utilisateurs: { [key: number]: Utilisateur } | null;
}
