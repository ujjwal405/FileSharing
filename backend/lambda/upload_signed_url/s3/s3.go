package s3

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type Mys3 struct {
	s3pClient *s3.PresignClient
}

func LoadDefaultConfig(env, region string) (*Mys3, error) {
	// Get ENVIRONMENT variable

	var cfg aws.Config
	var err error

	if env == "dev" {
		// Use LocalStack
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					if service == s3.ServiceID {
						return aws.Endpoint{URL: "http://localhost:4566"}, nil // LocalStack S3 Endpoint
					}
					return aws.Endpoint{}, &aws.EndpointNotFoundError{}
				},
			)),
			config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     "test", // Dummy credentials for LocalStack
					SecretAccessKey: "test",
					SessionToken:    "",
					Source:          "Hardcoded",
				}, nil
			})),
			config.WithRegion(region), // LocalStack allows any region

		)
		log.Println("Using LocalStack S3 Configuration")
	} else {
		// Use AWS S3 (Production)
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(region),
		)
		log.Println("Using AWS S3 Configuration")
	}

	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)
	return &Mys3{s3pClient: presignClient}, nil
}

func (s *Mys3) GetUploadSignedURL(ctx context.Context, bucketName, key string, expiration time.Duration) (string, error) {
	req, err := s.s3pClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expiration))

	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func GenerateUid() string {
	return uuid.NewString()
}
