package handlers

import (
	"github.com/grandleemon/go-test-app.git/internal/db/auth"
	"net/http"
)

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, err = auth.GetSessionByToken(cookie.Value)

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Access granted"))
}
