package services

import (
	"fmt"
	"mime/multipart"

	"veterinary-employee/settings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	S3Client   IAmazonS3
	BucketName string
}

func (s *S3Service) UploadFile(fileName, userId string, body multipart.File) (string, error) {
	_, err := s.S3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(fmt.Sprintf("%s/%s", userId, fileName)),
		Body:   aws.ReadSeekCloser(body),
	})

	if err != nil {
		return "", err
	}

	s3Endpoint := settings.InitializeAWS().S3.Endpoint

	return fmt.Sprintf("%s/%s/%s/%s", s3Endpoint, s.BucketName, userId, fileName), nil
}

func (s *S3Service) DeleteFile(key string) error {
	if _, err := s.S3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(key),
	}); err != nil {
		return err
	}

	return nil
}
