package handlers

import (
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver/forms"
	"github.com/Sergey-pr/movie-games-tg/utils"
	"io"
	"net/http"
	"os"
)

// BotUpdates is handling bot commands
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
	OrPanic(processor.ProcessMsg(ctx, &form))

	Ok(w)
}

// BotImage returns image by id
func BotImage(w http.ResponseWriter, r *http.Request) {
	imageId := GetImageId(r)

	img := ObjOrPanic(os.Open(fmt.Sprintf("card_files/%s", imageId)))
	defer OrPanic(img.Close())
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	ObjOrPanic(io.Copy(w, img))

}
