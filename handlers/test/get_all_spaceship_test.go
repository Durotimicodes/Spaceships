package test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/durotimicodes/xanda_task_R3_D3/cmd"
	mockdatabase "github.com/durotimicodes/xanda_task_R3_D3/cmd/database/mock"
	"github.com/durotimicodes/xanda_task_R3_D3/handlers"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllSpaceship(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockSpaceshipRepository(ctrl)
	h := &handlers.Handler{
		DB: store,
	}

	//start test server and send request
	route, _ := cmd.StartApi(h)

	spaceship, err := generateAllSpaceship()
	
	byteKey := []byte(fmt.Sprintf("%v", spaceship))
	bodyJSON, err := json.Marshal(byteKey)
	if err != nil {
		t.Fail()
	}

	t.Run("getting products", func(t *testing.T) {
		store.EXPECT().GetAll().Times(1).Return(spaceship, nil)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/spaceships")
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(string(bodyJSON)))

		require.NoError(t, err)
		route.ServeHTTP(recorder, req)
		fmt.Println(recorder.Body.String())
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "name")
	})

}

func generateAllSpaceship() (map[string]interface{}, error) {

	spaceship := generateRandomSpaceship()

	var err error
	data := map[string]any{"data": spaceship}
	if err != nil {
		log.Println("unable to get data", err)
		return nil, err
	}

	return data, nil
}
