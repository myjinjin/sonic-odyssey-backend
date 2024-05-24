package entities

import "time"

type CollectionMusicMapping struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	CollectionID uint `gorm:"index"`
	MusicID      uint `gorm:"index"`

	CreatedAt time.Time
}
