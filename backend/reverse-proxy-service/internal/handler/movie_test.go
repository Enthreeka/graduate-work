package handler

import (
	"context"
	"github.com/stretchr/testify/assert"

	redisMock "github.com/Enthreeka/reverse-proxy-service/internal/repo/redis/mocks"
	"github.com/Enthreeka/reverse-proxy-service/pkg/grpc/client"
	"github.com/Enthreeka/reverse-proxy-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	"go.uber.org/mock/gomock"
	"testing"
)

func BenchmarkSearch_Test(b *testing.B) {
	ctrl := gomock.NewController(b)

	redisRepoMock := redisMock.NewMockRepo(ctrl)

	log := logger.New("[backend-reverse-proxy-service]", true)

	clientElastic := client.NewGrpcClient(log, "localhost:9001", client.NewGatewayClientWrapper)
	ce, err := clientElastic.Connect()
	if err != nil {
		log.Fatal("failed to connect to clientElastic: %v", err)
	}

	clientAggregator := client.NewGrpcClient(log, "localhost:9002", client.NewAggregatorClientWrapper)
	ca, err := clientAggregator.Connect()
	if err != nil {
		log.Fatal("failed to connect to clientElastic: %v", err)
	}

	h := Handler{
		ClientElastic:    ce.(pb.GatewayClient),
		ClientAggregator: ca.(pb.AggregatorClient),
		Log:              log,
		RedisRepo:        redisRepoMock,
	}

	ctx := context.Background()
	b.Run("search in cache", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(true, nil).Times(1)
			req := &pb.SearchMovieRequest{Query: "harry"}
			_, err = h.SearchMovie(ctx, req)
			if err != nil {
				b.Error(err)
			}
		}
		b.StopTimer()
		b.ReportAllocs()
	})

	b.Run("search in elasticsearch", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(false, nil).Times(1)
			req := &pb.SearchMovieRequest{Query: "harry"}
			_, err = h.SearchMovie(ctx, req)
			if err != nil {
				b.Error(err)
			}
		}
		b.StopTimer()
		b.ReportAllocs()
	})
}

func TestSearchMovie(t *testing.T) {
	ctrl := gomock.NewController(t)

	log := logger.New("[backend-reverse-proxy-service]", true)

	clientElastic := client.NewGrpcClient(log, "localhost:9001", client.NewGatewayClientWrapper)
	ce, err := clientElastic.Connect()
	if err != nil {
		log.Fatal("failed to connect to clientElastic: %v", err)
	}

	clientAggregator := client.NewGrpcClient(log, "localhost:9002", client.NewAggregatorClientWrapper)
	ca, err := clientAggregator.Connect()
	if err != nil {
		log.Fatal("failed to connect to clientElastic: %v", err)
	}

	ctx := context.Background()

	t.Run("Searching in database", func(t *testing.T) {
		t.Run("set cache", func(t *testing.T) {
			redisRepoMock := redisMock.NewMockRepo(ctrl)
			redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(false, nil).AnyTimes()

			h := Handler{
				ClientElastic:    ce.(pb.GatewayClient),
				ClientAggregator: ca.(pb.AggregatorClient),
				Log:              log,
				RedisRepo:        redisRepoMock,
			}

			_, err := h.SearchMovie(ctx, &pb.SearchMovieRequest{
				Query: "Harry Potter",
				Cache: true,
			})
			assert.Nil(t, err)
		})

		t.Run("search elastic", func(t *testing.T) {
			redisRepoMock := redisMock.NewMockRepo(ctrl)
			redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(false, nil).AnyTimes()

			h := Handler{
				ClientElastic:    ce.(pb.GatewayClient),
				ClientAggregator: ca.(pb.AggregatorClient),
				Log:              log,
				RedisRepo:        redisRepoMock,
			}

			_, err := h.SearchMovie(ctx, &pb.SearchMovieRequest{Query: "Harry Potter"})
			assert.Nil(t, err)
		})

		t.Run("search aggregator", func(t *testing.T) {
			redisRepoMock := redisMock.NewMockRepo(ctrl)
			redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()

			h := Handler{
				ClientElastic:    ce.(pb.GatewayClient),
				ClientAggregator: ca.(pb.AggregatorClient),
				Log:              log,
				RedisRepo:        redisRepoMock,
			}

			_, err := h.SearchMovie(ctx, &pb.SearchMovieRequest{Query: "Harry Potter"})
			assert.Nil(t, err)
		})
	})
}
