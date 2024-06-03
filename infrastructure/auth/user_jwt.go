package auth

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/hash"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
)

type UserJWT interface {
	Authenticator(c *gin.Context) (interface{}, error)
	PayloadFunc(data interface{}) jwt.MapClaims
	IdentityHandler(c *gin.Context) interface{}
	Authorizator(data interface{}, c *gin.Context) bool
	Unauthorized(c *gin.Context, code int, message string)
	LoginResponse(c *gin.Context, code int, token string, time time.Time)
}

type userJWT struct {
	userRepo repositories.UserRepository
}

func NewUserJWT(userRepo repositories.UserRepository) UserJWT {
	return &userJWT{userRepo}
}

type LoginRequest struct {
	Email    string `json:"email" example:"odyssey@example.com"`
	Password string `json:"password" example:"Example123!"`
}

type UserPayload struct {
	UserID uint `json:"user_id"`
}
type LoginResponse struct {
	ExpiresAt time.Time `json:"expires_at" example:"2024-05-30T09:00:00Z"`
	Token     string    `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

type UnauthorizedResponse struct {
	Error string `json:"error" example:"incorrect Username or Password"`
}

func (u *userJWT) Authenticator(c *gin.Context) (interface{}, error) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	email := req.Email
	password := req.Password

	user, err := u.userRepo.FindByEmailHash(hash.SHA256EmailHasher().HashEmail(email))
	if err != nil {
		return "", jwt.ErrFailedAuthentication
	}

	if !hash.BCryptPasswordHasher().CheckPasswordHash(password, user.PasswordHash) {
		return "", jwt.ErrFailedAuthentication
	}

	userPayload := &UserPayload{UserID: user.ID}
	return userPayload, nil
}

func (u *userJWT) PayloadFunc(data interface{}) jwt.MapClaims {
	if payload, ok := data.(*UserPayload); ok {
		return jwt.MapClaims{
			identityKey: &payload,
		}
	}
	return jwt.MapClaims{}
}

func (u *userJWT) IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	if payload, ok := claims[identityKey]; ok {
		return payload.(map[string]interface{})
	}
	return nil
}

func (u *userJWT) Authorizator(data interface{}, c *gin.Context) bool {
	return true
}

func (u *userJWT) Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, UnauthorizedResponse{Error: message})
}

// LoginResponse godoc
// @Summary      User Login
// @Description  Responds with a JWT token and expiration time upon successful login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param 		 request body   LoginRequest	true "Login Request"
// @Success      200  {object}  LoginResponse
// @Failure      400  {object}  UnauthorizedResponse
// @Failure      401  {object}  UnauthorizedResponse
// @Router       /api/v1/auth/sign-in [post]
func (u *userJWT) LoginResponse(c *gin.Context, code int, token string, time time.Time) {
	c.JSON(code, LoginResponse{ExpiresAt: time, Token: token})
}
