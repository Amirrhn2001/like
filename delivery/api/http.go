package api

import (
	"encoding/json"
	"io"
	"like/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler interface {
	Like(c *gin.Context)
	Dislike(c *gin.Context)
	GetLikesByRefId(c *gin.Context)
	GetDislikesByRefId(c *gin.Context)
	GetLikesCountByRefId(c *gin.Context)
	GetDislikesCountRefId(c *gin.Context)
	DeleteLikeDislikeByRefIdAndUserId(c *gin.Context)
}

type handler struct {
	Service domain.Service
}

func NewHandler(service domain.Service) Handler {
	return &handler{Service: service}
}

func (h handler) Like(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Err := domain.CreateError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	like := &domain.LikeDislike{}
	err = json.Unmarshal(body, like)
	if err != nil {
		Err := domain.CreateError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	data, Err := h.Service.CreateLikeService(like)
	response := domain.CreteResponse(data, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h handler) Dislike(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Err := domain.CreateError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	dislike := &domain.LikeDislike{}
	err = json.Unmarshal(body, dislike)
	if err != nil {
		Err := domain.CreateError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	data, Err := h.Service.CreateDislikeService(dislike)
	response := domain.CreteResponse(data, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h handler) GetLikesByRefId(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	refId := c.Param("ref-uuid")
	data, Err := h.Service.GetLikesByRefIdService(refId)
	response := domain.CreteResponse(data, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) GetDislikesByRefId(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	refId := c.Param("ref-uuid")
	data, Err := h.Service.GetDislikesByRefIdService(refId)
	response := domain.CreteResponse(data, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) GetLikesCountByRefId(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	refId := c.Param("ref-uuid")
	data, Err := h.Service.GetLikeCountByRefIdService(refId)
	response := domain.CreteResponse(data, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) GetDislikesCountRefId(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	refId := c.Param("ref-uuid")
	data, Err := h.Service.GetDislikeCountByRefIdService(refId)
	response := domain.CreteResponse(data, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) DeleteLikeDislikeByRefIdAndUserId(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	refId := c.Param("ref-uuid")
	userId := c.Param("user-uuid")
	Err := h.Service.DeleteLikeDislikeByRefIdAndUserIdService(refId, userId)
	response := domain.CreteResponse(nil, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
