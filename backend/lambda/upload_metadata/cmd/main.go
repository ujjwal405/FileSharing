package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	my_dynamo "github.com/ujjwal405/FileSharing/upload_metadata/dynamo_db"
	"github.com/ujjwal405/FileSharing/upload_metadata/metadata"

	"github.com/ujjwal405/FileSharing/upload_metadata/handler"
)

var lambdaHandler *handler.UploadHandler

func init() {

	dbClient, err := my_dynamo.NewDynamoClient()
	if err != nil {
		log.Fatalf("unable to load cognito config, %v", err)
	}

	lambdaHandler = handler.NewLambdaHandler(dbClient)
}

func HandleUploadMetadata(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req metadata.FileInfo
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
		log.Printf("failed to handle uploadmetadata: %v", err.Error())
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
		"msg": "successfully added file metadata",
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
	lambda.Start(HandleUploadMetadata)
}
