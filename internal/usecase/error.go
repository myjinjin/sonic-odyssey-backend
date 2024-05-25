package usecase

import "errors"

var (
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrNicknameAlreadyExists = errors.New("nickname already exists")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrUserNotFound          = errors.New("user not found")

	ErrPasswordTooShort      = errors.New("password must be at least 8 characters long")
	ErrPasswordNoUppercase   = errors.New("password must contain at least one uppercase letter")
	ErrPasswordNoLowercase   = errors.New("password must contain at least one lowercase letter")
	ErrPasswordNoNumber      = errors.New("password must contain at least one number")
	ErrPasswordNoSpecialChar = errors.New("password must contain at least one special character")
)
