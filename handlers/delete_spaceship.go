package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
	"github.com/go-chi/chi"
)

// Delete spaceship end-point
func (h Handler) DeleteSpaceshipHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	spaceshipID, err := strconv.Atoi(id)

	helpers.HandlerErr(err)

	resp, err := h.repository.DeleteSpaceship(spaceshipID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error: ": err.Error()})
		return
	}

	w.WriteHeader(200)
	helpers.ApiResponse(resp, w)
}
