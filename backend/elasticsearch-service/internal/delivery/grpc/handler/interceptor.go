package handler

import (
	"context"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"google.golang.org/grpc"
)

func LoggerUnaryInterceptorServer(log *logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {

		if info.FullMethod == "/proxy_proto.Gateway/BulkAPI" {
			log.Info("Received request - [request: a lot of messages], [method: %s]", info.FullMethod)
		} else {
			log.Info("Received request - [request: %v], [method: %s]", req, info.FullMethod)
		}

		resp, err := handler(ctx, req)
		return resp, err
	}
}
