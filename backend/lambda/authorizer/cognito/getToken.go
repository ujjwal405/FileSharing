package my_cognito

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (c *CognitoClient) GetToken(ctx context.Context, refreshToken string) (string, string, error) {
	input := &cognito.InitiateAuthInput{
		AuthFlow: "REFRESH_TOKEN_AUTH",
		ClientId: aws.String(c.appClientID),
		AuthParameters: map[string]string{
			"REFRESH_TOKEN": refreshToken,
		},
	}

	// Call InitiateAuth to get new tokens
	resp, err := c.cognitoClient.InitiateAuth(ctx, input)
	if err != nil {
		return "", "", fmt.Errorf("failed to refresh tokens: %v", err)
	}

	// Extract the new tokens from the response
	if resp.AuthenticationResult == nil {
		return "", "", fmt.Errorf("no authentication result returned")
	}

	return *resp.AuthenticationResult.IdToken,
		*resp.AuthenticationResult.AccessToken,
		nil

}
