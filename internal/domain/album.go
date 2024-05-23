package domain

import "time"

type Album struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255)"`
	ArtistID    uint      `gorm:"index"`
	ReleaseDate time.Time `gorm:"type:date"`
	ImageURL    string    `gorm:"type:varchar(255)"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Music []Music `gorm:"foreignKey:AlbumID"`
}
