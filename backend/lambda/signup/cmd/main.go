package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ujjwal405/FileSharing/signup/apiError"
	my_cognito "github.com/ujjwal405/FileSharing/signup/cognito"
	"github.com/ujjwal405/FileSharing/signup/helper"

	"github.com/ujjwal405/FileSharing/signup/handler"
	user "github.com/ujjwal405/FileSharing/signup/user"
)

var lambdaHandler *handler.UserHandler

func init() {

	cClient, err := my_cognito.NewCognitoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}
	validator := helper.NewValidator()

	lambdaHandler = handler.NewUserHandler(cClient, validator)
}

func handleUserSignUp(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req user.SignUpRequest
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
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

	if err = lambdaHandler.SignUpUser(ctx, req); err != nil {
		log.Printf("failed to handle Signup User: %v", err.Error())
		if apiErr, ok := err.(apiError.APIError); ok {
			return events.APIGatewayProxyResponse{
				StatusCode: apiErr.StatusCode,
				Headers: map[string]string{
					"Content-Type":                     "application/json",
					"Access-Control-Allow-Origin":      "https://fileshare.ujjwalsilwal123.com.np",
					"Access-Control-Allow-Credentials": "true",
					"Access-Control-Allow-Headers":     "Content-Type, Authorization, X-Id-Token",
					"Access-Control-Allow-Methods":     "GET,POST,OPTIONS",
				},
				Body: apiErr.Error(),
			}, nil
		} else {

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
	}
	responseBody, _ := json.Marshal(map[string]string{
		"msg": "successfully signedup",
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
	lambda.Start(handleUserSignUp)
}
