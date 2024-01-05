package main

import (
	"github.com/Enthreeka/elasticsearch-service/internal/app"
	"github.com/Enthreeka/elasticsearch-service/internal/config"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
)

func main() {
	log := logger.New()

	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed load config: %v", err)
	}

	if err := app.Run(log, cfg); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}
