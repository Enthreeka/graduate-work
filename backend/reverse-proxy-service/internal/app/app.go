package app

import (
	"context"
	"github.com/Enthreeka/reverse-proxy-service/internal/config"
	"github.com/Enthreeka/reverse-proxy-service/internal/handler"
	redisRepository "github.com/Enthreeka/reverse-proxy-service/internal/repo/redis"
	"github.com/Enthreeka/reverse-proxy-service/pkg/grpc/client"
	"github.com/Enthreeka/reverse-proxy-service/pkg/logger"
	"github.com/Enthreeka/reverse-proxy-service/pkg/redis"
	pb "github.com/Entreeka/proto-proxy/go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config, log *logger.Logger) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	rds, err := redis.New(ctx, cfg.Redis.Host, cfg.Redis.Password, cfg.Redis.MinIdleCons, cfg.Redis.Db)
	if err != nil {
		log.Fatal("failed to run redis: %v", err)
	}

	redisRepo := redisRepository.NewRedisRepo(rds)

	clientElastic := client.NewGrpcClient(log, cfg.GRPC.ElasticsearchService, client.NewGatewayClientWrapper)
	ce, err := clientElastic.Connect()
	if err != nil {
		log.Fatal("failed to connect to clientElastic: %v", err)
	}
	clientElastic.Ping(ctx)

	clientAggregator := client.NewGrpcClient(log, cfg.GRPC.AggregatorService, client.NewAggregatorClientWrapper)
	ca, err := clientAggregator.Connect()
	if err != nil {
		log.Fatal("failed to connect to clientAggregator: %v", err)
	}
	clientAggregator.Ping(ctx)

	defer func() {
		rds.Close()
		clientAggregator.Close()
		clientElastic.Close()
	}()

	h := handler.Handler{
		Log:              log,
		RedisRepo:        redisRepo,
		ClientElastic:    ce.(pb.GatewayClient),
		ClientAggregator: ca.(pb.AggregatorClient),
	}

	mux := runtime.NewServeMux(
		//runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		//runtime.WithErrorHandler(),
		runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := make(map[string]string)

			md["url"] = r.URL.Path
			md["x-request-id"] = r.Header.Get("x-request-id")

			return metadata.New(md)
		}),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)

	err = pb.RegisterGatewayHandlerServer(ctx, mux, &h)
	if err != nil {
		log.Fatal("RegisterGatewayHandlerServer: %v", err)
	}

	srv := &http.Server{
		Addr:    cfg.Gateway.Port,
		Handler: handler.MiddlewareLogger(log, handler.EnableCors(mux)),
	}

	log.Info("Server listening at [http://localhost%s]", cfg.Gateway.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("failed to lister and server: %v", err)
	}
}
