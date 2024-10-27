package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

type config struct {
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	DATABASE_NAME     string

	HTTP_PORT string
}

func NewConfig(log *slog.Logger) *config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file: %s", err)
	}

	config := &config{
		DATABASE_HOST:     viper.GetString("DATABASE_HOST"),
		DATABASE_PORT:     viper.GetString("DATABASE_PORT"),
		DATABASE_USERNAME: viper.GetString("DATABASE_USERNAME"),
		DATABASE_PASSWORD: viper.GetString("DATABASE_PASSWORD"),
		DATABASE_NAME:     viper.GetString("DATABASE_NAME"),
		HTTP_PORT:         viper.GetString("HTTP_PORT"),
	}

	return config
}
