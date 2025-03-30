package dynamo_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	secret_manager "github.com/ujjwal405/FileSharing/logout/secret_manager"
)

type DynamoClient struct {
	dbClient  *dynamodb.Client
	tableName string
}

func NewDynamoClient() (*DynamoClient, error) {
	secretIDs := []string{"DYNAMO_REGION", "DYNAMO_USER_TABLE_NAME"}
	secrets, err := secret_manager.GetSecrets(secretIDs)
	if err != nil {
		return nil, err
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(secrets["DYNAMO_REGION"]))
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoClient{
		dbClient:  client,
		tableName: secrets["DYNAMO_USER_TABLE_NAME"],
	}, nil
}
