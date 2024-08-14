package router

import (
	"github.com/grandleemon/go-test-app.git/internal/router/todos"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	todosroutes.Register(mux)

	return mux
}
