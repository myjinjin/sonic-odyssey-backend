package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type UserSocialAccountRepository interface {
	Create(userSocialAccount *entities.UserSocialAccount) error
	FindByID(id uint) (*entities.UserSocialAccount, error)
	FindByUserIDAndProvider(userID uint, provider string) (*entities.UserSocialAccount, error)
	Update(userSocialAccount *entities.UserSocialAccount) error
	Delete(id uint) error
}
