package app

import (
	"github.com/akxcix/butler/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", handlers.HealthCheck)
	return r
}
