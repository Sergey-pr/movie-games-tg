package handlers

import (
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
