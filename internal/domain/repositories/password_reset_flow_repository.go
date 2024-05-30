package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type PasswordResetFlowRepository interface {
	Create(flow *entities.PasswordResetFlow) error
	FindByFlowID(flowID string) (*entities.PasswordResetFlow, error)
	FindByUserID(userID uint) (*entities.PasswordResetFlow, error)
	DeleteByFlowID(flowID string) error
}
