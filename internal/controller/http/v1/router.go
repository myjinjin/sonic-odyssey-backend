package v1

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/myjinjin/sonic-odyssey-backend/docs"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
)

func SetupRouter(userUsecase usecase.UserUsecase) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// 요청 시작 시간 기록
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 요청 처리
		c.Next()

		// 요청 종료 시간 기록
		end := time.Now()
		latency := end.Sub(start)

		// 로그 필드 설정
		fields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		// 로그 출력
		logging.Log().Info("Request", fields...)
	})

	userController := NewUserController(userUsecase)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/users", userController.SignUp)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
