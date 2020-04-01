# Logiciel pour l'intendance en collectivité - Trame

### Objectif principal : Faciliter et automatiser le travail de l'intendance, de la conception des menus jusqu'à la commande.

## Définition des concepts

- **Ingrédient** : désigne un composant d'une recette (ex : poivron, lait, etc.) On lui associe une unité (ex : kg, L, unité, ...), nécessaire pour le calcul des quantités.

  - _Remarque: un même aliment peut se décliner en deux ingrédients, si plusieurs calibrages sont nécessaires. Par exemple, jus d'orange en brique de 25 cl (pour un goûter) / jus d'orange en vrac (petit déjeuner). Le jus d'orange en vrac n'est pas compatible avec des briques de 25cl, donc il faut deux ingrédients différents._

  - _Remarque: Dans un deuxième temps, on peut associer à chaque ingrédient une catégorie (féculent, laitage, viande, etc.) et un nombre de calories, pour analyser et proposer des menus équilibrés._

- **Recette** : une liste d'ingrédient / quantité pour une personne. Peut définir un mode d'emploi.

- **Menu** : les recettes (et ingrédients supplémentaires éventuels) se regroupent en un menu.

Ces concepts forment une partie _abstraite_, déconnectée des fournisseurs, et _universelle_ : les recettes et menus sont directement partageables et ré-utilisables entre plusieurs intendants / séjours. Les concepts suivant sont quant à eux dépendant du contexte:

- **Séjour**: un séjour est une suite de journées (successives). Le séjour défini une date de début, qui permet de prévoir les commandes, en fonction du délai et des jours d'ouvertures des fournisseurs. Un séjour peut définir une liste de _groupes_ (nom et nombre de personnes) utiles dans la définition des repas.

- **Repas**: un repas est lié à un séjour. Il contient une liste de groupes, ainsi qu'un nombre de personnes bonus. Il définit aussi le menu utilisé pour ce repas.

- **Fournisseur** : commerçant/magasin chez qui les produits sont commandés. Définit des jours d'ouvertures, et un délai de livraison.

- **Produit**: ce qu'on achète réellement (une ligne dans une commande). Est lié à un fournisseur, et idéalement possède une référence (qui permettrait d'éditer directement une commande). Définit aussi un conditionnement (quantité minimale qu'on peut acheter), et un prix.

- **Commande**: une liste de produit et quantité, associée à un jour (de livraison). Elle se divise ensuite entre plusieurs fournisseurs et jours de commande. Elle représente une série de commandes réelles à effectuer (ou déjà effectuées).

## Principe de fonctionnement pour l'utilisateur

L'intendant commence par définir les ingrédients / recettes / menus qu'il veut utiliser (typiquement en amont du camp). Le point clé : il **associe aussi à chaque ingrédient un produit** (donc un fournisseur).

Remarque : _Plusieurs produits peuvent être associés au même ingrédient (différents fournisseurs/marques/conditionnements). Un ingrédient est partageable entre plusieurs intendants, donc la liste des produits s'allongera au fur et à mesure, et on stocke pour chaque intendant une valeur par défaut (voir plus loin)_.

Remarque: _L'association ingrédient -> produit peut être repoussée, mais sera nécessaire pour pouvoir exporter une commande._

Il définit ensuite les dates de son séjour, les groupes du séjour et forme les journées. Il valide ensuite les commandes nécessaires, puis peut exporter une liste de commandes effectives (jour par jour).

## Unités

Pour simplifier la gestion des recettes, un ingrédient défini son unité, et est utilisé ensuite en précisant un nombre qui se réfère à cette unité.
Les produits associés à un ingrédient doivent être de la même unité.
L'unité "Piece" est un cas particulier. Elle permet de ne pas vraiment spécifier d'unité, mais de renvoyer directement au produit.
Les ingrédients à la pièce peuvent définir un conditionnement, qui devra être vérifié par tout produit associé.
Exemples :

- l'ingrédient "Jus d'orange - brique 25cl" aura typiquement l'unité "Piece", et le conditionnement 0.25 L
- l'ingrédient "Paille" aura l'unité "Pièce" mais pas de conditionnement.

## Organisation des tables (base de données)

Voir la documentation du package `models`.

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
