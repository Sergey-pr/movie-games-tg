package muxserver

import (
	"context"
	"github.com/Sergey-pr/movie-games-tg/models"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var user *models.User

		ctx = context.WithValue(ctx, models.UserContextKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
