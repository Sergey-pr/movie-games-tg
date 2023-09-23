package serializers

import (
	"github.com/Sergey-pr/movie-games-tg/models"
)

type user struct {
	Id         int    `json:"id"`
	TelegramId int    `json:"tg_id"`
	Name       string `json:"name"`
}

func User(obj *models.User) *user {
	return &user{
		Id:         obj.Id,
		TelegramId: obj.TelegramId,
		Name:       obj.Name,
	}
}
