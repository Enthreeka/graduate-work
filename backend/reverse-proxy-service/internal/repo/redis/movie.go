package redis

import (
	"context"
	"errors"
	"github.com/Enthreeka/reverse-proxy-service/pkg/redis"
	rds "github.com/redis/go-redis/v9"
)

type Repo interface {
	IsExist(ctx context.Context, query string) (bool, error)
}

type redisRepo struct {
	*redis.Redis
}

func NewRedisRepo(redis *redis.Redis) Repo {
	return &redisRepo{
		redis,
	}
}

func (r *redisRepo) IsExist(ctx context.Context, query string) (bool, error) {
	_, err := r.Rds.Get(ctx, query).Result()
	if errors.Is(err, rds.Nil) {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}
