package handlers

import (
	"encoding/json"
	"github.com/grandleemon/go-test-app.git/internal/db"
	"github.com/grandleemon/go-test-app.git/internal/models"
	"log"
	"net/http"
)

func GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := db.GetAllTodos()

	log.Println(err)

	if err != nil {
		http.Error(w, "Failed to fetch todos2222", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := db.CreateTodo(todo)

	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	todo.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
