package usecase

import (
	"errors"

	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
)

var (
	ErrEmailAlreadyExists        = errors.New("email already exists")
	ErrNicknameAlreadyExists     = errors.New("nickname already exists")
	ErrInvalidPassword           = errors.New("invalid password")
	ErrUserNotFound              = errors.New("user not found")
	ErrPasswordResetFlowNotFound = errors.New("password reset flow not found")
	ErrPasswordResetFlowExpired  = errors.New("password reset flow is expired")

	ErrPasswordTooShort      = errors.New("password must be at least 8 characters long")
	ErrPasswordNoUppercase   = errors.New("password must contain at least one uppercase letter")
	ErrPasswordNoLowercase   = errors.New("password must contain at least one lowercase letter")
	ErrPasswordNoNumber      = errors.New("password must contain at least one number")
	ErrPasswordNoSpecialChar = errors.New("password must contain at least one special character")

	ErrHashingPassword = errors.New("failed to hash password")
	ErrEncryptingEmail = errors.New("failed to encrypt email")
	ErrDecryptingEmail = errors.New("failed to decrypt email")

	ErrSendingEmail = errors.New("failed to send email")

	ErrCreatingRecord = repositories.ErrCreate
	ErrFindingRecord  = repositories.ErrFind
	ErrUpdatingRecord = repositories.ErrUpdate
	ErrDeletingRecord = repositories.ErrDelete
)
