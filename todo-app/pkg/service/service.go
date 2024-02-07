package service

import "github.com/spargapees/todo-app/todo-app/pkg/repository"

type Service interface {
}

type service struct {
	repos repository.Repository
}

func NewService(repos repository.Repository) Service {
	return &service{
		repos: repos,
	}
}
