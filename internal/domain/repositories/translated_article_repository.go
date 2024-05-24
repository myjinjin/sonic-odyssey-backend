package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type TranslatedArticleRepository interface {
	Create(translatedArticle *entities.TranslatedArticle) error
	FindByID(id uint) (*entities.TranslatedArticle, error)
	FindByOriginalURL(originalURL string) (*entities.TranslatedArticle, error)
	Update(translatedArticle *entities.TranslatedArticle) error
	Delete(id uint) error
}
