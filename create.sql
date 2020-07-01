-- DO NOT EDIT - autogenerated by structgen 
CREATE TABLE commandes (
	id serial PRIMARY KEY,
	id_utilisateur integer NOT NULL,
	date_emission timestamp NOT NULL,
	tag varchar NOT NULL
);

CREATE TABLE commande_produits (
	id_commande integer NOT NULL,
	id_produit integer NOT NULL,
	quantite integer NOT NULL
);

CREATE TABLE defaut_produits (
	id_utilisateur integer NOT NULL,
	id_ingredient integer NOT NULL,
	id_fournisseur integer NOT NULL,
	id_produit integer NOT NULL
);

CREATE TABLE fournisseurs (
	id serial PRIMARY KEY,
	nom varchar NOT NULL,
	lieu varchar NOT NULL
);

CREATE TABLE groupes (
	id serial PRIMARY KEY,
	id_sejour integer NOT NULL,
	nom varchar NOT NULL,
	nb_personnes integer NOT NULL,
	couleur varchar NOT NULL
);

CREATE TABLE ingredients (
	id serial PRIMARY KEY,
	nom varchar NOT NULL,
	unite varchar NOT NULL CHECK (unite IN ('Kg', 'L', 'P')),
	categorie varchar NOT NULL,
	callories jsonb NOT NULL,
	conditionnement jsonb NOT NULL
);

CREATE TABLE ingredient_produits (
	id_ingredient integer NOT NULL,
	id_produit integer NOT NULL,
	id_utilisateur integer NOT NULL
);

CREATE TABLE lien_ingredients (
	id_ingredient integer NOT NULL,
	quantite real NOT NULL,
	cuisson varchar NOT NULL
);

CREATE TABLE livraisons (
	id serial PRIMARY KEY,
	id_fournisseur integer NOT NULL,
	nom varchar NOT NULL,
	jours_livraison boolean [7] NOT NULL CHECK (array_length(jours_livraison, 1) = 7),
	delai_commande integer NOT NULL,
	anticipation integer NOT NULL
);

CREATE TABLE menus (
	id serial PRIMARY KEY,
	id_utilisateur integer,
	commentaire varchar NOT NULL
);

CREATE TABLE menu_ingredients (
	id_menu integer NOT NULL,
	id_ingredient integer NOT NULL,
	quantite real NOT NULL,
	cuisson varchar NOT NULL
);

CREATE TABLE menu_recettes (
	id_menu integer NOT NULL,
	id_recette integer NOT NULL
);

CREATE TABLE produits (
	id serial PRIMARY KEY,
	id_livraison integer NOT NULL,
	nom varchar NOT NULL,
	conditionnement jsonb NOT NULL,
	prix real NOT NULL,
	reference_fournisseur varchar NOT NULL,
	colisage integer NOT NULL
);

CREATE TABLE recettes (
	id serial PRIMARY KEY,
	id_utilisateur integer,
	nom varchar NOT NULL,
	mode_emploi varchar NOT NULL
);

CREATE TABLE recette_ingredients (
	id_recette integer NOT NULL,
	id_ingredient integer NOT NULL,
	quantite real NOT NULL,
	cuisson varchar NOT NULL
);

CREATE TABLE repass (
	id serial PRIMARY KEY,
	id_sejour integer NOT NULL,
	offset_personnes integer NOT NULL,
	jour_offset integer NOT NULL,
	horaire varchar NOT NULL CHECK (
		horaire IN ('cinquieme', 'diner', 'gouter', 'midi', 'matin')
	),
	anticipation integer NOT NULL
);

CREATE TABLE repas_groupes (
	id_repas integer NOT NULL,
	id_groupe integer NOT NULL
);

CREATE TABLE repas_ingredients (
	id_repas integer NOT NULL,
	id_ingredient integer NOT NULL,
	quantite real NOT NULL,
	cuisson varchar NOT NULL
);

CREATE TABLE repas_recettes (
	id_repas integer NOT NULL,
	id_recette integer NOT NULL
);

CREATE TABLE sejours (
	id serial PRIMARY KEY,
	id_utilisateur integer NOT NULL,
	date_debut timestamp NOT NULL,
	nom varchar NOT NULL
);

CREATE TABLE sejour_fournisseurs (
	id_utilisateur integer NOT NULL,
	id_sejour integer NOT NULL,
	id_fournisseur integer NOT NULL
);

CREATE TABLE utilisateurs (
	id serial PRIMARY KEY,
	password varchar NOT NULL,
	mail varchar NOT NULL,
	prenom_nom varchar NOT NULL
);

CREATE TABLE utilisateur_fournisseurs (
	id_utilisateur integer NOT NULL,
	id_fournisseur integer NOT NULL
);

ALTER TABLE
	commandes
ADD
	FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;

ALTER TABLE
	commande_produits
ADD
	FOREIGN KEY(id_commande) REFERENCES commandes;

ALTER TABLE
	commande_produits
ADD
	FOREIGN KEY(id_produit) REFERENCES produits;

ALTER TABLE
	defaut_produits
ADD
	FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;

ALTER TABLE
	defaut_produits
ADD
	FOREIGN KEY(id_ingredient) REFERENCES ingredients;

ALTER TABLE
	defaut_produits
ADD
	FOREIGN KEY(id_fournisseur) REFERENCES fournisseurs;

ALTER TABLE
	defaut_produits
ADD
	FOREIGN KEY(id_produit) REFERENCES produits;

ALTER TABLE
	groupes
