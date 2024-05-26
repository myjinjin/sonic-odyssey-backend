package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
)

type ControllerError struct {
	Err        error
	StatusCode int
}

func (e ControllerError) Error() string {
	return e.Err.Error()
}

func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *ControllerError:
		c.JSON(e.StatusCode, gin.H{"error": e.Err.Error()})
	case error:
		switch e {
		case usecase.ErrEmailAlreadyExists, usecase.ErrNicknameAlreadyExists,
			usecase.ErrPasswordTooShort, usecase.ErrPasswordNoUppercase,
			usecase.ErrPasswordNoLowercase, usecase.ErrPasswordNoNumber,
			usecase.ErrPasswordNoSpecialChar:
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
	}
}
