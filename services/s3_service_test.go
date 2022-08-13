package services

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"veterinary-employee/services/mocks"
	"veterinary-employee/settings"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUploadFile(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAmazonS3 := mocks.NewMockIAmazonS3(mockController)
	s3Service := S3Service{
		BucketName: settings.InitializeAWS().S3.BucketName,
		S3Client:   mockAmazonS3,
	}

	t.Run("Should create the full path of the file", func(t *testing.T) {
		putObjectOutput := s3.PutObjectOutput{}
		mockAmazonS3.
			EXPECT().
			PutObject(gomock.Any()).
			Return(&putObjectOutput, nil).
			Times(1)

		fileName := "myavatar.png"
		userId := "dummyid"

		image, _ := os.Open("../assets/avatar.png")
		defer image.Close()

		fullPath, err := s3Service.UploadFile(fileName, userId, image)
		expected := fmt.Sprintf(
			"%s/%s/%s/%s",
			settings.InitializeAWS().S3.Endpoint,
			settings.InitializeAWS().S3.BucketName,
			userId,
			fileName,
		)

		assert.NoError(t, err)
		assert.Equal(t, expected, fullPath)
	})

	t.Run("Should return error and empty string as the path", func(t *testing.T) {
		mockAmazonS3.
			EXPECT().
			PutObject(gomock.Any()).
			Return(nil, errors.New("dummy error")).
			Times(1)

		fileName := "myavatar.png"
		userId := "dummyid"

		image, _ := os.Open("../assets/avatar.png")
		defer image.Close()

		fullPath, err := s3Service.UploadFile(fileName, userId, image)

		assert.Error(t, err)
		assert.Equal(t, "", fullPath)
	})
}
