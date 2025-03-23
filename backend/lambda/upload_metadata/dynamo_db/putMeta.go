package dynamo_db

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db *DynamoClient) PutMetaData(ctx context.Context, s3Filename, email, filename, createdAt string) error {

	item := map[string]types.AttributeValue{
		"s3filename":      &types.AttributeValueMemberS{Value: s3Filename},
		"email":           &types.AttributeValueMemberS{Value: email},
		"uploadTimestamp": &types.AttributeValueMemberN{Value: strconv.FormatInt(time.Now().Unix(), 10)},
		"filename":        &types.AttributeValueMemberS{Value: filename},
		"created_at":      &types.AttributeValueMemberS{Value: createdAt},
	}

	// Prepare PutItem input
	input := &dynamodb.PutItemInput{
		TableName: aws.String(db.tableName), // Replace with your table name
		Item:      item,
	}
	_, err := db.dbClient.PutItem(ctx, input)

	if err != nil {
		return err
	}
	return nil
}
