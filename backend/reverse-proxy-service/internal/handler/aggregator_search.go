package handler

import (
	"context"
	pb "github.com/Entreeka/proto-proxy/go"
)

func (h *handler) aggregatorSearch(ctx context.Context, query string) (*pb.SearchMoviePostgresResponse, error) {
	h.log.Info("Searching in aggregator...")

	response, err := h.clientAggregator.SearchMoviePostgres(ctx, &pb.SearchMoviePostgresRequest{
		Query: query,
	})
	if err != nil {
		h.log.Error("Error in searching movie in aggregator", err)
		return nil, err
	}

	return response, err
}
