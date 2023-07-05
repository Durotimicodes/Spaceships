package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
	"github.com/durotimicodes/xanda_task_R3_D3/models"
	"github.com/durotimicodes/xanda_task_R3_D3/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h Handler) UpdateSpaceshipHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "ID")
	spaceshipID, err := strconv.Atoi(id)
	helpers.HandlerErr(err)

	body := readBody(r)
	var spaceship models.Spaceship

	err = json.Unmarshal(body, &spaceship)
	helpers.HandlerErr(err)

	//get the particular spacship
	m, err := service.GetSpaceship(spaceshipID)

	if err != nil {
		w.WriteHeader(500)
		render.JSON(w, r, err)
		return
	}

	newSpaceship := h.repository.UpdateSpaceship(
		name: spaceship.Name,
		class: spaceship.Class,
		status: spaceship.Status,
		crew: spaceship.Crew,
		value: spaceship.Value,
	)

	updateSpaceship := h.repository.UpdateSpaceship(id, newSpaceship)

	resp, err := service.UpdateSpaceship(spaceshipID, body)

	if err != nil {
		w.WriteHeader(500)
		render.JSON(w, r, err)
		return
	}

	w.WriteHeader(200)
	ApiResponse(resp, w)
}
