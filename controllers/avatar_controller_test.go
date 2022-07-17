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

func TestAvatarCreate(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAvatarRepository := mocks.NewMockIAvatarRepository(mockController)
	controller := AvatarController{
		Repository: mockAvatarRepository,
	}

	t.Run("Should return error when no send file field", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/api/v1/employees/avatar", nil)
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

		response := controller.Create(c)

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
			"/api/v1/employees/avatar",
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

		response := controller.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
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
			Create(gomock.Any(), gomock.Any()).
			Return(mockAvatar, errors.New("dummy error")).
			Times(1)

		response := controller.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when uploading an avatar", func(t *testing.T) {
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
			Create(gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(1)

		response := controller.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestAvatarUpdateMe(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAvatarRepository := mocks.NewMockIAvatarRepository(mockController)
	controller := AvatarController{
		Repository: mockAvatarRepository,
	}

	t.Run("Should return error when no send file field", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodPatch, "/api/v1/employees/avatar/me", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(0)

		response := controller.UpdateMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()
		request := httptest.NewRequest(
			http.MethodPatch,
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
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(0)

		response := controller.UpdateMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()
		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/employees/avatar/me",
			bytes.NewReader(body.Bytes()),
		)
		recorder := httptest.NewRecorder()

		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(mockAvatar, errors.New("dummy error")).
			Times(1)

		response := controller.UpdateMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when updating the avatar", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()
		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/employees/avatar/me",
			bytes.NewReader(body.Bytes()),
		)
		recorder := httptest.NewRecorder()

		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatar := models.Avatar{}
		mockAvatarRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(mockAvatar, nil).
			Times(1)

		response := controller.UpdateMe(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestAvatarDeleteMe(t *testing.T) {
	mockController := gomock.NewController(t)
	mockAvatarRepository := mocks.NewMockIAvatarRepository(mockController)
	controller := AvatarController{
		Repository: mockAvatarRepository,
	}

	t.Run("Should return error when object id is not valid", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/avatar/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "")

		mockAvatarRepository.
			EXPECT().
			Delete(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(0)

		response := controller.DeleteMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/avatar/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatarRepository.
			EXPECT().
			Delete(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := controller.DeleteMe(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when deleting an avatar", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/avatar/me", nil)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "62d24f2801ad56f85d5fd0f2")

		mockAvatarRepository.
			EXPECT().
			Delete(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := controller.DeleteMe(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}
