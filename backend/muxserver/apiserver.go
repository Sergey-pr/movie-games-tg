package muxserver

import (
	"context"
	"github.com/Sergey-pr/movie-games-tg/muxserver/handlers"
	muxhandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"time"
)

// ApiServer is a server object with Run and Shutdown methods
type ApiServer struct {
	router     *mux.Router
	httpServer *http.Server
}

// NewApiServer returns ApiServer object with set up handlers
func NewApiServer() *ApiServer {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(panicMiddleware)
	apiRouter := router.PathPrefix("/api").Subrouter()

	public := apiRouter.PathPrefix("/public").Subrouter()
	public.HandleFunc("/login/", handlers.Login).Methods(http.MethodPost).Name("public:login")
	public.HandleFunc("/bot-updates/", handlers.BotUpdates).Methods(http.MethodPost).Name("public:bot_updates")
	public.HandleFunc("/bot-image/{image_id}/", handlers.BotImage).Methods(http.MethodGet).Name("public:bot_image")

	private := apiRouter.PathPrefix("").Subrouter()
	private.Use(authMiddleware)

	private.HandleFunc("/user/", handlers.UserInfo).Methods(http.MethodGet).Name("")
	private.HandleFunc("/user/lang/", handlers.UserChangeLang).Methods(http.MethodPost).Name("")
	private.HandleFunc("/user/answer/", handlers.UserProcessAnswer).Methods(http.MethodPost).Name("")

	private.HandleFunc("/cards/", handlers.CardsList).Methods(http.MethodGet).Name("")

	private.HandleFunc("/leaderboard/", handlers.Leaderboard).Methods(http.MethodGet).Name("")

	return &ApiServer{
		router: router,
	}
}

// Run starts the server
func (s *ApiServer) Run(addr string) {
	go func() {
		s.httpServer = &http.Server{
			Handler: cors.AllowAll().Handler(muxhandlers.LoggingHandler(os.Stdout, s.router)),
			Addr:    addr,
		}

		log.Printf("server started at: %s\n", addr)

		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()
}

// Shutdown turns off the server with a 5-second delay to cancel shutdown
func (s *ApiServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = s.httpServer.Shutdown(ctx)
}
