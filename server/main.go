package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/benoitkugler/intendance/server/controller"
	"github.com/benoitkugler/intendance/server/views"

	"github.com/benoitkugler/intendance/logs"
	"github.com/benoitkugler/intendance/server/models"
	"github.com/labstack/echo"
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
	e.Debug = *dev
	routes(e)

	var adress string
	if *dev {
		adress = "localhost:1323"
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
}
