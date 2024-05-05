package service

import (
	"context"
	"fmt"
	"github.com/Enthreeka/elasticsearch-service/internal/repo"
	"github.com/Enthreeka/elasticsearch-service/pkg/serialize"
	pb "github.com/Entreeka/proto-proxy/go"
	"google.golang.org/protobuf/types/known/anypb"
)

type elasticService struct {
	elasticRepo repo.Elastic
}

func NewElasticService(elasticRepo repo.Elastic) ElasticService {
	return &elasticService{
		elasticRepo: elasticRepo,
	}
}

func (e *elasticService) Index(ctx context.Context, data *pb.Movie) error {
	return e.elasticRepo.Index(ctx, data)
}

func (e *elasticService) IndexBulkAPI(ctx context.Context, data []*pb.Movie) error {
	return e.elasticRepo.IndexWithBulk(ctx, data)
}

func (e *elasticService) Search(ctx context.Context, query string) (*pb.SearchMovieResponse, error) {
	searchResponse, err := e.elasticRepo.SearchIndex(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	if searchResponse == nil {
		return nil, fmt.Errorf("nil searchResponse")
	}

	return searchResponse, nil
}

func (e *elasticService) GetAllDocument(ctx context.Context) (*pb.GetAllMovieResponse, error) {
	return e.elasticRepo.QueryAllDataInIndex(ctx)
}

func (e *elasticService) GetDocumentByID(ctx context.Context, id int) (*pb.GetMovieByIDResponse, error) {
	return e.elasticRepo.QueryByDocumentID(ctx, id)
}

func (e *elasticService) GetIndexInfo(ctx context.Context, index []string) (map[string]*anypb.Any, error) {
	interfaceMap, err := e.elasticRepo.GetIndexInfo(ctx, index)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	indices := make(map[string]*anypb.Any, len(interfaceMap))
	for key, value := range interfaceMap {
		if value == nil {
			continue
		}
		anyValue, err := serialize.ConvertInterfaceToAny(value)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		indices[key] = anyValue
	}

	return indices, nil
}
