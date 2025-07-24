package google

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	ssm "github.com/ujjwal405/FileSharing/google_callback/ssm"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleConfig struct {
	OAuthConfig *oauth2.Config
}

func InitGoogleConfig() (*GoogleConfig, error) {

	secretIDs := []string{"/myapp/google/googleClientID", "/myapp/google/googleClientSecret", "/myapp/google/googleRedirectURL"}
	secrets, err := ssm.GetParameters(secretIDs)
	if err != nil {
		return nil, err
	}
	oconfig := &oauth2.Config{
		ClientID:     secrets["/myapp/google/googleClientID"],
		ClientSecret: secrets["/myapp/google/googleClientSecret"],
		RedirectURL:  secrets["/myapp/google/googleRedirectURL"], // Should be your callback Lambda URL
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return &GoogleConfig{OAuthConfig: oconfig}, nil

}

func (c *GoogleConfig) Callback(ctx context.Context, code string) (string, string, string, error) {
	token, err := c.OAuthConfig.Exchange(ctx, code)
	if err != nil {
		return "", "", "", err
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return "", "", "", err
	}

	defer resp.Body.Close()
	userdata, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", err
	}
	var data googleInfo
	json.Unmarshal(userdata, &data)

	return data.Email, data.GivenName, data.FamilyName, nil

}
