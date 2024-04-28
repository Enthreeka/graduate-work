package app

import (
	"context"
	"github.com/Enthreeka/proxy-service/internal/config"
	"github.com/Enthreeka/proxy-service/internal/handler"
	"github.com/Enthreeka/proxy-service/pkg/grpc/client"
	"github.com/Enthreeka/proxy-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

func Run(cfg *config.Config, log *logger.Logger) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//rds, err := redis.New(ctx, cfg)
	//if err != nil {
	//	log.Fatal("failed to run redis: %v", err)
	//}
	//
	//redisRepo := redisClient.NewRedisRepo(rds)

	clientElastic := client.NewGrpcClient(log, cfg.GRPC.ElasticsearchService)

	ce, err := clientElastic.Connect()
	if err != nil {
		log.Fatal("failed to connect to clientElastic: %v", err)
	}

	defer func() {
		clientElastic.Close()
	}()

	h := handler.Handler{
		Log:           log,
		RedisRepo:     nil,
		ClientElastic: ce.(pb.GatewayClient),
	}

	mux := runtime.NewServeMux()
	err = pb.RegisterGatewayHandlerServer(ctx, mux, &h)
	if err != nil {
		log.Fatal("RegisterGatewayHandlerServer: %v", err)
	}

	log.Info("Server listening at %s", cfg.Gateway.Port)
	if err := http.ListenAndServe(cfg.Gateway.Port, mux); err != nil {
		log.Fatal("failed to lister and server: %v", err)
	}
}
