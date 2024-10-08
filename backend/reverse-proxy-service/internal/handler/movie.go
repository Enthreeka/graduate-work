package handler

import (
	"context"
	"github.com/Enthreeka/reverse-proxy-service/internal/repo/redis"
	"github.com/Enthreeka/reverse-proxy-service/pkg/logger"
	"github.com/Enthreeka/reverse-proxy-service/pkg/metric"
	"github.com/Enthreeka/reverse-proxy-service/pkg/utils"
	pb "github.com/Entreeka/proto-proxy/go"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
)

const (
	getMovieApi    = "/v1/api/movie/get/"
	searchMovieApi = "/v1/api/movie/search"
	bulkMovieApi   = "/v1/api/movie/bulk"
	createMovieApi = "/v1/api/movie/create"
)

type handler struct {
	pb.UnimplementedGatewayServer
	pb.UnimplementedAggregatorServer

	log              *logger.Logger
	redisRepo        redis.Repo
	clientElastic    pb.GatewayClient
	clientAggregator pb.AggregatorClient
}

func NewHandler(log *logger.Logger,
	redisRepo redis.Repo,
	clientElastic pb.GatewayClient,
	clientAggregator pb.AggregatorClient,
) *handler {
	return &handler{
		log:              log,
		redisRepo:        redisRepo,
		clientElastic:    clientElastic,
		clientAggregator: clientAggregator,
	}
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

		return s, SwitchToGrpcStatus(grpcStatusCode, s.Err())
	}

	return nil, nil
}

func (h *handler) GetAllMovie(ctx context.Context, _ *pb.GetAllMovieRequest) (*pb.GetAllMovieResponse, error) {
	response := new(pb.GetAllMovieResponse)

	response, err := h.clientElastic.GetAllMovie(ctx, &pb.GetAllMovieRequest{})
	if s, err := ErrorWrapper(err); err != nil {
		h.log.Error("GetAllMovie: error: %v, message: %s, ", err, s.Message())
		return nil, err
	}

	h.log.Info("Get all movie by client elastic completed successfully")
	return response, nil
}

