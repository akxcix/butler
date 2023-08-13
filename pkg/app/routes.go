package app

import (
	"github.com/akxcix/butler/pkg/handlers"
	waitlistHandlers "github.com/akxcix/butler/pkg/handlers/waitlist"
	"github.com/akxcix/butler/pkg/services/waitlist"

	"github.com/go-chi/chi/v5"
)

func createRoutes(waitlistService *waitlist.Service) *chi.Mux {
	r := chi.NewRouter()

	// global middlewares
	r.Use(handlers.LogRequest)

	// general routes
	r.Get("/health", handlers.HealthCheck)

	waitlistHandlers := waitlistHandlers.New(waitlistService)
	r.Post("/waitlist", waitlistHandlers.PostWaitlist)

	return r
}
