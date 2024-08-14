package authroutes

import (
	"github.com/grandleemon/go-test-app.git/internal/handlers/auth"
	"net/http"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/users/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			authhandlers.Register(w, r)
		}
	})

	mux.HandleFunc("/api/users/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			authhandlers.Login(w, r)
		}
	})
}
