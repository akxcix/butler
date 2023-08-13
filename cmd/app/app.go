package app

import (
	"fmt"

	"github.com/akxcix/butler/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Application struct {
	Config *config.Config
}

func Run() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	config, err := config.Read("./config.yml")
	if err != nil {
		log.Fatal().Err(err)
	}

	app := Application{
		Config: config,
	}

	log.Info().Msg(fmt.Sprintf("Running application at port %s\n", app.Config.Server.Port))
}
