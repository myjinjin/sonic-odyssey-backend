package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type UserFollowRepository interface {
	Create(userFollow *entities.UserFollow) error
	FindByID(id uint) (*entities.UserFollow, error)
	FindByFollowerIDAndFollowingID(followerID, followingID uint) (*entities.UserFollow, error)
	FindFollowersByUserID(userID uint) ([]*entities.UserFollow, error)
	FindFollowingsByUserID(userID uint) ([]*entities.UserFollow, error)
	CountFollowersByUserID(userID uint) (int64, error)
	CountFollowingsByUserID(userID uint) (int64, error)
	Delete(id uint) error
	DeleteByFollowerIDAndFollowingID(followerID, followingID uint) error
}
