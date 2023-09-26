package serializers

import (
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/utils"
)

type card struct {
	Id            int         `json:"id"`
	Category      string      `json:"category"`
	NameRu        string      `json:"name_ru"`
	DescRu        string      `json:"desc_ru"`
	QuoteRu       string      `json:"quote_ru"`
	NameEn        string      `json:"name_en"`
	DescEn        string      `json:"desc_en"`
	QuoteEn       string      `json:"quote_en"`
	FactsRu       utils.JSONB `json:"facts_ru"`
	FactsEn       utils.JSONB `json:"facts_en"`
	AnswersRu     utils.JSONB `json:"answers_ru"`
	AnswersEn     utils.JSONB `json:"answers_en"`
	DrawingUrl    string      `json:"drawing_url"`
	PixelatedUrl  string      `json:"pixelated_url"`
	ScreenshotUrl string      `json:"screenshot_url"`
	BackgroundUrl string      `json:"bg_url"`
}

func Card(obj *models.Card) *card {
	return &card{
		Id:            obj.Id,
		Category:      obj.Category,
		NameRu:        obj.NameRu,
		DescRu:        obj.DescRu,
		QuoteRu:       obj.QuoteRu,
		NameEn:        obj.NameEn,
		DescEn:        obj.DescEn,
		QuoteEn:       obj.QuoteEn,
		FactsRu:       obj.FactsRu,
		FactsEn:       obj.FactsEn,
		AnswersRu:     obj.AnswersRu,
		AnswersEn:     obj.AnswersEn,
		DrawingUrl:    obj.DrawingUrl,
		PixelatedUrl:  obj.PixelatedUrl,
		ScreenshotUrl: obj.ScreenshotUrl,
		BackgroundUrl: obj.BackgroundUrl,
	}
}

func Cards(objs []*models.Card) []*card {
	cards := make([]*card, len(objs))

	for i, obj := range objs {
		cards[i] = Card(obj)
	}
	return cards
}
