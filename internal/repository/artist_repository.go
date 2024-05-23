package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type ArtistRepository interface {
	Create(artist *domain.Artist) error
	FindByID(id uint) (*domain.Artist, error)
	FindByName(name string) ([]*domain.Artist, error)
	Update(artist *domain.Artist) error
	Delete(id uint) error
}
