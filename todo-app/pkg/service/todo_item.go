package service

import todoapp "github.com/spargapees/todo-app/todo-app"

func (s *service) CreateItem(userId int, listId int, item todoapp.TodoItem) (int, error) {
	_, err := s.repos.GetListById(userId, listId)
	if err != nil {
		return -1, err
	}

	return s.repos.CreateItem(listId, item)
}

func (s *service) GetAllItems(userId, listId int) ([]todoapp.TodoItem, error) {
	return s.repos.GetAllItems(userId, listId)
}

func (s *service) GetItemById(userId, itemId int) (todoapp.TodoItem, error) {
	return s.repos.GetItemById(userId, itemId)
}

func (s *service) DeleteItem(userId, itemId int) error {
	return s.repos.DeleteItem(userId, itemId)
}

func (s *service) UpdateItem(userId, itemId int, input todoapp.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repos.UpdateItem(userId, itemId, input)
}
