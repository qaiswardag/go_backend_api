package tokengen

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateRandomToken(length int) string {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate the token: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
