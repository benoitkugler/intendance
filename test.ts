/* Do not change, this code is generated from Golang structs */

export interface Callories {}
export interface IngredientRecette {
    id: number;
    nom: string;
    unite: string;
    categorie: string;
    callories: Callories;
    quantite: number;
    cuisson: string;
}
export interface Recette {
    id: number;
    id_proprietaire: number;
    nom: string;
    mode_emploi: string;
    ingredients: IngredientRecette[] | null;
}
export interface Horaire {}
export interface Menu {
    id: number;
    id_proprietaire: number;
    commentaire: string;
    recettes: Recette[] | null;
    ingredients: IngredientRecette[] | null;
    nb_personnes: number;
    horaire: Horaire;
}
export interface Journee {
    jour_offset: number;
    menus: Menu[] | null;
}
export interface Time {}
export interface Sejour {
    id: number;
    id_proprietaire: number;
    date_debut: Time;
    nom: string;
    journees: Journee[] | null;
}
export interface AgendaUtilisateur {
    sejours: Sejour[] | null;
}
