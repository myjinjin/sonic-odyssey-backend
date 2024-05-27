package postgresql

import (
	"errors"

	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entities.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return repositories.ErrCreate
	}
	return nil
}

func (r *UserRepository) FindByID(id uint) (*entities.User, error) {
	user := new(entities.User)
	err := r.db.First(user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrNotFound
		}
		return nil, repositories.ErrFind
	}
	return user, nil
}

func (r *UserRepository) FindByEmailHash(hashedEmail string) (*entities.User, error) {
	user := new(entities.User)
	if err := r.db.Where("email_hash = ?", hashedEmail).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrNotFound
		}
		return nil, repositories.ErrFind
	}
	return user, nil
}

func (r *UserRepository) FindByNickname(nickname string) (*entities.User, error) {
	user := new(entities.User)
	err := r.db.Where("nickname = ?", nickname).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrNotFound
		}
		return nil, repositories.ErrFind
	}
	return user, nil
}

func (r *UserRepository) Update(user *entities.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return repositories.ErrUpdate
	}
	return nil
}

func (r *UserRepository) Delete(id uint) error {
	if err := r.db.Delete(&entities.User{}, id).Error; err != nil {
		return repositories.ErrDelete
	}
	return nil
}
