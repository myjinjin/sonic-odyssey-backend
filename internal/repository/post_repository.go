package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type PostRepository interface {
	Create(post *domain.Post) error
	FindByID(id uint) (*domain.Post, error)
	FindByUserID(userID uint, offset, limit int) ([]*domain.Post, error)
	FindByGenreCommunityID(genreCommunityID uint, offset, limit int) ([]*domain.Post, error)
	CountByGenreCommunityID(genreCommunityID uint) (int64, error)
	FindAll(offset, limit int) ([]*domain.Post, error)
	CountAll() (int64, error)
	SearchByTitle(title string, offset, limit int) ([]*domain.Post, error)
	SearchByContent(content string, offset, limit int) ([]*domain.Post, error)
	Update(post *domain.Post) error
	Delete(id uint) error
	CountLikesAndDislikesByID(id uint) (likes, dislikes int64, err error)
}
