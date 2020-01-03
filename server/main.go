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

func setup(e *echo.Echo, dev bool) string {
	var adress string
	if dev {
		adress = "localhost:1323"
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowMethods:  append(middleware.DefaultCORSConfig.AllowMethods, http.MethodOptions),
			AllowHeaders:  []string{"Authorization", "Content-Type", "Access-Control-Allow-Origin"},
			ExposeHeaders: []string{"Content-Disposition"},
		}))
		fmt.Println("CORS activé.")
	} else {
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
		"/agenda",
		"/menus",
	} {
		e.GET(route, views.Accueil)
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

	e.GET("/api/agenda", views.GetAgenda)

	e.PUT("/api/sejours", views.CreateSejour)
	e.POST("/api/sejours", views.UpdateSejour)
	e.DELETE("/api/sejours", views.DeleteSejour)

	e.PUT("/api/sejours/repas", views.CreateRepas)
	e.POST("/api/sejours/repas", views.UpdateRepas)
	e.DELETE("/api/sejours/repas", views.DeleteRepas)

	e.POST("/api/resolution", views.ResoudIngredients)
}
