package test

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/durotimicodes/xanda_task_R3_D3/cmd"
	mockdatabase "github.com/durotimicodes/xanda_task_R3_D3/cmd/database/mock"
	"github.com/durotimicodes/xanda_task_R3_D3/handlers"
	"github.com/durotimicodes/xanda_task_R3_D3/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSpaceship(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockSpaceshipRepository(ctrl)
	h := &handlers.Handler{
		DB: store,
	}

	//start test server and send request
	route, _ := cmd.StartApi(h)

	spaceship := generateRandomSpaceship()

	bodyJSON, err := json.Marshal(*spaceship)
	if err != nil {
		t.Fail()
	}

	t.Run("getting product by ID", func(t *testing.T) {
		id := spaceship.ID
		store.EXPECT().GetSingleSpaceship(id).Times(1).Return(spaceship, nil)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/spaceship/%d", id)
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(string(bodyJSON)))
		log.Println(string(bodyJSON))
		log.Println("real id", id)

		require.NoError(t, err)
		route.ServeHTTP(recorder, req)
		fmt.Println(recorder.Body.String())
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), string(bodyJSON))
	})

}

func generateRandomSpaceship() *models.Spaceship {

	randomNum := uint(rand.Intn(1000))
	randomFloat := rand.Float32()

	armanment := models.Armament{
		ID:          randomNum,
		SpaceshipID: randomNum,
		Title:       gofakeit.Company(),
		Qty:         strconv.Itoa(int(randomNum)),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return &models.Spaceship{
		ID:        randomNum,
		Name:      gofakeit.Name(),
		Armaments: []models.Armament{armanment},
		Crew:      int(randomNum),
		Image:     gofakeit.ImageURL(int(randomNum), int(randomNum)),
		Value:     randomFloat,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
