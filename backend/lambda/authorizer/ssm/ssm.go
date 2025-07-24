package myssm

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var ssmClient *ssm.Client

func init() {
	region := MustEnv("SECRET_REGION")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	ssmClient = ssm.NewFromConfig(cfg)
}

func MustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Env var %s must be set", key)
	}
	return value
}

func GetParameters(paramNames []string) (map[string]string, error) {
	params := make(map[string]string)

	input := &ssm.GetParametersInput{
		Names:          paramNames,
		WithDecryption: aws.Bool(false), // Needed if using SecureString
	}

	output, err := ssmClient.GetParameters(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	for _, param := range output.Parameters {
		params[*param.Name] = *param.Value
	}

	// Check for any invalid parameters
	if len(output.InvalidParameters) > 0 {
		log.Printf("Warning: Some parameters not found: %v", output.InvalidParameters)
	}

	return params, nil
}
