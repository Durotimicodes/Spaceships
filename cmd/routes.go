package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
	"github.com/durotimicodes/xanda_task_R3_D3/handlers"
	"github.com/durotimicodes/xanda_task_R3_D3/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func StartApi() {

	const webPort = ":3300"

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(middleware.Timeout(60 * time.Second))

	repository := service.NewMySqlDB(database.DB)

	handler := handlers.NewHandler(repository)

	r.Get("/spaceships", handler.GetAllSpaceShipsHandler)
	r.Get("/spaceship/{ID}", handler.GetSpaceShipHandler)
	r.Post("/spaceship/create", handler.CreateSpaceshipHandler)
	r.Delete("/spaceship/delete/{ID}", handler.DeleteSpaceshipHandler)
	r.Put("/spaceship/update/{ID}", handlers.UpdateSpaceshipHandler)

	log.Printf("Starting the server on port %s", webPort)
	log.Fatal(http.ListenAndServe(webPort, r))
}
