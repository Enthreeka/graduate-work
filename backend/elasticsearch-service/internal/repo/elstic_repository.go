package repo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	client "github.com/Enthreeka/elasticsearch-service/pkg/elasticsearch"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"github.com/Enthreeka/elasticsearch-service/pkg/serialize"
	pb "github.com/Entreeka/proto-proxy/go"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"strconv"
)

type elasticRepo struct {
	elastic *client.Elastic
	log     *logger.Logger

	searchFields []string
}

const indexMovie = "movie"

func NewElasticRepo(elasticsearch *client.Elastic, log *logger.Logger, searchFields []string) Elastic {
	return &elasticRepo{
		elastic:      elasticsearch,
		log:          log,
		searchFields: searchFields,
	}
}

// Index - создание нового экземпляра сущности в индексе
func (e *elasticRepo) Index(ctx context.Context, data *pb.Movie) error {
	buf, err := serialize.Marshal(data)
	if err != nil {
		return fmt.Errorf("вывести ошибки при сериализации данных")
	}

	response, err := e.elastic.Index(
		indexMovie,
		bytes.NewReader(buf),
		e.elastic.Index.WithDocumentID(strconv.Itoa(int(data.Id))),
		e.elastic.Index.WithContext(ctx),
		e.elastic.Index.WithHuman(),
		e.elastic.Index.WithErrorTrace(),
	)
	defer func() {
		err := response.Body.Close()
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}()

	if err != nil {
		return fmt.Errorf("вывести ошибки при создании индекса")
	}

	if response.IsError() {
		return fmt.Errorf("вывести ошибки при возвращении ответа")
	}

	e.log.Info("document indexed: %s", response.String())
	return nil
}

// https://youtu.be/j0RiWwef8Z8?t=2582
func (e *elasticRepo) IndexWithBulk(ctx context.Context, data []*pb.Movie) error {
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:     e.elastic.Client,
		Index:      indexMovie,
		NumWorkers: 5,
		ErrorTrace: true,
	})

	if err != nil {
		return fmt.Errorf("вывести ошибки при Bulk API: %v", err)
	}

	for _, document := range data {
		bufferDocument, err := serialize.Marshal(document)
		if err != nil {
			return fmt.Errorf("вывести ошибки при сериализации данных: %v", err)
		}

		err = bulkIndexer.Add(
			ctx,
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(int(document.Id)),
				Body:       bytes.NewReader(bufferDocument), // Replace
			})

		if err != nil {
			return fmt.Errorf("вывести ошибки при созданни индексов через Bulk API: %v", err)
		}
	}

	defer func() {
		err := bulkIndexer.Close(ctx)
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}()

	biStats := bulkIndexer.Stats()
	e.log.Info("document indexed: %d,document failed: %d, document added: %d",
		biStats.NumIndexed, biStats.NumFailed, biStats.NumAdded)

	if biStats.NumFailed > 0 {
		return fmt.Errorf("biStats.NumFailed > 0")
	}

	return nil
}

// QueryByDocumentID default look up with documentID
func (e *elasticRepo) QueryByDocumentID(ctx context.Context, documentID int) (*pb.GetMovieByIDResponse, error) {
	response, err := e.elastic.Get(
		indexMovie,
		strconv.Itoa(documentID),
		e.elastic.Get.WithErrorTrace(),
		e.elastic.Get.WithContext(ctx),
	)

	if err != nil {
		return nil, fmt.Errorf("вывести ошибку о получении индекса")
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}()

	//getResponse := entity.GetResponse{}
	getResponse := new(pb.GetMovieByIDResponse)

	err = json.NewDecoder(response.Body).Decode(&getResponse) // Replace
	if err != nil {
		return nil, fmt.Errorf("вывести ошибку о сериализации")
	}

	if response.IsError() {
		return nil, fmt.Errorf("вывести ошибки при возвращении ответа: %s", response.Status())
	}

	movieName := getResponse.XSource.Title
	e.log.Info("Get movie with %d:%s", documentID, movieName)

	return getResponse, nil
}

