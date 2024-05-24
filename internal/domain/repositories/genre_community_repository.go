package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type GenresCommunityRepository interface {
	Create(genreCommunity *entities.GenreCommunity) error
	FindByID(id uint) (*entities.GenreCommunity, error)
	FindByGenreID(genreID uint) ([]*entities.GenreCommunity, error)
	Update(genresCommunity *entities.GenreCommunity) error
	Delete(id uint) error
}
