package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Email        string `gorm:"type:varchar(255);unique;"`
	PasswordHash string `gorm:"type:varchar(255)"`
	Name         string `gorm:"type:varchar(50)"`
	Nickname     string `gorm:"type:varchar(50);unique"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserProfiles       UserProfile
	UserSocialAccounts []UserSocialAccount `gorm:"foreignKey:UserID"`
	Followers          []UserFollow        `gorm:"foreignKey:FollowingID"`
	Following          []UserFollow        `gorm:"foreignKey:FollowerID"`
	UserTopsters       []UserTopster
	UserLikes          []UserLike        `gorm:"foreignKey:UserID"`
	MusicCollections   []MusicCollection `gorm:"foreignKey:UserID"`
}
