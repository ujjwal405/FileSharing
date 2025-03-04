package handler

import (
	"context"

	my_dynamo "github.com/ujjwal405/FileSharing/upload_metadata/dynamo_db"
	"github.com/ujjwal405/FileSharing/upload_metadata/metadata"
)

type UploadHandler struct {
	dynamoClient *my_dynamo.DynamoClient
}

func NewLambdaHandler(dbClient *my_dynamo.DynamoClient) *UploadHandler {
	return &UploadHandler{

		dynamoClient: dbClient,
	}
}

func (h *UploadHandler) HandleConformPassword(ctx context.Context, info metadata.FileInfo) error {

	if err := h.dynamoClient.PutMetaData(ctx, info.S3FileName, info.Email, info.FileName); err != nil {
		return err
	}
	return nil
}
