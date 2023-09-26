package handlers

import (
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver/forms"
	"github.com/Sergey-pr/movie-games-tg/utils"
	"net/http"
)

// BotUpdates ...
func BotUpdates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var form forms.BotUpdate
	OrPanic(ValidateForm(r, &form))

	user, err := models.LoginUser(ctx, form.Message.From.Id)
	if err != nil {
		user = &models.User{
			TelegramId: form.Message.From.Id,
			Name:       form.Message.From.FirstName,
			UserName:   form.Message.From.Username,
			Language:   form.Message.From.LanguageCode,
		}
		OrPanic(user.Save(ctx))
	}

	if !user.IsAdmin {
		var msg string
		if user.Language == "ru" {
			msg = "Эта команда только для админов"
		} else {
			msg = "This command is only for admins"
		}
		OrPanic(utils.SendBotMessage(form.Message.Chat.Id, msg))
		Ok(w)
		return
	}

	processor := ObjOrPanic(user.GetBotProcessor(ctx, form.Message.Chat.Id))
	OrPanic(processor.ProcessMsg(ctx, form.Message.Text))

	Ok(w)
}
