package redis

import (
	"context"
	"github.com/Enthreeka/proxy-service/pkg/redis"
)

type Repo interface {
	IsExist(ctx context.Context) (bool, error)
}

type RedisRepo struct {
	redis *redis.Redis
}

func NewRedisRepo(redis *redis.Redis) Repo {
	return &RedisRepo{
		redis: redis,
	}
}

func (r *RedisRepo) IsExist(ctx context.Context) (bool, error) {

	return false, nil
}
