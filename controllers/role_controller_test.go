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

func TestRoleGetAll(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRoleRepository := mocks.NewMockIRoleRepository(mockController)
	controller := RoleController{
		Repository: mockRoleRepository,
	}

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/roles", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockRoleRepository.
			EXPECT().
			GetAll(gomock.Any()).
			Return([]models.Role{}, errors.New("dummy error")).
			Times(1)

		response := controller.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when listing roles", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/roles", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockRoleRepository.
			EXPECT().
			GetAll(gomock.Any()).
			Return([]models.Role{}, nil).
			Times(1)

		response := controller.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestRoleCreate(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRoleRepository := mocks.NewMockIRoleRepository(mockController)
	controller := RoleController{
		Repository: mockRoleRepository,
	}

	t.Run("Should return error when binding role", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/roles",
			strings.NewReader("bad body"),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockRoleRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(models.Role{}, nil).
			Times(0)

		response := controller.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		body := `{"name": "dummy role"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/roles",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockRoleRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(models.Role{}, errors.New("dummy error")).
			Times(1)

		response := controller.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when creating a role", func(t *testing.T) {
		e := echo.New()
		body := `{"name": "dummy role"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/roles",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockRoleRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(models.Role{}, nil).
			Times(1)

		response := controller.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestRoleUpdate(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRoleRepository := mocks.NewMockIRoleRepository(mockController)
	controller := RoleController{
		Repository: mockRoleRepository,
	}

	t.Run("Should return error when binding role", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/roles",
			strings.NewReader("bad body"),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockRoleRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Role{}, nil).
			Times(0)

		response := controller.Update(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()
		body := `{"name": "dummy role"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/roles",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("")

		mockRoleRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Role{}, nil).
			Times(0)

		response := controller.Update(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		body := `{"name": "dummy role"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/roles",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockRoleRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Role{}, errors.New("dummy error")).
			Times(1)

		response := controller.Update(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when updating the role", func(t *testing.T) {
		e := echo.New()
		body := `{"name": "dummy role"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/roles",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockRoleRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Role{}, nil).
			Times(1)

		response := controller.Update(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
