package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (c *CognitoClient) Authenticate(ctx context.Context, Email, Password string) (*Token, error) {
	authInput := &cognito.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME": Email,
			"PASSWORD": Password,
		},
		ClientId: aws.String(c.appClientID),
	}

	result, err := c.cognitoClient.InitiateAuth(ctx, authInput)
	if err != nil {
		return nil, err
	}
	var token Token
	token.AccessToken = *result.AuthenticationResult.AccessToken
	token.RefreshToken = *result.AuthenticationResult.RefreshToken
	token.IDToken = *result.AuthenticationResult.IdToken
	return &token, nil
}
