package domain

import "time"

type MusicGenreMapping struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	MusicID uint `gorm:"index"`
	GenreID uint `gorm:"index"`

	CreatedAt time.Time
}
