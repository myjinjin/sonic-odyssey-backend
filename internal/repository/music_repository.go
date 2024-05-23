package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type MusicRepository interface {
	Create(music *domain.Music) error
	FindByID(id uint) (*domain.Music, error)
	FindByTitle(title string) ([]*domain.Music, error)
	FindByAlbumID(albumID uint) ([]*domain.Music, error)
	FindByGenreID(genreID uint, offset, limit int) ([]*domain.Music, error)
	FindByArtistID(artistID uint, offset, limit int) ([]*domain.Music, error)
	FindBySpotifyID(spotifyID string) (*domain.Music, error)
	FindByLastfmID(lastfmID string) (*domain.Music, error)
	SearchByTitle(title string, offset, limit int) ([]*domain.Music, error)
	SearchByArtist(artistName string, offset, limit int) ([]*domain.Music, error)
	SearchByAlbum(albumName string, offset, limit int) ([]*domain.Music, error)
	Update(music *domain.Music) error
	Delete(id uint) error
	CountLikesAndDislikesByID(id uint) (likes, dislikes int64, err error)
}
