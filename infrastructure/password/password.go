package password

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type bcryptPasswordHasher struct{}

func BCryptPasswordHasher() PasswordHasher {
	return &bcryptPasswordHasher{}
}

func (b bcryptPasswordHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (b bcryptPasswordHasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
