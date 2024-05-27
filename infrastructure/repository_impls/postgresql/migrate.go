package postgresql

import (
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.User{},
		&entities.Album{},
		&entities.Artist{},
		&entities.CollectionMusicMapping{},
		&entities.Genre{},
		&entities.GenreCommunity{},
		&entities.Music{},
		&entities.MusicArtistMapping{},
		&entities.MusicCollection{},
		&entities.MusicGenreMapping{},
		&entities.Post{},
		&entities.Comment{},
		&entities.TopsterAlbum{},
		&entities.TranslatedArticle{},
		&entities.UserFollow{},
		&entities.UserLike{},
		&entities.UserProfile{},
		&entities.UserSocialAccount{},
		&entities.UserTopster{},
	)
}
