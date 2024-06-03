package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/auth"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"

	"github.com/stretchr/testify/assert"
)

func TestUserController_SignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		defer func() { mockUserUsecase.Mock.ExpectedCalls = nil }()
		mockInput := usecase.SignUpInput{
			Email:    "test@example.com",
			Password: "Password123!",
			Name:     "Test User",
			Nickname: "testuser",
		}
		mockOutput := &usecase.SignUpOutput{
			UserID: 1,
		}
		mockUserUsecase.On("SignUp", mockInput).Return(mockOutput, nil)

		reqBody, _ := json.Marshal(SignUpRequest{
			Email:    "test@example.com",
			Password: "Password123!",
			Name:     "Test User",
			Nickname: "testuser",
		})

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		reqBody, _ := json.Marshal(SignUpRequest{
			Email: "invalid_email",
		})

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockUserUsecase.AssertNotCalled(t, "SignUp")
	})

	t.Run("EmailAlreadyExists", func(t *testing.T) {
		defer func() { mockUserUsecase.Mock.ExpectedCalls = nil }()
		mockInput := usecase.SignUpInput{
			Email:    "test@example.com",
			Password: "Password123!",
			Name:     "Test User",
			Nickname: "testuser",
		}
		mockUserUsecase.On("SignUp", mockInput).Return(nil, usecase.ErrEmailAlreadyExists)

		reqBody, _ := json.Marshal(SignUpRequest{
			Email:    "test@example.com",
			Password: "Password123!",
			Name:     "Test User",
			Nickname: "testuser",
		})

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})
}

func TestUserController_SendPasswordRecoveryEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		defer func() { mockUserUsecase.Mock.ExpectedCalls = nil }()

		ts := httptest.NewServer(testRouter)
		defer ts.Close()

		mockInput := "test@example.com"
		mockUserUsecase.On("SendPasswordRecoveryEmail", ts.URL, mockInput).Return(nil)

		reqBody, _ := json.Marshal(SendPasswordRecoveryEmailRequest{
			Email: "test@example.com",
		})

		req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/users/password/recovery", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		reqBody, _ := json.Marshal(SendPasswordRecoveryEmailRequest{
			Email: "invalid_email",
		})

		ts := httptest.NewServer(testRouter)
		defer ts.Close()

		req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/users/password/recovery", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockUserUsecase.AssertNotCalled(t, "SendPasswordRecoveryEmail")
	})

	t.Run("UserNotFound", func(t *testing.T) {
		defer func() { mockUserUsecase.Mock.ExpectedCalls = nil }()

		ts := httptest.NewServer(testRouter)
		defer ts.Close()

		mockInput := "test@example.com"
		mockUserUsecase.On("SendPasswordRecoveryEmail", ts.URL, mockInput).Return(usecase.ErrUserNotFound)

		reqBody, _ := json.Marshal(SendPasswordRecoveryEmailRequest{
			Email: "test@example.com",
		})

		req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/users/password/recovery", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})
}

func TestUserController_ResetPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		defer func() { mockUserUsecase.Mock.ExpectedCalls = nil }()
		password := "newPassword123!"
		flowID := "flow123"
		mockUserUsecase.On("ResetPassword", password, flowID).Return(nil)

		reqBody, _ := json.Marshal(ResetPasswordRequest{
			Password: password,
			FlowID:   flowID,
		})

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/password/reset", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		reqBody, _ := json.Marshal(ResetPasswordRequest{
			Password: "short",
		})

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/password/reset", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockUserUsecase.AssertNotCalled(t, "ResetPassword")
	})

	t.Run("FlowNotFound", func(t *testing.T) {
		defer func() { mockUserUsecase.Mock.ExpectedCalls = nil }()
		password := "newPassword123!"
		flowID := "flow123"
		mockUserUsecase.On("ResetPassword", password, flowID).Return(usecase.ErrPasswordResetFlowNotFound)

		reqBody, _ := json.Marshal(ResetPasswordRequest{
			Password: password,
			FlowID:   flowID,
		})

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/password/reset", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("FlowExpired", func(t *testing.T) {
		defer func() { mockUserUsecase.Mock.ExpectedCalls = nil }()
		password := "newPassword123!"
		flowID := "flow123"
		mockUserUsecase.On("ResetPassword", password, flowID).Return(usecase.ErrPasswordResetFlowExpired)

		reqBody, _ := json.Marshal(ResetPasswordRequest{
			Password: password,
			FlowID:   flowID,
		})

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/password/reset", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})
}

func TestUserController_GetMyUserInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		defer func() { mockUserUsecase.Mock.ExpectedCalls = nil }()
		userID := uint(1)
		user := &usecase.GetUserByIDOutput{
			ID:              userID,
			Email:           "test@example.com",
			Name:            "John Doe",
			Nickname:        "johndoe",
			ProfileImageURL: "https://example.com/profile.jpg",
			Bio:             "Test bio",
			Website:         "https://example.com",
		}
		mockUserUsecase.On("GetUserByID", userID).Return(user, nil)

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/me", nil)
		token, _, _ := testUserJwtAuth.TokenGenerator(&auth.UserPayload{UserID: userID})
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUserUsecase.AssertExpectations(t)

		var res GetMyUserInfoResponse
		err := json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, user.ID, res.UserID)
		assert.Equal(t, user.Email, res.Email)
		assert.Equal(t, user.Name, res.Name)
		assert.Equal(t, user.Nickname, res.Nickname)
		assert.Equal(t, user.ProfileImageURL, res.ProfileImageURL)
		assert.Equal(t, user.Bio, res.Bio)
		assert.Equal(t, user.Website, res.Website)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/me", nil)
		req.Header.Set("Authorization", "Bearer invalid_token")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		mockUserUsecase.AssertNotCalled(t, "GetUserByID")
	})

	t.Run("UserNotFound", func(t *testing.T) {
		defer func() { mockUserUsecase.Mock.ExpectedCalls = nil }()
		userID := uint(1)
		mockUserUsecase.On("GetUserByID", userID).Return(nil, usecase.ErrUserNotFound)

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/me", nil)
		token, _, _ := testUserJwtAuth.TokenGenerator(&auth.UserPayload{UserID: userID})
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockUserUsecase.AssertExpectations(t)
	})
}
