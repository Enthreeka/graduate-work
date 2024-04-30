package repo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Enthreeka/elasticsearch-service/internal/entity"
	client "github.com/Enthreeka/elasticsearch-service/pkg/elasticsearch"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"github.com/Enthreeka/elasticsearch-service/pkg/serialize"
	pb "github.com/Entreeka/proto-proxy/go"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"strconv"
)

type Elastic interface {
	Index(ctx context.Context, data *pb.Movie) error
	IndexWithBulk(ctx context.Context, data []*pb.Movie) error
	QueryByDocumentID(ctx context.Context, documentID int) error
	SearchIndex(ctx context.Context) (*entity.SearchResponse, error)
	QueryAllDataInIndex(ctx context.Context) (*pb.GetAllMovieResponse, error)
}

type elasticRepo struct {
	Elastic *client.Elastic
	log     *logger.Logger
}

const indexMovie = "movie"

func NewElasticRepo(elasticsearch *client.Elastic, log *logger.Logger) Elastic {
	return &elasticRepo{
		elasticsearch,
		log,
	}
}

func (e *elasticRepo) Index(ctx context.Context, data *pb.Movie) error {
	buf, err := serialize.Marshal(data)
	if err != nil {
		return errors.New("вывести ошибки при сериализации данных")
	}

	response, err := e.Elastic.Index(
		indexMovie,
		bytes.NewReader(buf),
		e.Elastic.Index.WithDocumentID(strconv.Itoa(int(data.Id))),
		e.Elastic.Index.WithContext(ctx),
		e.Elastic.Index.WithHuman(),
	)
	defer func() {
		err := response.Body.Close()
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}()

	if err != nil {
		return errors.New("вывести ошибки при создании индекса")
	}

	if response.IsError() {
		return errors.New("вывести ошибки при возвращении ответа")
	}

	e.log.Info("document indexed: %s", response.String())
	return nil
}

// https://youtu.be/j0RiWwef8Z8?t=2582
func (e *elasticRepo) IndexWithBulk(ctx context.Context, data []*pb.Movie) error {
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:     e.Elastic.Client,
		Index:      indexMovie,
		NumWorkers: 5,
	})

	if err != nil {
		return errors.New("вывести ошибки при Bulk API")
	}

	for _, document := range data {
		bufferDocument, err := serialize.Marshal(document)
		if err != nil {
			return errors.New("вывести ошибки при сериализации данных")
		}

		err = bulkIndexer.Add(ctx,
			esutil.BulkIndexerItem{
				Action:     indexMovie,
				DocumentID: strconv.Itoa(int(document.Id)),
				Body:       bytes.NewReader(bufferDocument), // Replace
			})

		if err != nil {
			return errors.New("вывести ошибки при созданни индексов через Bulk API")
		}
	}

	defer func(ctx context.Context) {
		err := bulkIndexer.Close(ctx)
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}(ctx)

	biStats := bulkIndexer.Stats()
	e.log.Info("document indexed:", biStats.NumIndexed)

	if biStats.NumFailed > 0 {
		return errors.New("biStats.NumFailed > 0")
	}

	return nil
}

// QueryByDocumentID default look up with documentID
func (e *elasticRepo) QueryByDocumentID(ctx context.Context, documentID int) error {
	response, err := e.Elastic.Get(indexMovie, strconv.Itoa(documentID))
	if err != nil {
		return errors.New("вывести ошибку о получении индекса")
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}()

	getResponse := entity.GetResponse{}

	err = json.NewDecoder(response.Body).Decode(&getResponse) // Replace
	if err != nil {
		return errors.New("вывести ошибку о сериализации")
	}

	test := getResponse.Source.Title
	e.log.Info("test with %d:%s", documentID, test)

	if response.IsError() {
		return errors.New("вывести ошибки при возвращении ответа")
	}

	return nil
}

func (e *elasticRepo) SearchIndex(ctx context.Context) (*entity.SearchResponse, error) {
	var searchBuffer bytes.Buffer
	search := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{},
			},
		},
	}

	err := json.NewEncoder(&searchBuffer).Encode(search)
	if err != nil {
		return nil, errors.New("вывести ошибку о сериализации")
	}

	response, err := e.Elastic.Search(
		e.Elastic.Search.WithContext(ctx),
		e.Elastic.Search.WithIndex(indexMovie),
		e.Elastic.Search.WithBody(&searchBuffer),
		e.Elastic.Search.WithTrackTotalHits(true),
		e.Elastic.Search.WithPretty(),
	)

	defer func() {
		err := response.Body.Close()
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}()

	searchResponse := &entity.SearchResponse{}
	err = json.NewDecoder(response.Body).Decode(&searchResponse) // Replace
	if err != nil {
		return nil, errors.New("вывести ошибку о сериализации")
	}

	if searchResponse.Hits.Total.Value > 0 {
		var testTitles []string
		for _, testTitle := range searchResponse.Hits.Hits {
			testTitles = append(testTitles, testTitle.Source.Title)
		}
		//fmt.Printf(strings.Join(testTitles, ","))
	}

	return searchResponse, nil
}

func (e *elasticRepo) QueryAllDataInIndex(ctx context.Context) (*pb.GetAllMovieResponse, error) {
	response, err := e.Elastic.Search(
		e.Elastic.Search.WithContext(ctx),
		e.Elastic.Search.WithIndex(indexMovie),
		e.Elastic.Search.WithPretty(),
	)
	if err != nil {
		return nil, errors.New("вывести ошибку о получении индекса")
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}()

	getResponse := new(pb.GetAllMovieResponse)
	err = json.NewDecoder(response.Body).Decode(&getResponse) // Replace
	if err != nil {
		return nil, fmt.Errorf("вывести ошибку о сериализации %v", err)
	}

	e.log.Info("getting numbers of movie: %d", getResponse.Hits.Total.Value)
	return getResponse, nil
}
