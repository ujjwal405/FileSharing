package dynamo_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ssm "github.com/ujjwal405/FileSharing/upload_metadata/ssm"
)

type DynamoClient struct {
	dbClient  *dynamodb.Client
	tableName string
}

func NewDynamoClient() (*DynamoClient, error) {
	secretIDs := []string{"/myapp/dynamo/dynamoRegion", "/myapp/dynamo/dynamoFileName"}
	secrets, err := ssm.GetParameters(secretIDs)
	if err != nil {
		return nil, err
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(secrets["/myapp/dynamo/dynamoRegion"]))
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoClient{
		dbClient:  client,
		tableName: secrets["/myapp/dynamo/dynamoUserTableName"],
	}, nil
}
