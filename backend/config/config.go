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
	RestListen      string `envconfig:"REST_LISTEN" required:"true" default:"0.0.0.0:8888"`
	BaseUrl         string `envconfig:"BASE_URL" required:"true" default:"http://localhost:8888/api"`
	BaseFrontendUrl string `envconfig:"BASE_FRONTEND_URL" required:"true" default:"http://localhost:3000"`

	Dns          string `envconfig:"DATABASE" required:"true"`
	MaxIdleConns int    `envconfig:"DATABASE_MAX_IDLE_CONNS" required:"15"`
	MaxOpenConns int    `envconfig:"DATABASE_OPEN_IDLE_CONNS" required:"10"`
}

func getConfig() *config {
	var c config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatalln(err)
	}

	return &c
}
