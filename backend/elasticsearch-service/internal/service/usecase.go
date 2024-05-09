package service

import (
	"context"
	pb "github.com/Entreeka/proto-proxy/go"
	"google.golang.org/protobuf/types/known/anypb"
)

type ElasticService interface {
	Index(ctx context.Context, data *pb.Movie) error
	IndexBulkAPI(ctx context.Context, data []*pb.Movie) error
	Search(ctx context.Context, query string, cache bool, redisKey string) (*pb.SearchMovieResponse, error)
	GetAllDocument(ctx context.Context) (*pb.GetAllMovieResponse, error)
	GetDocumentByID(ctx context.Context, id int) (*pb.GetMovieByIDResponse, error)
	GetIndexInfo(ctx context.Context, index []string) (map[string]*anypb.Any, error)
}
