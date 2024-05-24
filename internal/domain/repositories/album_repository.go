package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type AlbumRepository interface {
	Create(album *entities.Album) error
	FindByID(id uint) (*entities.Album, error)
	FindByArtistID(artistID uint, offset, limit int) ([]*entities.Album, error)
	SearchByName(name string, offset, limit int) ([]*entities.Album, error)
	Update(album *entities.Album) error
	Delete(id uint) error
}
