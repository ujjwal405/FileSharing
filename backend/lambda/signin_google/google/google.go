package google

import (
	"github.com/ujjwal405/FileSharing/signin_google/secret_manager"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func InitGoogleConfig() (*oauth2.Config, error) {

	secretIDs := []string{"GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "GOOGLE_REDIRECT_URL"}
	secrets, err := secret_manager.GetSecrets(secretIDs)
	if err != nil {
		return nil, err
	}
	return &oauth2.Config{
		ClientID:     secrets["GOOGLE_CLIENT_ID"],
		ClientSecret: secrets["GOOGLE_CLIENT_SECRET"],
		RedirectURL:  secrets["GOOGLE_REDIRECT_URL"], // Should be your callback Lambda URL
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}, nil

}
