package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type ArtistRepository interface {
	Create(artist *entities.Artist) error
	FindByID(id uint) (*entities.Artist, error)
	FindByName(name string) ([]*entities.Artist, error)
	Update(artist *entities.Artist) error
	Delete(id uint) error
}
