package handlers

import (
	"encoding/json"
	"github.com/Sergey-pr/movie-games-tg/models"
	"net/http"
)

// OrPanic is a shortcut to panic
func OrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// ObjOrPanic is a shortcut to panic
func ObjOrPanic[T any](o T, err error) T {
	if err != nil {
		panic(err)
	}
	return o
}

// JsonResponse is a shortcut to response data from serialized struct
func JsonResponse(w http.ResponseWriter, i interface{}, statusCode ...int) {
	var status = http.StatusOK
	if len(statusCode) > 0 {
		status = statusCode[0]
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if i != nil {
		_ = json.NewEncoder(w).Encode(i)
	}
}

// GetUser is a shortcut to get user from context
func GetUser(r *http.Request) *models.User {
	return r.Context().Value(models.UserContextKey).(*models.User)
}

// CompletedResponse is a shortcut to completed response
func CompletedResponse(w http.ResponseWriter) {
	JsonResponse(w, map[string]string{"status": "completed"})
}
