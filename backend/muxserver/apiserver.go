package muxserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	muxhandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ApiServer struct {
	name         string
	router       *mux.Router
	writeTimeout time.Duration
	readTimeout  time.Duration
	httpServer   *http.Server
}

func NewApiServer() *ApiServer {

	rootRouter := mux.NewRouter().StrictSlash(true)
	rootRouter.Use(authMiddleware)

	_ = rootRouter.PathPrefix("/api").Subrouter()

	return &ApiServer{
		name:   "API",
		router: rootRouter,
	}
}

func (s *ApiServer) Run(addr string) {
	go func() {
		s.httpServer = &http.Server{
			ReadTimeout:  s.readTimeout,
			WriteTimeout: s.writeTimeout,
			Handler:      cors.AllowAll().Handler(muxhandlers.LoggingHandler(os.Stdout, s.router)),
			Addr:         addr,
		}

		log.Printf("%s server listen: %s\n", s.name, addr)

		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()
}

func (s *ApiServer) Shutdown() {
	log.Printf("%s server shutting down\n", s.name)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_ = s.httpServer.Shutdown(ctx)

}
