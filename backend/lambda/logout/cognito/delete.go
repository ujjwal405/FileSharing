package cognito

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (c *CognitoClient) DeleteUser(ctx context.Context, email string) error {
	input := &cognito.AdminDeleteUserInput{
		UserPoolId: aws.String(c.userPoolId),
		Username:   aws.String(email),
	}

	_, err := c.cognitoClient.AdminDeleteUser(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
