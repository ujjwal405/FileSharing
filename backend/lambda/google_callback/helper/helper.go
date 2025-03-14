package helper

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
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

func VerifyStateToken(stateToken, jwtSecret string) error {
	// Parse and validate the token
	token, err := jwt.Parse(stateToken, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token uses the correct signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, UnexpectedSigningMethod()
		}
		return jwtSecret, nil
	})

	if err != nil {
		return err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return InvalidStateToken()
	}

	// Check expiry
	if exp, ok := claims["exp"].(float64); ok {
		if int64(exp) < time.Now().Unix() {
			return TokenExpired()
		}
	} else {
		return InvalidExpirationTime()
	}

	return nil
}
