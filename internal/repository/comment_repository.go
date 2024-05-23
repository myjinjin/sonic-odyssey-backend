package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type CommentRepository interface {
	Create(comment *domain.Comment) error
	FindByID(id uint) (*domain.Comment, error)
	FindByUserID(userID uint, offset, limit int) ([]*domain.Comment, error)
	FindByPostID(postID uint, offset, limit int) ([]*domain.Comment, error)
	Update(comment *domain.Comment) error
	Delete(id uint) error
}
