package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	my_cognito "github.com/ujjwal405/FileSharing/confirm_password/cognito"
	"github.com/ujjwal405/FileSharing/confirm_password/confirmation"

	"github.com/ujjwal405/FileSharing/confirm_password/handler"
)

var lambdaHandler *handler.PasswordHandler

func init() {

	cClient, err := my_cognito.NewCognitoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}

	lambdaHandler = handler.NewLambdaHandler(cClient)
}

func HandleConfirmPassword(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req confirmation.Confirmation
	if err := json.Unmarshal([]byte(event.Body), &req); err != nil {
		log.Printf("failed to parse request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Content-Type":                     "application/json",
				"Access-Control-Allow-Origin":      "https://fileshare.ujjwalsilwal123.com.np",
				"Access-Control-Allow-Credentials": "true",
				"Access-Control-Allow-Headers":     "Content-Type, Authorization, X-Id-Token",
				"Access-Control-Allow-Methods":     "GET,POST,OPTIONS",
			},
			Body: "Invalid request body",
		}, nil
	}

	if err := lambdaHandler.HandleConformPassword(ctx, req); err != nil {
		log.Printf("failed to handle forget password: %v", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type":                     "application/json",
				"Access-Control-Allow-Origin":      "https://fileshare.ujjwalsilwal123.com.np",
				"Access-Control-Allow-Credentials": "true",
				"Access-Control-Allow-Headers":     "Content-Type, Authorization, X-Id-Token",
				"Access-Control-Allow-Methods":     "GET,POST,OPTIONS",
			},
			Body: "Internal Server Error",
		}, nil
	}

	responseBody, _ := json.Marshal(map[string]string{
		"msg": "Successfuly changed password",
	})
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "https://fileshare.ujjwalsilwal123.com.np",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Headers":     "Content-Type, Authorization, X-Id-Token",
			"Access-Control-Allow-Methods":     "GET,POST,OPTIONS",
		},
		Body: string(responseBody),
	}, nil
}

func main() {
	lambda.Start(HandleConfirmPassword)
}
