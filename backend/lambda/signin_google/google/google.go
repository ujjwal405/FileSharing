package google

import (
	ssm "github.com/ujjwal405/FileSharing/signin_google/ssm"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func InitGoogleConfig() (*oauth2.Config, error) {

	secretIDs := []string{"/myapp/google/googleClientID", "/myapp/google/googleClientSecret", "/myapp/google/googleRedirectURL"}
	secrets, err := ssm.GetParameters(secretIDs)
	if err != nil {
		return nil, err
	}
	return &oauth2.Config{
		ClientID:     secrets["/myapp/google/googleClientID"],
		ClientSecret: secrets["/myapp/google/googleClientSecret"],
		RedirectURL:  secrets["/myapp/google/googleRedirectURL"], // Should be your callback Lambda URL
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}, nil

}
