package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
)

type UserController interface {
	SignUp(c *gin.Context)
}

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &userController{
		userUsecase: userUsecase,
	}
}

func (u *userController) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c, ErrInvalidRequestBody)
		return
	}

	input := usecase.SignUpInput{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	}

	output, err := u.userUsecase.SignUp(input)
	if err != nil {
		HandleError(c, err)
		return
	}

	res := SignUpResponse{
		UserID: output.UserID,
	}

	c.JSON(http.StatusCreated, res)
}
