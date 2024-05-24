package entities

import "time"

type UserFollow struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	FollowerID  uint `gorm:"index"`
	FollowingID uint `gorm:"index"`

	CreatedAt time.Time
}
