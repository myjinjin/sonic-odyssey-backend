package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type MusicArtistMappingRepository interface {
	Create(musicArtistMapping *domain.MusicArtistMapping) error
	FindByID(id uint) (*domain.MusicArtistMapping, error)
	FindByMusicID(musicID uint) ([]*domain.MusicArtistMapping, error)
	FindByArtistID(artistID uint) ([]*domain.MusicArtistMapping, error)
	Delete(id uint) error
}
