package redis

import (
	"context"
	"github.com/Enthreeka/reverse-proxy-service/pkg/logger"
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

func New(ctx context.Context, host, password string, minIdleCons, db int) (*Redis, error) {
	// TODO сделать настройки
	client := redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     password,
		MinIdleConns: minIdleCons,
		DB:           db,
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
