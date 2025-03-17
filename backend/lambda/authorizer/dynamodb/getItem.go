package dynamo_db

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db *DynamoClient) GetUserExpiresAt(ctx context.Context, email string) (int64, string, error) {

	key := map[string]types.AttributeValue{
		"email": &types.AttributeValueMemberS{Value: email},
	}

	resp, err := db.dbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(db.tableName),
		Key:       key,
	})
	if err != nil {
		return 0, "", fmt.Errorf("failed to get item from DynamoDB: %v", err)
	}

	// Check if the item exists
	if len(resp.Item) == 0 {
		return 0, "", fmt.Errorf("no item found for email: %s", email)
	}

	expiresAtAttr, ok := resp.Item["expires_at"]
	if !ok {
		return 0, "", fmt.Errorf("expires_at attribute not found for email: %s", email)
	}

	expiresAtStr, ok := expiresAtAttr.(*types.AttributeValueMemberN)
	if !ok {
		return 0, "", fmt.Errorf("expires_at is not a number for email: %s", email)
	}

	expiresAt, err := strconv.ParseInt(expiresAtStr.Value, 10, 64)
	if err != nil {
		return 0, "", fmt.Errorf("failed to parse expires_at for email %s: %v", email, err)
	}

	refreshTokenAttr, ok := resp.Item["refresh_token"]
	if !ok {
		return 0, "", fmt.Errorf("refresh_token attribute not found for email: %s", email)
	}

	refreshTokenStr, ok := refreshTokenAttr.(*types.AttributeValueMemberS)
	if !ok {
		return 0, "", fmt.Errorf("refresh_token is not a string for email: %s", email)
	}

	return expiresAt, refreshTokenStr.Value, nil
}
