package handler

import (
	"context"
	"github.com/Enthreeka/reverse-proxy-service/internal/repo/redis"
	"github.com/Enthreeka/reverse-proxy-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
)

const getMovieApi = "/v1/api/movie/get/"

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
	h.Log.Info("CreateNewMovie request: %v", req.Movie)

	if err := req.Validate(); err != nil {
		h.Log.Error("CreateNewMovie: validation error: %v", err)
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	response, err := h.ClientElastic.CreateNewMovie(ctx, req)
	if s, err := ErrorWrapper(err); err != nil {
		h.Log.Error("CreateNewMovie: error: %v, message: %v, internal_error: %v", err, s.Message(), s.Err())
		return nil, err
	}

	return response, nil
}

func (h *Handler) GetIndices(ctx context.Context, _ *pb.GetIndicesRequest) (*pb.GetIndicesResponse, error) {
	response := new(pb.GetIndicesResponse)

	response, err := h.ClientElastic.GetIndices(ctx, &pb.GetIndicesRequest{})
	if s, err := ErrorWrapper(err); err != nil {
		h.Log.Error("CreateNewMovie: error: %v, message: %v, internal_error: %v", err, s.Message(), s.Err())
		return nil, err
	}

	return response, nil
}

// https://blog.logrocket.com/guide-to-grpc-gateway/
func (h *Handler) GetMovieByID(ctx context.Context, _ *pb.GetMovieByIDRequest) (*pb.GetMovieByIDResponse, error) {
	response := new(pb.GetMovieByIDResponse)

	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		h.Log.Error("metadata from incoming context is empty in GetMovieByID")
		return nil, SwitchToGrpcStatus(http.StatusInternalServerError, "metadata from incoming context is empty")
	}

	url := md.Get("url")[0]
	_, movieID, exist := strings.Cut(url, getMovieApi)
	if !exist {
		h.Log.Error("in url = [%s] not found some value", url)
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	movieIDInt, err := strconv.Atoi(movieID)
	if err != nil {
		h.Log.Error("movieID not [int] data type: %v", err)
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	h.Log.Info("GetMovieByID: redirect to elasticsearch-service - movie_id: %d", movieIDInt)
	response, err = h.ClientElastic.GetMovieByID(ctx, &pb.GetMovieByIDRequest{
		MovieId: int64(movieIDInt),
	})
	if s, err := ErrorWrapper(err); err != nil {
		h.Log.Error("GetMovieByID: error: %v, message: %v, internal_error: %v", err, s.Message(), s.Err())
		return nil, err
	}

	return response, nil
}

func (h *Handler) SearchMovie(ctx context.Context, req *pb.SearchMovieRequest) (*pb.SearchMovieResponse, error) {
	response := new(pb.SearchMovieResponse)

	if req.GetQuery() == "" {
		h.Log.Error("SearchMovie: empty query request")
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	movie, exist, err := h.RedisRepo.GetMovie(ctx, req.Query)
	if err != nil {
		h.Log.Error("SearchMovie: failed to search movie in redis - error: %v, movie_exist: %v", err, exist)
		return nil, SwitchToGrpcStatus(http.StatusInternalServerError, err)
	}
	if !exist {
		response, err = h.ClientElastic.SearchMovie(ctx, req)
		if s, err := ErrorWrapper(err); err != nil {
			h.Log.Error("SearchMovie: error: %v, message: %v, internal_error: %v", err, s.Message(), s.Err())
			return nil, err
		}
	} else {
		response.Movie = movie
		response.Status = "delivered from Redis"
	}

	return response, nil
}

func (h *Handler) UpdateMovieData(ctx context.Context, req *pb.UpdateMovieDataRequest) (*pb.UpdateMovieDataResponse, error) {
	response := new(pb.UpdateMovieDataResponse)

	if err := req.Validate(); err != nil {
		h.Log.Error("UpdateMovieData: validation error: %v", err)
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	response, err := h.ClientElastic.UpdateMovieData(ctx, req)
	if s, err := ErrorWrapper(err); err != nil {
		h.Log.Error("UpdateMovieData: error: %v, message: %v, internal_error: %v", err, s.Message(), s.Err())
		return nil, err
	}

	return response, nil
}

func (h *Handler) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	response := new(pb.DeleteMovieResponse)

	if req.MovieId == 0 && req.Title == "" {
		h.Log.Error("DeleteMovie: empty data in request")
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	response, err := h.ClientElastic.DeleteMovie(ctx, req)
	if s, err := ErrorWrapper(err); err != nil {
		h.Log.Error("DeleteMovie: error: %v, message: %v, internal_error: %v", err, s.Message(), s.Err())
		return nil, err
	}

	return response, nil
}
