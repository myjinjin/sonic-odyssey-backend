package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
)

type UserController interface {
	SignUp(c *gin.Context)
	SendPasswordRecoveryEmail(c *gin.Context)
}

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &userController{
		userUsecase: userUsecase,
	}
}

// SignUp godoc
// @Summary User SignUp
// @Description User SignUp
// @Tags users
// @Accept json
// @Produce json
// @Param request body SignUpRequest true "SignUp Request"
// @Success 201 {object} SignUpResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/users [post]
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

// SendPasswordRecoveryEmail godoc
// @Summary Send password recovery email
// @Description 비밀번호 복구 이메일 전송
// @Tags users
// @Accept json
// @Produce json
// @Param request body SendPasswordRecoveryEmailRequest true "SendPasswordRecoveryEmailRequest Request"
// @Success 201 {object} SendPasswordRecoveryEmailResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/users/password/recovery [post]
func (u *userController) SendPasswordRecoveryEmail(c *gin.Context) {
	var req SendPasswordRecoveryEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c, ErrInvalidRequestBody)
		return
	}

	bashURL := getBaseURL(c)
	if err := u.userUsecase.SendPasswordRecoveryEmail(bashURL, req.Email); err != nil {
		HandleError(c, err)
		return
	}

	res := SendPasswordRecoveryEmailResponse{}
	c.JSON(http.StatusOK, res)
}

func getBaseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + c.Request.Host
}
