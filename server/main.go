package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/benoitkugler/intendance/server/controller"
	"github.com/benoitkugler/intendance/server/views"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	dev := flag.Bool("dev", false, "localhost + DB dev")
	flag.Parse()

	db, err := models.ConnectDB(logs.DB_DEV)
	if err != nil {
		log.Fatal(err)
	}
	views.Server = controller.NewServer(db, *dev)

	e := echo.New()

	adress := setup(e, *dev)

	// erreurs explicites, même en production
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		err = echo.NewHTTPError(400, err.Error())
		e.DefaultHTTPErrorHandler(err, c)
	}

	e.Group("/static", middleware.Gzip()).Static("/*", "server/static")
	routes(e)

	e.Logger.Fatal(e.Start(adress))
}

func autoriseCORS(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:  append(middleware.DefaultCORSConfig.AllowMethods, http.MethodOptions),
		AllowHeaders:  []string{"Authorization", "Content-Type", "Access-Control-Allow-Origin"},
		ExposeHeaders: []string{"Content-Disposition"},
	}))
	fmt.Println("CORS activé.")
}

func setup(e *echo.Echo, dev bool) string {
	var adress string
	if dev {
		adress = "localhost:1323"
		autoriseCORS(e)
	} else {
		autoriseCORS(e) //FIXME:
		if err := views.Server.PingDB(); err != nil {
			log.Fatalf("DB not responding : %s", err)
		}
		fmt.Println("DB OK.")

		host := os.Getenv("IP")
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal("Didn't find 'PORT' env. variable.", err)
		}
		adress = fmt.Sprintf("%s:%d", host, port)
	}
	return adress
}

func routes(e *echo.Echo) {
	for _, route := range []string{
		"/",
		"/sejours",
		"/agenda",
		"/recap",
		"/menus",
		"/fournisseurs",
	} {
		e.GET(route, views.Accueil, NoCache)
	}

	e.POST("/api/loggin", views.Loggin)

	e.GET("/api/utilisateurs", views.GetUtilisateurs)

	e.GET("/api/ingredients", views.GetIngredients)
	e.PUT("/api/ingredients", views.CreateIngredient)
	e.POST("/api/ingredients", views.UpdateIngredient)
	e.DELETE("/api/ingredients", views.DeleteIngredient)

	e.GET("/api/recettes", views.GetRecettes)
	e.PUT("/api/recettes", views.CreateRecette)
	e.POST("/api/recettes", views.UpdateRecette)
	e.DELETE("/api/recettes", views.DeleteRecette)

	e.GET("/api/menus", views.GetMenus)
	e.PUT("/api/menus", views.CreateMenu)
	e.POST("/api/menus", views.UpdateMenu)
	e.DELETE("/api/menus", views.DeleteMenu)

	e.GET("/api/sejours", views.GetSejours)
	e.PUT("/api/sejours", views.CreateSejour)
	e.POST("/api/sejours", views.UpdateSejour)
	e.DELETE("/api/sejours", views.DeleteSejour)

	e.PUT("/api/groupes", views.CreateGroupe)
	e.POST("/api/groupes", views.UpdateGroupe)
	e.DELETE("/api/groupes", views.DeleteGroupe)

	e.POST("/api/sejours/fournisseurs", views.UpdateSejourFournisseurs)

	e.PUT("/api/sejours/repas", views.CreateRepas)
	e.POST("/api/sejours/repas", views.UpdateRepas)
	e.DELETE("/api/sejours/repas", views.DeleteRepas)

	e.PUT("/api/sejours/assistant", views.AssistantCreateRepas)

	e.POST("/api/resolution", views.ResoudIngredients)

	e.GET("/api/fournisseurs", views.GetFournisseurs)
	e.PUT("/api/fournisseurs", views.CreateFournisseur)
	e.POST("/api/fournisseurs", views.UpdateFournisseur)
	e.DELETE("/api/fournisseurs", views.DeleteFournisseur)

	e.PUT("/api/livraisons", views.CreateLivraison)
	e.POST("/api/livraisons", views.UpdateLivraison)
	e.DELETE("/api/livraisons", views.DeleteLivraison)

	e.GET("/api/ingredient-produit", views.GetIngredientProduits)
	e.POST("/api/ingredient-produit", views.AjouteIngredientProduit)
	e.POST("/api/ingredient-produit-defaut", views.SetDefautProduit)
	e.POST("/api/produits", views.UpdateProduit)
	e.DELETE("/api/produits", views.DeleteProduit)

	e.POST("/api/commande", views.EtablitCommande)
}

// Empêche le navigateur de mettre en cache
// pour avoir les dernières versions des fichiers statiques
// (essentiellement les builds .js)
func NoCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "no-store")
		c.Response().Header().Set("Expires", "0")
		return next(c)
	}
}
