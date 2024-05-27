package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
)

func SetupRouter(userUsecase usecase.UserUsecase) *gin.Engine {
	r := gin.Default()

	userController := NewUserController(userUsecase)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/users", userController.SignUp)
	}

	return r
}
