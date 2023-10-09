package models

import (
	"context"
	"errors"
	"github.com/Sergey-pr/movie-games-tg/persist"
	"github.com/Sergey-pr/movie-games-tg/utils/jwt"
	"github.com/doug-martin/goqu/v9"
)

const (
	UsersTableName = "users"
	UserContextKey = "users"
)

type User struct {
	Id         int     `db:"id" goqu:"skipupdate,skipinsert"`
	TelegramId int     `db:"telegram_id"`
	Name       string  `db:"name"`
	LastName   *string `db:"last_name"`
	UserName   string  `db:"username"`
	Language   string  `db:"language"`
	IsAdmin    bool    `db:"is_admin"`
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

// GetUsersCache returns cache of users by user id
func GetUsersCache(ctx context.Context, ids []int) (map[int]*User, error) {
	var objs []*User
	err := persist.Db.From(UsersTableName).Where(
		goqu.Ex{"id": ids},
	).ScanStructsContext(ctx, &objs)
	if err != nil {
		return nil, err
	}
	data := make(map[int]*User)
	for _, obj := range objs {
		data[obj.Id] = obj
	}
	return data, nil
}

// GetUserByTelegramId return user object by its telegram id
func GetUserByTelegramId(ctx context.Context, telegramId int) (*User, error) {
	var obj User
	exists, err := persist.Db.From(UsersTableName).Where(
		goqu.Ex{"telegram_id": telegramId},
	).ScanStructContext(ctx, &obj)
	if err != nil {
		return nil, err
	}
	if exists == false {
		return nil, nil
	}
	return &obj, nil
}

// GetUserById return user object by user id
func GetUserById(ctx context.Context, userId int) (*User, error) {
	var obj User
	exists, err := persist.Db.From(UsersTableName).Where(
		goqu.Ex{"id": userId},
	).ScanStructContext(ctx, &obj)
	if err != nil {
		return nil, err
	}
	if exists == false {
		return nil, nil
	}
	return &obj, nil
}

// GetBotProcessor return bot processor object by chat id
func (obj *User) GetBotProcessor(ctx context.Context, chatId int) (*CardProcessor, error) {
	var processor CardProcessor
	exists, err := persist.Db.From(CardProcessorsTableName).Where(
		goqu.Ex{"user_id": obj.Id},
	).ScanStructContext(ctx, &processor)
	if err != nil {
		return nil, err
	}
	if !exists {
		processor = CardProcessor{
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
		var card Card
		exists, err = persist.Db.From(CardsTableName).Where(
			goqu.Ex{"id": processor.CardId},
		).ScanStructContext(ctx, &card)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, errors.New("card not found")
		}
		processor.Card = &card
	}
	return &processor, nil
}

// GetJwtToken return JWT token to use as auth method
func (obj *User) GetJwtToken() (string, error) {
	return jwt.GetJwtToken(&jwt.Claims{
		User: jwt.UserClaims{
			Id:         obj.Id,
			TelegramId: obj.TelegramId,
			Name:       obj.Name,
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

// create private method to create new user DB record
func (obj *User) create(ctx context.Context) error {
	insert := persist.Db.Insert(UsersTableName).
		Rows(obj).
		Returning("*").Executor()

	if _, err := insert.ScanStructContext(ctx, obj); err != nil {
		return err
	}
	return nil
}

// update private method to update user record in DB
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

// Delete user from DB
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
