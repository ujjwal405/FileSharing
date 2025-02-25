package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/ujjwal405/FileSharing/code_verification/user"
)

func (c *CognitoClient) ConfirmAccount(ctx context.Context, usr user.UserConfirmation) error {
	confirmationInput := &cognito.ConfirmSignUpInput{
		Username:         aws.String(usr.Email),
		ConfirmationCode: aws.String(usr.Code),
		ClientId:         aws.String(c.appClientID),
	}
	_, err := c.cognitoClient.ConfirmSignUp(ctx, confirmationInput)
	if err != nil {
		return err
	}
	return nil

}
