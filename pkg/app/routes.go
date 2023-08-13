package app

import (
	"github.com/akxcix/butler/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func setRouter(app application) application {
	r := chi.NewRouter()
	r.Get("/health", handlers.HealthCheck)
	app.Router = r

	return app
}
