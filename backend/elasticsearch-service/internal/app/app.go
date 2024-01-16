package app

import (
	"context"
	"fmt"
	"github.com/Enthreeka/elasticsearch-service/internal/config"
	"github.com/Enthreeka/elasticsearch-service/internal/delivery/http"
	"github.com/Enthreeka/elasticsearch-service/internal/repo"
	"github.com/Enthreeka/elasticsearch-service/internal/service"
	"github.com/Enthreeka/elasticsearch-service/pkg/elasticsearch"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"os"
	"os/signal"
	"syscall"
)

func Run(log *logger.Logger, cfg *config.Config) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	es, err := elasticsearch.New(cfg.JSON.Elasticsearch.Addr)
	if err != nil {
		log.Fatal("error creating the client: %s", err)
	}

	log.Info("Connected to elasticsearch addr: %v", es.Transport.(*elastictransport.Client).URLs())
	log.Info("Connected to kibana url: %s", cfg.JSON.Kibana.Addr)

	elasticRepo := repo.NewElasticRepo(es, log)

	elasticUsecase := service.NewElasticService(elasticRepo)

	server := http.NewServer(log, http.Services{Elastic: elasticUsecase}, http.ServerOption{
		Addr: fmt.Sprintf(":%s", cfg.HTTPServer.Port),
	})

	go func() {
		log.Info("Starting http server on %s: %s:%s", cfg.HTTPServer.TypeServer, cfg.HTTPServer.Hostname, cfg.HTTPServer.Port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("listen: %s", err)
		}
	}()

	<-ctx.Done()
	server.Shutdown(ctx)
	return nil
}
