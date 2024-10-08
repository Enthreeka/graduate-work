package app

import (
	"context"
	"github.com/Enthreeka/elasticsearch-service/internal/config"
	"github.com/Enthreeka/elasticsearch-service/internal/delivery/grpc/handler"
	"github.com/Enthreeka/elasticsearch-service/internal/repo"
	"github.com/Enthreeka/elasticsearch-service/internal/service"
	"github.com/Enthreeka/elasticsearch-service/pkg/elasticsearch"
	"github.com/Enthreeka/elasticsearch-service/pkg/grpc/client"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

//func Run(log *logger.Logger, cfg *config.Config) error {
//	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
//	defer cancel()
//
//	es, err := elasticsearch.New(cfg.App.Elasticsearch.Addr)
//	if err != nil {
//		log.Fatal("error creating the client: %s", err)
//	}
//
//	log.Info("Connected to elasticsearch addr: %v", es.Transport.(*elastictransport.Client).URLs())
//	log.Info("Connected to kibana url: %s", cfg.App.Kibana.Addr)
//
//	elasticRepo := repo.NewElasticRepo(es, log)
//
//	ElasticService := service.NewElasticService(elasticRepo)
//
//	server := http.NewServer(log, http.Services{elastic: ElasticService}, http.ServerOption{
//		Addr: fmt.Sprintf(":%s", cfg.HTTPServer.Port),
//	})
//
//	go func() {
//		log.Info("Starting http server on %s: %s:%s", cfg.HTTPServer.TypeServer, cfg.HTTPServer.Hostname, cfg.HTTPServer.Port)
//		if err := server.ListenAndServe(); err != nil {
//			log.Fatal("listen: %s", err)
//		}
//	}()
//
//	<-ctx.Done()
//	server.Shutdown(ctx)
//	return nil
//}

type App struct {
	elasticService service.ElasticService
	elasticRepo    repo.Elastic
}

func NewApp() *App { return &App{} }

func (a *App) Run(log *logger.Logger, cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	es, err := elasticsearch.New(cfg.App.Elasticsearch.Addr)
	if err != nil {
		log.Fatal("error creating the client: %s", err)
	}

	log.Info("Connected to elasticsearch addr: %v", es.Transport.(*elastictransport.Client).URLs())
	log.Info("Connected to kibana addr: [%s]", cfg.App.Kibana.Addr)

	clientAggregator := client.NewGrpcClient(log, cfg.App.GRPC.AggregatorService, client.NewAggregatorClientWrapper)
	ca, err := clientAggregator.Connect()
	if err != nil {
		log.Fatal("failed to connect to clientAggregator: %v", err)
	}
	clientAggregator.Ping(ctx)

	a.elasticRepo = repo.NewElasticRepo(es, log, cfg.App.Elasticsearch.SearchFields)
	a.elasticService = service.NewElasticService(a.elasticRepo, ca.(pb.AggregatorClient), log)

	lis, err := net.Listen("tcp", cfg.App.GRPC.Addr)
	if err != nil {
		log.Fatal("failed to listen grpc addr: %s, error: %v", cfg.App.GRPC.Addr, err)
	}

	movieHandler := handler.NewMovieHandler(a.elasticService, log)

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			handler.LoggerUnaryInterceptorServer(log),
			grpc_recovery.UnaryServerInterceptor(),
			apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
		),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGatewayServer(s, movieHandler)

	log.Info("Server listening on address: [http://%s]", cfg.App.GRPC.Addr)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to server grpc service on addr: %s, error: %v", cfg.App.GRPC.Addr, err)
	}

	<-ctx.Done()
	s.GracefulStop()
}
