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

CREATE TABLE fournisseurs (
	id serial PRIMARY KEY,
	nom varchar NOT NULL,
	delai_commande integer NOT NULL,
	jours_livraison boolean[7] NOT NULL CHECK (array_length(jours_livraison, 1) = 7)
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
	unite varchar NOT NULL,
	categorie varchar NOT NULL,
	callories jsonb NOT NULL,
	conditionnement jsonb NOT NULL
);

CREATE TABLE ingredient_produits (
	id_ingredient integer NOT NULL,
	id_produit integer NOT NULL,
	id_utilisateur integer NOT NULL
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
	id_fournisseur integer NOT NULL,
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
	id_menu integer,
	offset_personnes integer NOT NULL,
	jour_offset integer NOT NULL,
	horaire varchar NOT NULL
);

CREATE TABLE repas_groupes (
	id_repas integer NOT NULL,
	id_groupe integer NOT NULL
);

CREATE TABLE sejours (
	id serial PRIMARY KEY,
	id_utilisateur integer NOT NULL,
	date_debut timestamp NOT NULL,
	nom varchar NOT NULL
);

CREATE TABLE utilisateurs (
	id serial PRIMARY KEY,
	password varchar NOT NULL,
	mail varchar NOT NULL,
	prenom_nom varchar NOT NULL
);

ALTER TABLE commandes ADD FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;
ALTER TABLE groupes ADD FOREIGN KEY(id_sejour) REFERENCES sejours;
ALTER TABLE ingredient_produits ADD FOREIGN KEY(id_ingredient) REFERENCES ingredients;
ALTER TABLE ingredient_produits ADD FOREIGN KEY(id_produit) REFERENCES produits;
ALTER TABLE menu_ingredients ADD FOREIGN KEY(id_menu) REFERENCES menus;
ALTER TABLE repas_groupes ADD FOREIGN KEY(id_repas) REFERENCES repass;
ALTER TABLE repass ADD FOREIGN KEY(id_sejour) REFERENCES sejours;
ALTER TABLE menu_recettes ADD UNIQUE(id_menu, id_recette);
ALTER TABLE repass ADD FOREIGN KEY(id_menu) REFERENCES menus;
ALTER TABLE repas_groupes ADD UNIQUE(id_repas, id_groupe);
ALTER TABLE menu_recettes ADD FOREIGN KEY(id_recette) REFERENCES recettes;
ALTER TABLE produits ADD FOREIGN KEY(id_fournisseur) REFERENCES fournisseurs;
ALTER TABLE repas_groupes ADD FOREIGN KEY(id_groupe) REFERENCES groupes;
ALTER TABLE ingredients ADD UNIQUE(nom);
ALTER TABLE recettes ADD FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;
ALTER TABLE recette_ingredients ADD FOREIGN KEY(id_ingredient) REFERENCES ingredients;
ALTER TABLE menu_ingredients ADD UNIQUE(id_menu, id_ingredient);
ALTER TABLE utilisateurs ADD UNIQUE(mail);
ALTER TABLE recette_ingredients ADD UNIQUE(id_recette, id_ingredient);
ALTER TABLE commande_produits ADD FOREIGN KEY(id_produit) REFERENCES produits;
ALTER TABLE menus ADD FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;
ALTER TABLE menu_ingredients ADD FOREIGN KEY(id_ingredient) REFERENCES ingredients;
ALTER TABLE menu_recettes ADD FOREIGN KEY(id_menu) REFERENCES menus;
ALTER TABLE recette_ingredients ADD FOREIGN KEY(id_recette) REFERENCES recettes;
ALTER TABLE sejours ADD FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;
ALTER TABLE fournisseurs ADD UNIQUE(nom);
ALTER TABLE ingredient_produits ADD UNIQUE(id_ingredient, id_produit);
ALTER TABLE commande_produits ADD FOREIGN KEY(id_commande) REFERENCES commandes;
ALTER TABLE ingredient_produits ADD FOREIGN KEY(id_utilisateur) REFERENCES utilisateurs;
ALTER TABLE commande_produits ADD UNIQUE(id_commande, id_produit);
