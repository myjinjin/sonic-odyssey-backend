package v1

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/auth"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/logging"
	"github.com/myjinjin/sonic-odyssey-backend/internal/controller/http/mocks"
	mocks2 "github.com/myjinjin/sonic-odyssey-backend/internal/usecase/mocks"
	"go.uber.org/zap"
)

var (
	mockUserRepo     *mocks2.UserRepository
	mockUserUsecase  *mocks.UserUsecase
	mockMusicUsecase *mocks.MusicUsecase
	userJwt          auth.UserJWT
	testUserJwtAuth  *auth.JWTMiddleware
	testRouter       *gin.Engine
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	var err error
	mockUserRepo = new(mocks2.UserRepository)
	mockUserUsecase = new(mocks.UserUsecase)
	mockMusicUsecase = new(mocks.MusicUsecase)
	userJwt = auth.NewUserJWT(mockUserRepo)
	testUserJwtAuth, err = auth.NewJWTMiddleware(
		auth.WithKey([]byte(os.Getenv("JWT_SECRET_KEY"))),
		auth.WithPayloadFunc(userJwt.PayloadFunc),
		auth.WithIdentityHandler(userJwt.IdentityHandler),
		auth.WithAuthenticator(userJwt.Authenticator),
		auth.WithAuthorizator(userJwt.Authorizator),
		auth.WithUnauthorized(userJwt.Unauthorized),
		auth.WithLoginResponse(userJwt.LoginResponse),
	)
	if err != nil {
		logging.Log().Fatal("failed to create jwt auth middleware.", zap.Error(err))
	}
	testRouter = SetupRouter(mockUserUsecase, mockMusicUsecase, testUserJwtAuth)
	os.Exit(m.Run())
}
