# Api du server

## Identification

- /login **POST**:

  - email: string
  - password: string

  _returns_ {valid: bool, id_utilisateur: string, token: string}

### Les url suivantes demande une authentification de la forme {username: id_utilisateur, password: token}

### Ingredients

- /ingredients **PUT**
  { models.Ingredient }
  _returns_ Ingredient

- /ingredients **POST**
  { Ingredient } // `unite` et `conditionnement` sous condition

  _returns_ AgendaUtilisateur

- /ingredients **DELETE** // sous condition
  - id: int
  - remove_liens_produits: bool

### Recettes

- /recettes **PUT**:
  { models.Recette }
  _returns_ Recette

- /recettes **POST**:

  - recette : Recette
  - ingredients: []RecetteIngredient

  _returns_ AgendaUtilisateur

- /recettes **DELETE**:
  - id: int

### Menus

- /menus **PUT**:
  { models.Menu }
  _returns_ Menu

- /menus **POST**:

  - menu: Menu
  - recettes: []MenuRecette
  - ingredients []RecetteIngredient

  _returns_ AgendaUtilisateur

- /menus **DELETE**:
  - id: int

### Sejours

- /agenda **GET**:

  _returns_ AgendaUtilisateur

- /sejours **PUT**:
  { models.Sejour }
  _returns_ Sejour

- /sejours **POST**:
  { models.Sejour }

- /sejours \*_DELETE_:
  id: int

- /sejours/repas **PUT**:
  _returns_ models.Repas

- /sejours/repas **POST**: // modifie les champs du menu lié au séjour
  { []models.Repas }

- /sejours/repas **DELETE**:
  id: int
