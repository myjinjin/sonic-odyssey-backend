package domain

import "time"

type TranslatedArticle struct {
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	OriginalURL       string `gorm:"type:varchar(255)"`
	TranslatedTitle   string `gorm:"type:varchar(255)"`
	TranslatedContent string

	CreatedAt time.Time
	UpdatedAt time.Time
}
