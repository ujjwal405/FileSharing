package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (c *CognitoClient) SetEmailVerified(ctx context.Context, email string) error {
	input := &cognito.AdminUpdateUserAttributesInput{
		UserPoolId: aws.String(c.userPoolID),
		Username:   aws.String(email),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
	}
	_, err := c.cognitoClient.AdminUpdateUserAttributes(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
