package verifier

import "github.com/golang-jwt/jwt/v5"

type IDToken struct {
	TokenUse    string `json:"token_use"`
	Username    string `json:"cognito:username"`
	GoogleLogin string `json:"custom:google_login"`
	jwt.RegisteredClaims
}

type AccessToken struct {
	TokenUse string `json:"token_use"`
	jwt.RegisteredClaims
}
