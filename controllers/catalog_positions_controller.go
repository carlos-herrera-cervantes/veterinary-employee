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

type CatalogPositionsController struct {
	Repository repositories.ICatalogPositionsRepository
}

func (co CatalogPositionsController) GetAll(c echo.Context) error {
	var pager types.Pager

	if err := c.Bind(&pager); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	catalog, err := co.Repository.GetAll(ctx, pager.Offset, pager.Limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalDocs, err := co.Repository.CountDocuments(ctx, bson.M{})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetResult(&pager, totalDocs, catalog))
}

func (co CatalogPositionsController) GetById(c echo.Context) error {
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	position, err := co.Repository.Get(c.Request().Context(), bson.M{"_id": objectId})

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, position)
}

func (co CatalogPositionsController) Create(c echo.Context) error {
	var position models.CatalogPosition

	if err := c.Bind(&position); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := position.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, err := co.Repository.Get(c.Request().Context(), bson.M{"name": bson.M{"$regex": "^" + position.Name, "$options": "i"}}); err == nil {
		return echo.NewHTTPError(http.StatusConflict, "A position with the same name already exists")
	}

	if _, err := co.Repository.Create(c.Request().Context(), position); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, position)
}

func (co CatalogPositionsController) UpdateById(c echo.Context) error {
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var partialPosition models.PartialCatalogPosition

	if err := c.Bind(&partialPosition); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if partialPosition.Name != "" {
		if _, err := co.Repository.Get(c.Request().Context(), bson.M{"name": bson.M{"$regex": "^" + partialPosition.Name, "$options": "i"}}); err == nil {
			return echo.NewHTTPError(http.StatusConflict, "A position with the same name already exists")
		}
	}

	_ = partialPosition.Validate()
	position, err := co.Repository.Update(c.Request().Context(), bson.M{"_id": objectId}, partialPosition)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, position)
}
