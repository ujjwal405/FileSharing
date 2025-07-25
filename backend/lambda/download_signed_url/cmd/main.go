package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	apierror "github.com/ujjwal405/FileSharing/download_signed_url/apiError"
	"github.com/ujjwal405/FileSharing/download_signed_url/helper"
	mys3 "github.com/ujjwal405/FileSharing/download_signed_url/s3"
)

const (
	expiration = 300 * time.Second
)

var (
	MyS3Client *mys3.Mys3
	bucketName string
)

func MustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Env var %s must be set", key)
	}
	return value
}

func init() {

	bucketName = MustEnv("BUCKET_NAME")
	env := MustEnv("ENVIRONMENT")
	region := MustEnv("REGION")
	client, err := mys3.LoadDefaultConfig(env, region)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	MyS3Client = client

}

func handleDownloadSignedURL(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody mys3.RequestBody

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

	keyID, err := helper.IsCodeExpired(requestBody.KeyID)
	if err != nil {
		if apiErr, ok := err.(apierror.APIError); ok {
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
			log.Printf("failed to handle download signed url: %v", err)
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
	url, err := MyS3Client.GetDownloadSignedURL(ctx, bucketName, keyID, expiration)
	if err != nil {
		log.Printf("failed to get presigned URL: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type":                     "application/json",
				"Access-Control-Allow-Origin":      "https://fileshare.ujjwalsilwal123.com.np",
				"Access-Control-Allow-Credentials": "true",
				"Access-Control-Allow-Headers":     "Content-Type, Authorization, X-Id-Token",
				"Access-Control-Allow-Methods":     "GET,POST,OPTIONS",
			},
			Body: "Failed to get presigned URL",
		}, nil
	}
	responseBody, err := json.Marshal(map[string]string{
		"url": url,
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
	lambda.Start(handleDownloadSignedURL)
}
