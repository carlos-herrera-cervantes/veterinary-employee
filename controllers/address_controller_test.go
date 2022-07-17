package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"veterinary-employee/models"
	"veterinary-employee/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddressGetMe(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAddressRepository := mocks.NewMockIAddressRepository(mockController)
	controller := AddressController{
		Repository: mockAddressRepository,
	}

	t.Run("Should return 200 when address exist", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/address/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(mockAddress, nil).
			Times(1)

		response := controller.GetMe(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/address/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(mockAddress, nil).
			Times(0)

		response := controller.GetMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/address/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(mockAddress, errors.New("dummy error")).
			Times(1)

		response := controller.GetMe(c)

		assert.Error(t, response)
	})
}

func TestAddressCreate(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAddressRepository := mocks.NewMockIAddressRepository(mockController)
	controller := AddressController{
		Repository: mockAddressRepository,
	}

	t.Run("Should return error when binding the request body", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/address",
			strings.NewReader("bad body"),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(mockAddress, nil).
			Times(0)

		response := controller.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()
		body := `{"number": "25"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/address",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(mockAddress, nil).
			Times(0)

		response := controller.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		body := `{"number": "25"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/address",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(mockAddress, errors.New("dummy error")).
			Times(1)

		response := controller.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when creating an address", func(t *testing.T) {
		e := echo.New()
		body := `{"number": "25"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/address",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(mockAddress, nil).
			Times(1)

		response := controller.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestAddressUpdateMe(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAddressRepository := mocks.NewMockIAddressRepository(mockController)
	controller := AddressController{
		Repository: mockAddressRepository,
	}

	t.Run("Should return error when binding the request body", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/employees/address/me",
			strings.NewReader("bad body"),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(mockAddress, nil).
			Times(0)

		response := controller.UpdateMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()
		body := `{"number": "25"}`

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/employees/address/me",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(mockAddress, nil).
			Times(0)

		response := controller.UpdateMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		body := `{"number": "25"}`

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/employees/address/me",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(mockAddress, errors.New("dummy error")).
			Times(1)

		response := controller.UpdateMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when updating an address", func(t *testing.T) {
		e := echo.New()
		body := `{"number": "25"}`

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/employees/address/me",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAddress := models.Address{}
		mockAddressRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(mockAddress, nil).
			Times(1)

		response := controller.UpdateMe(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
