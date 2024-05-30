package entities

import "time"

type PasswordResetFlow struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	UserID    uint       `gorm:"not null"`
	User      User       `gorm:"foreignKey:UserID"`
	FlowID    string     `gorm:"type:varchar(255);unique;not null"` // 고유한 비밀번호 재설정 식별자
	ExpiresAt *time.Time `gorm:"not null"`                          // 비밀번호 재설정 링크 만료 시간

	CreatedAt time.Time
}