func (h *handler) CreateNewMovie(ctx context.Context, req *pb.CreateNewMovieRequest) (*pb.CreateNewMovieResponse, error) {
	response := new(pb.CreateNewMovieResponse)
	h.log.Info("CreateNewMovie request: %v", req.Movie)

	if err := req.Validate(); err != nil {
		h.log.Error("CreateNewMovie: validation error: %v", err)
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	response, err := h.clientElastic.CreateNewMovie(ctx, req)
	if s, err := ErrorWrapper(err); err != nil {
		h.log.Error("CreateNewMovie: error: %v, message: %s, ", err, s.Message())
		return nil, err
	}

	metric.CreateMovieHits.WithLabelValues(createMovieApi).Add(1)

	h.log.Info("Creating movie in client elastic completed successfully")
	return response, nil
}

func (h *handler) GetIndices(ctx context.Context, req *pb.GetIndicesRequest) (*pb.GetIndicesResponse, error) {
	response := new(pb.GetIndicesResponse)

	response, err := h.clientElastic.GetIndices(ctx, req)
	if s, err := ErrorWrapper(err); err != nil {
		h.log.Error("CreateNewMovie: error: %v, message: %v, internal_error: %v", err, s.Message(), s.Err())
		return nil, err
	}

	return response, nil
}

// https://blog.logrocket.com/guide-to-grpc-gateway/
func (h *handler) GetMovieByID(ctx context.Context, _ *pb.GetMovieByIDRequest) (*pb.GetMovieByIDResponse, error) {
	response := new(pb.GetMovieByIDResponse)

	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		h.log.Error("metadata from incoming context is empty in GetMovieByID")
		return nil, SwitchToGrpcStatus(http.StatusInternalServerError, "metadata from incoming context is empty")
	}

	url := md.Get("url")[0]
	_, movieID, exist := strings.Cut(url, getMovieApi)
	if !exist {
		h.log.Error("in url = [%s] not found some value", url)
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	movieIDInt, err := strconv.Atoi(movieID)
	if err != nil {
		h.log.Error("movieID not [int] data type: %v", err)
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	h.log.Info("GetMovieByID: redirect to elasticsearch-service - movie_id: %d", movieIDInt)
	response, err = h.clientElastic.GetMovieByID(ctx, &pb.GetMovieByIDRequest{
		MovieId: int64(movieIDInt),
	})
	if s, err := ErrorWrapper(err); err != nil {
		h.log.Error("GetMovieByID: error: %v, message: %s, ", err, s.Message())
		return nil, err
	}

	h.log.Info("Getting movie by id from client elastic completed successfully")
	return response, nil
}

func (h *handler) SearchMovie(ctx context.Context, req *pb.SearchMovieRequest) (*pb.SearchMovieResponse, error) {
	response := new(pb.SearchMovieResponse)

	if req.GetQuery() == "" {
		h.log.Error("SearchMovie: empty query request")
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	h.log.Info("SearchMovie: got message - %v", req)

	redisKey := utils.KeyGenerator(searchMovieApi, req.Query)
	isExist, err := h.redisRepo.IsExist(ctx, redisKey)
	if err != nil {
		h.log.Error("SearchMovie: failed to search movie in redis - error: %v, movie_exist: %v", err, isExist)
	}

	h.log.Info("in redis - %v, key - %s", isExist, redisKey)
	// Проверка на наличие кэша в Redis
	if !isExist || err != nil {
		if req.Cache == true {
			req.RedisKey = redisKey
		}
		// В случае отсутствия в кэше идет поиск в Elasticsearch
		response, err = h.clientElastic.SearchMovie(ctx, req)
		if s, err := ErrorWrapper(err); err != nil {
			if s.Code() == codes.Unavailable {
				h.log.Info("Searching in aggregator...")

				// Если elasticsearch-service недоступен, то ищем в PostgreSQL в aggregator-service
				responsePostgres, err := h.aggregatorSearch(ctx, req.Query)
				if _, err := ErrorWrapper(err); err != nil {
					h.log.Error("SearchMovie in postgres: error: %v", err)
					return nil, err
				}

				if responsePostgres.Hits.Total.Value == 0 || responsePostgres.GetHits() == nil {
					h.log.Error("empty response from aggregator")
					return nil, SwitchToGrpcStatus(http.StatusBadRequest)
				}

				metric.ResponseSearchMovie.WithLabelValues("aggregator-service").Add(float64(responsePostgres.Hits.Total.Value))

				return &pb.SearchMovieResponse{
					Hits:   responsePostgres.Hits,
					Status: "aggregator-service",
				}, nil
			}

			h.log.Error("SearchMovie: error in elasticsearch-service: %v", err)
			return nil, err
		}
		response.Status = "elasticsearch-service"
	} else {
		// В случае наличия кеша, отправляемся в aggregator-service, где сначала получим значения из кеша,
		// потом будем искать через where id = $1 в PostgreSQL
		responseAggregator, err := h.clientAggregator.SearchMovieAggregator(ctx, &pb.SearchMovieAggregatorRequest{
			RedisKey: redisKey,
			Query:    req.Query,
		})
		if _, err := ErrorWrapper(err); err != nil {
			h.log.Error("SearchMovie: error: %v", err)
			return nil, err
		}
		response.Hits = responseAggregator.Hits
		response.Status = "aggregator-service"
	}

	metric.ResponseSearchMovie.WithLabelValues(response.Status).Add(float64(response.Hits.Total.Value))

	return response, nil
}

func (h *handler) UpdateMovieData(ctx context.Context, req *pb.UpdateMovieDataRequest) (*pb.UpdateMovieDataResponse, error) {
	response := new(pb.UpdateMovieDataResponse)

	if err := req.Validate(); err != nil {
		h.log.Error("UpdateMovieData: validation error: %v", err)
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	response, err := h.clientElastic.UpdateMovieData(ctx, req)
	if s, err := ErrorWrapper(err); err != nil {
		h.log.Error("UpdateMovieData: error: %v, message: %s, ", err, s.Message())
		return nil, err
	}

	return response, nil
}

func (h *handler) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	response := new(pb.DeleteMovieResponse)

	if req.MovieId == 0 && req.Title == "" {
		h.log.Error("DeleteMovie: empty data in request")
		return nil, SwitchToGrpcStatus(http.StatusBadRequest)
	}

	response, err := h.clientElastic.DeleteMovie(ctx, req)
	if s, err := ErrorWrapper(err); err != nil {
		h.log.Error("DeleteMovie: error: %v, message: %s, ", err, s.Message())
		return nil, err
	}

	return response, nil
}

func (h *handler) BulkAPI(ctx context.Context, req *pb.BulkAPIRequest) (*pb.BulkAPIResponse, error) {
	response := new(pb.BulkAPIResponse)
	var builder strings.Builder

	var g errgroup.Group
	g.Go(func() error {
		res, err := h.clientElastic.BulkAPI(ctx, req)
		if err != nil {
			h.log.Error("BulkAPI: error: %v", err)
			return err
		}

		builder.WriteString("Elasticsearch-service: ")
		builder.WriteString(res.Status)
		builder.WriteString("|")
		return nil
	})

	g.Go(func() error {
		res, err := h.clientAggregator.CreateMoviePostgres(ctx, &pb.CreateMoviePostgresRequest{
			Movie: req.Movie,
		})
		if err != nil {
			h.log.Error("CreateMoviePostgres: error: %v", err)
			return err
		}

		builder.WriteString("Aggregator-service: ")
		builder.WriteString(res.Status)
		builder.WriteString("|")
		return nil
	})

	err := g.Wait()
	if _, err := ErrorWrapper(err); err != nil {
		h.log.Error("CreateMoviePostgres: error: %v, success service: %s, ", err, builder.String())

		errMsg := err.Error()
		if !strings.Contains(errMsg, "Elasticsearch-service") && !strings.Contains(errMsg, "Aggregator-service") {
			return nil, err
		}
	}

	metric.CreateMovieHits.WithLabelValues(bulkMovieApi).Add(float64(len(req.Movie)))

	response.Status = builder.String()
	return response, nil
}
