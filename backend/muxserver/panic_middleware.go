package muxserver

import (
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/muxserver/handlers"
	"github.com/lib/pq"
	"net/http"
)

// panicMiddleware is a middleware that returns errors in response
func panicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			recoveredError := recover()
			if recoveredError != nil {
				var errorText interface{}
				// Switch by error type to get error text
				switch err := recoveredError.(type) {
				case *pq.Error:
					errorText = fmt.Sprintf("%s: %s", err.Code, err.Message)
				case string:
					errorText = err
				case error:
					errorText = err.Error()
				default:
					errorText = "Undefined error"
				}
				// Respond with error text
				handlers.JsonResponse(w, map[string]string{"error": fmt.Sprintf("%v", errorText)}, http.StatusOK)
			}
		}()
		h.ServeHTTP(w, req)
	})
}
