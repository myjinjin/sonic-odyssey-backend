package repositories

import "github.com/myjinjin/sonic-odyssey-backend/internal/domain/entities"

type UserLikeRepository interface {
	Create(userLike *entities.UserLike) error
	FindByID(id uint) (*entities.UserLike, error)
	FindByUserIDAndMusicID(userID, musicID uint) (*entities.UserLike, error)
	FindByUserIDAndPostID(userID, postID uint) (*entities.UserLike, error)
	FindByUserIDAndCommentID(userID, commentID uint) (*entities.UserLike, error)
	CountLikesAndDislikesByMusicID(musicID uint) (likes int64, dislikes int64, err error)
	CountLikesAndDislikesByPostID(postID uint) (likes int64, dislikes int64, err error)
	CountLikesAndDislikesByCommentID(commentID uint) (likes int64, dislikes int64, err error)
	Update(userLike *entities.UserLike) error
	Delete(id uint) error
}
