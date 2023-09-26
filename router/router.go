package router

import (
	"net/http"
	"time"

	"shortner/controllers"
	"shortner/util"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
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

	router.Use(httprate.LimitByIP(100, 1*time.Minute))
	router.Get(BASE_ROUTE+"/url/{id}", router_cntrl.Controller.HandleGetUrl)
	router.Post(BASE_ROUTE+"/url/new", router_cntrl.Controller.HandleNewUrl)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		http_util.RespondWithJSON(w, http_util.OK(struct{ Status string }{Status: "OK"}))
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http_util.RespondWithJSON(w, http_util.OK("<3"))
	})

	return router
}
