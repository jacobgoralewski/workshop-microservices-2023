package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	TodosServiceURL string `mapstructure:"todos_service_url"`
	UsersServiceURL string `mapstructure:"users_service_url"`
}

func GetConfig() AppConfig {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.BindEnv("todos_service_url", "TODOS_SERVICE_URL")
	viper.BindEnv("users_service_url", "USERS_SERVICE_URL")

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
