package handler

import (
	"context"

	cognito "github.com/ujjwal405/FileSharing/code_verification/cognito"
	"github.com/ujjwal405/FileSharing/code_verification/user"
)

type VerifyHandler struct {
	cognitoClient *cognito.CognitoClient
}

func NewVerifyHandler(cClient *cognito.CognitoClient) *VerifyHandler {
	return &VerifyHandler{
		cognitoClient: cClient,
	}
}
func (h *VerifyHandler) VerifyHandler(ctx context.Context, confirmInput user.UserConfirmation) error {
	if err := h.cognitoClient.ConfirmAccount(ctx, confirmInput); err != nil {
		return err
	}

	if err := h.cognitoClient.SetEmailVerified(ctx, confirmInput.Email); err != nil {
		return err
	}
	return nil

}
