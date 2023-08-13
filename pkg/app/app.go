package app

import (
	"fmt"
	"net/http"

	"github.com/akxcix/butler/pkg/config"
	"github.com/akxcix/butler/pkg/repositories/waitlist"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Application struct {
	Config       *config.Config
	WaitlistRepo *waitlist.Database
	Router       *chi.Mux
}

func ReadConfigs(app Application) Application {
	config, err := config.Read("./config.yml")
	if err != nil {
		log.Fatal().Err(err)
	}

	app.Config = config
	return app
}

func addRepositories(app Application) Application {
	conf := app.Config
	if conf == nil {
		log.Fatal().Msg("Conf is nil")
	}

	waitlistRepo := waitlist.New(conf.Database)
	app.WaitlistRepo = waitlistRepo

	return app
}

func Run() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	app := Application{}
	app = ReadConfigs(app)
	app = addRepositories(app)
	app = SetRouter(app)

	addr := fmt.Sprintf("%s:%s", app.Config.Server.Host, app.Config.Server.Port)
	log.Info().Msg(fmt.Sprintf("Running application at %s", addr))
	http.ListenAndServe(addr, app.Router)
}