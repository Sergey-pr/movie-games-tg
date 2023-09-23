package models

import (
	"context"
	"github.com/Sergey-pr/movie-games-tg/persist"
	"github.com/doug-martin/goqu/v9"
)

const (
	UsersTableName = "users"
	UserContextKey = "users"
)

type User struct {
	Id int `db:"id" goqu:"skipupdate,skipinsert"`
}

func (obj *User) GetId() int {
	return obj.Id
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
