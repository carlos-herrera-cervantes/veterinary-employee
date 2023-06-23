package controllers

import (
	"net/http"

	"veterinary-employee/models"
	"veterinary-employee/repositories"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoleController struct {
	Repository repositories.IRoleRepository
}

func (co *RoleController) GetAll(c echo.Context) error {
	roles, err := co.Repository.GetAll(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, roles)
}

func (co *RoleController) Create(c echo.Context) error {
	var role models.Role

	if err := c.Bind(&role); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := role.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, err := co.Repository.Create(c.Request().Context(), role); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, role)
}

func (co *RoleController) Update(c echo.Context) error {
	var partialRole models.PartialRole

	if err := c.Bind(&partialRole); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_ = partialRole.ValidatePartial()

	role, err := co.Repository.Update(
		c.Request().Context(),
		bson.M{"_id": objectId},
		partialRole,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, role)
}
