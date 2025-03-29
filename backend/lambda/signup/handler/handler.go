package handler

import (
	"context"

	apiError "github.com/ujjwal405/FileSharing/signup/apiError"
	cognito "github.com/ujjwal405/FileSharing/signup/cognito"
	"github.com/ujjwal405/FileSharing/signup/helper"
	"github.com/ujjwal405/FileSharing/signup/user"
)

type UserHandler struct {
	cognitoclient *cognito.CognitoClient
	validator *helper.Validator
}

func NewUserHandler(c *cognito.CognitoClient,v *helper.Validator) *UserHandler {
	return &UserHandler{
		cognitoclient: c,
		validator:v,
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
	if err:=h.validator.ValidatePassword(req.Password);err!=nil{
		return apiError.MinimumLength()
	}

	if err = h.cognitoclient.Signup(ctx, req.Email, req.Password, req.FirstName, req.LastName); err != nil {
		return err
	}
	return nil
}
