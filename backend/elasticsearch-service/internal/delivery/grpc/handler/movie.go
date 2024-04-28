package handler

import (
	"context"
	"github.com/Enthreeka/elasticsearch-service/internal/service"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
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

func (m *movieHandler) GetSearchInfo(context.Context, *pb.GetSearchInfoRequest) (*pb.GetSearchInfoResponse, error) {
	response := new(pb.GetSearchInfoResponse)
	return response, nil
}
