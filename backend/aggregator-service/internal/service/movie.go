package service

import (
	"context"
	"encoding/json"
	"github.com/Enthreeka/aggregator-service/internal/repo"
	"github.com/Enthreeka/aggregator-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	"github.com/pkg/errors"
)

type movieService struct {
	log          *logger.Logger
	redisRepo    repo.StorageRedis
	postgresRepo repo.StoragePostgres
}

func NewMovieService(
	redisRepo repo.StorageRedis,
	postgresRepo repo.StoragePostgres,
	log *logger.Logger,
) AggregateService {
	return &movieService{
		redisRepo:    redisRepo,
		postgresRepo: postgresRepo,
		log:          log,
	}
}

func (m *movieService) SetDataRedis(ctx context.Context, key string, value []int64) error {
	byteValue, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "failed to marshal value")
	}

	if err := m.redisRepo.Set(ctx, key, byteValue); err != nil {
		return errors.Wrap(err, "failed to set value")
	}

	return nil
}

func (m *movieService) GetDataRedis(ctx context.Context, key string) ([]int64, error) {
	byteValue, err := m.redisRepo.Get(ctx, key)
	if err != nil {
		m.log.Error("GetDataRedis: %v", err)
		return nil, errors.Wrap(err, "failed to get value")
	}

	var value []int64
	if err := json.Unmarshal(byteValue, &value); err != nil {
		m.log.Error("GetDataRedis: %v", err)
		return nil, errors.Wrap(err, "failed to unmarshal value")
	}
	return value, nil
}

func (m *movieService) GetMovie(ctx context.Context, id []int64) (*pb.TotalHits, error) {
	movie, err := m.postgresRepo.GetByArrID(ctx, id)
	if err != nil {
		m.log.Error("GetMovie: %v", err)
		return nil, errors.Wrap(err, "failed to get movie")
	}

	hits := make([]*pb.Hits, len(movie))
	totalHits := &pb.TotalHits{
		Hits: hits,
	}
	total := &pb.Total{}

	for key, value := range movie {
		hit := &pb.Hits{XSource: &pb.Movie{}}
		hit.XSource = value

		totalHits.Hits[key] = hit
		total.Value++
	}
	totalHits.Total = total

	m.log.Info("GetMovie - totalHits.Hits: %v", len(totalHits.Hits))
	return totalHits, nil
}

func (m *movieService) StorageMovie(ctx context.Context, movieArr []*pb.Movie) error {
	m.log.Info("StorageMovie - get movie:%d", len(movieArr))
	var (
		existsCount int
		addedCount  int
	)

	for _, movie := range movieArr {

		isExist, err := m.postgresRepo.IfExistsMovie(ctx, movie.Id)
		if err != nil {
			return errors.Wrap(err, "failed to check if movie exists")
		}
		if !isExist {
			movieByte, err := json.Marshal(movie)
			if err != nil {
				m.log.Error("StorageMovie: %v", err)
				return errors.Wrap(err, "failed to marshal movie")
			}

			err = m.postgresRepo.CreateMovie(ctx, movie.Id, movie.Title, movieByte)
			if err != nil {
				m.log.Error("StorageMovie: %v", err)
				return errors.Wrap(err, "failed to save movie")
			}
			addedCount++
		} else {
			existsCount++
		}
	}

	m.log.Info("StorageMovie: addedCount:%d, existsCount: %d", addedCount, existsCount)
	return nil
}
