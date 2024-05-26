package entities

import "time"

type UserLike struct {
	ID        uint  `gorm:"primaryKey;autoIncrement"`
	UserID    uint  `gorm:"index"`
	MusicID   *uint `gorm:"index"`
	PostID    *uint `gorm:"index"`
	CommentID *uint `gorm:"index"`
	Liked     bool  `gorm:"not null"`

	CreatedAt time.Time
}
