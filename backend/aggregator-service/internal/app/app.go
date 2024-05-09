package app

import (
	"context"
	"github.com/Enthreeka/aggregator-service/internal/config"
	handler "github.com/Enthreeka/aggregator-service/internal/delivery/grpc"
	pgRepo "github.com/Enthreeka/aggregator-service/internal/repo/postgres"
	rdsRepo "github.com/Enthreeka/aggregator-service/internal/repo/redis"
	"github.com/Enthreeka/aggregator-service/internal/service"
	"github.com/Enthreeka/aggregator-service/pkg/logger"
	"github.com/Enthreeka/aggregator-service/pkg/postgres"
	"github.com/Enthreeka/aggregator-service/pkg/redis"
	pb "github.com/Entreeka/proto-proxy/go"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
}

func NewApp() *App { return &App{} }

func (a *App) Run(log *logger.Logger, cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	psql, err := postgres.New(ctx, 5, cfg.App.Postgres.LocalAddr)
	if err != nil {
		log.Fatal("failed to connect PostgreSQL: %v", err)
	}
	log.Info("database PostgreSQL connected on address: [%s]", cfg.App.Postgres.LocalAddr)

	rds, err := redis.New(ctx, cfg.App.Redis.Host, cfg.App.Redis.Password, cfg.App.Redis.MinIdleCons, cfg.App.Redis.Db)
	if err != nil {
		log.Fatal("failed to run redis: %v", err)
	}
	log.Info("database Redis connected on port: [%s]", cfg.App.Redis.Host)

	defer func() {
		psql.Close()
		rds.Close()
	}()

	postgresStorage := pgRepo.NewPostgresRepo(psql)
	redisStorage := rdsRepo.NewRedisRepo(rds)

	movieService := service.NewMovieService(redisStorage, postgresStorage, log)

	lis, err := net.Listen("tcp", cfg.App.GRPC.Addr)
	if err != nil {
		log.Fatal("failed to listen grpc addr: %s, error: %v", cfg.App.GRPC.Addr, err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			handler.LoggerUnaryInterceptorServer(log),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
				log.Error("Panic occurred %v:", p)
				return status.Errorf(codes.Internal, "Internal server error")
			})),
		),
	}

	h := handler.NewHandler(log, movieService)

	s := grpc.NewServer(opts...)
	pb.RegisterAggregatorServer(s, h)

	log.Info("Server listening on address: [http://%s]", cfg.App.GRPC.Addr)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to server grpc service on addr: %s, error: %v", cfg.App.GRPC.Addr, err)
	}

	<-ctx.Done()
	s.GracefulStop()
}
