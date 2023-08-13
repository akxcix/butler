package app

import (
	"fmt"
	"net/http"

	"github.com/akxcix/butler/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Application struct {
	Config *config.Config
	Router *chi.Mux
}

func Run() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	config, err := config.Read("./config.yml")
	if err != nil {
		log.Fatal().Err(err)
	}

	r := NewRouter()

	app := Application{
		Config: config,
		Router: r,
	}

	addr := fmt.Sprintf("%s:%s", app.Config.Server.Host, app.Config.Server.Port)
	log.Info().Msg(fmt.Sprintf("Running application at %s", addr))
	http.ListenAndServe(addr, r)
}
