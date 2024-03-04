package service

import (
	"time"

	todoapp "github.com/spargapees/todo-app/todo-app"
	"github.com/spargapees/todo-app/todo-app/pkg/repository"
)

const (
	salt       = "akj474o2yojhgjkg08y24"
	signingKey = "qrugbjkfnl#%&7842KUHDFL868HDSslavfd"
	tokenTTL   = 12 * time.Hour
)

type Service interface {
	CreateUser(user todoapp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	CreateList(userId int, list todoapp.TodoList) (int, error)
	GetAllLists(userId int) ([]todoapp.TodoList, error)
	GetListById(userId, listId int) (todoapp.TodoList, error)
}

type service struct {
	repos repository.Repository
}

func NewService(repos repository.Repository) Service {
	return &service{
		repos: repos,
	}
}
