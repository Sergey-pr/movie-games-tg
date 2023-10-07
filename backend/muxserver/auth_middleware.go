package muxserver

import (
	"context"
	"errors"
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/utils/jwt"
	"net/http"
	"strings"
)

// authMiddleware is a middleware that validates jwt token and adds user to context
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.RequestURI, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		ctx := r.Context()
		var user *models.User
		// Get token from headers
		jwtToken := r.Header.Get(jwt.TokenKey)
		if len(jwtToken) > 0 {
			// If token expired get new token
			newTknStr, err := jwt.RenewJwtToken(jwtToken)
			if err != nil {
				panic(err)
			}
			// Parse token to get user id
			claims, err := jwt.ParseJwtToken(newTknStr)
			if err != nil {
				panic(err)
			}
			// Get user object by id from database
			user, err = models.GetUserById(ctx, claims.User.Id)
			if err != nil {
				panic(err)
			}
		}
		// Check if we found a user by JWT token
		if user == nil {
			panic(errors.New("not authorized error"))
		}
		// Add user to context
		ctx = context.WithValue(ctx, models.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
