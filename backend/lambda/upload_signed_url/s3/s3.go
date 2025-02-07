package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type Mys3 struct {
	s3pClient *s3.PresignClient
}

func LoadDefaultConfig() (*Mys3, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)
	return &Mys3{s3pClient: presignClient}, nil

}

func (s *Mys3) GetUploadSignedURL(ctx context.Context, bucketName string, expiration time.Duration) (string, error) {
	key := generateUid()
	req, err := s.s3pClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expiration))

	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func generateUid() string {
	return uuid.NewString()
}
