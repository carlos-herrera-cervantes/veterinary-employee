package controllers

import (
	"net/http"

	"veterinary-employee/models"
	"veterinary-employee/repositories"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressController struct {
	Repository repositories.IAddressRepository
}

func (co *AddressController) GetMe(c echo.Context) error {
	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	address, err := co.Repository.Get(c.Request().Context(), bson.M{"employee_id": objectId})

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, address)
}

func (co *AddressController) Create(c echo.Context) error {
	var address models.Address

	if err := c.Bind(&address); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	address.EmployeeId = objectId
	_ = address.Validate()

	insertResult, err := co.Repository.Create(c.Request().Context(), address)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, insertResult)
}

func (co *AddressController) UpdateMe(c echo.Context) error {
	var partialAddress models.PartialAddress

	if err := c.Bind(&partialAddress); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userId := c.Request().Header.Get("user-id")
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_ = partialAddress.ValidatePartial()

	address, err := co.Repository.Update(
		c.Request().Context(),
		bson.M{"employee_id": objectId},
		partialAddress,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, address)
}
