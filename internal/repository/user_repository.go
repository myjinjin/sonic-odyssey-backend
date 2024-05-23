package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByNickname(nickname string) (*domain.User, error)
	Update(user *domain.User) error
	DeletE(id uint) error
}
