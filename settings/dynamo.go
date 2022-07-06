package settings

import "os"

type dynamo struct {
	AWSRegion    string
	AWSSecretId  string
	AWSSecretKey string
	AWSEndpoint  string
}

func InitializeDynamo() dynamo {
	return dynamo{
		AWSRegion:    os.Getenv("DYNAMO_DEFAULT_DB_REGION"),
		AWSSecretId:  os.Getenv("DYNAMO_SECRET_ID"),
		AWSSecretKey: os.Getenv("DYNAMO_SECRET_KEY"),
		AWSEndpoint:  os.Getenv("DYNAMO_ENDPOINT"),
	}
}
