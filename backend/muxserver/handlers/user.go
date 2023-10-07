package handlers

import (
	"encoding/json"
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver/forms"
	"github.com/Sergey-pr/movie-games-tg/muxserver/serializers"
	"net/http"
)

// UserInfo returns current user data
func UserInfo(w http.ResponseWriter, r *http.Request) {
	user := GetUser(r)
	JsonResponse(w, serializers.User(user))
}

// UserChangeLang changes language for current user
func UserChangeLang(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := GetUser(r)
	// Fill form struct with data from request and validate it
	var form forms.UserLang
	OrPanic(json.NewDecoder(r.Body).Decode(&form))
	// Change language and save user data
	user.Language = form.Language
	OrPanic(user.Save(ctx))
	// Return updated user data
	JsonResponse(w, serializers.User(user))
}

// UserProcessAnswer save user answer to card
func UserProcessAnswer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := GetUser(r)
	// Fill form struct with data from request and validate it
	var form forms.UserAnswer
	OrPanic(json.NewDecoder(r.Body).Decode(&form))
	// We only save first answer from user
	// If the answer from this user to this card already exist, we do nothing
	answer := ObjOrPanic(models.GetAnswerByCardIdAndUserId(ctx, form.CardId, user.Id))
	if answer == nil {
		// If answer doesn't exist, we create it
		answer = &models.Answer{
			UserId: user.Id,
			CardId: form.CardId,
			Points: form.Points,
		}
		OrPanic(answer.Save(ctx))
	}
	CompletedResponse(w)
}
