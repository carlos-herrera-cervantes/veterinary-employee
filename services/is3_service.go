package services

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3"
)

//go:generate mockgen -destination=./mocks/iamazon_s3.go -package=mocks --build_flags=--mod=mod . IAmazonS3
type IAmazonS3 interface {
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
	DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
}

//go:generate mockgen -destination=./mocks/is3_service.go -package=mocks --build_flags=--mod=mod . IS3Service
type IS3Service interface {
	UploadFile(fileName, userId string, body multipart.File) (string, error)
	DeleteFile(key string) error
}
