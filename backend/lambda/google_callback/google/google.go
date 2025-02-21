package google

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	secret_manager "github.com/ujjwal405/FileSharing/google_callback/secret_manager"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleConfig struct {
	OAuthConfig *oauth2.Config
}

func initGoogleConfig() (*GoogleConfig, error) {

	secretIDs := []string{"GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "GOOGLE_REDIRECT_URL"}
	secrets, err := secret_manager.GetSecrets(secretIDs)
	if err != nil {
		return nil, err
	}
	oconfig := &oauth2.Config{
		ClientID:     secrets["GOOGLE_CLIENT_ID"],
		ClientSecret: secrets["GOOGLE_CLIENT_SECRET"],
		RedirectURL:  secrets["GOOGLE_REDIRECT_URL"], // Should be your callback Lambda URL
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return &GoogleConfig{OAuthConfig: oconfig}, nil

}

func (c *GoogleConfig) Callback(ctx context.Context, code string) (string, error) {
	token, err := c.OAuthConfig.Exchange(ctx, code)
	if err != nil {
		return "", err
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	userdata, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var data map[string]interface{}
	json.Unmarshal(userdata, &data)
	email := data["email"].(string)
	return email, nil
}
