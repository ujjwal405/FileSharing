package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	my_cognito "github.com/ujjwal405/FileSharing/logout/cognito"
	my_db "github.com/ujjwal405/FileSharing/logout/dynamodb"
	"github.com/ujjwal405/FileSharing/logout/handler"
)

var lambdaHandler *handler.LogoutHandler

func init() {

	cClient, err := my_cognito.NewCognitoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}
	dClient, err := my_db.NewDynamoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}

	lambdaHandler = handler.NewLambdaHandler(cClient, dClient)
}

func handleLogout(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	authorizer := event.RequestContext.Authorizer
	username, _ := authorizer["username"].(string)
	googleLogin, _ := authorizer["google_login"].(string)
	if err := lambdaHandler.LogoutHandler(ctx, username, googleLogin); err != nil {
		log.Printf("failed to handle Logout User: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Internal Server Error",
		}, nil
	}

	responseBody, _ := json.Marshal(map[string]string{
		"msg": "successfully logout",
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
	lambda.Start(handleLogout)
}
