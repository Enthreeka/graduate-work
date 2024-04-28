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

	newApp := app.NewApp()

	newApp.Run(log, cfg)
}
