package config

var AppConfig mainConfig

type mainConfig struct {
	Host       string
	RedirectAd string
}

func GenerateConfig(host string, reditectAd string) {
	AppConfig = mainConfig{Host: host, RedirectAd: reditectAd}
}
