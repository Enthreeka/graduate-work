package grpc

import (
	"context"
	"github.com/Enthreeka/aggregator-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
)

type handler struct {
	pb.UnimplementedGatewayServer

	log *logger.Logger
}

func NewHandler(log *logger.Logger) *handler {
	return &handler{
		log: log,
	}
}

func (h *handler) SearchMovie(context.Context, *pb.SearchMovieRequest) (*pb.SearchMovieResponse, error) {

	return nil, nil
}
