package handler

import (
	"context"
	"errors"

	apierror "github.com/ujjwal405/FileSharing/authorizer/apiError"
	my_cognito "github.com/ujjwal405/FileSharing/authorizer/cognito"
	my_dynamo "github.com/ujjwal405/FileSharing/authorizer/dynamodb"
	"github.com/ujjwal405/FileSharing/authorizer/helper"
	"github.com/ujjwal405/FileSharing/authorizer/verifier"
)

type AuthorizerHandler struct {
	cognito  my_cognito.CognitoClient
	verifier *verifier.Verifier
	dynamo   *my_dynamo.DynamoClient
}

func NewLambdaHandler(cognito my_cognito.CognitoClient, verifier *verifier.Verifier, dynamo *my_dynamo.DynamoClient) *AuthorizerHandler {
	return &AuthorizerHandler{
		cognito:  cognito,
		verifier: verifier,
		dynamo:   dynamo,
	}
}

func (h *AuthorizerHandler) Authorize(ctx context.Context, accessToken, idToken string) (*NewToken, error) {

	idTokenClaims, err := h.verifier.VerifyIDToken(idToken)
	if err != nil && !errors.Is(err, apierror.ErrTokenExpired) {
		return nil, err
	}
	_, err = h.verifier.VerifyAccessToken(accessToken)
	if err != nil && !errors.Is(err, apierror.ErrTokenExpired) {
		return nil, err
	}
	expires_at, refreshToken, err := h.dynamo.GetUserExpiresAt(ctx, idTokenClaims.Username)
	if err != nil {
		return nil, err
	}
	if helper.IsExpired(expires_at) && errors.Is(err, apierror.ErrTokenExpired) {
		if idTokenClaims.GoogleLogin == "true" {
			if err := h.cognito.DeleteUser(ctx, idTokenClaims.Username); err != nil {
				return nil, err
			}

		}
		return nil, apierror.TokenExpired()
	} else if !helper.IsExpired(expires_at) && errors.Is(err, apierror.ErrTokenExpired) {
		newIDToken, newAccesToken, err := h.cognito.GetToken(ctx, refreshToken)
		if err != nil {
			return nil, err
		}
		var newToken NewToken
		newToken.AccessToken = newAccesToken
		newToken.IdToken = newIDToken
		return &newToken, nil
	}
	return nil, nil

}
