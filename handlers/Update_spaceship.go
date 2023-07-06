package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/durotimicodes/xanda_task_R3_D3/models"

	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
	"github.com/go-chi/chi"
)

// Update spaceship end-point
func (h Handler) UpdateSpaceshipHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "ID")
	spaceshipID, err := strconv.Atoi(id)
	helpers.HandlerErr(err)

	body := helpers.ReadBody(r)
	var spaceshipReq models.CreateSpaceshipRequest

	err = json.Unmarshal(body, &spaceshipReq)
	helpers.HandlerErr(err)

	spaceshipModel := helpers.ConvertRequestToModel(&spaceshipReq)
	updateSpaceship, err := h.repository.UpdateSpaceship(
		spaceshipID,
		spaceshipModel,
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error: ": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	helpers.ApiResponse(updateSpaceship, w)
}
