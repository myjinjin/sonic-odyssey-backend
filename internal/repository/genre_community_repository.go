package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type GenresCommunityRepository interface {
	Create(genreCommunity *domain.GenreCommunity) error
	FindByID(id uint) (*domain.GenreCommunity, error)
	FindByGenreID(genreID uint) ([]*domain.GenreCommunity, error)
	Update(genresCommunity *domain.GenreCommunity) error
	Delete(id uint) error
}
