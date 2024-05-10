package handler

import (
	"context"
	pb "github.com/Entreeka/proto-proxy/go"
)

func (h *Handler) aggregatorSearch(ctx context.Context, query string) (*pb.SearchMoviePostgresResponse, error) {
	h.Log.Info("Searching in aggregator...")

	response, err := h.ClientAggregator.SearchMoviePostgres(ctx, &pb.SearchMoviePostgresRequest{
		Query: query,
	})
	if err != nil {
		h.Log.Error("Error in searching movie in aggregator", err)
		return nil, err
	}

	return response, err
}
