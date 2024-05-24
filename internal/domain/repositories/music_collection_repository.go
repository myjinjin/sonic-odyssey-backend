package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type MusicCollectionRepository interface {
	Create(musicCollection *entities.MusicCollection) error
	FindByID(id uint) (*entities.MusicCollection, error)
	FindByUserID(userID uint, offset, limit int) ([]*entities.MusicCollection, error)
	Update(musicCollection *entities.MusicCollection) error
	Delete(id uint) error
}
