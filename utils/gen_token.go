package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomString(length int) string {
	tokenBytes := make([]byte, length)
	_, _ = rand.Read(tokenBytes) // Ignoring the error

	token := hex.EncodeToString(tokenBytes)
	return token
}
