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

func TestCatalogPositionsController_GetAll(t *testing.T) {
	gomockController := gomock.NewController(t)
	mockCatalogPositionsRepository := mocks.NewMockICatalogPositionsRepository(gomockController)
	catalogPositionsController := CatalogPositionsController{
		Repository: mockCatalogPositionsRepository,
	}

	t.Run("Should return bad request when pager is not valid", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/catalog-positions", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockCatalogPositionsRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.CatalogPosition{}, nil).
			Times(0)
		mockCatalogPositionsRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(0), nil).
			Times(0)

		response := catalogPositionsController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when listing positions fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/catalog-positions?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockCatalogPositionsRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.CatalogPosition{}, errors.New("dummy error")).
			Times(1)
		mockCatalogPositionsRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(0), nil).
			Times(0)

		response := catalogPositionsController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when counting positions fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/catalog-positions?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockCatalogPositionsRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.CatalogPosition{}, nil).
			Times(1)
		mockCatalogPositionsRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(0), errors.New("dummy error")).
			Times(1)

		response := catalogPositionsController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when listing positions", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/catalog-positions?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockCatalogPositionsRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.CatalogPosition{}, nil).
			Times(1)
		mockCatalogPositionsRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(0), nil).
			Times(1)

		response := catalogPositionsController.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestCatalogPositionsController_GetById(t *testing.T) {
	gomockController := gomock.NewController(t)
	mockCatalogPositionsRepository := mocks.NewMockICatalogPositionsRepository(gomockController)
	catalogPositionsController := CatalogPositionsController{
		Repository: mockCatalogPositionsRepository,
	}

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/catalog-positions",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("bad_object_id")

		mockCatalogPositionsRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, nil).
			Times(0)

		response := catalogPositionsController.GetById(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/catalog-positions",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockCatalogPositionsRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, errors.New("dummy error")).
			Times(1)

		response := catalogPositionsController.GetById(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when profile exist", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/catalog-positions",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockCatalogPositionsRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, nil).
			Times(1)

		response := catalogPositionsController.GetById(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestCatalogPositionsController_Create(t *testing.T) {
	gomockController := gomock.NewController(t)
	mockCatalogPositionsRepository := mocks.NewMockICatalogPositionsRepository(gomockController)
	catalogPositionsController := CatalogPositionsController{
		Repository: mockCatalogPositionsRepository,
	}

	t.Run("Should return error when binding position fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/catalog-positions",
			strings.NewReader("bad body"),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockCatalogPositionsRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, nil).
			Times(0)

		response := catalogPositionsController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when request body is not valid", func(t *testing.T) {
		e := echo.New()
		body := `{"name": "dummy position"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/catalog-positions",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockCatalogPositionsRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, errors.New("dummy error")).
			Times(0)

		response := catalogPositionsController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		body := `{"name": "dummy position", "description": "dummy description"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/catalog-positions",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockCatalogPositionsRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, errors.New("dummy error")).
			Times(1)
		mockCatalogPositionsRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, errors.New("dummy error")).
			Times(1)

		response := catalogPositionsController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when process completes successfully", func(t *testing.T) {
		e := echo.New()
		body := `{"name": "dummy position", "description": "dummy description"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/catalog-positions",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockCatalogPositionsRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, errors.New("dummy error")).
			Times(1)
		mockCatalogPositionsRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, nil).
			Times(1)

		response := catalogPositionsController.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestCatalogPositionsController_UpdateById(t *testing.T) {
	gomockController := gomock.NewController(t)
	mockCatalogPositionsRepository := mocks.NewMockICatalogPositionsRepository(gomockController)
	catalogPositionsController := CatalogPositionsController{
		Repository: mockCatalogPositionsRepository,
	}

	t.Run("Should return error when object ID is not valid", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/catalog-positions",
			strings.NewReader("bad body"),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("bad_object_id")

		mockCatalogPositionsRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, nil).
			Times(0)

		response := catalogPositionsController.UpdateById(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when binding position", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/catalog-positions",
			strings.NewReader("bad body"),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockCatalogPositionsRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, nil).
			Times(0)

		response := catalogPositionsController.UpdateById(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/catalog-positions",
			strings.NewReader(`{"description": "dummy description"}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockCatalogPositionsRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, errors.New("dummy error")).
			Times(1)

		response := catalogPositionsController.UpdateById(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process completes successfully", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/catalog-positions",
			strings.NewReader(`{"description": "dummy description"}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockCatalogPositionsRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.CatalogPosition{}, nil).
			Times(1)

		response := catalogPositionsController.UpdateById(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
