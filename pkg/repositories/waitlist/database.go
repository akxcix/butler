package waitlist

import (
	"github.com/akxcix/nomadcore/pkg/config"
	"github.com/akxcix/nomadcore/pkg/repositories"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sqlx.DB
}

func New(conf *config.DatabaseConfig) *Database {
	log.Info().Msg("Connecting to database")

	dsn := repositories.FormatPostgresDSN(
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DatabaseName,
	)
	conn := sqlx.MustConnect("postgres", dsn)
	log.Info().Msg("Connected to database")

	return &Database{db: conn}
}
