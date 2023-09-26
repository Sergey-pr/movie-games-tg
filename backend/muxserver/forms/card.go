package forms

import "github.com/Sergey-pr/movie-games-tg/utils"

type Card struct {
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
