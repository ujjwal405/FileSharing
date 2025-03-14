package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateStateToken creates a signed JWT as the state token
func GenerateStateToken(jwtSecret string) (string, error) {
	// Create claims with expiry
	claims := jwt.MapClaims{
		"exp": time.Now().Add(5 * time.Minute).Unix(), // Token expires in 5 minutes
		"iat": time.Now().Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	return token.SignedString(jwtSecret)
}
