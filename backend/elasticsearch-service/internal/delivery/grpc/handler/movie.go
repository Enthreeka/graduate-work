package handler

import (
	"context"
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

	response, err := h.elasticService.SearchAll(ctx)
	if err != nil {
		h.log.Error("CreateNewMovie: elasticService.Index: %v", err)
		return nil, HandleError(codes.Internal, err)
	}

	return response, nil
}

func (h *movieHandler) CreateNewMovie(ctx context.Context, req *pb.CreateNewMovieRequest) (*pb.CreateNewMovieResponse, error) {
	response := new(pb.CreateNewMovieResponse)
	h.log.Info("CreateNewMovie: got message - %v", req.Movie)

	if err := h.elasticService.Index(context.Background(), req.Movie); err != nil {
		h.log.Error("CreateNewMovie: elasticService.Index: %v", err)
		return nil, HandleError(codes.Internal, err)
	}

	return response, nil
}

func (h *movieHandler) GetIndices(ctx context.Context, _ *pb.GetIndicesRequest) (*pb.GetIndicesResponse, error) {
	response := new(pb.GetIndicesResponse)

	return response, nil
}

func (h *movieHandler) GetMovieByID(ctx context.Context, req *pb.GetMovieByIDRequest) (*pb.GetMovieByIDResponse, error) {
	response := new(pb.GetMovieByIDResponse)

	return response, nil
}

func (h *movieHandler) SearchMovie(ctx context.Context, req *pb.SearchMovieRequest) (*pb.SearchMovieResponse, error) {
	response := new(pb.SearchMovieResponse)

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
