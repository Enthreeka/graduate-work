package handler

import (
	"context"
	"errors"
	"github.com/Enthreeka/elasticsearch-service/internal/service"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type movieHandler struct {
	pb.UnimplementedGatewayServer

	elasticService service.ElasticService
	log            *logger.Logger
}

func NewMovieHandler(elasticService service.ElasticService, log *logger.Logger) *movieHandler {
	return &movieHandler{
		elasticService: elasticService,
		log:            log,
	}
}

func HandleError(c codes.Code, err error) error {
	s := status.New(c, err.Error())
	return s.Err()
}

func (h *movieHandler) GetAllMovie(ctx context.Context, _ *pb.GetAllMovieRequest) (*pb.GetAllMovieResponse, error) {
	response := new(pb.GetAllMovieResponse)

	response, err := h.elasticService.GetAllDocument(ctx)
	if err != nil {
		h.log.Error("GetAllMovie: elasticService.GetAllDocument: %v", err)
		return nil, HandleError(codes.Internal, err)
	}

	return response, nil
}

func (h *movieHandler) CreateNewMovie(ctx context.Context, req *pb.CreateNewMovieRequest) (*pb.CreateNewMovieResponse, error) {
	response := new(pb.CreateNewMovieResponse)
	h.log.Info("CreateNewMovie: got message - %v", req.Movie)

	if err := h.elasticService.Index(ctx, req.Movie); err != nil {
		h.log.Error("CreateNewMovie: elasticService.Index: %v", err)
		return nil, HandleError(codes.Internal, err)
	}

	return response, nil
}

func (h *movieHandler) GetIndices(ctx context.Context, req *pb.GetIndicesRequest) (*pb.GetIndicesResponse, error) {
	response := new(pb.GetIndicesResponse)

	indicesInfoMap, err := h.elasticService.GetIndexInfo(ctx, req.GetIndicesName())
	if err != nil {
		h.log.Error("GetIndices: elasticService.GetIndexInfo: %v", err)
		return nil, HandleError(codes.Internal, err)
	}

	response.IndexInfo = &pb.IndexInfo{}
	response.IndexInfo.Indices = indicesInfoMap

	return response, nil
}

func (h *movieHandler) GetMovieByID(ctx context.Context, req *pb.GetMovieByIDRequest) (*pb.GetMovieByIDResponse, error) {
	response := new(pb.GetMovieByIDResponse)
	h.log.Info("GetMovieByID: got message - %v", req.GetMovieId())

	response, err := h.elasticService.GetDocumentByID(ctx, int(req.MovieId))
	if err != nil {
		h.log.Error("GetMovieByID: elasticService.GetDocumentByID: %v", err)
		return nil, HandleError(codes.Internal, err)
	}

	return response, nil
}

func (h *movieHandler) SearchMovie(ctx context.Context, req *pb.SearchMovieRequest) (*pb.SearchMovieResponse, error) {
	response := new(pb.SearchMovieResponse)

	if req.Query == "" {
		return nil, HandleError(codes.InvalidArgument, errors.New("bad request: query is empty"))
	}

	response, err := h.elasticService.Search(ctx, req.Query, req.Cache, req.RedisKey)
	if err != nil {
		h.log.Error("SearchMovie: elasticService.Search: %v", err)
		return nil, HandleError(codes.Internal, err)
	}

	if response.Hits.Total.Value == 0 {
		h.log.Error("SearchMovie: hits.Total.Value == 0, query: %s", req.Query)
		return nil, HandleError(codes.NotFound, errors.New("movie not found"))
	}

	h.log.Info("SearchMovie: hits.Total.Value: %d, query: %s", response.Hits.Total.Value, req.Query)
	return response, nil
}

func (h *movieHandler) UpdateMovieData(ctx context.Context, req *pb.UpdateMovieDataRequest) (*pb.UpdateMovieDataResponse, error) {
	response := new(pb.UpdateMovieDataResponse)

	return response, nil
}

func (h *movieHandler) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	response := new(pb.DeleteMovieResponse)

	return response, nil
}

func (h *movieHandler) BulkAPI(ctx context.Context, req *pb.BulkAPIRequest) (*pb.BulkAPIResponse, error) {
	response := new(pb.BulkAPIResponse)

	err := h.elasticService.IndexBulkAPI(ctx, req.Movie)
	if err != nil {
		h.log.Error("BulkAPI: elasticService.Search: %v", err)
		return nil, HandleError(codes.Internal, err)
	}

	response.Status = "Bulk api completed successfully"
	return response, nil
}
