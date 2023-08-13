package test

import "testing"

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
// 	"github.com/durotimicodes/xanda_task_R3_D3/handlers"
// 	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
// 	"github.com/durotimicodes/xanda_task_R3_D3/models"
// 	"github.com/durotimicodes/xanda_task_R3_D3/repository"
// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/v5"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/gorm"
// )

func TestGetSpaceship(t *testing.T) {
	// 	database.InitDatabase()
	// 	db := database.GetDB()
	// 	req, w := setGetBookRouter(db)

	// 	a := assert.New(t)
	// 	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	// 	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	// 	body, err := ioutil.ReadAll(w.Body)
	// 	if err != nil {
	// 		a.Error(err)
	// 	}

	// 	actual := models.Spaceship{}
	// 	if err := json.Unmarshal(body, &actual); err != nil {
	// 		a.Error(err)
	// 	}

	// 	expected := models.Spaceship{}
	// 	a.Equal(expected, actual)

	// }

	// func setGetBookRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	// 	r := chi.NewRouter()

	// 	repository := repository.NewMySqlDB(database.DB)
	// 	handler := handlers.NewHandler(repository)

	// 	r.Get("/{ID}", handler.GetSpaceShipHandler)
	// 	req, err := http.NewRequest(http.MethodGet, "/{ID}", nil)
	// 	helpers.HandlerErr(err)

	// 	req.Header.Set("Content-Type", "application/json")
	// 	w := httptest.NewRecorder()
	// 	r.ServeHTTP(w, req)
	// 	return req, w
}
