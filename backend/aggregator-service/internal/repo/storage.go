package repo

import (
	"context"
	pb "github.com/Entreeka/proto-proxy/go"
)

type StorageRedis interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
	Del(ctx context.Context, key string) error
}

type StoragePostgres interface {
	SearchByText(ctx context.Context, text string) ([]*pb.Movie, error)
	GetByArrID(ctx context.Context, id []int64) ([]*pb.Movie, error)
	CreateMovie(ctx context.Context, id int64, title string, movie []byte) error
	IfExistsMovie(ctx context.Context, id int64) (bool, error)
}
