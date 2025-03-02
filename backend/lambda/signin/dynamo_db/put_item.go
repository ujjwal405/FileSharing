package dynamo_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db *DynamoClient) PutUser(ctx context.Context, email, token string, expires_at int64) error {
	item := map[string]types.AttributeValue{
		"email":         &types.AttributeValueMemberS{Value: email}, // Primary Key
		"refresh_token": &types.AttributeValueMemberS{Value: token},
		"expires_at":    &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", expires_at)}, // Convert int64 to string
	}
	_, err := db.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.tableName), // Replace with your table name
		Item:      item,
	})

	if err != nil {
		return err
	}
	return nil
}
