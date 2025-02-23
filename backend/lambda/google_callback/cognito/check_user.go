package cognito

import (
	"context"
	"errors"

	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (c *CognitoClient) CheckUser(ctx context.Context, email string) (bool, error) {
	input := &cognito.AdminGetUserInput{
		UserPoolId: &c.userPoolID,
		Username:   &email,
	}
	_, err := c.cognitoClient.AdminGetUser(ctx, input)
	var notFoundErr *types.UserNotFoundException
	if errors.As(err, &notFoundErr) {
		return false, nil // User does not exist
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
