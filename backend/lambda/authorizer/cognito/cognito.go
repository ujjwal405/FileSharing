package my_cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	secret_manager "github.com/ujjwal405/FileSharing/authorizer/secret_manager"
)

type CognitoClient struct {
	cognitoClient *cognito.Client
	appClientID   string
	userPoolID    string
}

func NewCognitoClient() (*CognitoClient, error) {

	secretIDs := []string{"COGNITO_REGIONS", "APP_CLIENT_IDS", "USER_POOL_IDS"}
	secrets, err := secret_manager.GetSecrets(secretIDs)
	if err != nil {
		return nil, err
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(secrets["COGNITO_REGIONS"]))
	if err != nil {
		return nil, err
	}
	client := cognito.NewFromConfig(cfg)
	return &CognitoClient{
		cognitoClient: client,
		appClientID:   secrets["APP_CLIENT_IDS"],
		userPoolID:    secrets["USER_POOL_IDS"],
	}, nil

}
