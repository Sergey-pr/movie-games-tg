package handlers

import (
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver/serializers"
	"net/http"
)

func Leaderboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := ObjOrPanic(models.GetLeaderboardData(ctx))

	Resp(w, ObjOrPanic(serializers.Leaderboards(ctx, data)))
}
