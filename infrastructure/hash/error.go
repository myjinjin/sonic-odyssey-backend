package hash

import "errors"

var (
	ErrPasswordTooLong = errors.New("password is too long")
	ErrHashingFailure  = errors.New("failed to hash password")
)
