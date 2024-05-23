package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type MusicGenreMappingRepository interface {
	Create(musicGenreMapping *domain.MusicGenreMapping) error
	FindByID(id uint) (*domain.MusicGenreMapping, error)
	FindByMusicID(musicID uint) ([]*domain.MusicGenreMapping, error)
	FindByGenreID(genreID uint) ([]*domain.MusicGenreMapping, error)
	Delete(id uint) error
}
