package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type PostRepository interface {
	Create(post *entities.Post) error
	FindByID(id uint) (*entities.Post, error)
	FindByUserID(userID uint, offset, limit int) ([]*entities.Post, error)
	FindByGenreCommunityID(genreCommunityID uint, offset, limit int) ([]*entities.Post, error)
	CountByGenreCommunityID(genreCommunityID uint) (int64, error)
	FindAll(offset, limit int) ([]*entities.Post, error)
	CountAll() (int64, error)
	SearchByTitle(title string, offset, limit int) ([]*entities.Post, error)
	SearchByContent(content string, offset, limit int) ([]*entities.Post, error)
	Update(post *entities.Post) error
	Delete(id uint) error
	CountLikesAndDislikesByID(id uint) (likes, dislikes int64, err error)
}
