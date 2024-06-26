package v1

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/myjinjin/sonic-odyssey-backend/docs"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/auth"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
)

func SetupRouter(userUsecase usecase.UserUsecase, musicUsecase usecase.MusicUsecase, jwtAuth *auth.JWTMiddleware) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		fields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		logging.Log().Info("Request", fields...)
	})

	userController := NewUserController(userUsecase, jwtAuth)
	musicController := NewMusicController(musicUsecase, jwtAuth)

	apiV1 := r.Group("/api/v1")
	{
		userGroup := apiV1.Group("/users")
		{
			userGroup.POST("", userController.SignUp)
			userGroup.POST("/password/recovery", userController.SendPasswordRecoveryEmail)
			userGroup.POST("/password/reset", userController.ResetPassword)
			userGroup.GET("/me", jwtAuth.MiddlewareFunc(), userController.GetMyUserInfo)
			userGroup.PATCH("/me", jwtAuth.MiddlewareFunc(), userController.PatchMyUser)
			userGroup.PUT("/me/password", jwtAuth.MiddlewareFunc(), userController.UpdatePassword)
		}

		authGroup := apiV1.Group("/auth")
		{
			authGroup.POST("/sign-in", jwtAuth.LoginHandler)
		}

		musicGroup := apiV1.Group("/music")
		{
			musicGroup.GET("/tracks", jwtAuth.MiddlewareFunc(), musicController.SearchTrack)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
