package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type MusicCollectionRepository interface {
	Create(musicCollection *domain.MusicCollection) error
	FindByID(id uint) (*domain.MusicCollection, error)
	FindByUserID(userID uint, offset, limit int) ([]*domain.MusicCollection, error)
	Update(musicCollection *domain.MusicCollection) error
	Delete(id uint) error
}
