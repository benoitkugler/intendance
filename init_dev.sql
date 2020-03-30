-- à lancer après create.sql
INSERT INTO fournisseurs (nom, lieu)
    VALUES ('Pomona', 'Chamaloc');

INSERT INTO fournisseurs (nom, lieu)
    VALUES ('SUPER U', 'Roumagnan');

INSERT INTO fournisseurs (nom, lieu)
    VALUES ('Primeur', 'Chamaloc');

INSERT INTO livraisons (id_fournisseur, nom, jours_livraison, delai_commande, anticipation)
    VALUES (1, 'Fruits', '{t,t,t,f,f,t,t}', 4, 1);

INSERT INTO livraisons (id_fournisseur, nom, jours_livraison, delai_commande, anticipation)
    VALUES (1, 'viandes', '{t,t,t,f,f,t,t}', 2, 1);

INSERT INTO livraisons (id_fournisseur, nom, jours_livraison, delai_commande, anticipation)
    VALUES (2, 'Fruits', '{f,f,f,f,f,t,t}', 1, 1);

INSERT INTO livraisons (id_fournisseur, nom, jours_livraison, delai_commande, anticipation)
    VALUES (NULL, 'Pain', '{t,t,t,t,t,t,t}', 1, 0);

INSERT INTO utilisateurs (PASSWORD, mail, prenom_nom)
    VALUES ('', 'test@intendance.fr', 'test');

INSERT INTO utilisateurs (PASSWORD, mail, prenom_nom)
    VALUES ('', 'debug', 'debug');

INSERT INTO utilisateur_fournisseurs
    VALUES (2, 1);

INSERT INTO utilisateur_fournisseurs
    VALUES (2, 2);

INSERT INTO utilisateur_fournisseurs
    VALUES (2, 3);

--
-- Data for Name: produits; Type: TABLE DATA; Schema: public; Owner: intendance
--

INSERT INTO public.produits
    VALUES (1, 2, 1, 'Orange (jus)', '{"unite": "L", "quantite": 2}', 5, '', 0);

INSERT INTO public.produits
    VALUES (2, 1, 1, 'C', '{"unite": "Kg", "quantite": 2}', 0.45, '', 0);

INSERT INTO public.produits
    VALUES (3, 3, 1, 'Tomate espagne', '{"unite": "Kg", "quantite": 1}', 5, '', 0);

--
-- Data for Name: ingredients; Type: TABLE DATA; Schema: public; Owner: intendance
--

INSERT INTO public.ingredients
    VALUES (1, 'Choux fleur', 'Kg', '', '{}', '{"unite": "", "quantite": 0}');

INSERT INTO public.ingredients
    VALUES (2, 'Tomate', 'Kg', '', '{}', '{"unite": "", "quantite": 0}');

INSERT INTO public.ingredients
    VALUES (3, 'Jus d''orange', 'L', '', '{}', '{"unite": "", "quantite": 0}');

INSERT INTO public.ingredients
    VALUES (4, 'Curcuma', 'Kg', '', '{}', '{"unite": "", "quantite": 0}');

--
-- Data for Name: sejours; Type: TABLE DATA; Schema: public; Owner: intendance
--

INSERT INTO public.sejours
    VALUES (1, 2, '2020-03-25 15:45:16.563', 'C2');

--
-- Data for Name: groupes; Type: TABLE DATA; Schema: public; Owner: intendance
--

INSERT INTO public.groupes
    VALUES (1, 1, 'Marins', 6, '#3D48D1');

INSERT INTO public.groupes
    VALUES (2, 1, 'Poussins', 8, '#3DD1C5');

SELECT
    setval(pg_get_serial_sequence('fournisseurs', 'id'), coalesce(max(id), 0) + 1, FALSE)
FROM
    fournisseurs;

SELECT
    setval(pg_get_serial_sequence('livraisons', 'id'), coalesce(max(id), 0) + 1, FALSE)
FROM
    livraisons;

SELECT
    setval(pg_get_serial_sequence('utilisateurs', 'id'), coalesce(max(id), 0) + 1, FALSE)
FROM
    utilisateurs;

SELECT
    setval(pg_get_serial_sequence('produits', 'id'), coalesce(max(id), 0) + 1, FALSE)
FROM
    produits;

SELECT
    setval(pg_get_serial_sequence('ingredients', 'id'), coalesce(max(id), 0) + 1, FALSE)
FROM
    ingredients;

SELECT
    setval(pg_get_serial_sequence('sejours', 'id'), coalesce(max(id), 0) + 1, FALSE)
FROM
    sejours;

SELECT
    setval(pg_get_serial_sequence('groupes', 'id'), coalesce(max(id), 0) + 1, FALSE)
FROM
    groupes;

