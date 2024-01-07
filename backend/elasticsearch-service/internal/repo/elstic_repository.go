package repo

import (
	"bytes"
	"context"
	"errors"
	"github.com/Enthreeka/elasticsearch-service/internal/entity"
	client "github.com/Enthreeka/elasticsearch-service/pkg/elasticsearch"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"github.com/Enthreeka/elasticsearch-service/pkg/serialize"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"strconv"
)

type Elastic interface {
	Index(ctx context.Context, data entity.Test) error
	Search() (*entity.Test, error)
}

type elasticRepo struct {
	Elastic *client.Elastic
	log     *logger.Logger
}

func NewElasticRepo(elasticsearch *client.Elastic, log *logger.Logger) *elasticRepo {
	return &elasticRepo{
		elasticsearch,
		log,
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

	e.log.Info("document indexed:", response.String())
	return nil
}

func (e *elasticRepo) IndexWithBulk(ctx context.Context, data []entity.Test) error {
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:     e.Elastic.Client,
		Index:      "index",
		NumWorkers: 5,
	})

	if err != nil {
		return errors.New("вывести ошибки при Bulk API")
	}

	for documentID, document := range data {
		bufferDocument, err := serialize.Marshal(document)
		if err != nil {
			return errors.New("вывести ошибки при сериализации данных")
		}

		err = bulkIndexer.Add(ctx,
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(documentID),
				Body:       bytes.NewReader(bufferDocument),
			})

		if err != nil {
			return errors.New("вывести ошибки при созданни индексов через Bulk API")
		}
	}

	defer bulkIndexer.Close(ctx)
	biStats := bulkIndexer.Stats()
	e.log.Info("document indexed:", biStats.NumIndexed)

	if biStats.NumFailed > 0 {
		return errors.New("biStats.NumFailed > 0")
	}

	return nil
}

func (e *elasticRepo) Search() (*entity.Test, error) {
	return nil, nil
}
