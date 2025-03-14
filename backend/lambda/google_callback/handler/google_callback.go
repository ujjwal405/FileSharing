package handler

import (
	"context"

	my_cognito "github.com/ujjwal405/FileSharing/google_callback/cognito"
	my_dynamo "github.com/ujjwal405/FileSharing/google_callback/dynamo_db"
	my_google "github.com/ujjwal405/FileSharing/google_callback/google"
	"github.com/ujjwal405/FileSharing/google_callback/helper"
)

type LambdaHandler struct {
	googleOauthConfig *my_google.GoogleConfig
	cognitoClient     *my_cognito.CognitoClient
	dynamoClient      *my_dynamo.DynamoClient
	secretKey         string
}

func NewLambdaHandler(googleOauthConfig *my_google.GoogleConfig, cognitoClient *my_cognito.CognitoClient, dynamoClient *my_dynamo.DynamoClient, secretKey string) *LambdaHandler {
	return &LambdaHandler{
		googleOauthConfig: googleOauthConfig,
		cognitoClient:     cognitoClient,
		dynamoClient:      dynamoClient,
		secretKey:         secretKey,
	}
}

func (h *LambdaHandler) HandleGoogleCallback(ctx context.Context, code string, stateToken string) (string, string, error) {
	if err := helper.VerifyStateToken(stateToken, h.secretKey); err != nil {
		return "", "", err
	}
	email, givenName, familyName, err := h.googleOauthConfig.Callback(ctx, code)
	if err != nil {
		return "", "", err
	}
	exist, err := h.cognitoClient.CheckUser(ctx, email)
	if err != nil {
		return "", "", err
	}
	if exist {
		return "", "", helper.UserAlreadyExistsError()
	}
	temp_pass := helper.GenerateTemporaryPassword()
	err = h.cognitoClient.CreateUser(ctx, email, givenName, familyName, temp_pass)
	if err != nil {
		return "", "", err
	}
	user_token, err := h.cognitoClient.Authenticate(ctx, email, temp_pass)
	if err != nil {
		return "", "", err
	}
	expires_at := helper.GenerateExpiryTime()

	if err := h.dynamoClient.PutUser(ctx, email, user_token.RefreshToken, expires_at); err != nil {
		if delErr := h.cognitoClient.DeleteUser(ctx, email); delErr != nil {
			err = delErr
		}
		return "", "", err
	}

	return user_token.IDToken, user_token.AccessToken, nil

}
