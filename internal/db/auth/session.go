package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"github.com/grandleemon/go-test-app.git/internal/db"
	"github.com/grandleemon/go-test-app.git/internal/models"
	"log"
	"time"
)

func GenerateSessionToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

func CreateSession(userID int) (*models.Session, error) {
	token, err := GenerateSessionToken()

	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(24 * time.Hour)

	session := &models.Session{
		UserID:       userID,
		SessionToken: token,
		ExpiresAt:    expiresAt,
	}

	query := `INSERT INTO sessions (user_id, session_token, expires_at) 
              VALUES ($1, $2, $3) RETURNING id, created_at`

	log.Println(userID)

	err = db.DbConn.QueryRow(context.Background(), query, userID, token, expiresAt).Scan(&session.ID, &session.CreatedAt)

	log.Println(err)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func GetSessionByToken(token string) (*models.Session, error) {
	session := &models.Session{}

	log.Println(token)

	query := "SELECT id, user_id, session_token, created_at, expires_at FROM sessions WHERE session_token = $1"

	err := db.DbConn.QueryRow(context.Background(), query, token).Scan(&session.ID, &session.UserID, &session.SessionToken, &session.CreatedAt, &session.ExpiresAt)

	log.Println(err)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("session not found")
		}

		return nil, err
	}

	if session.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("session expired")
	}

	return session, nil
}

func DeleteSession(token string) error {
	query := "DELETE FROM sessions WHERE session_token = $1"

	_, err := db.DbConn.Exec(context.Background(), query, token)

	return err
}
