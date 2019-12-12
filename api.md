# Api du server

## Identification

- /login **POST**: 
    - email: string
    - password: string 

    *returns* {valid: bool, id_utilisateur: string, token: string}

### Les url suivantes demande une authentification de la forme {username: id_utilisateur, password: token}

- /agenda **GET**:

    *returns* AgendaUtilisateur

### Ingredients

- /ingredients **PUT**

    *returns* Ingredient

- /ingredients **POST** 
    - id: int
    - nom: string 
    - categorie string    
	- callories Callories 

    // sous condition
	- unite: Unite  
    - conditionnement: Conditionnement 

    *returns* AgendaUtilisateur

- /ingredients **DELETE** // sous condition
    - id: int 
    - remove_liens_produits: bool

### Recettes

- /recettes **PUT**:

    *returns* Recette
    
- /recettes **POST**:
    -  id: int
    -  ingredients: []{id_ingredient, quantite}
    - nom: string
    - mode_emploi: string

    *returns* AgendaUtilisateur

- /recettes **DELETE**:
    - id_recette: int
