package repository

import (
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain"
)

type AlbumRepository interface {
	Create(album *domain.Album) error
	FindByID(id uint) (*domain.Album, error)
	FindByArtistID(artistID uint, offset, limit int) ([]*domain.Album, error)
	SearchByName(name string, offset, limit int) ([]*domain.Album, error)
	Update(album *domain.Album) error
	Delete(id uint) error
}
