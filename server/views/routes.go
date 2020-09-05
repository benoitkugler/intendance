package views

import (
	"fmt"
	"strconv"

	"github.com/avct/uasurfer"
	"github.com/benoitkugler/intendance/server/controller"
	"github.com/benoitkugler/intendance/server/models"
	"github.com/labstack/echo"
)

// Server expose l'API du serveur via des handler HTTP
type Server struct {
	controller.Server
}

// return the query parameter 'id'
func parseId(idS string) (int64, error) {
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Impossible de décrypter l'ID reçu %s : %s", idS, err)
	}
	return id, nil
}

func Accueil(c echo.Context) error {
	ua := uasurfer.Parse(c.Request().UserAgent())
	if ua.Browser.Name == uasurfer.BrowserIE && ua.Browser.Version.Major < 12 {
		return c.HTML(200, `Ce portail ne supporte pas Internet Explorer. 
			<br/> Veuillez nous excuser pour le désagrement occasioné. <br/>
			Plusieurs très bons navigateurs libres et gratuits sont disponibles (Mozilla Firefox, Google Chrome, ...).
			`)
	}
	return c.File("server/static/app/index.html")
}

// -------------------------------- Loggin --------------------------------
func (s Server) Loggin(c echo.Context) error {
	var params controller.InLoggin
	if err := c.Bind(&params); err != nil {
		return err
	}
	out, err := s.Server.Loggin(params.Mail, params.Password)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

// --------------------------------------------------------------------------
// ------------------------------ Utilisateurs ------------------------------
// --------------------------------------------------------------------------
func (s Server) GetUtilisateurs(c echo.Context) error {
	out, err := s.Server.LoadUtilisateurs()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

// --------------------------------------------------------------------------
// ------------------------------ Ingredients -------------------------------
// --------------------------------------------------------------------------

func (s Server) GetIngredients(c echo.Context) error {
	out, err := s.Server.LoadIngredients()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) CreateIngredient(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var ingredientIn models.Ingredient
	if err := c.Bind(&ingredientIn); err != nil {
		return err
	}
	newIngredient, err := ct.CreateIngredient()
	ingredientIn.Id = newIngredient.Id
	if err != nil {
		return err
	}
	ingredientIn, err = ct.UpdateIngredient(ingredientIn)
	if err != nil {
		return err
	}
	return c.JSON(200, ingredientIn)
}

func (s Server) UpdateIngredient(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var ig models.Ingredient
	if err := c.Bind(&ig); err != nil {
		return err
	}
	ig, err := ct.UpdateIngredient(ig)
	if err != nil {
		return err
	}
	return c.JSON(200, ig)
}

func (s Server) DeleteIngredient(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	idS := c.QueryParam("id")
	id, err := parseId(idS)
	if err != nil {
		return err
	}
	checkS := c.QueryParam("check_produits")
	if err = ct.DeleteIngredient(id, checkS != ""); err != nil {
		return err
	}
	out, err := s.Server.LoadIngredients()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

// --------------------------------------------------------------------------
// ------------------------------ Recettes ----------------------------------
// --------------------------------------------------------------------------

func (s Server) GetRecettes(c echo.Context) error {
	out, err := s.Server.LoadRecettes()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) CreateRecette(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var recetteIn controller.RecetteComplet
	if err := c.Bind(&recetteIn); err != nil {
		return err
	}
	newRecette, err := ct.CreateRecette()
	if err != nil {
		return err
	}
	// on utilise l'id fourni par la recette créée
	recetteIn.Id = newRecette.Id
	recetteIn, err = ct.UpdateRecette(recetteIn)
	if err != nil {
		return err
	}
	return c.JSON(200, recetteIn)
}

func (s Server) UpdateRecette(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var recette controller.RecetteComplet
	if err := c.Bind(&recette); err != nil {
		return err
	}
	recette, err := ct.UpdateRecette(recette)
	if err != nil {
		return err
	}
	return c.JSON(200, recette)
}

func (s Server) DeleteRecette(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	idS := c.QueryParam("id")
	id, err := parseId(idS)
	if err != nil {
		return err
	}
	if err = ct.DeleteRecette(id); err != nil {
		return err
	}
	out, err := s.Server.LoadRecettes()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

// --------------------------------------------------------------------------
// -------------------------------- Menus -----------------------------------
// --------------------------------------------------------------------------

func (s Server) GetMenus(c echo.Context) error {
	out, err := s.Server.LoadMenus()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) CreateMenu(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var menuIn controller.MenuComplet
	if err := c.Bind(&menuIn); err != nil {
		return err
	}
	newMenu, err := ct.CreateMenu()
	if err != nil {
		return err
	}

	// on utilise l'id fourni par le menu créé
	menuIn.Id = newMenu.Id
	menuIn, err = ct.UpdateMenu(menuIn)
	if err != nil {
		return err
	}
	return c.JSON(200, menuIn)
}

func (s Server) UpdateMenu(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var menu controller.MenuComplet
	if err := c.Bind(&menu); err != nil {
		return err
	}
	menu, err := ct.UpdateMenu(menu)
	if err != nil {
		return err
	}
	return c.JSON(200, menu)
}

func (s Server) DeleteMenu(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	idS := c.QueryParam("id")
	id, err := parseId(idS)
	if err != nil {
		return err
	}
	if err = ct.DeleteMenu(id); err != nil {
		return err
	}
	out, err := s.Server.LoadMenus()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

// --------------------------------------------------------------------------
// ---------------------------- Sejours et repas ----------------------------
// --------------------------------------------------------------------------

func (s Server) GetSejours(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	out, err := ct.LoadSejoursUtilisateur()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) CreateSejour(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var sejourIn models.Sejour
	if err := c.Bind(&sejourIn); err != nil {
		return err
	}
	newSejour, err := ct.CreateSejour()
	if err != nil {
		return err
	}
	sejourIn.Id = newSejour.Id // on garde les valeurs d'entrée
	sejourIn, err = ct.UpdateSejour(sejourIn)
	if err != nil {
		return err
	}
	return c.JSON(200, sejourIn)
}

func (s Server) UpdateSejour(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var sejour models.Sejour
	if err := c.Bind(&sejour); err != nil {
		return err
	}
	sejour, err := ct.UpdateSejour(sejour)
	if err != nil {
		return err
	}
	return c.JSON(200, sejour)
}

func (s Server) DeleteSejour(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	idS := c.QueryParam("id")
	id, err := parseId(idS)
	if err != nil {
		return err
	}
	if err = ct.DeleteSejour(id); err != nil {
		return err
	}
	out, err := ct.LoadSejoursUtilisateur()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) CreateGroupe(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var groupe models.Groupe
	if err := c.Bind(&groupe); err != nil {
		return err
	}
	newGroupe, err := ct.CreateGroupe(groupe.IdSejour)
	if err != nil {
		return err
	}
	groupe.Id = newGroupe.Id // on garde les valeurs d'entrée
	groupe, err = ct.UpdateGroupe(groupe)
	if err != nil {
		return err
	}
	return c.JSON(200, groupe)
}

func (s Server) UpdateGroupe(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var sejour models.Groupe
	if err := c.Bind(&sejour); err != nil {
		return err
	}
	sejour, err := ct.UpdateGroupe(sejour)
	if err != nil {
		return err
	}
	return c.JSON(200, sejour)
}

func (s Server) DeleteGroupe(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	idS := c.QueryParam("id")
	id, err := parseId(idS)
	if err != nil {
		return err
	}
	var out controller.OutDeleteGroupe
	if out.NbRepas, err = ct.DeleteGroupe(id); err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) CreateRepas(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var repasIn controller.RepasComplet
	if err := c.Bind(&repasIn); err != nil {
		return err
	}
	newRepas, err := ct.CreateRepas(repasIn.IdSejour)
	if err != nil {
		return err
	}
	repasIn.Id = newRepas.Id // on garde les valeurs d'entrée
	for i := range repasIn.Groupes {
		repasIn.Groupes[i].IdRepas = newRepas.Id
	}
	err = ct.UpdateManyRepas([]controller.RepasComplet{repasIn})
	if err != nil {
		return err
	}
	out, err := ct.LoadSejoursUtilisateur()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) UpdateRepas(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var repass []controller.RepasComplet
	if err := c.Bind(&repass); err != nil {
		return err
	}
	err := ct.UpdateManyRepas(repass)
	if err != nil {
		return err
	}
	out, err := ct.LoadSejoursUtilisateur()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) DeleteRepas(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	idS := c.QueryParam("id")
	id, err := parseId(idS)
	if err != nil {
		return err
	}
	if err = ct.DeleteRepas(id); err != nil {
		return err
	}
	out, err := ct.LoadSejoursUtilisateur()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

// -------------------------- Assistant de création --------------------------

func (s Server) AssistantCreateRepas(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var params controller.InAssistantCreateRepass
	if err := c.Bind(&params); err != nil {
		return err
	}
	if err := ct.InitiateRepas(params); err != nil {
		return err
	}
	out, err := ct.LoadSejoursUtilisateur()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

// --------------------------------------------------------------------------
// -------------------- Résolutions des ingrédients -------------------------
// --------------------------------------------------------------------------

func (s Server) ResoudIngredients(c echo.Context) error {
	var params controller.InResoudIngredients
	if err := c.Bind(&params); err != nil {
		return err
	}
	var (
		out []controller.DateIngredientQuantites
		err error
	)
	switch params.Mode {
	case "repas":
		var di controller.DateIngredientQuantites
		di.Ingredients, err = s.Server.ResoudIngredientsRepas(params.IdRepas, params.NbPersonnes)
		out = append(out, di)
	case "journees":
		out, err = s.Server.ResoudIngredientsJournees(params.IdSejour, params.JourOffsets)
	default:
		return fmt.Errorf("Mode de résolution inconnu : %s", params.Mode)
	}
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

// --------------------------------------------------------------------------
// ----------------------------- Produits -----------------------------------
// --------------------------------------------------------------------------

func (s Server) GetFournisseurs(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	out, err := ct.LoadFournisseurs()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) CreateFournisseur(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var fournisseur models.Fournisseur
	if err := c.Bind(&fournisseur); err != nil {
		return err
	}
	_, err := ct.CreateFournisseur(fournisseur)
	if err != nil {
		return err
	}

	out, err := ct.LoadFournisseurs()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) UpdateFournisseur(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var fournisseur models.Fournisseur
	if err := c.Bind(&fournisseur); err != nil {
		return err
	}
	fournisseur, err := ct.UpdateFournisseur(fournisseur)
	if err != nil {
		return err
	}
	return c.JSON(200, fournisseur)
}

func (s Server) DeleteFournisseur(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	idS := c.QueryParam("id")
	id, err := parseId(idS)
	if err != nil {
		return err
	}
	if err = ct.DeleteFournisseur(id); err != nil {
		return err
	}
	out, err := ct.LoadFournisseurs()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) UpdateSejourFournisseurs(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var params controller.InSejourFournisseurs
	if err := c.Bind(&params); err != nil {
		return err
	}
	if err := ct.UpdateSejourFournisseurs(params.IdSejour, params.IdsFournisseurs); err != nil {
		return err
	}
	out, err := ct.LoadSejoursUtilisateur()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) CreateLivraison(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var livraison models.Livraison
	if err := c.Bind(&livraison); err != nil {
		return err
	}
	livraison, err := ct.CreateLivraison(livraison)
	if err != nil {
		return err
	}
	return c.JSON(200, livraison)
}

func (s Server) UpdateLivraison(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var livraison models.Livraison
	if err := c.Bind(&livraison); err != nil {
		return err
	}
	livraison, err := ct.UpdateLivraison(livraison)
	if err != nil {
		return err
	}
	return c.JSON(200, livraison)
}

func (s Server) DeleteLivraison(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	idS := c.QueryParam("id")
	id, err := parseId(idS)
	if err != nil {
		return err
	}
	if err = ct.DeleteLivraison(id); err != nil {
		return err
	}
	out, err := ct.LoadFournisseurs()
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) GetIngredientProduits(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	idS := c.QueryParam("id")
	idIngredient, err := parseId(idS)
	if err != nil {
		return err
	}
	out, err := ct.GetIngredientProduits(idIngredient)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) AjouteIngredientProduit(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var params controller.InAjouteIngredientProduit
	if err := c.Bind(&params); err != nil {
		return err
	}
	_, err := ct.AjouteIngredientProduit(params.IdIngredient, params.Produit)
	if err != nil {
		return err
	}
	out, err := ct.GetIngredientProduits(params.IdIngredient)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

func (s Server) UpdateProduit(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var produit models.Produit
	if err := c.Bind(&produit); err != nil {
		return err
	}
	produit, err := ct.UpdateProduit(produit)
	if err != nil {
		return err
	}
	return c.JSON(200, produit)
}

func (s Server) DeleteProduit(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	idS := c.QueryParam("id")
	id, err := parseId(idS)
	if err != nil {
		return err
	}
	if err = ct.DeleteProduit(id); err != nil {
		return err
	}
	return c.NoContent(200)
}

func (s Server) SetDefautProduit(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var def controller.InSetDefautProduit
	if err := c.Bind(&def); err != nil {
		return err
	}
	err := ct.SetDefautProduit(def.IdIngredient, def.IdProduit, def.On)
	if err != nil {
		return err
	}

	out, err := ct.GetIngredientProduits(def.IdIngredient)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}

// --------------------------------------------------------------------------
// ----------------------------- Commandes ----------------------------------
// --------------------------------------------------------------------------

func (s Server) EtablitCommande(c echo.Context) error {
	ct := s.Server.NewRequeteContext(c)
	var params controller.InCommande
	if err := c.Bind(&params); err != nil {
		return err
	}
	out, err := ct.EtablitCommande(params.Ingredients, params.Contraintes)
	if err != nil {
		return err
	}
	return c.JSON(200, out)
}
