package dynamo_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db *DynamoClient) GetRefreshToken(ctx context.Context, email string) (string, error) {
	result, err := db.dbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(db.tableName),
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{Value: email},
		},
		ProjectionExpression: aws.String("refresh_token"), // Fetch only refresh_token
	})

	if err != nil {
		return "", err
	}

	// Extract refresh_token
	tokenAttr, ok := result.Item["refresh_token"].(*types.AttributeValueMemberS)
	if !ok {
		return "", fmt.Errorf("refresh_token not found for email %s", email)
	}

	return tokenAttr.Value, nil
}
