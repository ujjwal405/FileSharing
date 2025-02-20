package helper

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateStateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("Error generating state token:", err)
		return "default-state"
	}
	return base64.URLEncoding.EncodeToString(b)
}
