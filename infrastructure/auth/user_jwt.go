package auth

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/hash"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPayload struct {
	Nickname string `json:"nickname"`
}

type UserJWT interface {
	Authenticator(c *gin.Context) (interface{}, error)
	PayloadFunc(data interface{}) jwt.MapClaims
	IdentityHandler(c *gin.Context) interface{}
	Authorizator(data interface{}, c *gin.Context) bool
	Unauthorized(c *gin.Context, code int, message string)
	LoginResponse(c *gin.Context, code int, token string, time time.Time)
}

func NewUserJWT(userRepo repositories.UserRepository) UserJWT {
	return &userJWT{userRepo}
}

type userJWT struct {
	userRepo repositories.UserRepository
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

	return user, nil
}

func (u *userJWT) PayloadFunc(data interface{}) jwt.MapClaims {
	if user, ok := data.(*entities.User); ok {
		return jwt.MapClaims{
			identityKey: user.Nickname,
		}
	}
	return jwt.MapClaims{}
}

func (u *userJWT) IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &entities.User{
		Nickname: claims[identityKey].(string),
	}
}

func (u *userJWT) Authorizator(data interface{}, c *gin.Context) bool {
	return true
}

func (u *userJWT) Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func (u *userJWT) LoginResponse(c *gin.Context, code int, token string, time time.Time) {
	c.JSON(code, gin.H{"expires_at": time, "token": token})
}
