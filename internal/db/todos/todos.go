package todos

import (
	"context"
	"fmt"
	"github.com/grandleemon/go-test-app.git/internal/db"
	"github.com/grandleemon/go-test-app.git/internal/models"
)

func GetAll() ([]models.Todo, error) {
	var todos []models.Todo

	query := "SELECT * FROM todos"

	rows, err := db.DbConn.Query(context.Background(), query)

	if err != nil {
		return todos, fmt.Errorf("failed to fetch todos: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted); err != nil {
			return todos, fmt.Errorf("failed to scan todo: %w", err)
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func GetByID(id int) (models.Todo, error) {
	var todo models.Todo

	query := `SELECT id, title, description, is_completed FROM todos WHERE id = $1`

	row := db.DbConn.QueryRow(context.Background(), query, id)

	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func Create(todo models.Todo) (int, error) {
	var id int

	query := "INSERT INTO todos (title, description) VALUES ($1, $2) RETURNING id"

	err := db.DbConn.QueryRow(context.Background(), query, todo.Title, todo.Description).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error inserting todo: %w", err)
	}

	return id, nil
}

func Update(id int, todo models.Todo) error {
	query := "UPDATE todos SET title=$1, description=$2, is_completed=$3 WHERE id=$4"

	_, err := db.DbConn.Exec(context.Background(), query, todo.Title, todo.Description, todo.IsCompleted, id)

	return err
}

func DeleteById(id int) error {
	query := "DELETE FROM todos WHERE id = $1"

	_, err := db.DbConn.Exec(context.Background(), query, id)

	return err
}
