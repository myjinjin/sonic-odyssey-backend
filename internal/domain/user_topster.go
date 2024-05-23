package domain

import "time"

type UserTopster struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	UserID      uint   `gorm:"index"`
	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(500)"`
	GridSize    int
	ImageURL    string `gorm:"type:varchar(255)"`
	IsPublic    bool

	CreatedAt time.Time
	UpdatedAt time.Time

	TopsterAlbums []TopsterAlbum `gorm:"foreignKey:TopsterID"`
}
