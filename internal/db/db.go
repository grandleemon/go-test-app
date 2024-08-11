package db

import (
	"context"
	"fmt"
	"github.com/grandleemon/go-test-app.git/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

var dbConn *pgx.Conn

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connString := os.Getenv("DATABASE_URL")

	if connString == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	dbConn, err = pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Println("Connected to database")
}

func Close() {
	if dbConn != nil {
		dbConn.Close(context.Background())
	}
}

func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo

	query := "SELECT * FROM todos"

	rows, err := dbConn.Query(context.Background(), query)

	if err != nil {
		return todos, fmt.Errorf("failed to fetch todos: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted); err != nil {
			log.Println("ERROR", err)
			return todos, fmt.Errorf("failed to scan todo: %w", err)
		}

		todos = append(todos, todo)
	}

	log.Println("HERE", todos)

	return todos, nil
}

func CreateTodo(todo models.Todo) (int, error) {
	var id int
	query := "INSERT INTO todos (title, description) VALUES ($1, $2) RETURNING id"

	err := dbConn.QueryRow(context.Background(), query, todo.Title, todo.Description).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting todo: %w", err)
	}

	return id, nil
}
