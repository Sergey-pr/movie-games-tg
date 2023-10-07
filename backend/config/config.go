package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

var AppConfig *config

func init() {
	log.Println("Loading config...")
	AppConfig = getConfig()
}

type config struct {
	Debug           bool   `envconfig:"DEBUG" default:"true"`
	RestListen      string `envconfig:"REST_LISTEN" default:"0.0.0.0:8888"`
	BaseUrl         string `envconfig:"BASE_URL" default:"http://localhost:8888/api"`
	BaseFrontendUrl string `envconfig:"BASE_FRONTEND_URL" default:"http://localhost:3000"`

	SessionTTL int64  `envconfig:"SESSION_DAYS" default:"24"`
	JwtSecret  string `envconfig:"JWT_SECRET" default:"asdasklhfkdjfksh23423432hjgjhg"`

	Dns          string `envconfig:"DATABASE" required:"true"`
	MaxIdleConns int    `envconfig:"DATABASE_MAX_IDLE_CONNS" default:"15"`
	MaxOpenConns int    `envconfig:"DATABASE_OPEN_IDLE_CONNS" default:"10"`

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