ADD
	FOREIGN KEY(id_sejour) REFERENCES sejours;

ALTER TABLE
	ingredient_produits
ADD
	FOREIGN KEY(id_ingredient) REFERENCES ingredients;

ALTER TABLE
	ingredient_produits
ADD
	FOREIGN KEY(id_produit) REFERENCES produits;

ALTER TABLE
	ingredient_produits
ADD
	FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;

ALTER TABLE
	lien_ingredients
ADD
	FOREIGN KEY(id_ingredient) REFERENCES ingredients;

ALTER TABLE
	livraisons
ADD
	FOREIGN KEY(id_fournisseur) REFERENCES fournisseurs;

ALTER TABLE
	menus
ADD
	FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;

ALTER TABLE
	menu_ingredients
ADD
	FOREIGN KEY(id_menu) REFERENCES menus;

ALTER TABLE
	menu_ingredients
ADD
	FOREIGN KEY(id_ingredient) REFERENCES ingredients;

ALTER TABLE
	menu_recettes
ADD
	FOREIGN KEY(id_menu) REFERENCES menus;

ALTER TABLE
	menu_recettes
ADD
	FOREIGN KEY(id_recette) REFERENCES recettes;

ALTER TABLE
	produits
ADD
	FOREIGN KEY(id_livraison) REFERENCES livraisons;

ALTER TABLE
	recettes
ADD
	FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;

ALTER TABLE
	recette_ingredients
ADD
	FOREIGN KEY(id_recette) REFERENCES recettes;

ALTER TABLE
	recette_ingredients
ADD
	FOREIGN KEY(id_ingredient) REFERENCES ingredients;

ALTER TABLE
	repass
ADD
	FOREIGN KEY(id_sejour) REFERENCES sejours;

ALTER TABLE
	repas_groupes
ADD
	FOREIGN KEY(id_repas) REFERENCES repass;

ALTER TABLE
	repas_groupes
ADD
	FOREIGN KEY(id_groupe) REFERENCES groupes;

ALTER TABLE
	repas_ingredients
ADD
	FOREIGN KEY(id_repas) REFERENCES repass;

ALTER TABLE
	repas_ingredients
ADD
	FOREIGN KEY(id_ingredient) REFERENCES ingredients;

ALTER TABLE
	repas_recettes
ADD
	FOREIGN KEY(id_repas) REFERENCES repass;

ALTER TABLE
	repas_recettes
ADD
	FOREIGN KEY(id_recette) REFERENCES recettes;

ALTER TABLE
	sejours
ADD
	FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;

ALTER TABLE
	sejour_fournisseurs
ADD
	FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;

ALTER TABLE
	sejour_fournisseurs
ADD
	FOREIGN KEY(id_sejour) REFERENCES sejours;

ALTER TABLE
	sejour_fournisseurs
ADD
	FOREIGN KEY(id_fournisseur) REFERENCES fournisseurs;

ALTER TABLE
	utilisateur_fournisseurs
ADD
	FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;

ALTER TABLE
	utilisateur_fournisseurs
ADD
	FOREIGN KEY(id_fournisseur) REFERENCES fournisseurs;

ALTER TABLE
	utilisateurs UNIQUE(mail);

ALTER TABLE
	ingredients UNIQUE(nom);

ALTER TABLE
	recette_ingredients UNIQUE(id_recette, id_ingredient);

ALTER TABLE
	menu_ingredients UNIQUE(id_menu, id_ingredient);

ALTER TABLE
	menu_recettes UNIQUE(id_menu, id_recette);

ALTER TABLE
	sejours UNIQUE(id, id_utilisateur);

ALTER TABLE
	repas_ingredients UNIQUE(id_repas, id_ingredient);

ALTER TABLE
	repas_recettes UNIQUE(id_repas, id_recette);

ALTER TABLE
	repas_groupes UNIQUE(id_repas, id_groupe);

ALTER TABLE
	fournisseurs UNIQUE(nom);

ALTER TABLE
	utilisateur_fournisseurs UNIQUE(id_utilisateur, id_fournisseur);

ALTER TABLE
	sejour_fournisseurs UNIQUE(id_sejour, id_fournisseur);

ALTER TABLE
	sejour_fournisseurs FOREIGN KEY (id_utilisateur, id_sejour) REFERENCES sejours (id_utilisateur, id);

ALTER TABLE
	sejour_fournisseurs FOREIGN KEY (id_utilisateur, id_fournisseur) REFERENCES utilisateur_fournisseurs (id_utilisateur, id_fournisseur);

ALTER TABLE
	produits CHECK(prix >= 0);

ALTER TABLE
	produits UNIQUE(id_livraison, nom);

ALTER TABLE
	livraisons CHECK(anticipation >= 0);

ALTER TABLE
	livraisons CHECK(delai_commande >= 0);

ALTER TABLE
	livraisons UNIQUE(id_fournisseur, nom);

ALTER TABLE
	ingredient_produits UNIQUE(id_ingredient, id_produit);

ALTER TABLE
	defaut_produits UNIQUE(id_utilisateur, id_ingredient, id_fournisseur);

ALTER TABLE
	defaut_produits FOREIGN KEY (id_utilisateur, id_fournisseur) REFERENCES utilisateur_fournisseurs (id_utilisateur, id_fournisseur);

ALTER TABLE
	defaut_produits FOREIGN KEY (id_ingredient, id_produit) REFERENCES ingredient_produits (id_ingredient, id_produit);

ALTER TABLE
	commande_produits UNIQUE(id_commande, id_produit);