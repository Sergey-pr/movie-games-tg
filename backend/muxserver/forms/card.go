package forms

import "github.com/Sergey-pr/movie-games-tg/utils"

type Card struct {
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
}
