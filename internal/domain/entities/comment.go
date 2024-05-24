package entities

import "time"

type Comment struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	UserID  uint   `gorm:"index"`
	PostID  uint   `gorm:"index"`
	Content string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time

	UserLikes []UserLike `gorm:"foreignKey:CommentID"`
}
