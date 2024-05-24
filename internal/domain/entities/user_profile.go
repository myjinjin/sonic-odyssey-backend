package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	UserID          uint   `gorm:"index"`
	ProfileImageURL string `gorm:"type:varchar(255)"`
	Bio             string `gorm:"type:varchar(500)"`
	Website         string `gorm:"type:varchar(255)"`
	UserAgreedAt    time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
