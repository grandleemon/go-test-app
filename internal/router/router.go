package router

import (
	"github.com/grandleemon/go-test-app.git/internal/handlers"
	"github.com/grandleemon/go-test-app.git/internal/router/auth"
	"github.com/grandleemon/go-test-app.git/internal/router/todos"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	todosroutes.Register(mux)
	authroutes.Register(mux)

	mux.HandleFunc("/api/secure", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.SecureHandler(w, r)
		}
	})

	return mux
}
