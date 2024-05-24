package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type CollectionMusicMappingRepository interface {
	Create(collectionMusicMapping *entities.CollectionMusicMapping) error
	FindByID(id uint) (*entities.CollectionMusicMapping, error)
	FindByCollectionID(collectionID uint) ([]*entities.CollectionMusicMapping, error)
	FindByMusicID(musicID uint) ([]*entities.CollectionMusicMapping, error)
	Delete(id uint) error
}
