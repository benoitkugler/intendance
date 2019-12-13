// this file was automatically generated, DO NOT EDIT
// structs
// struct2ts:github.com/benoitkugler/intendance/server/datamodel.Callories
export interface Callories {
}

// struct2ts:github.com/benoitkugler/intendance/server/datamodel.Conditionnement
export interface Conditionnement {
	quantite: number;
	unite: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.IngredientRecette
export interface IngredientRecette {
	id: number;
	nom: string;
	unite: string;
	categorie: string;
	callories: Callories;
	conditionnement: Conditionnement;
	quantite: number;
	cuisson: string;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Recette
export interface Recette {
	id: number;
	id_proprietaire: number;
	nom: string;
	mode_emploi: string;
	ingredients: IngredientRecette[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/datamodel.Horaire
export interface Horaire {
	heure: number;
	minute: number;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Menu
export interface Menu {
	id: number;
	id_proprietaire: number;
	commentaire: string;
	recettes: Recette[] | null;
	ingredients: IngredientRecette[] | null;
	nb_personnes: number;
	horaire: Horaire;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Journee
export interface Journee {
	jour_offset: number;
	menus: Menu[] | null;
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.Sejour
export interface Sejour {
	id: number;
	id_proprietaire: number;
	date_debut: Date;
	nom: string;
	journees: { [key: number]: Journee };
}

// struct2ts:github.com/benoitkugler/intendance/server/controller.AgendaUtilisateur
export interface AgendaUtilisateur {
	sejours: Sejour[] | null;
}

