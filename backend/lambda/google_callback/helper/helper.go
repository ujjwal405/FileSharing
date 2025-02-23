package helper

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"
)

// GenerateTemporaryPassword generates a secure temporary password
func GenerateTemporaryPassword() string {
	const passwordLength = 8
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*()_+"
	var passwordBuilder strings.Builder

	// Generate a random password
	for i := 0; i < passwordLength; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		passwordBuilder.WriteByte(charset[index.Int64()])
	}

	return passwordBuilder.String()
}

func GenerateExpiryTime() int64 {
	return time.Now().Add(30 * 24 * time.Hour).Unix()
}
