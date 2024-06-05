package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/auth"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
	"github.com/myjinjin/sonic-odyssey-backend/pkg/utils"
)

type UserController interface {
	SignUp(c *gin.Context)
	SendPasswordRecoveryEmail(c *gin.Context)
	ResetPassword(c *gin.Context)
	GetMyUserInfo(c *gin.Context)
	PatchMyUser(c *gin.Context)
	UpdatePassword(c *gin.Context)
}

type userController struct {
	userUsecase usecase.UserUsecase
	jwtAuth     *auth.JWTMiddleware
}

func NewUserController(userUsecase usecase.UserUsecase, jwtAuth *auth.JWTMiddleware) UserController {
	return &userController{
		userUsecase: userUsecase,
		jwtAuth:     jwtAuth,
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

	baseURL := getBaseURL(c)
	if err := u.userUsecase.SendPasswordRecoveryEmail(baseURL, req.Email); err != nil {
		HandleError(c, err)
		return
	}

	res := SendPasswordRecoveryEmailResponse{}
	c.JSON(http.StatusOK, res)
}

// ResetPassword godoc
// @Summary Reset password
// @Description 비밀번호 재설정
// @Tags users
// @Accept json
// @Produce json
// @Param request body ResetPasswordRequest true "ResetPasswordRequest Request"
// @Success 201 {object} ResetPasswordResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/users/password/reset [post]
func (u *userController) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c, ErrInvalidRequestBody)
		return
	}

	err := u.userUsecase.ResetPassword(req.Password, req.FlowID)
	if err != nil {
		HandleError(c, err)
		return
	}

	res := ResetPasswordResponse{}
	c.JSON(http.StatusOK, res)
}

// GetMyUserInfo godoc
// @Summary      Get my user info
// @Description  JWT 인증 토큰 기반 내 유저 정보 조회
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  GetMyUserInfoResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/users/me [get]
func (u *userController) GetMyUserInfo(c *gin.Context) {
	payload := auth.GetUserPayload(c, u.jwtAuth.GinJWTMiddleware)
	userID := payload.UserID
	user, err := u.userUsecase.GetUserByID(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	res := GetMyUserInfoResponse{
		UserID:          user.ID,
		Email:           user.Email,
		Name:            user.Name,
		Nickname:        user.Nickname,
		ProfileImageURL: user.ProfileImageURL,
		Bio:             user.Bio,
		Website:         user.Website,
	}
	c.JSON(http.StatusOK, res)
}

// PatchMyUser godoc
// @Summary      Patch my user info
// @Description  JWT 인증 토큰 기반 내 유저 정보 업데이트
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param request body PatchMyUserRequest true "PatchMyUser Request"
// @Success      200  {object}  PatchMyUserResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/users/me [patch]
func (u *userController) PatchMyUser(c *gin.Context) {
	payload := auth.GetUserPayload(c, u.jwtAuth.GinJWTMiddleware)
	userID := payload.UserID
	_, err := u.userUsecase.GetUserByID(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	var req PatchMyUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c, ErrInvalidRequestBody)
		return
	}

	err = utils.ValidateRequest(&req)
	if err != nil {
		HandleError(c, err)
		return
	}

	err = u.userUsecase.PatchUser(userID, &usecase.PatchUserInput{
		Name:     req.Name,
		Nickname: req.Nickname,
		Bio:      req.Bio,
		Website:  req.Website,
	})
	if err != nil {
		HandleError(c, usecase.ErrUpdatingRecord)
		return
	}

	res := PatchMyUserResponse{}
	c.JSON(http.StatusOK, res)
}

// UpdatePassword godoc
// @Summary      Update my user password
// @Description  JWT 인증 토큰 기반 내 비밀번호 수정
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param request body UpdatePasswordRequest true "UpdatePassword Request"
// @Success      200  {object}  UpdatePasswordResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/users/me/password [put]
func (u *userController) UpdatePassword(c *gin.Context) {
	var req UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c, ErrInvalidRequestBody)
		return
	}

	userPayload := auth.GetUserPayload(c, u.jwtAuth.GinJWTMiddleware)

	input := usecase.UpdatePasswordInput{
		UserID:       userPayload.UserID,
		CurrPassword: req.CurrPassword,
		NewPassword:  req.NewPassword,
	}

	if err := u.userUsecase.UpdatePassword(input); err != nil {
		HandleError(c, err)
		return
	}

	res := UpdatePasswordResponse{}
	c.JSON(http.StatusOK, res)
}

func getBaseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + c.Request.Host
}
