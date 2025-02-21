package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	secret_manager "github.com/ujjwal405/FileSharing/google_callback/secret_manager"
)

type cognitoClient struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientID   string
	userPoolID    string
}

func NewCognitoClient() (*cognitoClient, error) {
	secretIDs := []string{"COGNITO_REGION", "APP_CLIENT_ID", "USER_POOL_ID"}
	secrets, err := secret_manager.GetSecrets(secretIDs)

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(secrets["COGNITO_REGION"]))
	if err != nil {
		return nil, err
	}
	client := cognito.NewFromConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &cognitoClient{
		cognitoClient: client,
		appClientID:   secrets["APP_CLIENT_ID"],
		userPoolID:    secrets["USER_POOL_ID"],
	}, nil
}
