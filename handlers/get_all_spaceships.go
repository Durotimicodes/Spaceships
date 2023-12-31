package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// Get all Spaceship handler and filter by either name, class or status
func (h Handler) GetAllSpaceShipsHandler(w http.ResponseWriter, r *http.Request) {
	var resp interface{}
	var err error

	urlQueryParams := r.URL.Query()

	if len(urlQueryParams) == 0 {
		resp, err = h.DB.GetAll()
		w.WriteHeader(http.StatusOK)
	}

	if nameQueryParams := urlQueryParams.Get("name"); nameQueryParams != "" {
		resp, err = h.DB.FilterAllByName(nameQueryParams)
	}

	if classQueryParams := urlQueryParams.Get("class"); classQueryParams != "" {
		resp, err = h.DB.FilterAllByClass(classQueryParams)
	}
	if statusQueryParams := urlQueryParams.Get("status"); statusQueryParams != "" {
		resp, err = h.DB.FilterAllByName(statusQueryParams)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	render.JSON(w, r, resp)
	w.WriteHeader(http.StatusOK)

}
