package controllers

import (
	"fmt"
	"net/http"
	"time"

	"veterinary-employee/models"
	"veterinary-employee/repositories"
	"veterinary-employee/services"
	"veterinary-employee/settings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvatarController struct {
	Repository repositories.IAvatarRepository
	S3Service  services.IS3Service
}

func (co *AvatarController) GetMe(c echo.Context) error {
	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	avatar, err := co.Repository.Get(c.Request().Context(), bson.M{"employee_id": objectId})

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	bucketName := settings.InitializeAWS().S3.BucketName
	s3Endpoint := settings.InitializeAWS().S3.Endpoint
	avatar.Path = fmt.Sprintf("%s/%s/%s", s3Endpoint, bucketName, avatar.Path)

	return c.JSON(http.StatusOK, avatar)
}

func (co *AvatarController) Create(c echo.Context) error {
	image, err := c.FormFile("image")

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if image.Size > settings.InitializeAWS().S3.MaxImageSize {
		return echo.NewHTTPError(http.StatusBadRequest, "Image is greather than 2MB")
	}

	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	src, err := image.Open()
	defer src.Close()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	url, err := co.S3Service.UploadFile(image.Filename, userId, src)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	avatar := models.Avatar{
		Path:       fmt.Sprintf("%s/%s", userId, image.Filename),
		EmployeeId: objectId,
	}
	_ = avatar.Validate()

	_, err = co.Repository.Create(c.Request().Context(), avatar)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	avatar.Path = url

	return c.JSON(http.StatusCreated, avatar)
}

func (co *AvatarController) UpdateMe(c echo.Context) error {
	image, err := c.FormFile("image")

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if image.Size > settings.InitializeAWS().S3.MaxImageSize {
		return echo.NewHTTPError(http.StatusBadRequest, "Image is greather than 2MB")
	}

	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	src, err := image.Open()
	defer src.Close()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	url, err := co.S3Service.UploadFile(image.Filename, userId, src)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	partialAvatar := map[string]interface{}{
		"path":       fmt.Sprintf("%s/%s", userId, image.Filename),
		"updated_at": time.Now(),
	}

	avatar, err := co.Repository.Update(
		c.Request().Context(),
		bson.M{"employee_id": objectId},
		partialAvatar,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	avatar.Path = url

	return c.JSON(http.StatusOK, avatar)
}

func (co *AvatarController) DeleteMe(c echo.Context) error {
	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	avatar, err := co.Repository.Get(ctx, bson.M{"employee_id": objectId})

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := co.S3Service.DeleteFile(avatar.Path); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := co.Repository.Delete(ctx, bson.M{"employee_id": objectId}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
