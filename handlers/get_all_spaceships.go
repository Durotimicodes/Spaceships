package handlers

import (
	"net/http"

	"github.com/durotimicodes/xanda_task_R3_D3/service"
	"github.com/go-chi/render"
)


// Assumption: That we can only apply a single filter at any point in time
func (h Handler) GetAllSpaceShipsHandler(w http.ResponseWriter, r *http.Request) {
	var m interface{}
	var err error

	urlQueryParams := r.URL.Query()

	if len(urlQueryParams) == 0 {
		m, err = service.GetAllSpaceShips()
		w.WriteHeader(200)
	}

	if nameQueryParams := urlQueryParams.Get("name"); nameQueryParams != "" {
		m, err = service.GetAllSpaceShipsByName(nameQueryParams)
	}

	if classQueryParams := urlQueryParams.Get("class"); classQueryParams != "" {
		m, err = service.GetAllSpaceShipsByClass(classQueryParams)

	}
	if statusQueryParams := urlQueryParams.Get("status"); statusQueryParams != "" {
		m, err = service.GetAllSpaceShipsByStatus(statusQueryParams)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		render.JSON(w, r, err)
		return
	}

	w.WriteHeader(200)
	render.JSON(w, r, m)
}


