package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	TodosServiceURL string `mapstructure:"todos_service_url"`
}

func GetConfig() AppConfig {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msgf("unable to decode into struct, %v", err)
	}

	appConf := AppConfig{}
	err = viper.Unmarshal(&appConf)
	if err != nil {
		log.Fatal().Err(err).Msgf("unable to decode into struct, %v", err)
	}
	log.Debug().Msgf("config: %+v", appConf)

	return appConf
}
