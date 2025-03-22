package dynamo_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ujjwal405/FileSharing/myfiles/apiError"
)

func (db *DynamoClient) GetTotalFiles(ctx context.Context, email string) (int64, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(db.tableName),
		KeyConditionExpression: aws.String("email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: email},
		},
		Select: "COUNT",
	}

	result, err := db.dbClient.Query(ctx, input)
	if err != nil {
		return 0, err
	}
	count := int64(result.Count)

	if count == 0 {
		return 0, apiError.FileNotFoundError()
	}

	return count, nil

}
