package handler

import (
	"context"
	"github.com/Enthreeka/proxy-service/internal/repo/redis"
	"github.com/Enthreeka/proxy-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Handler struct {
	pb.UnimplementedGatewayServer

	Log           *logger.Logger
	RedisRepo     redis.Repo
	ClientElastic pb.GatewayClient
}

func SwitchToGrpcStatus(statusCode int, addInfo ...any) error {
	switch statusCode {
	case http.StatusNotFound:
		return status.Errorf(codes.NotFound, "Movie not found")
	default:
		return status.Errorf(codes.Internal, "Internal server error")
	}

}

func (h *Handler) GetSearchInfo(ctx context.Context, _ *pb.GetSearchInfoRequest) (*pb.GetSearchInfoResponse, error) {
	response := new(pb.GetSearchInfoResponse)

	//exist, err := h.RedisRepo.IsExist(ctx)
	//if err != nil {
	//	h.log.Error("GetSearchInfo IN RedisRepo.IsExist: %v", err)
	//	return nil, SwitchToGrpcStatus(http.StatusInternalServerError)
	//}
	//
	//if exist {
	//	response, err = h.ClientElastic.GetSearchInfo(ctx, &pb.GetSearchInfoRequest{})
	//	if err != nil {
	//		h.log.Error("GetSearchInfo IN ClientSearchAggregator.Connect().GetSearchInfo: %v", err)
	//		return nil, SwitchToGrpcStatus(http.StatusInternalServerError)
	//	}
	//} else {
	//
	//}

	response, err := h.ClientElastic.GetSearchInfo(ctx, &pb.GetSearchInfoRequest{})
	if err != nil {
		h.Log.Error("GetSearchInfo IN ClientSearchAggregator.Connect().GetSearchInfo: error: %v", err)
		return nil, SwitchToGrpcStatus(http.StatusInternalServerError)
	}

	//s, exist := status.FromError(err)

	return response, nil
}
