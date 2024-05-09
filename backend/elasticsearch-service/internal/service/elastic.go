package service

import (
	"context"
	"fmt"
	"github.com/Enthreeka/elasticsearch-service/internal/repo"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"github.com/Enthreeka/elasticsearch-service/pkg/serialize"
	pb "github.com/Entreeka/proto-proxy/go"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/anypb"
)

type elasticService struct {
	elasticRepo repo.Elastic
	log         *logger.Logger

	clientAggregator pb.AggregatorClient
}

func NewElasticService(elasticRepo repo.Elastic,
	clientAggregator pb.AggregatorClient,
	log *logger.Logger,
) ElasticService {
	return &elasticService{
		elasticRepo:      elasticRepo,
		clientAggregator: clientAggregator,
		log:              log,
	}
}

func (e *elasticService) Index(ctx context.Context, data *pb.Movie) error {
	return e.elasticRepo.Index(ctx, data)
}

func (e *elasticService) IndexBulkAPI(ctx context.Context, data []*pb.Movie) error {
	return e.elasticRepo.IndexWithBulk(ctx, data)
}

func (e *elasticService) Search(ctx context.Context, query string, cache bool, redisKey string) (*pb.SearchMovieResponse, error) {
	searchResponse, err := e.elasticRepo.SearchIndex(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search index")
	}
	if searchResponse == nil || searchResponse.Hits.Total.Value == 0 {
		return nil, errors.New("no response from search")
	}

	if cache == true {
		go func(response *pb.SearchMovieResponse) {
			valueArr := make([]int64, response.Hits.Total.Value)
			for key, movie := range response.Hits.Hits {
				if movie.GetXSource() == nil {
					e.log.Error("source field not found in response: %s", movie.XId)
					continue
				}
				valueArr[key] = movie.GetXSource().Id
			}

			_, err = e.clientAggregator.SetCache(context.Background(), &pb.SetCacheRequest{
				Key:   redisKey,
				Value: valueArr,
			})
			if err != nil {
				e.log.Error("failed to set cache in clientAggregator: %v", err)
			}
			e.log.Info("successfully save to cache in clientAggregator - key: %s", redisKey)
		}(searchResponse)
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
