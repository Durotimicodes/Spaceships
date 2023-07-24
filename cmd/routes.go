package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/durotimicodes/xanda_task_R3_D3/cmd/database"
	"github.com/durotimicodes/xanda_task_R3_D3/handlers"
	"github.com/durotimicodes/xanda_task_R3_D3/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Routes
func StartApi() {

	const webPort = ":3400"

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	//group of middlewares
	// commonMiddlewares := []middlewares.NewMiddleware{
	// 	middlewares.CORSMiddleware,
	// 	middlewares.LoggingMiddleware,
	// 	middlewares.PanicRecoveryMiddleware,
	// 	middlewares.HeaderMiddleware,
	// }

	//repository
	repository := repository.NewMySqlDB(database.DB)
	handler := handlers.NewHandler(repository)

	//if the middleware is still alive
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/spaceships", handler.GetAllSpaceShipsHandler)

	r.Route("/spaceship", func(r chi.Router) {
		r.Get("/{ID}", handler.GetSpaceShipHandler)

		r.Post("/create", handler.CreateSpaceshipHandler)

		r.Delete("/delete/{ID}", handler.DeleteSpaceshipHandler)

		r.Put("/update/{ID}", handler.UpdateSpaceshipHandler)
	})

	log.Printf("Starting the server on port %s", webPort)
	log.Fatal(http.ListenAndServe(webPort, r))
}
