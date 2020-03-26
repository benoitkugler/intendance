-- à lancer après create.sql
INSERT INTO fournisseurs (nom, lieu, delai_commande, jours_livraison)
    VALUES ('Pomona', 'Chamaloc', 4, '{t,t,t,f,f,t,t}');

INSERT INTO fournisseurs (nom, lieu, delai_commande, jours_livraison)
    VALUES ('SUPER U', 'Roumagnan', 2, '{t,t,t,f,f,t,t}');

INSERT INTO fournisseurs (nom, lieu, delai_commande, jours_livraison)
    VALUES ('Primeur', 'Chamaloc', 1, '{f,f,f,f,f,t,t}');

INSERT INTO utilisateurs (PASSWORD, mail, prenom_nom)
    VALUES ('', 'test@intendance.fr', 'test');

INSERT INTO utilisateurs (PASSWORD, mail, prenom_nom)
    VALUES ('', 'debug', 'debug');

