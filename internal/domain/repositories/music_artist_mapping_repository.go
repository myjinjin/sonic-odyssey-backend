package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type MusicArtistMappingRepository interface {
	Create(musicArtistMapping *entities.MusicArtistMapping) error
	FindByID(id uint) (*entities.MusicArtistMapping, error)
	FindByMusicID(musicID uint) ([]*entities.MusicArtistMapping, error)
	FindByArtistID(artistID uint) ([]*entities.MusicArtistMapping, error)
	Delete(id uint) error
}
