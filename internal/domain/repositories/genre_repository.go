package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type GenreRepository interface {
	Create(genre *entities.Genre) error
	FindByID(id uint) (*entities.Genre, error)
	FindByName(name string) (*entities.Genre, error)
	Update(genre *entities.Genre) error
	Delete(id uint) error
}
