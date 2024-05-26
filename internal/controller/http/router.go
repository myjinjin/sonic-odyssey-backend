package http

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/myjinjin/sonic-odyssey-backend/internal/controller/http/v1"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
)

func SetupRouter(userUsecase usecase.UserUsecase) *gin.Engine {
	r := gin.Default()

	userController := v1.NewUserController(userUsecase)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/users", userController.SignUp)
	}

	return r
}
