# Logiciel pour l'intendance en collectivité - Trame

### Objectif principal : Faciliter et automatiser le travail de l'intendance, de la conception des menus jusqu'à la commande.

## Définition des concepts

- **Ingrédient** : désigne un composant d'une recette (ex : poivron, lait, etc.) On lui associe une unité (ex : kg, L, unité, ...), nécessaire pour le calcul des quantités.

  - _Remarque: un même aliment peut se décliner en deux ingrédients, si plusieurs calibrages sont nécessaires. Par exemple, jus d'orange en brique de 25 cl (pour un goûter) / jus d'orange en vrac (petit déjeuner). Le jus d'orange en vrac n'est pas compatible avec des briques de 25cl, donc il faut deux ingrédients différents._

  - _Remarque: Dans un deuxième temps, on peut associer à chaque ingrédient une catégorie (féculent, laitage, viande, etc.) et un nombre de calories, pour analyser et proposer des menus équilibrés._

- **Recette** : une liste d'ingrédient / quantité pour une personne. Peut définir un mode d'emploi.

- **Menu** : les recettes (et ingrédients supplémentaires éventuels) se regroupent en un menu.

- **Journée**: une journée est une liste de menus, auxquels on associe un nombre de personnes, et un moment dans la journée (label ou heure approximative). L'équilibre des repas se calcule au niveau d'une journée.

Ces concepts forment une partie _abstraite_, déconnectée des fournisseurs, et _universelle_ : les recettes et menus sont directement partageables et ré-utilisables entre plusieurs intendants / séjours. Les concepts suivant sont quant à eux dépendant du contexte:

- **Séjour**: un séjour est une suite de journées (successives). Le séjour défini une date de début, qui permet de prévoir les commandes, en fonction du délai et des jours d'ouvertures des fournisseurs.

- **Fournisseur** : commerçant/magasin chez qui les produits sont commandés. Définit des jours d'ouvertures, et un délai de livraison.

