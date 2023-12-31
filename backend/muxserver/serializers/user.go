package serializers

import (
	"github.com/Sergey-pr/movie-games-tg/models"
)

type user struct {
	Id         int    `json:"id"`
	TelegramId int    `json:"tg_id"`
	Name       string `json:"name"`
	Language   string `json:"language"`
}

// User returns serialized user data, so we could marshal it in response
func User(obj *models.User) *user {
	return &user{
		Id:         obj.Id,
		TelegramId: obj.TelegramId,
		Name:       obj.Name,
		Language:   obj.Language,
	}
}
