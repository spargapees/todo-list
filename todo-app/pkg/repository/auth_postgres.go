package repository

import (
	"fmt"

	todoapp "github.com/spargapees/todo-app/todo-app"
)

func (r *repository) CreateUser(user todoapp.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) GetUser(username, password string) (todoapp.User, error) {
	var user todoapp.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
