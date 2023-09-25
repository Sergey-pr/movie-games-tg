package handlers

import (
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver/forms"
	"github.com/Sergey-pr/movie-games-tg/muxserver/serializers"
	"net/http"
)

// Login user by email and password
func Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var form forms.LoginForm
	OrPanic(ValidateForm(r, &form))

	user, err := models.LoginUser(ctx, form.User.TelegramId)
	if err != nil {
		user = &models.User{
			TelegramId: form.User.TelegramId,
			Name:       form.User.Name,
			Language:   form.User.Language,
		}
		OrPanic(user.Save(ctx))
	}

	token, expirationTime, err := user.GetJwtToken()
	OrPanic(err)

	Resp(w, serializers.JwtToken{
		Token:   token,
		ExpTime: expirationTime.Unix(),
	})
}
