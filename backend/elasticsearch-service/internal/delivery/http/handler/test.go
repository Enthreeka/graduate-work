package handler

import (
	"github.com/Enthreeka/elasticsearch-service/internal/service"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"net/http"
)

type testHandler struct {
	elasticService service.ElasticUsecase

	log *logger.Logger
}

func NewTestHandler(elasticService service.ElasticUsecase, log *logger.Logger) *testHandler {
	return &testHandler{
		elasticService: elasticService,
		log:            log,
	}
}

func (t *testHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (t *testHandler) Get(w http.ResponseWriter, r *http.Request) {

}
