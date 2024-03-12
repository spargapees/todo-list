package repository

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	todoapp "github.com/spargapees/todo-app/todo-app"
)

func (r *repository) CreateItem(listId int, item todoapp.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return -1, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		return -1, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	return itemId, tx.Commit()
}

func (r *repository) GetAllItems(userId, listId int) ([]todoapp.TodoItem, error) {
	var items []todoapp.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON li.list_id = ul.list_id WHERE li.list_id = $1 AND ul.user_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *repository) GetItemById(userId, itemId int) (todoapp.TodoItem, error) {
	var item todoapp.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on ti.id = li.item_id INNER JOIN %s ul ON li.list_id = ul.list_id WHERE li.item_id = $1 AND ul.user_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil

}

func (r *repository) DeleteItem(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *repository) UpdateItem(userId, itemId int, input todoapp.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND ul.list_id = li.list_id AND ul.user_id = $%d AND ti.id = $%d",
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)

	args = append(args, userId, itemId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
