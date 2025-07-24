package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ujjwal405/FileSharing/getcode/helper"
	"github.com/ujjwal405/FileSharing/getcode/user"
)

func handleGetCode(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody user.Key

	err := json.Unmarshal([]byte(event.Body), &requestBody)
	if err != nil {
		log.Printf("failed to parse request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
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
	code := helper.GenerateUniqueCode(requestBody.KeyID)

	responseBody, err := json.Marshal(map[string]string{
		"code": code,
	})
	if err != nil {
		log.Printf("failed to marshal response: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
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
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
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
	lambda.Start(handleGetCode)
}
