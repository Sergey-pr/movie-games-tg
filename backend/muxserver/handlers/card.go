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
		Category:      form.Category,
		NameRu:        form.NameRu,
		DescRu:        form.DescRu,
		QuoteRu:       form.QuoteRu,
		NameEn:        form.NameEn,
		DescEn:        form.DescEn,
		QuoteEn:       form.QuoteEn,
		FactsRu:       form.FactsRu,
		FactsEn:       form.FactsEn,
		AnswersRu:     form.AnswersRu,
		AnswersEn:     form.AnswersEn,
		DrawingUrl:    form.DrawingUrl,
		PixelatedUrl:  form.PixelatedUrl,
		ScreenshotUrl: form.ScreenshotUrl,
		BackgroundUrl: form.BackgroundUrl,
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

	if form.DrawingUrl != "" && form.DrawingUrl != card.DrawingUrl {
		card.DrawingUrl = form.DrawingUrl
	}

	if form.PixelatedUrl != "" && form.PixelatedUrl != card.PixelatedUrl {
		card.PixelatedUrl = form.PixelatedUrl
	}

	if form.ScreenshotUrl != "" && form.ScreenshotUrl != card.ScreenshotUrl {
		card.ScreenshotUrl = form.ScreenshotUrl
	}

	if form.BackgroundUrl != "" && form.BackgroundUrl != card.BackgroundUrl {
		card.BackgroundUrl = form.BackgroundUrl
	}

	OrPanic(card.Save(ctx))

	Resp(w, serializers.Card(card))
}
