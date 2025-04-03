package config

import (
	"log"

	"github.com/spf13/viper"
)

type Envs struct {
	Port string `mapstructure:"PORT"`
}

func LoadEnvs() Envs {
	viper.AutomaticEnv() // Read envs

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	var envs Envs
	if err := viper.Unmarshal(&envs); err != nil {
		log.Fatal("Error loading envs:", err)
	}

	return envs
}
