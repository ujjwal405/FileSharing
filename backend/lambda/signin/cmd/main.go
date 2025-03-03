package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ujjwal405/FileSharing/signin/apiError"
	my_cognito "github.com/ujjwal405/FileSharing/signin/cognito"
	my_db "github.com/ujjwal405/FileSharing/signin/dynamo_db"

	"github.com/ujjwal405/FileSharing/signin/handler"
)

var lambdaHandler *handler.SignInHandler

func init() {

	cClient, err := my_cognito.NewCognitoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}
	dClient, err := my_db.NewDynamoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}

	lambdaHandler = handler.NewSignInHandler(cClient, dClient)
}

func handleUserSignIn(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req my_cognito.User
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		log.Printf("failed to parse request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Invalid request body",
		}, nil
	}

	accessToken, idToken, err := lambdaHandler.SignInUser(ctx, req)
	if err != nil {
		log.Printf("failed to handle Signin User: %v", err.Error())
		if apiErr, ok := err.(apiError.APIError); ok {
			return events.APIGatewayProxyResponse{
				StatusCode: apiErr.StatusCode,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: apiErr.Error(),
			}, nil
		} else {
			log.Printf("failed to handle user signin: %v", err)
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
		"access_token": accessToken,
		"id_token":     idToken,
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
	lambda.Start(handleUserSignIn)
}
