package service

import (
	"context"
	"github.com/Enthreeka/elasticsearch-service/internal/entity"
	pb "github.com/Entreeka/proto-proxy/go"
)

type ElasticService interface {
	Index(ctx context.Context, data *pb.Movie) error
	IndexBulkAPI(ctx context.Context, data []*pb.Movie) error
	Search(ctx context.Context) (*entity.SearchResponse, error)
	SearchAll(ctx context.Context) (*pb.GetAllMovieResponse, error)
}
