package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
	"github.com/durotimicodes/xanda_task_R3_D3/models"
	"github.com/go-chi/chi"
)

func (h Handler) UpdateSpaceshipHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "ID")
	spaceshipID, err := strconv.Atoi(id)
	helpers.HandlerErr(err)

	body := readBody(r)
	var spaceship models.Spaceship

	err = json.Unmarshal(body, &spaceship)
	helpers.HandlerErr(err)

	updateSpaceship, err := h.repository.UpdateSpaceship(
		spaceshipID,
		spaceship.Name,
		spaceship.Class,
		spaceship.Status,
		spaceship.Crew,
		spaceship.Value,
		spaceship.Armaments,
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error: ": err.Error()})
		return
	}

	ApiResponse(updateSpaceship, w)
}
