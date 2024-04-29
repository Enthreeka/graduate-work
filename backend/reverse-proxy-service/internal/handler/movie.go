package handler

import (
	"context"
	"github.com/Enthreeka/reverse-proxy-service/internal/repo/redis"
	"github.com/Enthreeka/reverse-proxy-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
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
	case http.StatusBadRequest:
		return status.Errorf(codes.InvalidArgument, "Incorrect input data")
	default:
		return status.Errorf(codes.Internal, "Internal server error: %v", addInfo)
	}
}

func ErrorWrapper(err error) (*status.Status, error) {
	if err == nil {
		return nil, nil
	}

	if s, exist := status.FromError(err); exist {
		grpcStatusCode, _ := strconv.Atoi(s.Code().String())

		return s, SwitchToGrpcStatus(grpcStatusCode)
	}

	return nil, nil
}

func (h *Handler) GetAllMovie(ctx context.Context, _ *pb.GetAllMovieRequest) (*pb.GetAllMovieResponse, error) {
	response := new(pb.GetAllMovieResponse)

	response, err := h.ClientElastic.GetAllMovie(ctx, &pb.GetAllMovieRequest{})
	if s, err := ErrorWrapper(err); err != nil {
		h.Log.Error("GetAllMovie: error: %v, message: %v, internal_error: %v", err, s.Message(), s.Err())
		return nil, err
	}

	return response, nil
}

func (h *Handler) CreateNewMovie(ctx context.Context, req *pb.CreateNewMovieRequest) (*pb.CreateNewMovieResponse, error) {
	response := new(pb.CreateNewMovieResponse)

	return response, nil
}

func (h *Handler) GetIndices(ctx context.Context, req *pb.GetIndicesRequest) (*pb.GetIndicesResponse, error) {
	response := new(pb.GetIndicesResponse)

	return response, nil
}

// https://blog.logrocket.com/guide-to-grpc-gateway/
func (h *Handler) GetMovieByID(ctx context.Context, req *pb.GetMovieByIDRequest) (*pb.GetMovieByIDResponse, error) {
	response := new(pb.GetMovieByIDResponse)

	//md, _ := metadata.FromIncomingContext(ctx)
	//
	//h.Log.Info("%v", md)
	//movieID, ok := md["movie_id"]
	//if !ok {
	//
	//	return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	//}

	//h.Log.Info("%v", movieID)

	return response, nil
}

func (h *Handler) SearchMovie(ctx context.Context, req *pb.SearchMovieRequest) (*pb.SearchMovieResponse, error) {
	response := new(pb.SearchMovieResponse)

	//exist, err := h.RedisRepo.IsExist(ctx)
	//if err != nil {
	//	h.Log.Error("GetSearchInfo IN RedisRepo.IsExist: %v", err)
	//	return nil, SwitchToGrpcStatus(http.StatusInternalServerError)
	//}
	//
	//if exist {
	//
	//} else {
	//
	//}

	return response, nil
}

func (h *Handler) UpdateMovieData(ctx context.Context, req *pb.UpdateMovieDataRequest) (*pb.UpdateMovieDataResponse, error) {
	response := new(pb.UpdateMovieDataResponse)

	return response, nil
}

func (h *Handler) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	response := new(pb.DeleteMovieResponse)

	return response, nil
}
