package test

import "testing"

// "bytes"
// "encoding/json"
// "net/http"
// "net/http/httptest"
// "testing"

// "github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
// "github.com/durotimicodes/xanda_task_R3_D3/handlers"
// "github.com/durotimicodes/xanda_task_R3_D3/helpers"
// "github.com/durotimicodes/xanda_task_R3_D3/models"
// "github.com/durotimicodes/xanda_task_R3_D3/repository"
// "github.com/stretchr/testify/assert"

func TestGetAllSpaceship(t *testing.T) {
	// r := helpers.SetUpRouter()

	// //repository
	// repository := repository.NewMySqlDB(database.DB)
	// handler := handlers.NewHandler(repository)

	// r.Get("/spaceships", handler.GetAllSpaceShipsHandler)

	// jsonValue, _ := json.MarshalIndent(map[string]bool{"success": true}, "", " ")
	// req, _ := http.NewRequest("GET", "/spaceships", bytes.NewBuffer(jsonValue))
	// w := httptest.NewRecorder()
	// r.ServeHTTP(w, req)

	// var spaceships []models.Spaceship
	// json.Unmarshal(w.Body.Bytes(), &spaceships)

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.NotEmpty(t, spaceships)
}
