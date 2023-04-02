package settings

import "os"

type aws struct {
	S3 s3
}

type s3 struct {
	BucketName   string
	Endpoint     string
	Region       string
	AccessKey    string
	SecretKey    string
	MaxImageSize int64
}

var singletonAws *aws

func InitializeAWS() *aws {
	if singletonAws != nil {
		return singletonAws
	}

	lock.Lock()
	defer lock.Unlock()

	singletonAws = &aws{
		S3: s3{
			BucketName:   os.Getenv("S3_BUCKET"),
			Endpoint:     os.Getenv("S3_ENDPOINT"),
			Region:       os.Getenv("S3_REGION"),
			AccessKey:    os.Getenv("S3_ACCESS_KEY"),
			SecretKey:    os.Getenv("S3_SECRET_KEY"),
			MaxImageSize: 2000000,
		},
	}

	return singletonAws
}