- **Produit**: ce qu'on achète réellement (une ligne dans une commande). Est lié à un fournisseur, et idéalement possède une référence (qui permettrait d'éditer directement une commande). Définit aussi un conditionnement (quantité minimale qu'on peut acheter), et un prix.

- **Commande**: une liste de produit et quantité, associée à un jour (de livraison). Elle se divise ensuite entre plusieurs fournisseurs et jours de commande. Elle représente une série de commandes réelles à effectuer (ou déjà effectuées).

## Principe de fonctionnement pour l'utilisateur

L'intendant commence par définir les ingrédients / recettes / menus qu'il veut utiliser (typiquement en amont du camp). Le point clé : il **associe aussi à chaque ingrédient un produit** (donc un fournisseur).

Remarque : _Plusieurs produits peuvent être associés au même ingrédient (différents fournisseurs/marques/conditionnements). Un ingrédient est partageable entre plusieurs intendants, donc la liste des produits s'allongera au fur et à mesure, et on stocke pour chaque intendant une valeur par défaut (voir plus loin)_.

Remarque: _L'association ingrédient -> produit peut être repoussée, mais sera nécessaire pour pouvoir exporter une commande._

Il définit ensuite les dates de son séjour et forme les journées. Il valide ensuite les commandes nécessaires, puis peut exporter une liste de commandes effectives (jour par jour).

## Unités

Pour simplifier la gestion des recettes, un ingrédient défini son unité, et est utilisé ensuite en précisant un nombre qui se réfère à cette unité.
Les produits associés à un ingrédient doivent être de la même unité.
L'unité "Piece" est un cas particulier. Elle permet de ne pas vraiment spécifier d'unité, mais de renvoyer directement au produit.
Les ingrédients à la pièce peuvent définir un conditionnement, qui devra être vérifié par tout produit associé.
Exemples :

- l'ingrédient "Jus d'orange - brique 25cl" aura typiquement l'unité "Piece", et le conditionnement 0.25 L
- l'ingrédient "Paille" aura l'unité "Pièce" mais pas de conditionnement.

## Organisation des tables (base de données)

Le logiciel est disponible sous la forme d'une **application web** accessible par mail/password.

- Table **utilisateurs** : id, password, mail (optionnel : nom_prenom)

- Table **utilisateurs_fournisseurs**: id_utilisateur, id_fournisseur

- Table **produits_par_defaut** : id_utilisateur, id_ingredient, id_produit

### Partie "abstraite" :

- Table **ingredients** : id, nom, unite, (optionnel : categorie, callories)

- Table **ingredients_produits** : id_ingredient, id_produit

* Table **recettes** : id, id_proprietaire, nom, mode_emploi

- Table **recettes_ingredients** : id_recette, id_ingredient, quantite, (optionnel: cuisson)

* Table **menus** : id, id_proprietaire

* Table **menus_recettes** : id_menu, id_recette

- Table **menus_ingredients** : id_menu, id_ingredient, quantite

Les ingrédients sont _partagés_ : n'importe quel intendant peut utiliser un ingrédient déjà défini, et ajouter un produit lié. En revanche, les recettes et menus ne sont modifiables que par son _propriétaire_ (mais copiables libremement). En cas de modification par le propriétaire, les intendants utilisant la ressource sont notifiés par mail et peuvent choisir d'accepter la modification ou de s'approprier la ressource en la copiant.
Les recettes et menus peuvent n'être liés à aucun propriétaire, et sont alors éditable par tout le monde.

### Journées

Le concept de journée nécessite d'être lié à la donnée du nombre de personnes pour chaque menu. Cela ne colle pas bien avec un schéma SQL classique. De plus, une journée n'a pas vraiment d'intérêt à être partagée : la modification sur une journée entrainerait celle sur une autre, ce qui est serait plutôt déroutant.
On propose donc de ne pas utiliser de table "journée", mais de construire (dynamiquement) les journées à partir de la table _sejour_menus_ (voir ci dessous). En revanche, le concept de journée sera bien présent pour l'utilisateur, pour organiser son emploi du temps, ou pour copier des journées déjà existantes.

- Table **sejours** : id, id_proprietaire, date_debut

- Table **repas** : id, id_sejour, id_menu, nb_personnes,jour_offset, horaire (matin, midi, goûter, soir, etc...)

Les séjours sont _privés_, mais les journées formées peuvent être copiées.

### Partie "concrète" :

- Table **fournisseurs** : id, nom, delai_commande, jours_livraison

- Table **produits** : id, nom, conditionnement (quantité + unité), colisage (quantité minimal), prix (pour le conditionnement), id_fournisseur, reference_fournisseur

- Table **commandes**: id, id_proprietaire, jour_livraison

- Table **commandes_produits**: id_commande, id_produit, quantite

## Intelligence apportée par le logiciel

En plus de structurer et de rendre partageable les recettes, le logiciel automatise les points suivants.

### Calcul (dynamique) des quantités en fonction du nombre de personnes

(qui peut être pondéré en fonction de l'âge)

### Répartition des produits par jour de commande

Une fois les journées et nombre de personnes définies, on obtient (via l'association _ingrédient_ -> _produit_) une commande (plusieurs en fait). On prend maintenant en compte le délai de livraison (et les jours ouvrables) de chaque fournisseurs : on classe chaque produit nécessaire par fournisseur, puis on soustrait au jour de livraison le délai (du fournisseur) - en prenant en compte les jours ouvrables - pour obtenir une liste de produit à commander tel jour.<br/>
_Exemple :_ Si notre commande contient les trois produits ci-dessous :

- 3 x Boite tomate - Pomona
- 2 x Thon - Pomona
- 5 x Orange - Super U

nécessaires pour le jeudi 12 octobre, et que Pomona livre avec 3 jours de délai, et Super U 4 (les deux étant fermés le week-end) les deux commandes effectives sont :

- à passer lundi 9 :
  - 3 x Boite tomate - Pomona
  - 2 x Thon - Pomona
- à passer vendredi 6 :
  - 5 x Orange - Super U

### Ajustement des commandes

Admettons qu'un ajustement de menu soit nécessaire en cours de séjour. Si aucune commande n'a été passé, pas de problème - il suffit d'effacer les anciennes commandes et d'en générer une nouvelle. En revanche, si une commande (effective) a déjà été passée, on peut proposer d'évaluer la différence entre l'ancienne commande et la nouvelle, et de n'ajouter que ce qui est nécessaire. <br/>
_Remarque :_ Ce cas semble justifier l'intérêt de stocker dans la base de données les commandes (en plus de servir de mémo à l'intendant)

### _(bonus)_ Edition d'un bon de commande (pour chaque fournisseur)

Cette fonction nécessite d'enregistrer une référence fournisseur pour chaque produit. Ensuite, on peut facilement ajouter la génération d'un bon de commande au format Excel (ou même Pdf si besoin).

### _(bonus)_ calcul du coût énergétique / équilibre d'un menu

### _(bonus)_ proposition (aléatoire) de menus équilibrés
