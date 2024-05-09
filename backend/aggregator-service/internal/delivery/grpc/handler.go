package grpc

import (
	"context"
	"errors"
	"github.com/Enthreeka/aggregator-service/internal/service"
	"github.com/Enthreeka/aggregator-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
)

type handler struct {
	pb.UnimplementedAggregatorServer

	log          *logger.Logger
	movieService service.AggregateService
}

func NewHandler(log *logger.Logger, movieService service.AggregateService) *handler {
	if movieService == nil {
		log.Fatal("movieService is required")
	}
	return &handler{
		log:          log,
		movieService: movieService,
	}
}

func (h *handler) SearchMovieAggregator(ctx context.Context, req *pb.SearchMovieAggregatorRequest) (*pb.SearchMovieAggregatorResponse, error) {
	h.log.Info("SearchMovieAggregator called: %v", req)

	idArr, err := h.movieService.GetDataRedis(ctx, req.RedisKey)
	if err != nil {
		h.log.Error("SearchMovieAggregator GetDataRedis: %v", err)
		return nil, err
	}
	h.log.Info("got idArr: %v", idArr)

	totalHits, err := h.movieService.GetMovie(ctx, idArr)
	if err != nil {
		h.log.Error("SearchMovieAggregator GetMovie: %v", err)
		return nil, err
	}

	return &pb.SearchMovieAggregatorResponse{
		Hits: totalHits,
	}, nil
}

func (h *handler) SetCache(ctx context.Context, req *pb.SetCacheRequest) (*pb.SetCacheResponse, error) {
	err := h.movieService.SetDataRedis(ctx, req.Key, req.Value)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handler) GetCache(ctx context.Context, req *pb.GetCacheRequest) (*pb.GetCacheResponse, error) {
	h.log.Info("GetCache called: %s", req.GetKey())

	redisValue, err := h.movieService.GetDataRedis(ctx, req.Key)
	if err != nil {
		h.log.Error("GetCache - movie service error: %v", err)
		return nil, err
	}

	return &pb.GetCacheResponse{
		Value: redisValue,
	}, nil
}

func (h *handler) CreateMoviePostgres(ctx context.Context, req *pb.CreateMoviePostgresRequest) (*pb.CreateMoviePostgresResponse, error) {
	if req.Movie == nil {
		return nil, errors.New("movie is required")
	}

	if err := h.movieService.StorageMovie(ctx, req.Movie); err != nil {
		return nil, err
	}

	return &pb.CreateMoviePostgresResponse{
		Status: "Movie in Postgres created",
	}, nil
}
