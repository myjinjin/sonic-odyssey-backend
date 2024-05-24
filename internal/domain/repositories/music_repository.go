package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type MusicRepository interface {
	Create(music *entities.Music) error
	FindByID(id uint) (*entities.Music, error)
	FindByTitle(title string) ([]*entities.Music, error)
	FindByAlbumID(albumID uint) ([]*entities.Music, error)
	FindByGenreID(genreID uint, offset, limit int) ([]*entities.Music, error)
	FindByArtistID(artistID uint, offset, limit int) ([]*entities.Music, error)
	FindBySpotifyID(spotifyID string) (*entities.Music, error)
	FindByLastfmID(lastfmID string) (*entities.Music, error)
	SearchByTitle(title string, offset, limit int) ([]*entities.Music, error)
	SearchByArtist(artistName string, offset, limit int) ([]*entities.Music, error)
	SearchByAlbum(albumName string, offset, limit int) ([]*entities.Music, error)
	Update(music *entities.Music) error
	Delete(id uint) error
	CountLikesAndDislikesByID(id uint) (likes, dislikes int64, err error)
}
