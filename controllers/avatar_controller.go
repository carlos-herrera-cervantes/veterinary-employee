package controllers

import (
	"fmt"
	"net/http"

	"veterinary-employee/models"
	"veterinary-employee/repositories"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvatarController struct {
	Repository repositories.IAvatarRepository
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

	return c.JSON(http.StatusOK, avatar)
}

func (co *AvatarController) Create(c echo.Context) error {
	image, err := c.FormFile("image")

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	avatar := models.Avatar{
		Path:       fmt.Sprintf("%s/%s", userId, image.Filename),
		EmployeeId: objectId,
	}

	insertResult, err := co.Repository.Create(c.Request().Context(), avatar)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, insertResult)
}

func (co *AvatarController) UpdateMe(c echo.Context) error {
	image, err := c.FormFile("image")

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	partialAvatar := map[string]string{
		"Path": fmt.Sprintf("%s/%s", userId, image.Filename),
	}

	avatar, err := co.Repository.Update(
		c.Request().Context(),
		bson.M{"employee_id": objectId},
		partialAvatar,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, avatar)
}

func (co *AvatarController) DeleteMe(c echo.Context) error {
	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := co.Repository.Delete(
		c.Request().Context(),
		bson.M{"employee_id": objectId},
	); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
