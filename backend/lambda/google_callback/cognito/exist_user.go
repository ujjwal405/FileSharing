package cognito

import (
	"context"

	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (c *CognitoClient) ExistUser(ctx context.Context, email string) error {
	input := &cognito.AdminCreateUserInput{
		UserPoolId: &c.userPoolID,
		Username:   &email,
	}
	_, err := c.cognitoClient.AdminCreateUser(ctx, input)
	return err
}
