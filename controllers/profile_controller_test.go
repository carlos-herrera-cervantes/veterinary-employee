package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"veterinary-employee/models"
	"veterinary-employee/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestProfileGetAll(t *testing.T) {
	mockController := gomock.NewController(t)
	mockProfileRepository := mocks.NewMockIProfileRepository(mockController)
	controller := ProfileController{
		Repository: mockProfileRepository,
	}

	t.Run("Should return error when validating the pager", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/profiles", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Profile{}, nil).
			Times(0)
		mockProfileRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(0), nil).
			Times(0)

		response := controller.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when listing profiles", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/employees/profiles?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Profile{}, errors.New("dummy error")).
			Times(1)
		mockProfileRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(0), nil).
			Times(0)

		response := controller.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when counting profiles", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/employees/profiles?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Profile{}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(0), errors.New("dummy error")).
			Times(1)

		response := controller.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when listing profiles", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/employees/profiles?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Profile{}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(0), nil).
			Times(1)

		response := controller.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestProfileGetById(t *testing.T) {
	mockController := gomock.NewController(t)
	mockProfileRepository := mocks.NewMockIProfileRepository(mockController)
	controller := ProfileController{
		Repository: mockProfileRepository,
	}

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/employees/profiles",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("bad_object_id")

		mockProfileRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(0)

		response := controller.GetById(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/employees/profiles",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockProfileRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, errors.New("dummy error")).
			Times(1)

		response := controller.GetById(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when profile exist", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/employees/profiles",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockProfileRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(1)

		response := controller.GetById(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestProfileGetMe(t *testing.T) {
	mockController := gomock.NewController(t)
	mockProfileRepository := mocks.NewMockIProfileRepository(mockController)
	controller := ProfileController{
		Repository: mockProfileRepository,
	}

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/profiles/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "")

		mockProfileRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(0)

		response := controller.GetMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/profiles/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockProfileRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, errors.New("dummy error")).
			Times(1)

		response := controller.GetMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when profile exist", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/profiles/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockProfileRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(1)

		response := controller.GetMe(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestProfileUpdateMe(t *testing.T) {
	mockController := gomock.NewController(t)
	mockProfileRepository := mocks.NewMockIProfileRepository(mockController)
	controller := ProfileController{
		Repository: mockProfileRepository,
	}

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()
		body := `{"last_name": "dummy name"}`

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/employees/profiles/me",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "")

		mockProfileRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(0)

		response := controller.UpdateMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when binding profile", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/employees/profiles/me",
			strings.NewReader("bad body"),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockProfileRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(0)

		response := controller.UpdateMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		body := `{"last_name": "dummy name"}`

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/employees/profiles/me",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockProfileRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Profile{}, errors.New("dummy error")).
			Times(1)

		response := controller.UpdateMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when updating the profile", func(t *testing.T) {
		e := echo.New()
		body := `{"last_name": "dummy name"}`

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/employees/profiles/me",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockProfileRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(1)

		response := controller.UpdateMe(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
