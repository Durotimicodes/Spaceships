package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/durotimicodes/xanda_task_R3_D3/models"
	"github.com/go-chi/chi"
)

func HandlerErr(err error) {
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func ReadBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	HandlerErr(err)
	return body
}

func ApiResponse(resp map[string]bool, w http.ResponseWriter) {
	if resp["success"] {
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


func CreateMockSpaceship() map[string]bool {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	Id := uint(r1.Intn(20))
	Qt := strconv.Itoa(r1.Intn(100))
	Crw := r1.Intn(60)
	Val := ((rand.Float32() * 5) + 100)

	armament := models.Armament{
		ID:          Id,
		SpaceshipID: Id,
		Title:       "007 Armament",
		Qty:         Qt,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	armaments := []models.Armament{
		armament,
	}

	spaceship := models.Spaceship{
		ID:        Id,
		Name:      "007 Terminator",
		Class:     "007",
		Armaments: armaments,
		Crew:      Crw,
		Image:     "https://www.architecturaldigest.com/story/tour-this-newly-unveiled-luxury-spaceship",
		Value:     Val,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := make(map[string]bool)

	if spaceship.IsValidSpaceship() {
		result = map[string]bool{"success": true}
	}

	return result
}


func SetUpRouter() *chi.Mux {
	r := chi.NewRouter()
	return r
}
