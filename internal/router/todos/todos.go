package todosroutes

import (
	"github.com/grandleemon/go-test-app.git/internal/handlers/todos"
	"net/http"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoshandlers.GetAll(w, r)
		case http.MethodPost:
			todoshandlers.Create(w, r)
		}
	})

	mux.HandleFunc("/api/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			todoshandlers.Update(w, r)
		case http.MethodDelete:
			todoshandlers.Delete(w, r)
		}
	})
}
