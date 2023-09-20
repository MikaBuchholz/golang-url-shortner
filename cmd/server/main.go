package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shortner/controllers"
	"shortner/db"
	"shortner/router"
	"shortner/services"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve(router_cntrl *router.RouterController) error {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: router_cntrl.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	port := os.Getenv("PORT")
	dsn := os.Getenv("DSN")

	dbConn, err := db.ConnectPostgres(dsn)

	if err != nil {
		log.Fatal("Failed DB connection")
	}

	svc := services.NewUrlShortnerService(dbConn.DB)
	controller := controllers.NewController(svc)
	router_cntrl := router.NewRouterController(controller)

	config := Config{Port: port}
	app := Application{Config: config}

	defer dbConn.DB.Close()

	err = app.Serve(router_cntrl)

	if err != nil {
		log.Fatal("Failed to serve app")
	}
}
