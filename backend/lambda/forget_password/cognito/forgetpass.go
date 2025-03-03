package cognito

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (c *CognitoClient) ForgetPassword(ctx context.Context, email string) error {
	input := &cognito.ForgotPasswordInput{
		ClientId: aws.String(c.appClientID),
		Username: aws.String(email),
	}
	_, err := c.cognitoClient.ForgotPassword(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to initiate forgot password: %w", err)
	}
	return nil

}
