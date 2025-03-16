package verifier

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ujjwal405/FileSharing/authorizer/cache"
)

type Verifier struct {
	cache      *cache.JWKSCache
	clientID   string
	region     string
	userPoolID string
}

func NewVerifier(cache *cache.JWKSCache, cleintID, region, userPoolID string) *Verifier {
	return &Verifier{
		cache:      cache,
		clientID:   cleintID,
		region:     region,
		userPoolID: userPoolID,
	}
}

func (v *Verifier) VerifyIDToken(tokenString string) (*IDToken, error) {
	claims := &IDToken{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("kid not found in token header")
		}

		key, exists := v.cache.GetKey(kid)
		if !exists {
			return nil, fmt.Errorf("key with kid %s not found in cache", kid)
		}

		var rsaKey rsa.PublicKey
		if err := key.Raw(&rsaKey); err != nil {
			return nil, fmt.Errorf("failed to convert JWK to RSA key: %v", err)
		}
		return &rsaKey, nil
	}, jwt.WithExpirationRequired())

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, fmt.Errorf("id token has expired")
		}
		return nil, fmt.Errorf("token verification failed: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Validate access token-specific claims
	expectedIssuer := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s", v.region, v.userPoolID)

	// Check issuer (iss)
	if claims.Issuer != expectedIssuer {
		return nil, fmt.Errorf("invalid issuer: got %s, expected %s", claims.Issuer, expectedIssuer)
	}

	// Check audience (aud)
	if !containsAudience(claims.Audience, v.clientID) {
		return nil, fmt.Errorf("invalid audience for id token: got %v, expected %s", claims.Audience, v.clientID)
	}

	// Check token_use
	if claims.TokenUse != "id" {
		return nil, fmt.Errorf("invalid token_use for access token: got %s, expected 'access'", claims.TokenUse)
	}

	return claims, nil

}

func (v *Verifier) VerifyAccessToken(tokenString string) (*AccessToken, error) {
	claims := &AccessToken{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("kid not found in token header")
		}

		key, exists := v.cache.GetKey(kid)
		if !exists {
			return nil, fmt.Errorf("key with kid %s not found in cache", kid)
		}

		var rsaKey rsa.PublicKey
		if err := key.Raw(&rsaKey); err != nil {
			return nil, fmt.Errorf("failed to convert JWK to RSA key: %v", err)
		}
		return &rsaKey, nil
	}, jwt.WithExpirationRequired()) // Automatically enforces expiration

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, fmt.Errorf("access token has expired")
		}
		return nil, fmt.Errorf("token verification failed: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Validate access token-specific claims
	expectedIssuer := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s", v.region, v.userPoolID)

	// Check issuer (iss)
	if claims.Issuer != expectedIssuer {
		return nil, fmt.Errorf("invalid issuer: got %s, expected %s", claims.Issuer, expectedIssuer)
	}

	// Check audience (aud)
	if !containsAudience(claims.Audience, v.clientID) {
		return nil, fmt.Errorf("invalid audience for access token: got %v, expected %s", claims.Audience, v.clientID)
	}

	// Check token_use
	if claims.TokenUse != "access" {
		return nil, fmt.Errorf("invalid token_use for access token: got %s, expected 'access'", claims.TokenUse)
	}

	return claims, nil

}

func containsAudience(audiences jwt.ClaimStrings, target string) bool {
	for _, aud := range audiences {
		if aud == target {
			return true
		}
	}
	return false
}
