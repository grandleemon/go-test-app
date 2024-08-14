package security

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(salt), nil
}
