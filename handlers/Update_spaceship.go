package handlers

import (
	"net/http"
	"strconv"

	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
	"github.com/durotimicodes/xanda_task_R3_D3/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h Handler) UpdateSpaceshipHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "ID")
	spaceshipID, err := strconv.Atoi(id)
	helpers.HandlerErr(err)

	body, err := service.GetSpaceship(spaceshipID)
	if err != nil {
		w.WriteHeader(500)
		render.JSON(w, r, err)
		return
	}

	resp, err := service.UpdateSpaceship(spaceshipID, body)

	if err != nil {
		w.WriteHeader(500)
		render.JSON(w, r, err)
		return
	}

	w.WriteHeader(200)
	ApiResponse(resp, w)
}
