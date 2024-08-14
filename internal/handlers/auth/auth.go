package authhandlers

import (
	"encoding/json"
	"github.com/grandleemon/go-test-app.git/internal/db/auth"
	"github.com/grandleemon/go-test-app.git/internal/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	userId, registerErr := auth.Register(user.Email, user.Password)

	if registerErr != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	session, sessionErr := auth.CreateSession(userId)

	if sessionErr != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.SessionToken,
		Expires:  session.ExpiresAt,
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	valid, loginErr := auth.Login(user.Email, user.Password)

	if loginErr != nil || !valid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Login successfully"))
}
