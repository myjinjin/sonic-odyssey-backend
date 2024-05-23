package domain

import "time"

type Post struct {
	ID               uint   `gorm:"primaryKey;autoIncrement"`
	UserID           uint   `gorm:"index"`
	GenreCommunityID uint   `gorm:"index"`
	Title            string `gorm:"type:varchar(100)"`
	Content          string

	CreatedAt time.Time
	UpdatedAt time.Time

	Comments  []Comment  `gorm:"foreignKey:PostID"`
	UserLikes []UserLike `gorm:"foreignKey:PostID"`
}
