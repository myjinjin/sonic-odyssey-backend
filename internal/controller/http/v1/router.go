package v1

import (
	"github.com/gin-gonic/gin"
	_ "github.com/myjinjin/sonic-odyssey-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
)

func SetupRouter(userUsecase usecase.UserUsecase) *gin.Engine {
	r := gin.Default()

	userController := NewUserController(userUsecase)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/users", userController.SignUp)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
