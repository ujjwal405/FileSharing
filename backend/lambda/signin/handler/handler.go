package handler

import (
	"context"

	cognito "github.com/ujjwal405/FileSharing/signin/cognito"
	"github.com/ujjwal405/FileSharing/signin/dynamo_db"
	"github.com/ujjwal405/FileSharing/signin/helper"
)

type SignInHandler struct {
	cognitoClient *cognito.CognitoClient
	dynamoClient  *dynamo_db.DynamoClient
}

func NewSignInHandler(cc *cognito.CognitoClient, dc *dynamo_db.DynamoClient) *SignInHandler {
	return &SignInHandler{
		cognitoClient: cc,
		dynamoClient:  dc,
	}
}

func (h *SignInHandler) SignInUser(ctx context.Context, usr cognito.User) (string, string, error) {
	token, err := h.cognitoClient.UserLogin(ctx, usr)
	if err != nil {
		return "", "", err
	}

	expires_at := helper.GenerateExpiryTime()
	if err := h.dynamoClient.PutUser(ctx, usr.Email, token.RefreshToken, expires_at); err != nil {
		return "", "", err
	}
	return token.AccessToken, token.IdToken, nil
}