// SearchIndex - нестрогий поиск по индексу
func (e *elasticRepo) SearchIndex(ctx context.Context, query string) (*pb.SearchMovieResponse, error) {
	var searchBuffer bytes.Buffer

	search := map[string]any{
		"query": map[string]any{
			"bool": map[string]any{
				"should": []map[string]any{
					{
						"multi_match": map[string]any{
							"query":     query,
							"operator":  "and",
							"fields":    e.searchFields,
							"fuzziness": "auto",
						},
					},
					//{
					//	"multi_match": map[string]any{
					//		"query":  e.keyboardLayoutManager.GetOppositeLayoutWord(term),
					//		"fields": e.searchFields,
					//	},
					//},
				},
				//"must_not": []map[string]any{
				//	{
				//		"range": map[string]any{
				//			"count_in_stock": map[string]any{
				//				"lt": 1,
				//			},
				//		},
				//	},
				//},
			},
		},
	}

	err := json.NewEncoder(&searchBuffer).Encode(search)
	if err != nil {
		return nil, fmt.Errorf("вывести ошибку о сериализации: %v", err)
	}

	response, err := e.elastic.Search(
		e.elastic.Search.WithContext(ctx),
		e.elastic.Search.WithErrorTrace(),
		e.elastic.Search.WithIndex(indexMovie),
		e.elastic.Search.WithBody(&searchBuffer),
		e.elastic.Search.WithTrackTotalHits(true),
		e.elastic.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("вывести ошибку о получении индекса: %v", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}()

	if response.IsError() {
		return nil, fmt.Errorf("error: %s", response.Status())
	}

	searchResponse := new(pb.SearchMovieResponse)
	err = json.NewDecoder(response.Body).Decode(&searchResponse) // Replace
	if err != nil {
		return nil, fmt.Errorf("вывести ошибку о сериализации: %v", err)
	}

	//if searchResponse.Hits.Total.Value > 0 {
	//	var testTitles []string
	//	for _, testTitle := range searchResponse.Hits.Hits {
	//		testTitles = append(testTitles, testTitle.Source.Title)
	//	}
	//	//fmt.Printf(strings.Join(testTitles, ","))
	//}

	return searchResponse, nil
}

// QueryAllDataInIndex - получение всех экземпляров, которые хранятся в индексе
func (e *elasticRepo) QueryAllDataInIndex(ctx context.Context) (*pb.GetAllMovieResponse, error) {
	response, err := e.elastic.Search(
		e.elastic.Search.WithContext(ctx),
		e.elastic.Search.WithErrorTrace(),
		e.elastic.Search.WithIndex(indexMovie),
		e.elastic.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("вывести ошибку о получении индекса")
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}()

	if response.IsError() {
		return nil, fmt.Errorf("error: %s", response.Status())
	}

	getResponse := new(pb.GetAllMovieResponse)
	err = json.NewDecoder(response.Body).Decode(&getResponse)
	if err != nil {
		return nil, fmt.Errorf("вывести ошибку о сериализации %v", err)
	}

	e.log.Info("getting numbers of movie: %d", getResponse.Hits.Total.Value)
	return getResponse, nil
}

// GetIndexInfo - получение информации об индексах
func (e *elasticRepo) GetIndexInfo(ctx context.Context, index []string) (map[string]interface{}, error) {
	response, err := e.elastic.Indices.Get(
		index,
		e.elastic.Indices.Get.WithContext(ctx),
		e.elastic.Indices.Get.WithErrorTrace(),
		e.elastic.Indices.Get.WithHuman(),
	)
	if err != nil {
		return nil, fmt.Errorf("вывести ошибку о получении индекса: %v", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			e.log.Error("failed to close elastic index response body: %v", err)
		}
	}()

	if response.IsError() {
		return nil, fmt.Errorf("failed go get index info: %s", response.Status())
	}

	var getResponse map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&getResponse)
	if err != nil {
		return nil, fmt.Errorf("вывести ошибку о сериализации %v", err)
	}

	return getResponse, nil
}
