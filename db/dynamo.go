package db

import (
	"sync"
	"veterinary-employee/settings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	data *Data
	once sync.Once
)

type Data struct {
	DB *dynamodb.DynamoDB
}

func initDB() {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(settings.InitializeDynamo().AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			settings.InitializeDynamo().AWSSecretId,
			settings.InitializeDynamo().AWSSecretKey,
			"",
		),
		Endpoint: aws.String(settings.InitializeDynamo().AWSEndpoint),
	})

	data = &Data{
		DB: dynamodb.New(sess),
	}
}

func New() *Data {
	once.Do(initDB)
	return data
}
