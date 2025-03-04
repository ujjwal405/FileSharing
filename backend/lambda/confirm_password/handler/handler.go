package handler

import (
	"context"

	my_cognito "github.com/ujjwal405/FileSharing/confirm_password/cognito"
	"github.com/ujjwal405/FileSharing/confirm_password/confirmation"
)

type PasswordHandler struct {
	cognitoClient *my_cognito.CognitoClient
}

func NewLambdaHandler(cognitoClient *my_cognito.CognitoClient) *PasswordHandler {
	return &PasswordHandler{

		cognitoClient: cognitoClient,
	}
}

func (h *PasswordHandler) HandleConformPassword(ctx context.Context, confirm confirmation.Confirmation) error {

	if err := h.cognitoClient.ConfirmForgotPassword(ctx, confirm.Email, confirm.ConfirmationCode, confirm.Password); err != nil {
		return err
	}
	return nil
}
