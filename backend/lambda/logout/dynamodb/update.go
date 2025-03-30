package dynamo_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db *DynamoClient) Update(ctx context.Context, email string, expiredTime int64) error {

	_, err := db.dbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(db.tableName),
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{Value: email},
		},
		UpdateExpression: aws.String("SET expires_at = :expired"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":expired": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", expiredTime)},
		},
		ConditionExpression: aws.String("attribute_exists(email)"), // Ensure user exists
	})

	if err != nil {
		return err
	}

	return nil
}
