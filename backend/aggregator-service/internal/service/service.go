package service

import (
	"context"
	pb "github.com/Entreeka/proto-proxy/go"
)

type AggregateService interface {
	SetDataRedis(ctx context.Context, key string, value []int64) error
	GetDataRedis(ctx context.Context, key string) ([]int64, error)
	GetMovie(ctx context.Context, id []int64) (*pb.TotalHits, error)
	StorageMovie(ctx context.Context, movieArr []*pb.Movie) error
}
