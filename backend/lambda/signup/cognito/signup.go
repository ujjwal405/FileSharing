package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (c *CognitoClient) Signup(ctx context.Context, email, password, firstName, lastName string) error {
	input := &cognito.SignUpInput{
		ClientId: aws.String(c.appClientID),
		Username: aws.String(email),
		Password: aws.String(password),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(email)},
			{Name: aws.String("custom:first_name"), Value: aws.String(firstName)},
			{Name: aws.String("custom:last_name"), Value: aws.String(lastName)},
			{Name: aws.String("custom:google_login"), Value: aws.String("false")},
		},
	}
	_, err := c.cognitoClient.SignUp(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
