package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type UserProfileRepository interface {
	Create(userProfile *entities.UserProfile) error
	FindByID(id uint) (*entities.UserProfile, error)
	FindByUserID(userID uint) (*entities.UserProfile, error)
	Update(userProfile *entities.UserProfile) error
	Delete(id uint) error
}
