package domain

import "time"

type GenresCommunity struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	GenreID     uint   `gorm:"index"`
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(255)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
