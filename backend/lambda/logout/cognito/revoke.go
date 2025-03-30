package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (c *CognitoClient) RevokeToken(ctx context.Context, refreshToken string) error {
	input := &cognito.RevokeTokenInput{
		Token:    aws.String(refreshToken),
		ClientId: aws.String(c.appClientID),
	}
	_, err := c.cognitoClient.RevokeToken(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
