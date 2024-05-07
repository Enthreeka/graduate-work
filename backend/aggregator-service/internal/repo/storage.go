package repo

import (
	"context"
	"github.com/Enthreeka/aggregator-service/internal/entity"
)

type StorageRedis interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
	Del(ctx context.Context, key string) error
}

type StoragePostgres interface {
	SearchByText(ctx context.Context, text string) ([]entity.Movie, error)
	GetByArrID(ctx context.Context, id []int) ([]entity.Movie, error)
}
