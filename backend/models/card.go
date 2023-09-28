package models

import (
	"context"
	"errors"
	"github.com/Sergey-pr/movie-games-tg/persist"
	"github.com/Sergey-pr/movie-games-tg/utils"
	"github.com/doug-martin/goqu/v9"
)

const (
	CardsTableName = "cards"
)

type Card struct {
	Id               int         `db:"id" goqu:"skipupdate,skipinsert"`
	Category         string      `db:"category"`
	NameRu           string      `db:"name_ru"`
	DescRu           string      `db:"desc_ru"`
	QuoteRu          string      `db:"quote_ru"`
	NameEn           string      `db:"name_en"`
	DescEn           string      `db:"desc_en"`
	QuoteEn          string      `db:"quote_en"`
	FactsRu          utils.JSONB `db:"facts_ru"`
	FactsEn          utils.JSONB `db:"facts_en"`
	AnswersEn        utils.JSONB `db:"answers_en"`
	AnswersRu        utils.JSONB `db:"answers_ru"`
	DrawingId        string      `db:"drawing_id"`
	PixelatedId      string      `db:"pixelated_id"`
	ScreenshotId     string      `db:"screenshot_id"`
	BackgroundColor1 string      `db:"bg_color_1"`
	BackgroundColor2 string      `db:"bg_color_2"`
	TextColor        string      `db:"text_color"`
	Completed        bool        `db:"completed"`
}

// GetCardById return card object by expression
func GetCardById(ctx context.Context, cardId int) (*Card, error) {
	var obj Card
	exists, err := persist.Db.From(CardsTableName).Where(
		goqu.Ex{"id": cardId},
	).ScanStructContext(ctx, &obj)
	if err != nil {
		return nil, err
	}
	if exists == false {
		return nil, errors.New("card not found")
	}
	return &obj, nil
}

// GetAllCards return all existing cards
func GetAllCards(ctx context.Context) ([]*Card, error) {
	var objs []*Card
	err := persist.Db.From(CardsTableName).Where(goqu.C("completed").IsTrue()).ScanStructsContext(ctx, &objs)
	if err != nil {
		return nil, err
	}
	return objs, nil
}

// Save card instance in DB
func (obj *Card) Save(ctx context.Context) error {
	var err error
	if obj.Id == 0 {
		err = obj.create(ctx)
	} else {
		err = obj.update(ctx)
	}
	if err != nil {
		return err
	}

	return nil
}

// createCard private method for create new cards DB record
func (obj *Card) create(ctx context.Context) error {
	insert := persist.Db.Insert(CardsTableName).
		Rows(obj).
		Returning("*").Executor()

	if _, err := insert.ScanStructContext(ctx, obj); err != nil {
		return err
	}
	return nil
}

// updateCard private method for update card record in DB
func (obj *Card) update(ctx context.Context) error {
	update := persist.Db.From(CardsTableName).
		Where(goqu.C("id").Eq(obj.Id)).Update().Set(obj).
		Executor()
	_, err := update.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete delete card from DB
func (obj *Card) Delete(ctx context.Context) error {
	_, err := persist.Db.From(CardsTableName).
		Where(goqu.Ex{"id": obj.Id}).
		Delete().
		Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
