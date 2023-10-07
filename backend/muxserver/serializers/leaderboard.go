package serializers

import (
	"context"
	"github.com/Sergey-pr/movie-games-tg/models"
)

type userData struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

// Leaderboards returns serialized users data with their highest scores
func Leaderboards(ctx context.Context, pointsData []*models.UserData) ([]*userData, error) {
	// Get all users ids form objects
	usersIds := make([]int, len(pointsData))
	for i, obj := range pointsData {
		usersIds[i] = obj.UserId
	}
	// Get usernames cache as map[userId]username to use later
	usersNamesCache, err := models.GetUsersNamesByIds(ctx, usersIds)
	if err != nil {
		return nil, err
	}
	// Serialize data getting user names from cache
	data := make([]*userData, len(pointsData))
	for i, userPoints := range pointsData {
		data[i] = &userData{
			Name:   usersNamesCache[userPoints.UserId],
			Points: userPoints.TotalPoints,
		}
	}
	return data, nil
}
