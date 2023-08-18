package domain

type Service interface {
	CreateLikeService(like *LikeDislike) (*LikeDislike, *Error)
	CreateDislikeService(dislike *LikeDislike) (*LikeDislike, *Error)
	GetLikesByRefIdService(refId string) ([]LikeDislike, *Error)
	GetDislikesByRefIdService(refId string) ([]LikeDislike, *Error)
	GetLikeCountByRefIdService(refId string) (int64, *Error)
	GetDislikeCountByRefIdService(refId string) (int64, *Error)
	DeleteLikeDislikeByRefIdAndUserIdService(refId, userId string) *Error
}