package cognito

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/ujjwal405/FileSharing/signin/apiError"
)

func (c *CognitoClient) UserLogin(ctx context.Context, usr User) (*Token, error) {
	authInput := &cognito.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME": usr.Email,
			"PASSWORD": usr.Password,
		},
		ClientId: aws.String(c.appClientID),
	}
	result, err := c.cognitoClient.InitiateAuth(ctx, authInput)
	if err != nil {
		var notAuthorizedErr *types.NotAuthorizedException
		if errors.As(err, &notAuthorizedErr) {
			return nil, apiError.InvalidCredentials()
		}
		//
		return nil, err
	}
	var token Token
	token.AccessToken = *result.AuthenticationResult.AccessToken
	token.IdToken = *result.AuthenticationResult.IdToken
	token.RefreshToken = *result.AuthenticationResult.RefreshToken
	return &token, nil
}
