package domain

import "time"

type Artist struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(255)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
