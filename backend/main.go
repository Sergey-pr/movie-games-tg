package main

import (
	"github.com/Sergey-pr/movie-games-tg/config"
	"github.com/Sergey-pr/movie-games-tg/muxserver"
	"github.com/Sergey-pr/movie-games-tg/utils"
	"os"
	"os/signal"
)

func main() {
	api := muxserver.NewApiServer()
	api.Run(config.AppConfig.RestListen)

	// Register callback for telegram bot
	err := utils.RegisterCallback()
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	if api != nil {
		api.Shutdown()
	}
}
