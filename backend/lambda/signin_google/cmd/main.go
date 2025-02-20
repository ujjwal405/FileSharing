package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ujjwal405/FileSharing/signin_google/google"
	"github.com/ujjwal405/FileSharing/signin_google/helper"
	"golang.org/x/oauth2"
)

var googleOauthConfig *oauth2.Config

func init() {
	OauthConfig, err := google.InitGoogleConfig()
	if err != nil {
		log.Fatalf("unable to load google config, %v", err)
	}
	googleOauthConfig = OauthConfig
}

func handleGoogleSignIn(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	state_token := helper.GenerateStateToken()
	url := googleOauthConfig.AuthCodeURL(state_token, oauth2.AccessTypeOffline)
	responseBody, err := json.Marshal(map[string]string{
		"url": url,
	})
	if err != nil {
		log.Printf("failed to marshal response: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Internal Server Error",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseBody),
	}, nil

}

func main() {
	lambda.Start(handleGoogleSignIn)
}
