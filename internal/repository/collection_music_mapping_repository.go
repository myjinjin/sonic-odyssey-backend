package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type CollectionMusicMappingRepository interface {
	Create(collectionMusicMapping *domain.CollectionMusicMapping) error
	FindByID(id uint) (*domain.CollectionMusicMapping, error)
	FindByCollectionID(collectionID uint) ([]*domain.CollectionMusicMapping, error)
	FindByMusicID(musicID uint) ([]*domain.CollectionMusicMapping, error)
	Delete(id uint) error
}
