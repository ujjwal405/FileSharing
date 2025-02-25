package handler

import (
	"context"

	apiError "github.com/ujjwal405/FileSharing/signup/apiError"
	cognito "github.com/ujjwal405/FileSharing/signup/cognito"
	"github.com/ujjwal405/FileSharing/signup/user"
)

type UserHandler struct {
	cognitoclient *cognito.CognitoClient
}

func NewUserHandler(c *cognito.CognitoClient) *UserHandler {
	return &UserHandler{
		cognitoclient: c,
	}
}

func (h *UserHandler) SignUpUser(ctx context.Context, req user.SignUpRequest) error {
	exist, err := h.cognitoclient.CheckUser(ctx, req.Email)
	if err != nil {
		return err
	}
	if exist {
		return apiError.UserAlreadyExistsError()
	}
	if err = h.cognitoclient.Signup(ctx, req.Email, req.Password, req.FirstName, req.LastName); err != nil {
		return err
	}
	return nil
}
