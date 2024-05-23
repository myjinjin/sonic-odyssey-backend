package domain

import "time"

type MusicArtistMapping struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	MusicID  uint `gorm:"index"`
	ArtistID uint `gorm:"index"`

	CreatedAt time.Time
}
