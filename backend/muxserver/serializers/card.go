package serializers

import (
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/utils"
)

type card struct {
	Id               int         `json:"id"`
	Category         string      `json:"category"`
	NameRu           string      `json:"name_ru"`
	DescRu           string      `json:"desc_ru"`
	QuoteRu          string      `json:"quote_ru"`
	NameEn           string      `json:"name_en"`
	DescEn           string      `json:"desc_en"`
	QuoteEn          string      `json:"quote_en"`
	FactsRu          utils.JSONB `json:"facts_ru"`
	FactsEn          utils.JSONB `json:"facts_en"`
	AnswersRu        utils.JSONB `json:"answers_ru"`
	AnswersEn        utils.JSONB `json:"answers_en"`
	DrawingId        string      `json:"drawing_id"`
	PixelatedId      string      `json:"pixelated_id"`
	ScreenshotId     string      `json:"screenshot_id"`
	BackgroundColor1 string      `json:"bg_color_1"`
	BackgroundColor2 string      `json:"bg_color_2"`
	TextColor        string      `json:"text_color"`
}

// Card returns serialized Card object
func Card(obj *models.Card) *card {
	return &card{
		Id:               obj.Id,
		Category:         obj.Category,
		NameRu:           obj.NameRu,
		DescRu:           obj.DescRu,
		QuoteRu:          obj.QuoteRu,
		NameEn:           obj.NameEn,
		DescEn:           obj.DescEn,
		QuoteEn:          obj.QuoteEn,
		FactsRu:          obj.FactsRu,
		FactsEn:          obj.FactsEn,
		AnswersRu:        obj.AnswersRu,
		AnswersEn:        obj.AnswersEn,
		DrawingId:        obj.DrawingId,
		PixelatedId:      obj.PixelatedId,
		ScreenshotId:     obj.ScreenshotId,
		BackgroundColor1: obj.BackgroundColor1,
		BackgroundColor2: obj.BackgroundColor2,
		TextColor:        obj.TextColor,
	}
}

// Cards return serialized Cards objects that are ready to marshal to response
func Cards(objs []*models.Card) []*card {
	cards := make([]*card, len(objs))
	for i, obj := range objs {
		cards[i] = Card(obj)
	}
	return cards
}
