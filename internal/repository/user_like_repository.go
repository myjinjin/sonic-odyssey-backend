package repository

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain"

type UserLikeRepository interface {
	Create(userLike *domain.UserLike) error
	FindByID(id uint) (*domain.UserLike, error)
	FindByUserIDAndMusicID(userID, musicID uint) (*domain.UserLike, error)
	FindByUserIDAndPostID(userID, postID uint) (*domain.UserLike, error)
	FindByUserIDAndCommentID(userID, commentID uint) (*domain.UserLike, error)
	CountLikesAndDislikesByMusicID(musicID uint) (likes int64, dislikes int64, err error)
	CountLikesAndDislikesByPostID(postID uint) (likes int64, dislikes int64, err error)
	CountLikesAndDislikesByCommentID(commentID uint) (likes int64, dislikes int64, err error)
	Update(userLike *domain.UserLike) error
	Delete(id uint) error
}
