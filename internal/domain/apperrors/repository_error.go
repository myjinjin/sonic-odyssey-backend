package apperrors

import "errors"

var (
	ErrEncrypt  = errors.New("failed to encrypt data")
	ErrDecrypt  = errors.New("failed to decrypt data")
	ErrCreate   = errors.New("failed to create entity")
	ErrFind     = errors.New("failed to find entity")
	ErrNotFound = errors.New("entity not found")
	ErrUpdate   = errors.New("failed to update entity")
	ErrDelete   = errors.New("failed to delete entity")
)
