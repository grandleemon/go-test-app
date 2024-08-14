package router

import (
	"github.com/grandleemon/go-test-app.git/internal/router/auth"
	"github.com/grandleemon/go-test-app.git/internal/router/todos"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	todosroutes.Register(mux)
	authroutes.Register(mux)

	return mux
}
