package todoshandlers

import (
	"encoding/json"
	"github.com/grandleemon/go-test-app.git/internal/db/todos"
	"github.com/grandleemon/go-test-app.git/internal/models"
	"net/http"
	"strconv"
	"strings"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	_todos, err := todos.GetAll()

	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(_todos)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := todos.Create(todo)

	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	todo.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func Update(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")

	idStr := pathSegments[len(pathSegments)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	existingTodo, existingTodoErr := todos.GetByID(id)
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

	if err := todos.Update(id, existingTodo); err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	updatedTodo, updatedTodoErr := todos.GetByID(id)
	if updatedTodoErr != nil {
		http.Error(w, "Failed to retrieve updated todo", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedTodo)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")

	idStr := pathSegments[len(pathSegments)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, existingTodoErr := todos.GetByID(id)

	if existingTodoErr != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	deleteTodoErr := todos.DeleteById(id)

	if deleteTodoErr != nil {
		http.Error(w, "Failed to delete todo:"+deleteTodoErr.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo successfully deleted"))
}
