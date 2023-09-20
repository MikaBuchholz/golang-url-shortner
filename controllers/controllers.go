package controllers

import (
	"net/http"
	"shortner/services"
	"shortner/util"
)

type Controller struct {
	UrlShortnerService services.UrlShortnerService
}

func NewController(urlShortnerService *services.UrlShortnerService) Controller {
	return Controller{
		UrlShortnerService: *urlShortnerService,
	}
}

func (cntrl *Controller) HandleGetUrl(w http.ResponseWriter, r *http.Request) {
	data := &services.GetUrlDTO{}

	if err := data.ReadAndValidate(r); err != nil {
		http_util.RespondWithJSON(w, http_util.BadRequest(err.Error()))
		return
	}

	repsonse_dto, err := cntrl.UrlShortnerService.GetUrl(r.Context(), *data)

	if err != nil {
		http_util.RespondWithJSON(w, http_util.InternalError(err.Error()))
	}

	http_util.RespondWithJSON(w, http_util.OK(repsonse_dto))
}
func (cntrl *Controller) HandleNewUrl(w http.ResponseWriter, r *http.Request) {
	data := &services.UploadUrlDTO{}

	if err := data.ReadAndValidate(r); err != nil {
		http_util.RespondWithJSON(w, http_util.BadRequest(err.Error()))
		return
	}

	repsonse_dto, err := cntrl.UrlShortnerService.NewUrl(r.Context(), *data)

	if err != nil {
		http_util.RespondWithJSON(w, http_util.InternalError(err.Error()))
	}

	http_util.RespondWithJSON(w, http_util.OK(repsonse_dto))
}
