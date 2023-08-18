package mongodb

import (
	"context"
	"like/domain"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func NewMongoRepository() (domain.Repository, error) {
	url := os.Getenv("mongoURL")
	timeout, err := strconv.Atoi(os.Getenv("timeout"))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout) * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	repo := &mongoRepository{
		client: client,
		database: os.Getenv("database"),
		timeout: time.Duration(timeout) * time.Second,
	}
	return repo, nil
}

func (r mongoRepository) CreateLikeDislikeRepository(likeDislike *domain.LikeDislike) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.LikeDislikeCollection)

	_, err := collection.InsertOne(ctx, likeDislike)
	if err != nil {
		return err
	}

	return nil
}

func (r mongoRepository) GetOneLikeDislikeRepository(filters []domain.Filter) (*domain.LikeDislike, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.LikeDislikeCollection)
	filter, err := domain.CreteMongoFilter(filters)
	if err != nil {
		return nil, err
	}

	likeDislike := &domain.LikeDislike{}
	err = collection.FindOne(ctx, filter).Decode(likeDislike)
	if err != nil {
		return nil, err
	}
	if likeDislike == (&domain.LikeDislike{}) {
		return nil, mongo.ErrNoDocuments
	}

	return likeDislike, nil
}

func (r mongoRepository) GetLikeDislikesRepository(filters []domain.Filter) ([]domain.LikeDislike, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.LikeDislikeCollection)
	filter, err := domain.CreteMongoFilter(filters)
	if err != nil {
		return nil, err
	}

	results, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var likeDislikes []domain.LikeDislike
	err = results.All(ctx, &likeDislikes)
	if err != nil {
		return nil, err
	}

	err = results.Close(ctx)
	if err != nil {
		return nil, err
	}

	if len(likeDislikes) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return likeDislikes, nil
}

func (r mongoRepository) GetLikeDislikeCountRepository(filters []domain.Filter) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.LikeDislikeCollection)
	filter, err := domain.CreteMongoFilter(filters)
	if err != nil {
		return 0, err
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r mongoRepository) UpdateOneLikeDislikeRepository(filters []domain.Filter, likeDislike *domain.LikeDislike) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.LikeDislikeCollection)
	filter, err := domain.CreteMongoFilter(filters)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": likeDislike})
	if err != nil {
		return err
	}

	return nil
}

func (r mongoRepository) DeleteOneLikeDislikeRepository(filters []domain.Filter) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.LikeDislikeCollection)
	filter, err := domain.CreteMongoFilter(filters)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"status": -1}})
	if err != nil {
		return err
	}

	return nil
}