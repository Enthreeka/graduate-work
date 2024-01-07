package repo

import (
	"bytes"
	"context"
	"errors"
	"github.com/Enthreeka/elasticsearch-service/internal/entity"
	"github.com/Enthreeka/elasticsearch-service/pkg/elasticsearch"
	"github.com/Enthreeka/elasticsearch-service/pkg/serialize"
	"log"
	"strconv"
)

type Elastic interface {
	Index(ctx context.Context, data entity.Test) error
	Search() (*entity.Test, error)
}

type elasticRepo struct {
	Elastic *elasticsearch.Elastic
}

func NewElasticRepo(elasticsearch *elasticsearch.Elastic) *elasticRepo {
	return &elasticRepo{
		elasticsearch,
	}
}

func (e *elasticRepo) Index(ctx context.Context, data entity.Test) error {
	buf, err := serialize.Marshal(data)
	if err != nil {
		return errors.New("вывести ошибки при сериализации данных")
	}

	response, err := e.Elastic.Index(
		"index",
		bytes.NewReader(buf),
		e.Elastic.Index.WithDocumentID(strconv.Itoa(data.ID)),
		e.Elastic.Index.WithContext(ctx),
		e.Elastic.Index.WithHuman(),
	)
	defer response.Body.Close()

	if err != nil {
		return errors.New("вывести ошибки при создании индекса")
	}

	if response.IsError() {
		return errors.New("вывести ошибки при возвращении ответа")
	}

	log.Println("document indexed:", response.String())
	return nil
}

func (e *elasticRepo) Search() (*entity.Test, error) {
	return nil, nil
}
