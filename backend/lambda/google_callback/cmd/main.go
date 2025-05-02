package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	my_cognito "github.com/ujjwal405/FileSharing/google_callback/cognito"
	my_dynamo "github.com/ujjwal405/FileSharing/google_callback/dynamo_db"
	my_google "github.com/ujjwal405/FileSharing/google_callback/google"
	"github.com/ujjwal405/FileSharing/google_callback/handler"
	"github.com/ujjwal405/FileSharing/google_callback/helper"
	secret_manager "github.com/ujjwal405/FileSharing/google_callback/secret_manager"
	"github.com/ujjwal405/FileSharing/google_callback/token"
)

var lambdaHandler *handler.LambdaHandler
var secretKey string

func init() {
	googleConfig, err := my_google.InitGoogleConfig()
	if err != nil {
		log.Fatalf("unable to load google config, %v", err)
	}
	cClient, err := my_cognito.NewCognitoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}
	dClient, err := my_dynamo.NewDynamoClient()
	if err != nil {
		log.Fatalf("unable to load dynamo config, %v", err)
	}
	secret, err := secret_manager.GetSecrets([]string{"SECRET_KEY"})
	if err != nil {
		log.Fatalf("unable to load secrets %v", err)
	}
	secretKey = secret["SECRET_KEY"]
	lambdaHandler = handler.NewLambdaHandler(googleConfig, cClient, dClient, secretKey)
}

func handleGoogleCallback(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var usrToken token.Token
	_ = json.Unmarshal([]byte(event.Body), &usrToken)
	IdToken, AccessToken, email, err := lambdaHandler.HandleGoogleCallback(ctx, usrToken.Code, usrToken.State)
	if err != nil {
		log.Printf("failed to handle google callback: %v", err.Error())
		if apiErr, ok := err.(helper.APIError); ok {
			return events.APIGatewayProxyResponse{
				StatusCode: apiErr.StatusCode,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: apiErr.Error(),
			}, nil
		} else {

			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: "Internal Server Error",
			}, nil
		}
	}
	responseBody, _ := json.Marshal(map[string]string{
		"id_token":     IdToken,
		"access_token": AccessToken,
		"email":        email,
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
	lambda.Start(handleGoogleCallback)
}
