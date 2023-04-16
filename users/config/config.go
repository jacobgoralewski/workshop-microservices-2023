package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Username string `mapstructure:"username"`
	Password string
	Port     string
	Host     string
	Name     string
}

type AppConfig struct {
	Db DbConfig `mapstructure:"db"`
}

func GetConfig() AppConfig {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.BindEnv("db.port", "DB_PORT")
	viper.BindEnv("db.host", "DB_HOST")
	viper.BindEnv("db.name", "DB_NAME")
	viper.BindEnv("db.username", "DB_USERNAME")
	viper.BindEnv("db.password", "DB_PASSWORD")

	err := viper.ReadInConfig()
	if err != nil {
		log.Error().Err(err).Msgf("unable to decode into struct, %v", err)
	}

	appConf := AppConfig{}
	err = viper.Unmarshal(&appConf)
	if err != nil {
		log.Fatal().Err(err).Msgf("unable to decode into struct, %v", err)
	}

	log.Debug().Msgf("config: %+v", appConf)

	return appConf
}
