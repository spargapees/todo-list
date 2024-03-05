package repository

import (
	"github.com/jmoiron/sqlx"
	todoapp "github.com/spargapees/todo-app/todo-app"
)

type Repository interface {
	CreateUser(user todoapp.User) (int, error)
	GetUser(username, password string) (todoapp.User, error)

	CreateList(userId int, list todoapp.TodoList) (int, error)
	GetAllLists(userId int) ([]todoapp.TodoList, error)
	GetListById(userId, listId int) (todoapp.TodoList, error)
	UpdateList(userId, listId int, input todoapp.UpdateListInput) error
	DeleteList(userId, listId int) error

	CreateItem(listId int, item todoapp.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]todoapp.TodoItem, error)
	GetItemById(userId, itemId int) (todoapp.TodoItem, error)
	DeleteItem(userid, itemId int) error
	UpdateItem(userId, listId int, input todoapp.UpdateItemInput) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
