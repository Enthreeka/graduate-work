package app

import (
	"context"
	"fmt"
	"github.com/Enthreeka/elasticsearch-service/internal/config"
	"github.com/Enthreeka/elasticsearch-service/internal/delivery/http"
	"github.com/Enthreeka/elasticsearch-service/pkg/elasticsearch"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(log *logger.Logger, cfg *config.Config) error {
	_, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	es, err := elasticsearch.New("", "")
	if err != nil {
		log.Fatal("error creating the client: %s", err)
	}

	server := http.NewServer(log, http.Services{Elastic: es}, http.ServerOption{
		Addr: fmt.Sprintf(":%s", cfg.HTTPServer.Port),
	})

	log.Info("Starting http server on %s: %s:%s", cfg.HTTPServer.TypeServer, cfg.HTTPServer.Hostname, cfg.HTTPServer.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("listen: %s", err)
	}

	return nil
}
