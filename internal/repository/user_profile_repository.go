package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type UserProfileRepository interface {
	Create(userProfile *domain.UserProfile) error
	FindByID(id uint) (*domain.UserProfile, error)
	FindByUserID(userID uint) (*domain.UserProfile, error)
	Update(userProfile *domain.UserProfile) error
	Delete(id uint) error
}
