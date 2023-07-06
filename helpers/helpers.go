package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/durotimicodes/xanda_task_R3_D3/models"
)

func HandlerErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func ReadBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	HandlerErr(err)

	return body
}

func ApiResponse(resp map[string]bool, w http.ResponseWriter) {
	if resp["success"] == true {
		json.NewEncoder(w).Encode(resp)
	}
}

func ConvertRequestToModel(req *models.CreateSpaceshipRequest) *models.Spaceship {
	s := &models.Spaceship{
		Name:      req.Name,
		Class:     req.Class,
		Status:    req.Status,
		Crew:      req.Crew,
		Value:     req.Value,
		Armaments: make([]models.Armament, 0),
	}

	for _, arm := range req.Armaments {
		s.Armaments = append(s.Armaments, models.Armament{
			Title: arm.Title,
			Qty:   arm.Qty,
		})
	}

	return s
}
