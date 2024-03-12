package service

import todoapp "github.com/spargapees/todo-app/todo-app"

func (s *service) CreateList(userId int, list todoapp.TodoList) (int, error) {
	return s.repos.CreateList(userId, list)
}

func (s *service) GetAllLists(userId int) ([]todoapp.TodoList, error) {
	return s.repos.GetAllLists(userId)
}

func (s *service) GetListById(userId, listId int) (todoapp.TodoList, error) {
	return s.repos.GetListById(userId, listId)
}

func (s *service) UpdateList(userId, listId int, input todoapp.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repos.UpdateList(userId, listId, input)
}

func (s *service) DeleteList(userId, listId int) error {
	return s.repos.DeleteList(userId, listId)
}
