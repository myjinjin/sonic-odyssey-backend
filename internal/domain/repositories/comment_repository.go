package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type CommentRepository interface {
	Create(comment *entities.Comment) error
	FindByID(id uint) (*entities.Comment, error)
	FindByUserID(userID uint, offset, limit int) ([]*entities.Comment, error)
	FindByPostID(postID uint, offset, limit int) ([]*entities.Comment, error)
	Update(comment *entities.Comment) error
	Delete(id uint) error
}
