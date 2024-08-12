package handlers

import (
	"encoding/json"
	"github.com/grandleemon/go-test-app.git/internal/db"
	"github.com/grandleemon/go-test-app.git/internal/models"
	"net/http"
	"strconv"
	"strings"
)

func GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := db.GetAllTodos()

	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")

	idStr := pathSegments[len(pathSegments)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	existingTodo, existingTodoErr := db.GetTodoByID(id)
	if existingTodoErr != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	var updates models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if updates.Title != "" {
		existingTodo.Title = updates.Title
	}
	if updates.Description != "" {
		existingTodo.Description = updates.Description
	}
	existingTodo.IsCompleted = updates.IsCompleted

	if err := db.UpdateTodo(id, existingTodo); err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	updatedTodo, updatedTodoErr := db.GetTodoByID(id)
	if updatedTodoErr != nil {
		http.Error(w, "Failed to retrieve updated todo", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")

	idStr := pathSegments[len(pathSegments)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, existingTodoErr := db.GetTodoByID(id)

	if existingTodoErr != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	deleteTodoErr := db.DeleteTodoById(id)

	if deleteTodoErr != nil {
		http.Error(w, "Failed to delete todo:"+deleteTodoErr.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo successfully deleted"))
}
