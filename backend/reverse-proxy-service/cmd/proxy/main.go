package main

import (
	"github.com/Enthreeka/reverse-proxy-service/internal/app"
	"github.com/Enthreeka/reverse-proxy-service/internal/config"
	"github.com/Enthreeka/reverse-proxy-service/pkg/logger"
)

const serviceName = "[REVERSE-PROXY] "

func main() {
	log := logger.New(serviceName)
	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed to get config", err)
	}

	app.Run(cfg, log)
}
