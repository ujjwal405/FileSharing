package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	my_cognito "github.com/ujjwal405/FileSharing/forget_password/cognito"
	"github.com/ujjwal405/FileSharing/forget_password/user"

	"github.com/ujjwal405/FileSharing/forget_password/handler"
)

var lambdaHandler *handler.PasswordHandler

func init() {

	cClient, err := my_cognito.NewCognitoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}

	lambdaHandler = handler.NewPasswordHandler(cClient)
}

func HandleForgetPassword(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req user.RecoveryInfo
	if err := json.Unmarshal([]byte(event.Body), &req); err != nil {
		log.Printf("failed to parse request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Invalid request body",
		}, nil
	}

	if err := lambdaHandler.HandleForgetPassword(ctx, req.Email); err != nil {
		log.Printf("failed to handle forget password: %v", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Internal Server Error",
		}, nil
	}

	responseBody, _ := json.Marshal(map[string]string{
		"msg": "Enter a code sent to email",
	})
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseBody),
	}, nil
}

func main() {
	lambda.Start(HandleForgetPassword)
}
