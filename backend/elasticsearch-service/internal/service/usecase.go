package service

import (
	"context"
	"errors"
	"github.com/Enthreeka/elasticsearch-service/internal/entity"
	"github.com/Enthreeka/elasticsearch-service/internal/repo"
)

type ElasticUsecase interface {
	Index(ctx context.Context, data *entity.Test) error
	IndexBulkAPI(ctx context.Context, data []*entity.Test) error
	Search(ctx context.Context) (*entity.SearchResponse, error)
}

type elasticService struct {
	elasticRepo repo.Elastic
}

func NewElasticService(elasticRepo repo.Elastic) ElasticUsecase {
	return &elasticService{
		elasticRepo: elasticRepo,
	}
}

func (e *elasticService) Index(ctx context.Context, data *entity.Test) error {
	return e.elasticRepo.Index(ctx, data)
}

func (e *elasticService) IndexBulkAPI(ctx context.Context, data []*entity.Test) error {
	return e.elasticRepo.IndexWithBulk(ctx, data)
}

func (e *elasticService) Search(ctx context.Context) (*entity.SearchResponse, error) {
	searchResponse, err := e.elasticRepo.SearchIndex(ctx)
	if err != nil {
		return nil, errors.New("error")
	}

	if searchResponse == nil {
		return nil, errors.New("nil error")
	}

	return searchResponse, nil
}
