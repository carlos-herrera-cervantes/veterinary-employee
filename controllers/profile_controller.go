package controllers

import (
	"net/http"

	"veterinary-employee/models"
	"veterinary-employee/repositories"
	"veterinary-employee/types"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileController struct {
	Repository repositories.IProfileRepository
}

func (co *ProfileController) GetAll(c echo.Context) error {
	var pager types.Pager

	if err := c.Bind(&pager); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	profiles, err := co.Repository.GetAll(ctx, pager.Offset, pager.Limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalDocs, err := co.Repository.CountDocuments(ctx, bson.D{})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetResult(&pager, totalDocs, profiles))
}

func (co *ProfileController) GetById(c echo.Context) error {
	userId := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	profile, err := co.Repository.Get(c.Request().Context(), bson.M{"employee_id": objectId})

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func (co *ProfileController) GetMe(c echo.Context) error {
	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	profile, err := co.Repository.Get(c.Request().Context(), bson.M{"employee_id": objectId})

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func (co *ProfileController) UpdateMe(c echo.Context) error {
	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var partialProfile models.PartialProfile

	if err := c.Bind(&partialProfile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	_ = partialProfile.ValidatePartial()

	profile, err := co.Repository.Update(
		c.Request().Context(),
		bson.M{"employee_id": objectId},
		partialProfile,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}
