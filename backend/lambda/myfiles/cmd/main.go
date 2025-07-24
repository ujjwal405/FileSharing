package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ujjwal405/FileSharing/myfiles/apiError"

	"github.com/ujjwal405/FileSharing/myfiles/dynamo_db"
	"github.com/ujjwal405/FileSharing/myfiles/file"
	"github.com/ujjwal405/FileSharing/myfiles/handler"
	"github.com/ujjwal405/FileSharing/myfiles/page"
)

var lambdaHandler *handler.GetFileHandler

func init() {
	dbClient, err := dynamo_db.NewDynamoClient()
	if err != nil {
		log.Fatalf("unable to load dynamo config, %v", err)
	}
	lambdaHandler = handler.NewLambdaHandler(dbClient)
}

func handleGetFiles(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	authContext := event.RequestContext.Authorizer
	username, _ := authContext["username"].(string)
	accessToken, accessTokenExists := authContext["access_token"].(string)
	idToken, idTokenExists := authContext["id_token"].(string)

	//pageStr := event.QueryStringParameters["page"]
	var pageStr page.Page
	_ = json.Unmarshal([]byte(event.Body), &pageStr)

	paginatedFiles, err := lambdaHandler.HandleGetFiles(ctx, pageStr.PageNo, username)
	if err != nil {
		log.Printf("failed to handle get files: %v", err.Error())
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

	var userToken *file.Token
	if accessTokenExists || idTokenExists {
		userToken = &file.Token{}
		if accessTokenExists {
			userToken.AccessToken = &accessToken
		}
		if idTokenExists {
			userToken.IDToken = &idToken
		}
	}
	response := file.Response{
		Files:     paginatedFiles,
		UserToken: userToken,
	}

	responseBody, _ := json.Marshal(response)
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
	lambda.Start(handleGetFiles)
}
