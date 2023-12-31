package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database/repository"
	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
	"github.com/durotimicodes/xanda_task_R3_D3/models"
)

type Handler struct {
	DB repository.SpaceshipRepository
}

func NewHandler(s repository.SpaceshipRepository) Handler {
	return Handler{
		DB: s,
	}
}

// Create spaceship handler
func (h Handler) CreateSpaceshipHandler(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)
	var spaceshipReq models.CreateSpaceshipRequest

	err := json.Unmarshal(body, &spaceshipReq)
	helpers.HandlerErr(err)

	spaceshipModel := helpers.ConvertRequestToModel(&spaceshipReq)
	createShip, err := h.DB.CreateSpaceship(spaceshipModel)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error: ": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	helpers.ApiResponse(createShip, w)
}
