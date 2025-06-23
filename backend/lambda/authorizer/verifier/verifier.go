package verifier

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	apierror "github.com/ujjwal405/FileSharing/authorizer/apiError"
	"github.com/ujjwal405/FileSharing/authorizer/cache"
	secretmanager "github.com/ujjwal405/FileSharing/authorizer/secret_manager"
)

type Verifier struct {
	cache      *cache.JWKSCache
	clientID   string
	region     string
	userPoolID string
}

func NewVerifier() (*Verifier, error) {

	secrets, err := secretmanager.GetSecrets([]string{"COGNITO_REGIONS", "APP_CLIENT_IDS", "USER_POOL_IDS"})
	if err != nil {
		return nil, err
	}
	mycache := cache.NewJWKSCache()
	return &Verifier{
		cache:      mycache,
		clientID:   secrets["APP_CLIENT_IDS"],
		region:     secrets["COGNITO_REGIONS"],
		userPoolID: secrets["USER_POOL_IDS"],
	}, nil
}

func (v *Verifier) VerifyIDToken(tokenString string) (*IDToken, error) {
	claims := &IDToken{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, apierror.InvalidSignRequest()
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, apierror.NewAPIError(400, "kid not found in token header")
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
			return nil, apierror.ErrTokenExpired
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
		return nil, apierror.NewAPIError(400, "unknown issuer")
	}

	// Check audience (aud)
	if !containsAudience(claims.Audience, v.clientID) {
		return nil, apierror.NewAPIError(400, "unknown audience")
	}

	// Check token_use
	if claims.TokenUse != "id" {
		return nil, apierror.NewAPIError(400, "unknown token use")
	}

	return claims, nil

}

func (v *Verifier) VerifyAccessToken(tokenString string) (*AccessToken, error) {
	claims := &AccessToken{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, apierror.InvalidSignRequest()
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, apierror.NewAPIError(400, "kid not found in token header")
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
			return nil, apierror.ErrTokenExpired
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
		return nil, apierror.NewAPIError(400, "unknown issuer")
	}

	// Check audience (aud)
	if !containsAudience(claims.Audience, v.clientID) {
		return nil, apierror.NewAPIError(400, "unknown audience")
	}

	// Check token_use
	if claims.TokenUse != "access" {
		return nil, apierror.NewAPIError(400, "unknown token use")
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
