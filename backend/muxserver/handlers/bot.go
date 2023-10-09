package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver/forms"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
)

// BotUpdates is handling bot commands
func BotUpdates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Fill form struct with data from request and validate it
	var form forms.BotUpdate
	OrPanic(json.NewDecoder(r.Body).Decode(&form))
	// Get or create user
	user, err := models.LoginUser(ctx, form.Message.From.Id)
	if err != nil || user == nil {
		user = &models.User{
			TelegramId: form.Message.From.Id,
			Name:       form.Message.From.FirstName,
			UserName:   form.Message.From.Username,
			Language:   form.Message.From.LanguageCode,
		}
		OrPanic(user.Save(ctx))
	}
	// Admin commands are working only for IsAdmin users
	if user.IsAdmin {
		processor := ObjOrPanic(user.GetBotProcessor(ctx, form.Message.Chat.Id))
		OrPanic(processor.ProcessMsg(ctx, &form))
		CompletedResponse(w)
		return
	}

	OrPanic(models.SendStartBotMessage(ctx, form.Message.Chat.Id, user.Language))
	CompletedResponse(w)
}

// BotImage returns image static link by id
func BotImage(w http.ResponseWriter, r *http.Request) {
	imageId := mux.Vars(r)["image_id"]
	// Open image by id, image filenames are id
	img := ObjOrPanic(os.Open(fmt.Sprintf("card_files/%s", imageId)))
	// Set content-type header
	w.Header().Set("Content-Type", "image/jpeg")
	ObjOrPanic(io.Copy(w, img))
}
