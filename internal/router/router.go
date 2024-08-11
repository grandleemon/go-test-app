package router

import (
	"github.com/grandleemon/go-test-app.git/internal/handlers"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAllTodosHandler(w, r)
		case http.MethodPost:
			handlers.CreateTodoHandler(w, r)
		}
	})

	return mux
}
