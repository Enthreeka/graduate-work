package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Enthreeka/reverse-proxy-service/pkg/redis"
	pb "github.com/Entreeka/proto-proxy/go"
	rds "github.com/redis/go-redis/v9"
)

type Repo interface {
	IsExist(ctx context.Context, query string) (bool, error)
	GetMovie(ctx context.Context, query string) (*pb.Movie, bool, error)
}

type redisRepo struct {
	*redis.Redis
}

func NewRedisRepo(redis *redis.Redis) Repo {
	return &redisRepo{
		redis,
	}
}

func (r *redisRepo) GetMovie(ctx context.Context, query string) (*pb.Movie, bool, error) {
	movie := new(pb.Movie)

	movieBytes, err := r.Rds.Get(ctx, query).Bytes()
	if errors.Is(err, rds.Nil) {
		return nil, false, nil
	}
	if err != nil {
		return nil, true, err
	}

	err = json.Unmarshal(movieBytes, movie)
	if err != nil {
		return nil, true, err
	}

	return movie, true, nil
}

func (r *redisRepo) IsExist(ctx context.Context, query string) (bool, error) {

	r.Rds.Get(ctx, query)

	return false, nil
}
