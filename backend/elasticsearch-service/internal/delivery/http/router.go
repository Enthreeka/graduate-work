package http

import (
	"github.com/Enthreeka/elasticsearch-service/internal/delivery/http/handler"
	"github.com/Enthreeka/elasticsearch-service/pkg/elasticsearch"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"github.com/go-chi/chi/v5"
)

type Services struct {
	Elastic *elasticsearch.Elastic
}

func Router(log *logger.Logger, service Services) *chi.Mux {
	mux := chi.NewMux()

	test := handler.NewTestHandler(service.Elastic, log)

	mux.Route("/", func(r chi.Router) {
		r.Route("api/v1", func(r chi.Router) {
			r.Post("/", test.Create)
		})
	})

	return mux
}
