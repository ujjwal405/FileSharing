package dynamo_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db *DynamoClient) GetPage(ctx context.Context, email string, page int, limit int32, scanForward bool) ([]map[string]types.AttributeValue, error) {
	if page < 1 {
		page = 1
	}
	exclusiveStartKey := map[string]types.AttributeValue{}
	for i := 1; i < page; i++ {
		input := &dynamodb.QueryInput{
			TableName:              aws.String(db.tableName),
			KeyConditionExpression: aws.String("email = :email"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":email": &types.AttributeValueMemberS{Value: email},
			},
			Limit:             aws.Int32(limit),
			ExclusiveStartKey: exclusiveStartKey,
			ScanIndexForward:  aws.Bool(scanForward),
		}
		result, err := db.dbClient.Query(ctx, input)
		if err != nil {
			return nil, err
		}
		if len(result.LastEvaluatedKey) == 0 {
			// No more items, so return empty for this page
			return []map[string]types.AttributeValue{}, nil
		}
		exclusiveStartKey = result.LastEvaluatedKey
	}
	// Fetch the desired page
	input := &dynamodb.QueryInput{
		TableName:              aws.String(db.tableName),
		KeyConditionExpression: aws.String("email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: email},
		},
		Limit:             aws.Int32(limit),
		ExclusiveStartKey: exclusiveStartKey,
		ScanIndexForward:  aws.Bool(scanForward),
	}
	result, err := db.dbClient.Query(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}
