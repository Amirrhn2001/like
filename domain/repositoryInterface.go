package domain

type Repository interface {
	CreateLikeDislikeRepository(likeDislike *LikeDislike) error
	GetOneLikeDislikeRepository(filters []Filter) (*LikeDislike, error)
	GetLikeDislikesRepository(filters []Filter) ([]LikeDislike, error)
	GetLikeDislikeCountRepository(filters []Filter) (int64, error)
	UpdateOneLikeDislikeRepository(filters []Filter, likeDislike *LikeDislike) error
	DeleteOneLikeDislikeRepository(filters []Filter) error
}