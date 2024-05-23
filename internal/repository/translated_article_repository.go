package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type TranslatedArticleRepository interface {
	Create(translatedArticle *domain.TranslatedArticle) error
	FindByID(id uint) (*domain.TranslatedArticle, error)
	FindByOriginalURL(originalURL string) (*domain.TranslatedArticle, error)
	Update(translatedArticle *domain.TranslatedArticle) error
	Delete(id uint) error
}
