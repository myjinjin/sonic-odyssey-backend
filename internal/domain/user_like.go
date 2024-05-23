package domain

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

/*
TODO: add constraint

ALTER TABLE user_likes ADD CONSTRAINT chk_user_likes_target CHECK (
    (MusicID IS NOT NULL AND PostID IS NULL AND CommentID IS NULL) OR
    (MusicID IS NULL AND PostID IS NOT NULL AND CommentID IS NULL) OR
    (MusicID IS NULL AND PostID IS NULL AND CommentID IS NOT NULL)
);
*/
