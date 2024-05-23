package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type UserTopsterRepository interface {
	Create(userTopster *domain.UserTopster) error
	FindByID(id uint) (*domain.UserTopster, error)
	FindByUserID(userID uint) ([]*domain.UserTopster, error)
	Update(userTopster *domain.UserTopster) error
	Delete(id uint) error
}
