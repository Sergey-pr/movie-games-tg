package serializers

import (
	"context"
	"github.com/Sergey-pr/movie-games-tg/models"
)

type userData struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

func Leaderboards(ctx context.Context, pointsData []*models.UserData) ([]*userData, error) {
	usersIds := make([]int, len(pointsData))
	for i, obj := range pointsData {
		usersIds[i] = obj.UserId
	}
	usersNamesCache, err := models.GetUsersNamesByIds(ctx, usersIds)
	if err != nil {
		return nil, err
	}

	data := make([]*userData, len(pointsData))
	for i, userPoints := range pointsData {
		data[i] = &userData{
			Name:   usersNamesCache[userPoints.UserId],
			Points: userPoints.TotalPoints,
		}
	}
	return data, nil
}
