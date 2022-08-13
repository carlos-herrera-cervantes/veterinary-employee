package routes

import (
	"veterinary-employee/controllers"
	"veterinary-employee/db"
	"veterinary-employee/repositories"
	"veterinary-employee/services"
	"veterinary-employee/settings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
)

func BootstrapAvatarRoutes(v *echo.Group) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(settings.InitializeAWS().S3.Region),
		Credentials: credentials.NewStaticCredentials(
			settings.InitializeAWS().S3.AccessKey,
			settings.InitializeAWS().S3.SecretKey,
			"",
		),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(settings.InitializeAWS().S3.Endpoint),
	})
	s3Client := s3.New(sess)
	s3Service := services.S3Service{
		S3Client:   s3Client,
		BucketName: settings.InitializeAWS().S3.BucketName,
	}
	controller := &controllers.AvatarController{
		Repository: &repositories.AvatarRepository{
			Data: db.New(),
		},
		S3Service: &s3Service,
	}

	v.GET("/employees/avatar/me", controller.GetMe)
	v.POST("/employees/avatar", controller.Create)
	v.PATCH("/employees/avatar/me", controller.UpdateMe)
	v.DELETE("/employees/avatar/me", controller.DeleteMe)
}
