package handler

import (
	"context"

	"github.com/ujjwal405/FileSharing/logout/cognito"
	dynamo_db "github.com/ujjwal405/FileSharing/logout/dynamodb"
	"github.com/ujjwal405/FileSharing/logout/helper"
)

type LogoutHandler struct {
	cognito *cognito.CognitoClient
	dynamo  *dynamo_db.DynamoClient
}

func NewLambdaHandler(cc *cognito.CognitoClient, dc *dynamo_db.DynamoClient) *LogoutHandler {
	return &LogoutHandler{
		cognito: cc,
		dynamo:  dc,
	}
}

func (h *LogoutHandler) LogoutHandler(ctx context.Context, email string, googleLogin string) error {
	token, err := h.dynamo.GetRefreshToken(ctx, email)
	if err != nil {
		return err
	}
	if err := h.cognito.RevokeToken(ctx, token); err != nil {
		return err
	}
	expiredTime := helper.GenerateTime()
	if err := h.dynamo.Update(ctx, email, expiredTime); err != nil {
		return err
	}
	if googleLogin == "true" {
		if err := h.cognito.DeleteUser(ctx, email); err != nil {
			return err
		}
	}
	return nil
}
