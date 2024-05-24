package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type TopsterAlbumRepository interface {
	Create(topsterAlbum *entities.TopsterAlbum) error
	FindByID(id uint) (*entities.TopsterAlbum, error)
	FindByTopsterID(topsterID uint) ([]*entities.TopsterAlbum, error)
	Update(topsterAlbum *entities.TopsterAlbum) error
	Delete(id uint) error
}
