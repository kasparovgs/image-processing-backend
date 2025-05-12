package service

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func CreateSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	sessionID := base64.URLEncoding.EncodeToString(b)
	return sessionID, nil
}
