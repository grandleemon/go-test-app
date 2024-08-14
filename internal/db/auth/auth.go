package auth

import (
	"context"
	_ "database/sql"
	"github.com/grandleemon/go-test-app.git/internal/db"
	"github.com/grandleemon/go-test-app.git/pkg/security"
)

func Register(email, password string) error {
	salt, err := security.GenerateSalt(16)

	if err != nil {
		return err
	}

	hashedPassword := security.HashPassword(password, salt)

	query := "INSERT INTO users (email, password, salt) VALUES ($1, $2, $3)"

	_, insertErr := db.DbConn.Exec(context.Background(), query, email, hashedPassword, salt)

	if insertErr != nil {
		return err
	}

	return nil
}

func Login(email, password string) (bool, error) {
	var hashedPassword, salt string

	query := "SELECT password, salt FROM users WHERE email = $1"

	err := db.DbConn.QueryRow(context.Background(), query, email).Scan(&hashedPassword, &salt)

	if err != nil {
		return false, err
	}

	valid := security.VerifyPassword(hashedPassword, salt, password)

	return valid, nil
}
