package models

import (
	"context"
	"errors"
	"github.com/Sergey-pr/movie-games-tg/persist"
	"github.com/Sergey-pr/movie-games-tg/utils"
	"github.com/Sergey-pr/movie-games-tg/utils/jwt"
	"github.com/doug-martin/goqu/v9"
	"time"
)

const (
	UsersTableName = "users"
	UserContextKey = "users"
)

type User struct {
	Id            int         `db:"id" goqu:"skipupdate,skipinsert"`
	TelegramId    int         `db:"telegram_id"`
	Name          string      `db:"name"`
	UserName      string      `db:"username"`
	Language      string      `db:"language"`
	AnsweredCards utils.JSONB `db:"answered_cards"`
	IsAdmin       bool        `db:"is_admin"`
}

// LoginUser find user in DB and check password
func LoginUser(ctx context.Context, telegramId int) (*User, error) {
	user, err := GetUserByTelegramId(ctx, telegramId)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, "user", user)
	return user, nil
}

// GetUserByTelegramId return user object by expression
func GetUserByTelegramId(ctx context.Context, telegramId int) (*User, error) {
	var obj User
	exists, err := persist.Db.From(UsersTableName).Where(
		goqu.Ex{"telegram_id": telegramId},
	).ScanStructContext(ctx, &obj)
	if err != nil {
		return nil, err
	}
	if exists == false {
		return nil, errors.New("user not found")
	}
	return &obj, nil
}

// GetUserById return user object by expression
func GetUserById(ctx context.Context, userId int) (*User, error) {
	var obj User
	exists, err := persist.Db.From(UsersTableName).Where(
		goqu.Ex{"id": userId},
	).ScanStructContext(ctx, &obj)
	if err != nil {
		return nil, err
	}
	if exists == false {
		return nil, errors.New("user not found")
	}
	return &obj, nil
}

func (obj *User) GetBotProcessor(ctx context.Context, chatId int) (*CardProcessor, error) {
	var processor *CardProcessor
	exists, err := persist.Db.From(CardProcessorsTableName).Where(
		goqu.Ex{"user_id": obj.Id},
	).ScanStructContext(ctx, &processor)
	if err != nil {
		return nil, err
	}
	if !exists {
		processor = &CardProcessor{
			UserId: obj.Id,
			User:   obj,
			State:  0,
		}
		err = processor.Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	processor.ChatId = chatId
	processor.User = obj
	if processor.CardId != nil {
		var card *Card
		exists, err = persist.Db.From(CardsTableName).Where(
			goqu.Ex{"id": processor.CardId},
		).ScanStructContext(ctx, &card)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, errors.New("card not found")
		}
		processor.Card = card
	}
	return processor, nil
}

func (obj *User) GetJwtToken() (string, time.Time, error) {
	return jwt.GetJwtToken(&jwt.Claims{
		User: jwt.UserClaims{
			Id:            obj.Id,
			TelegramId:    obj.TelegramId,
			Name:          obj.Name,
			AnsweredCards: obj.AnsweredCards,
		},
	})
}

// Save user instance in DB
func (obj *User) Save(ctx context.Context) error {
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

// createUser private method for create new users DB record
func (obj *User) create(ctx context.Context) error {
	insert := persist.Db.Insert(UsersTableName).
		Rows(obj).
		Returning("*").Executor()

	if _, err := insert.ScanStructContext(ctx, obj); err != nil {
		return err
	}
	return nil
}

// updateUser private method for update user record in DB
func (obj *User) update(ctx context.Context) error {
	update := persist.Db.From(UsersTableName).
		Where(goqu.C("id").Eq(obj.Id)).Update().Set(obj).
		Executor()
	_, err := update.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete delete user from DB
func (obj *User) Delete(ctx context.Context) error {
	_, err := persist.Db.From(UsersTableName).
		Where(goqu.Ex{"id": obj.Id}).
		Delete().
		Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
