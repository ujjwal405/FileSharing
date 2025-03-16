package cache

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type JWKSCache struct {
	keys map[string]jwk.Key
}

func NewJWKSCache() *JWKSCache {
	return &JWKSCache{
		keys: make(map[string]jwk.Key),
	}
}

func (c *JWKSCache) FetchJWKS(userPoolID, region string) error {
	url := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch JWKS: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var rawKeys struct {
		Keys []json.RawMessage `json:"keys"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&rawKeys); err != nil {
		return fmt.Errorf("failed to decode JWKS: %v", err)
	}

	for _, rawKey := range rawKeys.Keys {
		key, err := jwk.ParseKey(rawKey)
		if err != nil {
			return fmt.Errorf("failed to parse JWK: %v", err)
		}
		kid, _ := key.Get("kid")
		c.keys[kid.(string)] = key
	}

	return nil
}

func (c *JWKSCache) GetKey(kid string) (jwk.Key, bool) {
	key, exists := c.keys[kid]
	return key, exists
}
