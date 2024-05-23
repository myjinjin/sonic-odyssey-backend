package domain

import "time"

type MusicCollection struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	UserID      uint   `gorm:"index"`
	Name        string `gorm:"type:varchar(255)"`
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time

	CollectionMusicMapping []CollectionMusicMapping `gorm:"foreignKey:CollectionID"`
}
