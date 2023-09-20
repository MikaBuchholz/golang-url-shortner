package router

import (
	"net/http"

	"shortner/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const BASE_API_URL = "api"
const VERSION = "v1"
const BASE_ROUTE = "/" + BASE_API_URL + "/" + VERSION

type RouterController struct {
	Controller controllers.Controller
}

func NewRouterController(cntrl controllers.Controller) *RouterController {
	return &RouterController{
		Controller: cntrl,
	}
}

func (router_cntrl *RouterController) Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get(BASE_ROUTE+"/url/{id}", router_cntrl.Controller.HandleGetUrl)
	router.Post(BASE_ROUTE+"/url/new", router_cntrl.Controller.HandleNewUrl)

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("works"))
	})

	return router
}
