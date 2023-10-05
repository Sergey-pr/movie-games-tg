package muxserver

import (
	"context"
	"github.com/Sergey-pr/movie-games-tg/muxserver/handlers"
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
	rootRouter.Use(panicMiddleware)
	router := rootRouter.PathPrefix("/api").Subrouter()

	public := router.PathPrefix("/public").Subrouter()
	public.HandleFunc("/login/", handlers.Login).Methods(http.MethodPost).Name("public:login")
	public.HandleFunc("/bot-updates/", handlers.BotUpdates).Methods(http.MethodPost).Name("public:bot_updates")
	public.HandleFunc("/bot-image/{image_id}/", handlers.BotImage).Methods(http.MethodGet).Name("public:bot_image")

	private := router.PathPrefix("").Subrouter()
	private.Use(authMiddleware)

	private.HandleFunc("/user/", handlers.UserInfo).Methods(http.MethodGet).Name("")
	private.HandleFunc("/user/lang/", handlers.UserChangeLang).Methods(http.MethodPost).Name("")
	private.HandleFunc("/user/answer/", handlers.UserProcessAnswer).Methods(http.MethodPost).Name("")

	private.HandleFunc("/cards/", handlers.CardsList).Methods(http.MethodGet).Name("")
	private.HandleFunc("/cards/", handlers.CardCreate).Methods(http.MethodPost).Name("")
	private.HandleFunc("/cards/{id}/", handlers.CardDelete).Methods(http.MethodDelete).Name("")
	private.HandleFunc("/cards/{id}/", handlers.CardInfo).Methods(http.MethodGet).Name("")
	private.HandleFunc("/cards/{id}/", handlers.CardUpdate).Methods(http.MethodPost).Name("")

	private.HandleFunc("/leaderboard/", handlers.Leaderboard).Methods(http.MethodGet).Name("")

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
