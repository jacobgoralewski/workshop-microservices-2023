package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"todos/config"
)

func Init() *sqlx.DB {
	configuration := config.GetConfig()

	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			configuration.Db.Host,
			configuration.Db.Port,
			configuration.Db.Username,
			configuration.Db.Password,
			configuration.Db.Name,
		),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create db connection")
	}

	if err != nil {
		log.Err(err).Msg("failed to connect to teh database")
		panic("DB Connection Error") // todo
	}

	return db
}
