package router

import (
	"github.com/grandleemon/go-test-app.git/internal/handlers"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/todos/", handlers.GetAllTodos)

	return mux
}
