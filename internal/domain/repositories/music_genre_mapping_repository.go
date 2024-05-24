package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type MusicGenreMappingRepository interface {
	Create(musicGenreMapping *entities.MusicGenreMapping) error
	FindByID(id uint) (*entities.MusicGenreMapping, error)
	FindByMusicID(musicID uint) ([]*entities.MusicGenreMapping, error)
	FindByGenreID(genreID uint) ([]*entities.MusicGenreMapping, error)
	Delete(id uint) error
}
