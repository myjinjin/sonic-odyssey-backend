package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
	"github.com/myjinjin/sonic-odyssey-backend/pkg/utils"
)

var (
	ErrInvalidRequestBody = errors.New("invalid request body")
)

var errorStatusMap = map[error]int{
	usecase.ErrEmailAlreadyExists:        http.StatusBadRequest,
	usecase.ErrNicknameAlreadyExists:     http.StatusBadRequest,
	usecase.ErrInvalidPassword:           http.StatusBadRequest,
	usecase.ErrUserNotFound:              http.StatusBadRequest,
	usecase.ErrPasswordResetFlowNotFound: http.StatusBadRequest,
	usecase.ErrPasswordResetFlowExpired:  http.StatusBadRequest,
	usecase.ErrPassowrdNotMatched:        http.StatusBadRequest,

	usecase.ErrPasswordTooShort:      http.StatusBadRequest,
	usecase.ErrPasswordNoUppercase:   http.StatusBadRequest,
	usecase.ErrPasswordNoLowercase:   http.StatusBadRequest,
	usecase.ErrPasswordNoNumber:      http.StatusBadRequest,
	usecase.ErrPasswordNoSpecialChar: http.StatusBadRequest,

	usecase.ErrHashingPassword: http.StatusInternalServerError,
	usecase.ErrEncryptingEmail: http.StatusInternalServerError,
	usecase.ErrDecryptingEmail: http.StatusInternalServerError,

	usecase.ErrSendingEmail: http.StatusInternalServerError,

	usecase.ErrCreatingRecord: http.StatusInternalServerError,
	usecase.ErrFindingRecord:  http.StatusInternalServerError,
	usecase.ErrUpdatingRecord: http.StatusInternalServerError,
	usecase.ErrDeletingRecord: http.StatusInternalServerError,

	ErrInvalidRequestBody: http.StatusBadRequest,
}

func HandleError(c *gin.Context, err error) {
	if validationErrors, ok := err.(utils.ValidationErrors); ok {
		c.JSON(http.StatusBadRequest, NewErrorResponse(validationErrors))
		return
	}

	if status, ok := errorStatusMap[err]; ok {
		c.JSON(status, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
}
