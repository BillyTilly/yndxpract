package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

var AppConfig mainConfig

type mainConfig struct {
	Host    string `env:"SERVER_ADDRESS"`
	BaseURL string `env:"BASE_URL"`
}

func GetConfig() {
	flag.StringVar(&AppConfig.Host, "a", "localhost:8080", "host")
	flag.StringVar(&AppConfig.BaseURL, "b", "http://localhost:8080", "resulted host")

	flag.Parse()

	env.Parse(&AppConfig)
}
