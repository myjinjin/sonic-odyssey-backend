package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type UserSocialAccountRepository interface {
	Create(userSocialAccount *domain.UserSocialAccount) error
	FindByID(id uint) (*domain.UserSocialAccount, error)
	FindByUserIDAndProvider(userID uint, provider string) (*domain.UserSocialAccount, error)
	Update(userSocialAccount *domain.UserSocialAccount) error
	Delete(id uint) error
}
