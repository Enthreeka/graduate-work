package service

import (
	"context"
	"github.com/Enthreeka/elasticsearch-service/internal/entity"
)

type ElasticService interface {
	Index(ctx context.Context, data *entity.Test) error
	IndexBulkAPI(ctx context.Context, data []*entity.Test) error
	Search(ctx context.Context) (*entity.SearchResponse, error)
}
