package hash

import (
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

var (
	oncePassword sync.Once
	bcryptHasher PasswordHasher
)

type bcryptPasswordHasher struct{}

func BCryptPasswordHasher() PasswordHasher {
	oncePassword.Do(func() { bcryptHasher = &bcryptPasswordHasher{} })
	return bcryptHasher
}

func (b bcryptPasswordHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		switch err {
		case bcrypt.ErrPasswordTooLong:
			return "", ErrPasswordTooLong
		default:
			return "", ErrHashingFailure
		}
	}
	return string(bytes), nil
}

func (b bcryptPasswordHasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
