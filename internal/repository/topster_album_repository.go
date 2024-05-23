package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type TopsterAlbumRepository interface {
	Create(topsterAlbum *domain.TopsterAlbum) error
	FindByID(id uint) (*domain.TopsterAlbum, error)
	FindByTopsterID(topsterID uint) ([]*domain.TopsterAlbum, error)
	Update(topsterAlbum *domain.TopsterAlbum) error
	Delete(id uint) error
}
