package cognito

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"

	"github.com/ujjwal405/FileSharing/google_callback/helper"
	secret_manager "github.com/ujjwal405/FileSharing/google_callback/secret_manager"
)

type CognitoClient struct {
	cognitoClient *cognito.Client
	appClientID   string
	userPoolID    string
}

func NewCognitoClient() (*CognitoClient, error) {
	secretIDs := []string{"COGNITO_REGION", "APP_CLIENT_ID", "USER_POOL_ID"}
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
		userPoolID:    secrets["USER_POOL_ID"],
	}, nil
}

func (c *CognitoClient) CheckUser(ctx context.Context, email string) (bool, error) {
	input := &cognito.AdminGetUserInput{
		UserPoolId: &c.userPoolID,
		Username:   &email,
	}
	_, err := c.cognitoClient.AdminGetUser(ctx, input)
	var notFoundErr *types.UserNotFoundException
	if errors.As(err, &notFoundErr) {
		return false, nil // User does not exist
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *CognitoClient) ExistUser(ctx context.Context, email string) error {
	input := &cognito.AdminCreateUserInput{
		UserPoolId: &c.userPoolID,
		Username:   &email,
	}
	_, err := c.cognitoClient.AdminCreateUser(ctx, input)
	return err
}

func (c *CognitoClient) CreateUser(ctx context.Context, email, firstName, lastName string) error {
	input := &cognito.AdminCreateUserInput{
		UserPoolId: aws.String(c.userPoolID),
		Username:   aws.String(email), // Using email as the username
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(email)},
			{Name: aws.String("first_name"), Value: aws.String(firstName)},  // First Name
			{Name: aws.String("last_name"), Value: aws.String(lastName)},    // Last Name
			{Name: aws.String("email_verified"), Value: aws.String("true")}, // Mark email as verified
		},
		TemporaryPassword:  aws.String(helper.GenerateTemporaryPassword()),
		ForceAliasCreation: false,                           // No need to force alias creation since we're checking first
		MessageAction:      types.MessageActionTypeSuppress, // Suppresses email invitation (optional)
	}

	_, err := c.cognitoClient.AdminCreateUser(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (c *CognitoClient) Authenticate(email, password string) (Token, error) {
	authInput := &cognito.AdminInitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": user.Email,
			"PASSWORD": user.Password,
		}),
		ClientId: aws.String(c.appClientID),
	}

	result, err := c.cognitoClient.InitiateAuth(authInput)
	if err != nil {
		return nil, err
	}
	var token Token
	token.AccessToken = *result.AuthenticationResult.AccessToken
	token.RefreshToken = *result.AuthenticationResult.RefreshToken
	token.IDToken = *result.AuthenticationResult.IdToken
	return token, nil
}
