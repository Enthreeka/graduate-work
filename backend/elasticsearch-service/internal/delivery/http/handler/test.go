package handler

import (
	"context"
	"encoding/json"
	"github.com/Enthreeka/elasticsearch-service/internal/apperror"
	"github.com/Enthreeka/elasticsearch-service/internal/entity"
	"github.com/Enthreeka/elasticsearch-service/internal/service"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	"github.com/google/uuid"
	"net/http"
)

type testHandler struct {
	elasticService service.ElasticService

	log *logger.Logger
}

func NewTestHandler(elasticService service.ElasticService, log *logger.Logger) *testHandler {
	return &testHandler{
		elasticService: elasticService,
		log:            log,
	}
}

type TestRequest struct {
	Name        string `json:"name"`
	MAIL        string `json:"mail"`
	Address     string `json:"address"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t *testHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := new(TestRequest)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		t.log.Error("json.NewDecoder: %v", err)
		DecodingError(w)
		return
	}

	test := &entity.Test{
		ID:          uuid.New().ClockSequence(),
		Name:        data.Name,
		MAIL:        data.MAIL,
		Address:     data.Address,
		Title:       data.Title,
		Description: data.Description,
	}

	if err := t.elasticService.Index(context.Background(), &pb.Movie{}); err != nil {
		t.log.Error("elasticService.Index: %v", err)
		HandleError(w, err, apperror.ParseHTTPErrStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.SetIndent(" ", " ")
	e.Encode(test)
}

func (t *testHandler) Get(w http.ResponseWriter, r *http.Request) {
	data := new(TestRequest)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		t.log.Error("json.NewDecoder: %v", err)
		DecodingError(w)
		return
	}

	searchResponse, err := t.elasticService.Search(context.Background())
	if err != nil {
		t.log.Error("elasticService.Search: %v", err)
		HandleError(w, err, apperror.ParseHTTPErrStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.SetIndent(" ", " ")
	e.Encode(searchResponse)
}
