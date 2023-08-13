package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/durotimicodes/xanda_task_R3_D3/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
	ContentType string
	Handlers    map[string]func(w http.ResponseWriter, r *http.Request)
}

// Routes
func StartApi(h *handlers.Handler) (*chi.Mux, string) {

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

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))

	//if the middleware is still alive
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/spaceships", h.GetAllSpaceShipsHandler)

	r.Route("/spaceship", func(r chi.Router) {
		r.Get("/{ID}", h.GetSpaceShipHandler)

		r.Post("/create", h.CreateSpaceshipHandler)

		r.Delete("/delete/{ID}", h.DeleteSpaceshipHandler)

		r.Put("/update/{ID}", h.UpdateSpaceshipHandler)
	})

	log.Printf("Starting the server on port %s", webPort)

	log.Fatal(http.ListenAndServe(webPort, r))

	return r, webPort

}
