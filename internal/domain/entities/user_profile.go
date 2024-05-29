package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	UserID          uint   `gorm:"foreignKey:ID"`
	ProfileImageURL string `gorm:"type:varchar(255);default:''"`
	Bio             string `gorm:"type:varchar(500);default:''"`
	Website         string `gorm:"type:varchar(255);default:''"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
