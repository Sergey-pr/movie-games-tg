package handlers

import (
	"github.com/Sergey-pr/movie-games-tg/models"
	"github.com/Sergey-pr/movie-games-tg/muxserver/forms"
	"github.com/Sergey-pr/movie-games-tg/muxserver/serializers"
	"net/http"
)

func CardsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cards, err := models.GetAllCards(ctx)
	OrPanic(err)

	Resp(w, serializers.Cards(cards))
}

func CardCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var form forms.Card
	OrPanic(ValidateForm(r, &form))

	card := &models.Card{
		Category:         form.Category,
		NameRu:           form.NameRu,
		DescRu:           form.DescRu,
		QuoteRu:          form.QuoteRu,
		NameEn:           form.NameEn,
		DescEn:           form.DescEn,
		QuoteEn:          form.QuoteEn,
		FactsRu:          form.FactsRu,
		FactsEn:          form.FactsEn,
		AnswersRu:        form.AnswersRu,
		AnswersEn:        form.AnswersEn,
		DrawingId:        form.DrawingId,
		PixelatedId:      form.PixelatedId,
		ScreenshotId:     form.ScreenshotId,
		BackgroundColor1: form.BackgroundColor1,
		BackgroundColor2: form.BackgroundColor2,
	}
	OrPanic(card.Save(ctx))

	Resp(w, serializers.Card(card))
}

func CardDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cardId := GetId(r)

	card := ObjOrPanic(models.GetCardById(ctx, cardId))
	OrPanic(card.Delete(ctx))

	Ok(w)
}

func CardInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cardId := GetId(r)
	card := ObjOrPanic(models.GetCardById(ctx, cardId))

	Resp(w, serializers.Card(card))
}

func CardUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var form forms.Card
	OrPanic(ValidateForm(r, &form))

	cardId := GetId(r)
	card := ObjOrPanic(models.GetCardById(ctx, cardId))

	if form.Category != "" && form.Category != card.Category {
		card.Category = form.Category
	}

	if form.NameRu != "" && form.NameRu != card.NameRu {
		card.NameRu = form.NameRu
	}

	if form.DescRu != "" && form.DescRu != card.DescRu {
		card.DescRu = form.DescRu
	}

	if form.QuoteRu != "" && form.QuoteRu != card.QuoteRu {
		card.QuoteRu = form.QuoteRu
	}

	if form.NameEn != "" && form.NameEn != card.NameEn {
		card.NameEn = form.NameEn
	}

	if form.DescEn != "" && form.DescEn != card.DescEn {
		card.DescEn = form.DescEn
	}

	if form.QuoteEn != "" && form.QuoteEn != card.QuoteEn {
		card.QuoteEn = form.QuoteEn
	}

	if form.FactsRu != nil {
		card.FactsRu = form.FactsRu
	}

	if form.FactsEn != nil {
		card.FactsEn = form.FactsEn
	}

	if form.AnswersRu != nil {
		card.AnswersRu = form.AnswersRu
	}

	if form.AnswersEn != nil {
		card.AnswersEn = form.AnswersEn
	}

	if form.DrawingId != "" && form.DrawingId != card.DrawingId {
		card.DrawingId = form.DrawingId
	}

	if form.PixelatedId != "" && form.PixelatedId != card.PixelatedId {
		card.PixelatedId = form.PixelatedId
	}

	if form.ScreenshotId != "" && form.ScreenshotId != card.ScreenshotId {
		card.ScreenshotId = form.ScreenshotId
	}

	if form.BackgroundColor1 != "" && form.BackgroundColor1 != card.BackgroundColor1 {
		card.BackgroundColor1 = form.BackgroundColor1
	}

	if form.BackgroundColor2 != "" && form.BackgroundColor2 != card.BackgroundColor2 {
		card.BackgroundColor2 = form.BackgroundColor2
	}

	OrPanic(card.Save(ctx))

	Resp(w, serializers.Card(card))
}
