package secret_manager

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var svc *secretsmanager.Client

func init() {

	region := MustEnv("AWS_REGION")

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	svc = secretsmanager.NewFromConfig(config)

}

func MustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Env var %s must be set", key)
	}
	return value
}

func GetSecrets(secretIDs []string) (map[string]string, error) {
	secrets := make(map[string]string)

	for _, secretID := range secretIDs {
		input := &secretsmanager.GetSecretValueInput{
			SecretId:     aws.String(secretID),
			VersionStage: aws.String("AWSCURRENT"), // You can specify a version stage if needed
		}

		result, err := svc.GetSecretValue(context.TODO(), input)
		if err != nil {
			return nil, err // Return error if fetching any secret fails
		}

		secrets[secretID] = *result.SecretString // Store the retrieved secret
	}

	return secrets, nil // Return all retrieved secrets
}
