package handlers

import (
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver/serializers"
	"net/http"
)

// CardsList return list of all cards from database
func CardsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Get cards structs
	cards, err := models.GetAllCards(ctx)
	OrPanic(err)
	// Return serialized cards
	JsonResponse(w, serializers.Cards(cards))
}
