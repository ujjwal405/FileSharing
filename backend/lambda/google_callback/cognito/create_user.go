package cognito

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func (c *CognitoClient) CreateUser(ctx context.Context, email, firstName, lastName, temp_pass string) error {
	input := &cognito.AdminCreateUserInput{
		UserPoolId: aws.String(c.userPoolID),
		Username:   aws.String(email), // Using email as the username
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(email)},
			{Name: aws.String("first_name"), Value: aws.String(firstName)},  // First Name
			{Name: aws.String("last_name"), Value: aws.String(lastName)},    // Last Name
			{Name: aws.String("email_verified"), Value: aws.String("true")}, // Mark email as verified
			{Name: aws.String("google_login"), Value: aws.String("true")},
		},
		TemporaryPassword:  aws.String(temp_pass),
		ForceAliasCreation: false,                           // No need to force alias creation since we're checking first
		MessageAction:      types.MessageActionTypeSuppress, // Suppresses email invitation (optional)
	}

	_, err := c.cognitoClient.AdminCreateUser(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
