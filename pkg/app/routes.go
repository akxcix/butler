package app

import (
	"github.com/akxcix/butler/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func SetRouter(app Application) Application {
	r := chi.NewRouter()
	r.Get("/health", handlers.HealthCheck)
	app.Router = r

	return app
}
