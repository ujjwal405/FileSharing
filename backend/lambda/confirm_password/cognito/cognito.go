package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"

	secret_manager "github.com/ujjwal405/FileSharing/confirm_password/secret_manager"
)

type CognitoClient struct {
	cognitoClient *cognito.Client
	appClientID   string
}

func NewCognitoClient() (*CognitoClient, error) {
	secretIDs := []string{"COGNITO_REGION", "APP_CLIENT_ID"}
	secrets, err := secret_manager.GetSecrets(secretIDs)
	if err != nil {
		return nil, err
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(secrets["COGNITO_REGION"]))
	if err != nil {
		return nil, err
	}
	client := cognito.NewFromConfig(cfg)

	return &CognitoClient{
		cognitoClient: client,
		appClientID:   secrets["APP_CLIENT_ID"],
	}, nil
}
