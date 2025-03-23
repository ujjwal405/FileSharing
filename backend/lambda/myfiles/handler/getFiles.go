package handler

import (
	"context"
	"math"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	my_dynamo "github.com/ujjwal405/FileSharing/myfiles/dynamo_db"
	"github.com/ujjwal405/FileSharing/myfiles/file"
)

const (
	limit int32 = 7
)

type GetFileHandler struct {
	dynamoClient *my_dynamo.DynamoClient
}

func NewLambdaHandler(dbClient *my_dynamo.DynamoClient) *GetFileHandler {
	return &GetFileHandler{

		dynamoClient: dbClient,
	}
}

func (h *GetFileHandler) HandleGetFiles(ctx context.Context, pageStr string, email string) (file.PaginatedFiles, error) {
	totalItems, err := h.dynamoClient.GetTotalFiles(ctx, email)
	if err != nil {
		return file.PaginatedFiles{}, err
	}
	page := 1
	if pageStr != "" {
		pageInt, err := strconv.Atoi(pageStr)
		if err == nil && pageInt > 0 {
			page = pageInt
		}
	}

	itemsMap, err := h.dynamoClient.GetPage(ctx, email, page, limit, false)
	if err != nil {
		return file.PaginatedFiles{}, err
	}

	var items []file.FileItem
	for _, item := range itemsMap {
		s3Filename := item["s3filename"].(*types.AttributeValueMemberS).Value
		filename := item["filename"].(*types.AttributeValueMemberS).Value
		createdAt := item["created_at"].(*types.AttributeValueMemberS).Value
		items = append(items, file.FileItem{S3Filename: s3Filename, Filename: filename, CreatedAt: createdAt})
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	response := file.PaginatedFiles{
		Items:      items,
		TotalPages: totalPages,
	}
	return response, nil
}
