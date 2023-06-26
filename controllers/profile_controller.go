package controllers

import (
	"encoding/json"
	"net/http"
	"sync"

	"veterinary-employee/models"
	"veterinary-employee/repositories"
	"veterinary-employee/services"
	"veterinary-employee/settings"
	"veterinary-employee/types"

	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileController struct {
	Repository repositories.IProfileRepository
	KafkaService services.IKafkaService
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

	var wg sync.WaitGroup
	wg.Add(1)

	go co.emitProfileUpdateMessage(&wg, userId, partialProfile.Roles)

	wg.Wait()

	return c.JSON(http.StatusOK, profile)
}

func (co *ProfileController) UpdateById(c echo.Context) error {
	userId := c.Param("id")
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

	var wg sync.WaitGroup
	wg.Add(1)

	go co.emitProfileUpdateMessage(&wg, userId, partialProfile.Roles)

	wg.Wait()

	return c.JSON(http.StatusOK, profile)
}

func (co *ProfileController) emitProfileUpdateMessage(wg *sync.WaitGroup, employeeId string, roles []string) {
	defer wg.Done()

	if roles == nil || len(roles) == 0 {
		golog.Info("skipping profile update message - roles are empty")
		return
	}

	message, err := json.Marshal(types.ProfileUpdateMessage{
		EmployeeId: employeeId,
		Roles: roles,
	})

	if err != nil {
		golog.Error("error when marshaling profile update message: ", err.Error())
		return
	}

	if err := co.KafkaService.SendMessage(settings.InitializeKafka().Topics.ProfileUpdate, message); err != nil {
		golog.Error("error when sending a message for profile update", err.Error())
	}
}
