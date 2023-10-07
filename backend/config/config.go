package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

var AppConfig = getConfig()

type config struct {
	RestListen string `envconfig:"REST_LISTEN" default:"0.0.0.0:8888"`
	JwtSecret  string `envconfig:"JWT_SECRET" default:"jfhDFDFhfdhgklslf32487234DFSfdsfd"`

	DatabaseDns string `envconfig:"DATABASE" required:"true"`

	TelegramBotToken string `envconfig:"TELEGRAM_BOT_TOKEN" required:"true"`

	FrontendHostname string `envconfig:"FRONTEND_HOSTNAME" required:"true"`
	BackendHostname  string `envconfig:"BACKEND_HOSTNAME" required:"true"`
}

func getConfig() *config {
	var c config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatalln(err)
	}

	return &c
}
