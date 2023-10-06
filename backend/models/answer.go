package models

import (
	"context"
	"github.com/Sergey-pr/movie-games-tg/persist"
	"github.com/doug-martin/goqu/v9"
)

const (
	AnswersTableName = "answers"
)

type Answer struct {
	Id     int `db:"id" goqu:"skipupdate,skipinsert"`
	CardId int `db:"card_id"`
	UserId int `db:"user_id"`
	Points int `db:"points"`
}

type UserData struct {
	UserId      int `db:"user_id"`
	TotalPoints int `db:"total_points"`
}

// GetLeaderboardData return count of points per user
func GetLeaderboardData(ctx context.Context) ([]*UserData, error) {
	var data []*UserData
	err := persist.Db.From(AnswersTableName).
		Select(goqu.C("user_id"), goqu.SUM(goqu.I("points")).As("total_points")).
		GroupBy(goqu.C("user_id")).
		Order(goqu.C("total_points").Desc()).ScanStructsContext(ctx, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetAnswerByCardIdAndUserId return answer object by card id and user id
func GetAnswerByCardIdAndUserId(ctx context.Context, cardId int, userId int) (*Answer, error) {
	var obj Answer
	exists, err := persist.Db.From(AnswersTableName).Where(
		goqu.Ex{"card_id": cardId, "user_id": userId},
	).ScanStructContext(ctx, &obj)
	if err != nil {
		return nil, err
	}
	if exists == false {
		return nil, nil
	}
	return &obj, nil
}

// Save answer instance in DB
func (obj *Answer) Save(ctx context.Context) error {
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

// create private method to create new answer DB record
func (obj *Answer) create(ctx context.Context) error {
	insert := persist.Db.Insert(AnswersTableName).
		Rows(obj).
		Returning("*").Executor()

	if _, err := insert.ScanStructContext(ctx, obj); err != nil {
		return err
	}
	return nil
}

// update private method to update answer record in DB
func (obj *Answer) update(ctx context.Context) error {
	update := persist.Db.From(AnswersTableName).
		Where(goqu.C("id").Eq(obj.Id)).Update().Set(obj).
		Executor()
	_, err := update.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete delete answer from DB
func (obj *Answer) Delete(ctx context.Context) error {
	_, err := persist.Db.From(AnswersTableName).
		Where(goqu.Ex{"id": obj.Id}).
		Delete().
		Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
