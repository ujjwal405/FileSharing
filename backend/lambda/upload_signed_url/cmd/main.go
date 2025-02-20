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
	mys3 "github.com/ujjwal405/FileSharing/upload_signed_url/s3"
)

const (
	expiration = 120 * time.Second
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

func handleUploadSignedURL(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url, err := MyS3Client.GetUploadSignedURL(ctx, bucketName, expiration)
	if err != nil {
		log.Printf("failed to get presigned URL: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Internal Server Error",
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
	lambda.Start(handleUploadSignedURL)
}
