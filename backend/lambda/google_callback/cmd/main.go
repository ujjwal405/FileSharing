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
)

var lambdaHandler *handler.LambdaHandler

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
	lambdaHandler = handler.NewLambdaHandler(googleConfig, cClient, dClient)
}

func handleGoogleCallback(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	IdToken, AccessToken, err := lambdaHandler.HandleGoogleCallback(ctx, event.QueryStringParameters["code"])
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
