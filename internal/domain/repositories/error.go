package repositories

import "errors"

var (
	ErrCreate = errors.New("failed to create entity")
	ErrFind   = errors.New("failed to find entity")
	ErrUpdate = errors.New("failed to update entity")
	ErrDelete = errors.New("failed to delete entity")

	ErrNotFound = errors.New("entity not found")
)
