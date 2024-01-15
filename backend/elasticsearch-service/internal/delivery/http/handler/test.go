package handler

import (
	"github.com/Enthreeka/elasticsearch-service/pkg/elasticsearch"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"net/http"
)

type testHandler struct {
	es *elasticsearch.Elastic

	log *logger.Logger
}

func NewTestHandler(es *elasticsearch.Elastic, log *logger.Logger) *testHandler {
	return &testHandler{
		es: es,
	}
}

func (t *testHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (t *testHandler) Get(w http.ResponseWriter, r *http.Request) {

}
