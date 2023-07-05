package handlers

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/durotimicodes/xanda_task_R3_D3/handlers"
// )

// func TestGetAllSpaceship(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "/spaceships", nil)
// 	res := httptest.NewRecorder()

// 	want := "Hello, get all spaceships"

// 	handlers.GetAllSpaceShipsHandler(res, req)

// 	body := res.Body.String()
// 	if body != want {
// 		t.Errorf("got %q want %q", body, want)
// 	}
// }