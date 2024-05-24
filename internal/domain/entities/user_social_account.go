package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserSocialAccount struct {
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	UserID         uint   `gorm:"index"`
	Provider       string `gorm:"type:enum('GOOGLE','KAKAO','NAVER', 'SPOTIFY');not null"`
	ProviderUserID string `gorm:"type:varchar(100)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
