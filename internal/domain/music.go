package domain

import "time"

type Music struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Title     string `gorm:"type:varchar(255)"`
	AlbumID   uint   `gorm:"index"`
	SpotifyID string `gorm:"type:varchar(50)"`
	LastfmID  string `gorm:"type:varchar(50)"`

	CreatedAt time.Time
	UpdatedAt time.Time

	MusicGenreMapping      []MusicGenreMapping      `gorm:"foreignKey:MusicID"`
	MusicArtistMapping     []MusicArtistMapping     `gorm:"foreignKey:MusicID"`
	UserMusicLikes         []UserMusicLike          `gorm:"foreignKey:MusicID"`
	CollectionMusicMapping []CollectionMusicMapping `gorm:"foreignKey:MusicID"`
}
