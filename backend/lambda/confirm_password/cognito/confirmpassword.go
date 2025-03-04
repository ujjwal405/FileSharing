package cognito

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (c *CognitoClient) ConfirmForgotPassword(ctx context.Context, email, confirmationCode, newPassword string) error {
	input := &cognito.ConfirmForgotPasswordInput{
		ClientId:         aws.String(c.appClientID),
		Username:         aws.String(email),
		ConfirmationCode: aws.String(confirmationCode),
		Password:         aws.String(newPassword),
	}

	_, err := c.cognitoClient.ConfirmForgotPassword(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to confirm forgot password: %w", err)
	}
	return nil
}
