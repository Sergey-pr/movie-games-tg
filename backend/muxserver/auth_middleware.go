package muxserver

import (
	"context"
	"errors"
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/utils/jwt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := mux.CurrentRoute(r).GetName()

		if strings.HasPrefix(r.RequestURI, "/public") || strings.Contains(name, "public") || strings.HasPrefix(r.RequestURI, "/ws") {
			next.ServeHTTP(w, r)
			return
		}
		ctx := r.Context()

		var user *models.User

		jwtToken := r.Header.Get(jwt.TokenKey)
		if len(jwtToken) > 0 {
			newTknStr, _, err := jwt.RenewJwtToken(jwtToken)
			if err != nil {
				panic(err)
			}
			claims, err := jwt.ParseJwtToken(newTknStr)
			if err != nil {
				panic(err)
			}

			user, err = models.GetUserById(ctx, claims.User.Id)
			if err != nil {
				panic(err)
			}
		}

		if user == nil {
			panic(errors.New("not authorized"))
		}

		ctx = context.WithValue(ctx, models.UserContextKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
