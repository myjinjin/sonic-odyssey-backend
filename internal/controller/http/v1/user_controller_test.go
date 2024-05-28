package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/auth"
	"github.com/myjinjin/sonic-odyssey-backend/internal/controller/http/mocks"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
	mocks2 "github.com/myjinjin/sonic-odyssey-backend/internal/usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUserController_SignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUserRepo := new(mocks2.UserRepository)
	mockUserUsecase := new(mocks.UserUsecase)
	userJwt := auth.NewUserJWT(mockUserRepo)
	testUserJwtAuth, err := auth.NewJWTMiddleware(
		auth.WithKey([]byte(os.Getenv("JWT_SECRET_KEY"))),
		auth.WithPayloadFunc(userJwt.PayloadFunc),
		auth.WithIdentityHandler(userJwt.IdentityHandler),
		auth.WithAuthenticator(userJwt.Authenticator),
		auth.WithAuthorizator(userJwt.Authorizator),
		auth.WithUnauthorized(userJwt.Unauthorized),
		auth.WithLoginResponse(userJwt.LoginResponse),
	)
	if err != nil {
		t.Fatalf("failed to create jwt auth middleware: %v", err)
	}
	testRouter := SetupRouter(mockUserUsecase, testUserJwtAuth)

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
