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

	// erreurs explicites, même en production
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		err = echo.NewHTTPError(400, err.Error())
		e.DefaultHTTPErrorHandler(err, c)
	}

	routes(e)

	var adress string
	if *dev {
		adress = "localhost:1323"
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowMethods:  append(middleware.DefaultCORSConfig.AllowMethods, http.MethodOptions),
			AllowHeaders:  []string{"Authorization", "Content-Type", "Access-Control-Allow-Origin"},
			ExposeHeaders: []string{"Content-Disposition"},
		}))
		fmt.Println("CORS activé.")
	} else {
		host := os.Getenv("IP")
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal("Didn't find 'PORT' env. variable.", err)
		}
		adress = fmt.Sprintf("%s:%d", host, port)
	}
	e.Logger.Fatal(e.Start(adress))
}

func routes(e *echo.Echo) {
	e.GET("/api/agenda", views.GetAgenda)

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
}
