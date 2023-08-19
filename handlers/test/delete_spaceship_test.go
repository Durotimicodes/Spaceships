package test

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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

func TestDeleteSpaceship(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdatabase.NewMockSpaceshipRepository(ctrl)
	h := &handlers.Handler{
		DB: store,
	}

	//start test server and send request
	route, _ := cmd.StartApi(h)

	response := map[string]bool{"success": true}
	byteKey := []byte(fmt.Sprintf("%v", response))
	bodyJSON, err := json.Marshal(byteKey)
	if err != nil {
		t.Fail()
	}

	t.Run("delete product by ID", func(t *testing.T) {
		id := uint(rand.Intn(1000))
		store.EXPECT().DeleteSpaceship(id).Times(1).Return(response, nil)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/spaceship/delete/%d", id)
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(string(bodyJSON)))
		log.Println("THE ERROR", req)

		require.NoError(t, err)
		route.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
		assert.Contains(t, recorder.Body.String(), nil)
	})
}
