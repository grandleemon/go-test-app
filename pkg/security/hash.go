package security

import (
	"encoding/base64"
	"golang.org/x/crypto/argon2"
)

func HashPassword(password, salt string) string {
	hash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	return base64.RawStdEncoding.EncodeToString(hash)
}

func VerifyPassword(hashedPassword, salt, password string) bool {
	expectedHash := HashPassword(password, salt)
	return hashedPassword == expectedHash
}
