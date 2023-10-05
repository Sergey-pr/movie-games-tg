package handlers

import (
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver/forms"
	"github.com/Sergey-pr/movie-games-tg/muxserver/serializers"
	"net/http"
)

func UserInfo(w http.ResponseWriter, r *http.Request) {
	user := GetUser(r)
	Resp(w, serializers.User(user))
}

func UserChangeLang(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := GetUser(r)

	var form forms.UserLang
	OrPanic(ValidateForm(r, &form))

	user.Language = form.Language
	OrPanic(user.Save(ctx))

	Resp(w, serializers.User(user))
}

func UserProcessAnswer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := GetUser(r)

	var form forms.UserAnswer
	OrPanic(ValidateForm(r, &form))

	answer := ObjOrPanic(models.GetAnswerByCardIdAndUserId(ctx, form.CardId, user.Id))
	if answer == nil {
		answer = &models.Answer{
			UserId: user.Id,
			CardId: form.CardId,
			Points: form.Points,
		}
		OrPanic(answer.Save(ctx))
	}

	Ok(w)
}
