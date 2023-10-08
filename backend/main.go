package main

import (
	"github.com/Sergey-pr/movie-games-tg/config"
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver"
	"os"
	"os/signal"
)

func main() {
	api := muxserver.NewApiServer()
	api.Run(config.AppConfig.RestListen)

	// Register callback for telegram bot
	err := models.RegisterCallback()
	if err != nil {
		panic(err)
	}
	// Create to channel to catch interrupt command
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	if api != nil {
		api.Shutdown()
	}
}
