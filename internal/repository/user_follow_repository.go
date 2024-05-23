package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type UserFollowRepository interface {
	Create(userFollow *domain.UserFollow) error
	FindByID(id uint) (*domain.UserFollow, error)
	FindByFollowerIDAndFollowingID(followerID, followingID uint) (*domain.UserFollow, error)
	FindFollowersByUserID(userID uint) ([]*domain.UserFollow, error)
	FindFollowingsByUserID(userID uint) ([]*domain.UserFollow, error)
	CountFollowersByUserID(userID uint) (int64, error)
	CountFollowingsByUserID(userID uint) (int64, error)
	Delete(id uint) error
	DeleteByFollowerIDAndFollowingID(followerID, followingID uint) error
}
