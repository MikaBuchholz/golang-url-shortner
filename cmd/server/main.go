package main

import (
	"fmt"

	"log"
	"net/http"
	"shortner/controllers"
	"shortner/db"
	"shortner/router"
	"shortner/services"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Router *router.RouterController
}

func (app *Application) Serve() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: app.Router.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	envFile, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Could not find .env file")
	}

	port := envFile["PORT"]
	dsn := envFile["DSN"]

	dbConn, err := db.ConnectPostgres(dsn)

	if err != nil {
		log.Fatal("Failed DB connection")
	}

	svc := services.NewUrlShortnerService(dbConn.DB)
	controller := controllers.NewController(svc)
	router_cntrl := router.NewRouterController(controller)

	config := Config{Port: port}
	app := Application{Config: config, Router: router_cntrl}

	defer dbConn.DB.Close()

	err = app.Serve()

	if err != nil {
		log.Fatal("Failed to serve app")
	}
}
