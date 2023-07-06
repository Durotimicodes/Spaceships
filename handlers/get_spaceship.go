package handlers

import (
	"net/http"
	"strconv"

	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Get one spaceship end-point
func (h Handler) GetSpaceShipHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "ID")
	spaceshipID, err := strconv.Atoi(id)

	helpers.HandlerErr(err)

	m, err := h.repository.GetSingleSpaceship(spaceshipID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, m)

}
