package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	ssm "github.com/ujjwal405/FileSharing/logout/ssm"
)

type CognitoClient struct {
	cognitoClient *cognito.Client
	appClientID   string
	userPoolId    string
}

func NewCognitoClient() (*CognitoClient, error) {

	secretIDs := []string{"/myapp/cognito/region", "/myapp/cognito/appClientID", "/myapp/cognito/userPoolID"}
	secrets, err := ssm.GetParameters(secretIDs)
	if err != nil {
		return nil, err
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(secrets["/myapp/cognito/region"]))
	if err != nil {
		return nil, err
	}
	client := cognito.NewFromConfig(cfg)
	return &CognitoClient{
		cognitoClient: client,
		appClientID:   secrets["/myapp/cognito/appClientID"],
		userPoolId:    secrets["/myapp/cognito/userPoolID"],
	}, nil

}
