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
	"strings"
)

func Run(cfg *config.Config, log *logger.Logger) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rds, err := redis.New(ctx, cfg.Redis.Host, cfg.Redis.Password, cfg.Redis.MinIdleCons, cfg.Redis.Db)
	if err != nil {
		log.Fatal("failed to run redis: %v", err)
	}

	redisRepo := redisRepository.NewRedisRepo(rds)

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
		RedisRepo:     redisRepo,
		ClientElastic: ce.(pb.GatewayClient),
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
		Handler: handler.MiddlewareLogger(log, mux),
	}

	log.Info("Server listening at localhost%s", cfg.Gateway.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("failed to lister and server: %v", err)
	}
}

var allowedHeaders = map[string]struct{}{
	"x-request-id": {},
}

func isHeaderAllowed(s string) (string, bool) {
	// check if allowedHeaders contain the header
	if _, isAllowed := allowedHeaders[s]; isAllowed {
		// send uppercase header
		return strings.ToUpper(s), true
	}
	// if not in the allowed header, don't send the header
	return s, false
}
