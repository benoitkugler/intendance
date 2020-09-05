package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	server := views.Server{Server: controller.Server{DB: db, Dev: *dev}}

	e := echo.New()

	adress := setup(e, *dev, server.Server)

	// erreurs explicites, même en production
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		err = echo.NewHTTPError(400, err.Error())
		e.DefaultHTTPErrorHandler(err, c)
	}

	e.Group("/static", middleware.Gzip()).Static("/*", "server/static")
	routes(e, server)

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

func setup(e *echo.Echo, dev bool, s controller.Server) string {
	var adress string
	if dev {
		adress = "localhost:1323"
		autoriseCORS(e)
		id, token, err := s.GetDevToken()
		if err != nil {
			log.Fatalf("Can't get token : %s", err)
		}
		fmt.Printf("Dev: user %d, token %s\n", id, token)
	} else {
		autoriseCORS(e) //FIXME:
		if err := s.DB.Ping(); err != nil {
			log.Fatalf("DB not responding : %s", err)
		}
		fmt.Println("DB OK.")

		host := os.Getenv("IP")
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal("Didn't find 'PORT' env. variable.", err)
		}
		if strings.Count(host, ":") >= 2 { // ipV6 -> besoin de crochet
			host = "[" + host + "]"
		}
		adress = fmt.Sprintf("%s:%d", host, port)
	}
	return adress
}

func routes(e *echo.Echo, s views.Server) {
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

	e.POST("/api/loggin", s.Loggin)

	tokenMid := middleware.JWT(logs.PASSPHRASE)

	e.GET("/api/utilisateurs", s.GetUtilisateurs, tokenMid)

	e.GET("/api/ingredients", s.GetIngredients, tokenMid)
	e.PUT("/api/ingredients", s.CreateIngredient, tokenMid)
	e.POST("/api/ingredients", s.UpdateIngredient, tokenMid)
	e.DELETE("/api/ingredients", s.DeleteIngredient, tokenMid)

	e.GET("/api/recettes", s.GetRecettes, tokenMid)
	e.PUT("/api/recettes", s.CreateRecette, tokenMid)
	e.POST("/api/recettes", s.UpdateRecette, tokenMid)
	e.DELETE("/api/recettes", s.DeleteRecette, tokenMid)

	e.GET("/api/menus", s.GetMenus, tokenMid)
	e.PUT("/api/menus", s.CreateMenu, tokenMid)
	e.POST("/api/menus", s.UpdateMenu, tokenMid)
	e.DELETE("/api/menus", s.DeleteMenu, tokenMid)

	e.GET("/api/sejours", s.GetSejours, tokenMid)
	e.PUT("/api/sejours", s.CreateSejour, tokenMid)
	e.POST("/api/sejours", s.UpdateSejour, tokenMid)
	e.DELETE("/api/sejours", s.DeleteSejour, tokenMid)

	e.PUT("/api/groupes", s.CreateGroupe, tokenMid)
	e.POST("/api/groupes", s.UpdateGroupe, tokenMid)
	e.DELETE("/api/groupes", s.DeleteGroupe, tokenMid)

	e.POST("/api/sejours/fournisseurs", s.UpdateSejourFournisseurs, tokenMid)

	e.PUT("/api/sejours/repas", s.CreateRepas, tokenMid)
	e.POST("/api/sejours/repas", s.UpdateRepas, tokenMid)
	e.DELETE("/api/sejours/repas", s.DeleteRepas, tokenMid)

	e.PUT("/api/sejours/assistant", s.AssistantCreateRepas, tokenMid)

	e.POST("/api/resolution", s.ResoudIngredients, tokenMid)

	e.GET("/api/fournisseurs", s.GetFournisseurs, tokenMid)
	e.PUT("/api/fournisseurs", s.CreateFournisseur, tokenMid)
	e.POST("/api/fournisseurs", s.UpdateFournisseur, tokenMid)
	e.DELETE("/api/fournisseurs", s.DeleteFournisseur, tokenMid)

	e.PUT("/api/livraisons", s.CreateLivraison, tokenMid)
	e.POST("/api/livraisons", s.UpdateLivraison, tokenMid)
	e.DELETE("/api/livraisons", s.DeleteLivraison, tokenMid)

	e.GET("/api/ingredient-produit", s.GetIngredientProduits, tokenMid)
	e.POST("/api/ingredient-produit", s.AjouteIngredientProduit, tokenMid)
	e.POST("/api/ingredient-produit-defaut", s.SetDefautProduit, tokenMid)
	e.POST("/api/produits", s.UpdateProduit, tokenMid)
	e.DELETE("/api/produits", s.DeleteProduit, tokenMid)

	e.POST("/api/commande", s.EtablitCommande, tokenMid)
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
