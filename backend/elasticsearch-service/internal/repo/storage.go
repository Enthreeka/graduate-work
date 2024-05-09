package repo

import (
	"context"
	pb "github.com/Entreeka/proto-proxy/go"
)

type Elastic interface {
	Index(ctx context.Context, data *pb.Movie) error
	IndexWithBulk(ctx context.Context, data []*pb.Movie) error
	QueryByDocumentID(ctx context.Context, documentID int) (*pb.GetMovieByIDResponse, error)
	SearchIndex(ctx context.Context, query string) (*pb.SearchMovieResponse, error)
	QueryAllDataInIndex(ctx context.Context) (*pb.GetAllMovieResponse, error)
	GetIndexInfo(ctx context.Context, index []string) (map[string]interface{}, error)
	QueryByDocumentIDs(ctx context.Context, documentIDs []int) (*pb.SearchMovieResponse, error)
}
