package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
	"github.com/durotimicodes/xanda_task_R3_D3/models"
	"github.com/durotimicodes/xanda_task_R3_D3/service"
)

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandlerErr(err)

	return body
}

func ApiResponse(resp map[string]bool, w http.ResponseWriter) {
	if resp["success"] == true {
		json.NewEncoder(w).Encode(resp)
	}
}

type Handler struct {
	repository service.SpaceshipRepository
}

func NewHandler(s service.SpaceshipRepository) Handler {
	return Handler{
		repository: s,
	}
}

func (h Handler) CreateSpaceshipHandler(w http.ResponseWriter, r *http.Request) {

	body := readBody(r)
	var spaceship models.Spaceship
	
	err := json.Unmarshal(body, &spaceship)
	helpers.HandlerErr(err)
	
	createShip := h.repository.CreateSpaceship(
		spaceship.Name,
		spaceship.Class,
		spaceship.Status,
		spaceship.Crew,
		spaceship.Value,
	)

	ApiResponse(createShip, w)

}
