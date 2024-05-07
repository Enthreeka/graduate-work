package redis

import (
	"context"
	"errors"
	"github.com/Enthreeka/aggregator-service/internal/repo"
	"github.com/Enthreeka/aggregator-service/pkg/redis"
	clientRedis "github.com/redis/go-redis/v9"
	"time"
)

const timeExpiration = 60 * time.Hour

var ErrNotFound = errors.New("key not found")

type redisRepo struct {
	*redis.Redis
}

func NewRedisRepo(redis *redis.Redis) repo.StorageRedis {
	return &redisRepo{
		redis,
	}
}

func (r *redisRepo) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := r.Rds.Get(ctx, key).Result()
	if errors.Is(err, clientRedis.Nil) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func (r *redisRepo) Set(ctx context.Context, key string, value []byte) error {
	err := r.Rds.Set(ctx, key, value, timeExpiration).Err()
	return err
}

func (r *redisRepo) Del(ctx context.Context, key string) error {
	err := r.Rds.Del(ctx, key).Err()
	if errors.Is(err, clientRedis.Nil) {
		return nil
	}
	return err
}
