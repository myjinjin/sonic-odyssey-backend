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

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type SignUpResponse struct {
	UserID uint `json:"user_id"`
}

func (u *userController) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c, ControllerError{Err: err, StatusCode: http.StatusBadRequest})
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
		switch err {
		case usecase.ErrEmailAlreadyExists, usecase.ErrNicknameAlreadyExists:
			HandleError(c, ControllerError{Err: err, StatusCode: http.StatusBadRequest})
		default:
			HandleError(c, err)
		}
		return
	}

	res := SignUpResponse{
		UserID: output.UserID,
	}

	c.JSON(http.StatusCreated, res)
}
