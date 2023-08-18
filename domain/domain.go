package domain

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type domain struct {
	Repo Repository
}

func NewDomainService(repo Repository) Service {
	return &domain{Repo: repo}
}

func (d domain) CreateLikeService(like *LikeDislike) (*LikeDislike, *Error) {
	filters := []Filter{
		{Field: "ref_uuid", Value: like.RefId, Operator: "$eq"},
		{Field: "user_uuid", Value: like.UserId, Operator: "$eq"},
		{Field: "status", Value: -1, Operator: "$ne"},
	}
	likeDislike, err := d.Repo.GetOneLikeDislikeRepository(filters)
	if err == nil && likeDislike != nil && likeDislike.Status == 1 {
		return nil, CreateError(http.StatusBadRequest, "Already Liked")
	}
	if err == nil && likeDislike != nil && likeDislike.Status == 2 {
		likeDislike.UpdatedAt = time.Now().Unix()
		likeDislike.Status = 1
		err := d.Repo.UpdateOneLikeDislikeRepository(filters, likeDislike)
		if err != nil {
			return nil, CreateError(http.StatusBadRequest, err.Error())
		}
		return likeDislike, nil
	}

	like.Id = uuid.NewString()
	like.CreatedAt = time.Now().Unix()
	like.Status = 1
	err = d.Repo.CreateLikeDislikeRepository(like)
	if err != nil {
		return nil, CreateError(http.StatusBadRequest, err.Error())
	}

	return like, nil
}

func (d domain) CreateDislikeService(dislike *LikeDislike) (*LikeDislike, *Error) {
	filters := []Filter{
		{Field: "ref_uuid", Value: dislike.RefId, Operator: "$eq"},
		{Field: "user_uuid", Value: dislike.UserId, Operator: "$eq"},
		{Field: "status", Value: -1, Operator: "$ne"},
	}
	likeDislike, err := d.Repo.GetOneLikeDislikeRepository(filters)
	if err == nil && likeDislike != nil && likeDislike.Status == 2 {
		return nil, CreateError(http.StatusBadRequest, "Already Disliked")
	}
	if err == nil && likeDislike != nil && likeDislike.Status == 1 {
		likeDislike.UpdatedAt = time.Now().Unix()
		likeDislike.Status = 2
		err := d.Repo.UpdateOneLikeDislikeRepository(filters, likeDislike)
		if err != nil {
			return nil, CreateError(http.StatusBadRequest, err.Error())
		}
		return likeDislike, nil
	}

	dislike.Id = uuid.NewString()
	dislike.CreatedAt = time.Now().Unix()
	dislike.Status = 2
	err = d.Repo.CreateLikeDislikeRepository(dislike)
	if err != nil {
		return nil, CreateError(http.StatusBadRequest, err.Error())
	}

	return dislike, nil
}

func (d domain) GetLikesByRefIdService(refId string) ([]LikeDislike, *Error) {
	filters := []Filter{
		{Field: "ref_uuid", Value: refId, Operator: "$eq"},
		{Field: "status", Value: 1, Operator: "$eq"},
	}

	likes, err := d.Repo.GetLikeDislikesRepository(filters)
	if err != nil {
		return nil, CreateError(http.StatusBadRequest, err.Error())
	}

	return likes, nil
}

func (d domain) GetDislikesByRefIdService(refId string) ([]LikeDislike, *Error) {
	filters := []Filter{
		{Field: "ref_uuid", Value: refId, Operator: "$eq"},
		{Field: "status", Value: 2, Operator: "$eq"},
	}

	dislikes, err := d.Repo.GetLikeDislikesRepository(filters)
	if err != nil {
		return nil, CreateError(http.StatusBadRequest, err.Error())
	}

	return dislikes, nil
}

func (d domain) GetLikeCountByRefIdService(refId string) (int64, *Error) {
	filters := []Filter{
		{Field: "ref_uuid", Value: refId, Operator: "$eq"},
		{Field: "status", Value: 1, Operator: "$eq"},
	}

	count, err := d.Repo.GetLikeDislikeCountRepository(filters)
	if err != nil {
		return 0, CreateError(http.StatusBadRequest, err.Error())
	}

	return count, nil
}

func (d domain) GetDislikeCountByRefIdService(refId string) (int64, *Error) {
	filters := []Filter{
		{Field: "ref_uuid", Value: refId, Operator: "$eq"},
		{Field: "status", Value: 2, Operator: "$eq"},
	}

	count, err := d.Repo.GetLikeDislikeCountRepository(filters)
	if err != nil {
		return 0, CreateError(http.StatusBadRequest, err.Error())
	}

	return count, nil
}

func (d domain) DeleteLikeDislikeByRefIdAndUserIdService(refId, userId string) *Error {
	filters := []Filter{
		{Field: "ref_uuid", Value: refId, Operator: "$eq"},
		{Field: "user_uuid", Value: userId, Operator: "$eq"},
		{Field: "status", Value: -1, Operator: "$ne"},
	}

	err := d.Repo.DeleteOneLikeDislikeRepository(filters)
	if err != nil {
		return CreateError(http.StatusBadRequest, err.Error())
	}

	return nil
}