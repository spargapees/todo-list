package repository

import (
	"github.com/jmoiron/sqlx"
	todoapp "github.com/spargapees/todo-app/todo-app"
)

type Repository interface {
	CreateUser(user todoapp.User) (int, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
