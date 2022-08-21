package controllers

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"veterinary-employee/models"
	"veterinary-employee/repositories/mocks"
	mocksServices "veterinary-employee/services/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAvatarGetMe(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAvatarRepository := mocks.NewMockIAvatarRepository(mockController)
	controller := AvatarController{
		Repository: mockAvatarRepository,
	}

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/avatar/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(0)

		response := controller.GetMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/avatar/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(mockAvatar, errors.New("dummy error")).
			Times(1)

		response := controller.GetMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when avatar exist", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/avatar/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(1)

		response := controller.GetMe(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestAvatarGetById(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAvatarRepository := mocks.NewMockIAvatarRepository(mockController)
	controller := AvatarController{
		Repository: mockAvatarRepository,
	}

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(http.MethodGet, "/api/v1/employees/avatar", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("bad_object_id")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(0)

		response := controller.GetById(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/employees/avatar",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(mockAvatar, errors.New("dummy error")).
			Times(1)

		response := controller.GetById(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when avatar exist", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/employees/avatar",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("62d24f2801ad56f85d5fd0f2")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(1)

		response := controller.GetById(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestAvatarUpsert(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAvatarRepository := mocks.NewMockIAvatarRepository(mockController)
	mockS3Service := mocksServices.NewMockIS3Service(mockController)
	controller := AvatarController{
		Repository: mockAvatarRepository,
		S3Service:  mockS3Service,
	}

	t.Run("Should return error when no send file field", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/api/v1/employees/avatar/me", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(0)

		response := controller.Upsert(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()
		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/employees/avatar/me",
			bytes.NewReader(body.Bytes()),
		)
		recorder := httptest.NewRecorder()

		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(0)

		response := controller.Upsert(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when upload an image fails", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/api/v1/employees/avatar/me", body)
		recorder := httptest.NewRecorder()

		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockS3Service.
			EXPECT().
			UploadFile(gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", errors.New("dummy error")).
			Times(1)

		response := controller.Upsert(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/api/v1/employees/avatar/me", body)
		recorder := httptest.NewRecorder()

		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatarRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(0), errors.New("dummy error")).
			Times(1)
		mockS3Service.
			EXPECT().
			UploadFile(gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", nil).
			Times(1)

		response := controller.Upsert(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when avatar is new", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/api/v1/employees/avatar", body)
		recorder := httptest.NewRecorder()

		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(0), nil).
			Times(1)
		mockAvatarRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(1)
		mockS3Service.
			EXPECT().
			UploadFile(gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", nil).
			Times(1)

		response := controller.Upsert(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})

	t.Run("Should return 200 when avatar is updated", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/api/v1/employees/avatar", body)
		recorder := httptest.NewRecorder()

		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			CountDocuments(gomock.Any(), gomock.Any()).
			Return(int64(1), nil).
			Times(1)
		mockAvatarRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(1)
		mockS3Service.
			EXPECT().
			UploadFile(gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", nil).
			Times(1)

		response := controller.Upsert(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
