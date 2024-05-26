package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	FindByID(id uint) (*entities.User, error)
	FindByEmailHash(hashedEmail string) (*entities.User, error)
	FindByNickname(nickname string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uint) error
}
