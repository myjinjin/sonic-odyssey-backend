package postgresql

import (
	"errors"

	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type PasswordResetFlowRepository struct {
	db *gorm.DB
}

func NewPasswordResetFlowRepository(db *gorm.DB) repositories.PasswordResetFlowRepository {
	return &PasswordResetFlowRepository{db}
}

func (r *PasswordResetFlowRepository) Create(flow *entities.PasswordResetFlow) error {
	err := r.db.Create(flow).Error
	if err != nil {
		return repositories.ErrCreate
	}
	return nil
}

func (r *PasswordResetFlowRepository) FindByFlowID(flowID string) (*entities.PasswordResetFlow, error) {
	flow := new(entities.PasswordResetFlow)
	err := r.db.Where("flow_id = ?", flowID).Preload("User").First(&flow).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrNotFound
		}
		return nil, repositories.ErrFind
	}
	return flow, nil
}

func (r *PasswordResetFlowRepository) FindByUserID(userID uint) (*entities.PasswordResetFlow, error) {
	flow := new(entities.PasswordResetFlow)
	err := r.db.Where("user_id = ?", userID).First(&flow).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrNotFound
		}
		return nil, repositories.ErrFind
	}
	return flow, nil
}

func (r *PasswordResetFlowRepository) DeleteByFlowID(flowID string) error {
	if err := r.db.Where("flow_id = ?", flowID).Delete(&entities.PasswordResetFlow{}).Error; err != nil {
		return repositories.ErrDelete
	}
	return nil
}
