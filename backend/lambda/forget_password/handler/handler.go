package handler

import (
	"context"

	my_cognito "github.com/ujjwal405/FileSharing/forget_password/cognito"
)

type PasswordHandler struct {
	cognitoClient *my_cognito.CognitoClient
}

func NewPasswordHandler(cognitoClient *my_cognito.CognitoClient) *PasswordHandler {
	return &PasswordHandler{

		cognitoClient: cognitoClient,
	}
}

func (h *PasswordHandler) HandleForgetPassword(ctx context.Context, email string) error {

	if err := h.cognitoClient.ForgetPassword(ctx, email); err != nil {
		return err
	}
	return nil
}
