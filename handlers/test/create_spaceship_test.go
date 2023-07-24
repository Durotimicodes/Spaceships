package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
	"github.com/durotimicodes/xanda_task_R3_D3/handlers"
	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
	"github.com/durotimicodes/xanda_task_R3_D3/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateSpaceship(t *testing.T) {

	r := helpers.SetUpRouter()

	//repository
	repository := repository.NewMySqlDB(database.DB)
	handler := handlers.NewHandler(repository)

	r.Post("/spaceship/create", handler.CreateSpaceshipHandler)

	createSpaceship := helpers.CreateMockSpaceship()

	jsonValue, _ := json.MarshalIndent(createSpaceship, "", "")

	req, err := http.NewRequest("POST", "/spaceship/create", bytes.NewBuffer(jsonValue))
	helpers.HandlerErr(err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}
