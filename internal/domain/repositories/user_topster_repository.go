package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type UserTopsterRepository interface {
	Create(userTopster *entities.UserTopster) error
	FindByID(id uint) (*entities.UserTopster, error)
	FindByUserID(userID uint) ([]*entities.UserTopster, error)
	Update(userTopster *entities.UserTopster) error
	Delete(id uint) error
}
