package db

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"todos/config"
	"todos/models"
)

func Init() *gorm.DB {
	configuration := config.GetConfig()

	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: fmt.Sprintf(
					"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
					configuration.Db.Host,
					configuration.Db.Username,
					configuration.Db.Password,
					configuration.Db.Name,
					configuration.Db.Port,
				),
				PreferSimpleProtocol: true,
			},
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		log.Err(err).Msg("failed to connect to teh database")
		panic("DB Connection Error") // todo
	}

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		panic("DB Connection Error") // todo
	}

	return db
}
