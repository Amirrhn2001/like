package main

import (
	"fmt"
	"like/delivery/api"
	"like/domain"
	"like/repository/mongodb"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := loadEnvFile()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	repository, err := mongodb.NewMongoRepository()
	if err != nil {
		log.Fatal("Error connecting repository:", err)
	}
	service := domain.NewDomainService(repository)
	handler := api.NewHandler(service)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/api/like", handler.Like)
	r.POST("/api/dislike", handler.Dislike)
	r.GET("/api/likes/ref-uuid/:ref-uuid", handler.GetLikesByRefId)
	r.GET("/api/dislikes/ref-uuid/:ref-uuid", handler.GetDislikesByRefId)
	r.GET("/api/likes/ref-uuid/:ref-uuid/count", handler.GetLikesCountByRefId)
	r.GET("/api/dislikes/ref-uuid/:ref-uuid/count", handler.GetDislikesCountRefId)
	r.DELETE("/api/like-dislike/ref-uuid/:ref-uuid/user-uuid/:user-uuid", handler.DeleteLikeDislikeByRefIdAndUserId)

	Err := make(chan error, 1)
	go func() {
		fmt.Println("Listening on port", choosePort())
		Err <- http.ListenAndServe(choosePort(), r)
	}()

	fmt.Printf("Terminated %s", <-Err)
}

func loadEnvFile() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}

func choosePort() string {
	port := os.Getenv("port")
	if port == "" {
		return ":8080"
	}
	return fmt.Sprintf(":%s", port)
}
