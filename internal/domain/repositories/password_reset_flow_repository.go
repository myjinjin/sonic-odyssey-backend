package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type PasswordResetFlowRepository interface {
	Create(flow *entities.PasswordResetFlow) error
	FindByID(id uint) (*entities.PasswordResetFlow, error)
	FindByUserID(userID uint) (*entities.PasswordResetFlow, error)
	Delete(id uint) error
}
