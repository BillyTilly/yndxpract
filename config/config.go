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
	var host string
	var baseURL string

	flag.StringVar(&host, "a", "localhost:8080", "host")

	flag.Parse()

	env.Parse(&AppConfig)

	if AppConfig.Host == "" {
		AppConfig.Host = host
	}

	flag.StringVar(&baseURL, "b", "http://"+AppConfig.Host, "resulted host")

	flag.Parse()

	if AppConfig.BaseURL == "" {
		AppConfig.BaseURL = baseURL
	}
}
