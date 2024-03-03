package service

import (
	todoapp "github.com/spargapees/todo-app/todo-app"
	"github.com/spargapees/todo-app/todo-app/pkg/repository"
)

const salt = "akj474o2yojhgjkg08y24"

type Service interface {
	CreateUser(user todoapp.User) (int, error)
}

type service struct {
	repos repository.Repository
}

func NewService(repos repository.Repository) Service {
	return &service{
		repos: repos,
	}
}
