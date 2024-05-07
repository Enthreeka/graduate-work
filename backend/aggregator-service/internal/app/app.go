package app

import (
	"context"
	"github.com/Enthreeka/aggregator-service/internal/config"
	handler "github.com/Enthreeka/aggregator-service/internal/delivery/grpc"
	"github.com/Enthreeka/aggregator-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
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

	lis, err := net.Listen("tcp", cfg.App.GRPC.Addr)
	if err != nil {
		log.Fatal("failed to listen grpc addr: %s, error: %v", cfg.App.GRPC.Addr, err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			handler.LoggerUnaryInterceptorServer(log),
			grpc_recovery.UnaryServerInterceptor(),
		),
	}

	h := handler.NewHandler(log)

	s := grpc.NewServer(opts...)
	pb.RegisterGatewayServer(s, h)

	log.Info("Server listening on address: [http://%s]", cfg.App.GRPC.Addr)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to server grpc service on addr: %s, error: %v", cfg.App.GRPC.Addr, err)
	}

	<-ctx.Done()
	s.GracefulStop()
}
