package redis

import (
	"context"
	"github.com/Enthreeka/proxy-service/internal/config"
	"github.com/Enthreeka/proxy-service/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Rds *redis.Client
}

func (r *Redis) Close() {
	if r.Rds != nil {
		if err := r.Rds.Close(); err != nil {
			logger.Error("failed to close redis: %v", err)
		}
	}
}

func New(ctx context.Context, cfg *config.Config) (*Redis, error) {
	// TODO сделать настройки
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		//MinIdleConns: cfg.Redis.MinIdleConns,
		DB: cfg.Redis.Db,
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	rds := &Redis{
		Rds: client,
	}

	return rds, nil
}
