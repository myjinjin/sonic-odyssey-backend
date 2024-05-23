package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type GenreRepository interface {
	Create(genre *domain.Genre) error
	FindByID(id uint) (*domain.Genre, error)
	FindByName(name string) (*domain.Genre, error)
	Update(genre *domain.Genre) error
	Delete(id uint) error
}
