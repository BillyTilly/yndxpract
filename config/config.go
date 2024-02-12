package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

var AppConfig mainConfig

type mainConfig struct {
	Host    string `env:"SERVER_ADDRESS"`
	BaseUrl string `env:"BASE_URL"`
}

func GenerateConfig() {
	var host string
	var baseUrl string

	flag.StringVar(&host, "a", "localhost:8080", "host")
	flag.StringVar(&baseUrl, "b", "http://localhost:8080", "resulted host")

	flag.Parse()

	env.Parse(&AppConfig)

	if AppConfig.Host == "" {
		AppConfig.Host = host
	}

	if string(AppConfig.Host) != "localhost:8080" {
		AppConfig.BaseUrl = "http://" + host
		return
	}

	if AppConfig.BaseUrl == "" {
		AppConfig.BaseUrl = baseUrl
	}
}
