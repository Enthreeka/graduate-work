package handler

import (
	"context"
	redisMock "github.com/Enthreeka/reverse-proxy-service/internal/repo/redis/mocks"
	"github.com/Enthreeka/reverse-proxy-service/pkg/grpc/client"
	"github.com/Enthreeka/reverse-proxy-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	pbMock "github.com/Entreeka/proto-proxy/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func BenchmarkDatabaseSearch_Test(b *testing.B) {
	ctrl := gomock.NewController(b)
	ctx := context.Background()
	log := logger.New("[backend-reverse-proxy-service]", true)

	b.Run("search in elasticsearch", func(b *testing.B) {
		redisRepoMock := redisMock.NewMockRepo(ctrl)
		redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(false, nil).AnyTimes()

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

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			req := &pb.SearchMovieRequest{Query: "harr potr"}
			_, err = h.SearchMovie(ctx, req)
			if err != nil {
				b.Error(err)
			}
		}
		b.StopTimer()
		b.ReportAllocs()
	})

	b.Run("search in postgres with unavailable elasticsearch", func(b *testing.B) {
		redisRepoMock := redisMock.NewMockRepo(ctrl)
		redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(false, nil).AnyTimes()

		elasticMock := pbMock.NewMockGatewayClient(ctrl)
		elasticMock.EXPECT().SearchMovie(gomock.Any(), gomock.Any()).Return(nil, status.Error(codes.Unavailable, "service unavailable")).AnyTimes()

		clientAggregator := client.NewGrpcClient(log, "localhost:9002", client.NewAggregatorClientWrapper)
		ca, err := clientAggregator.Connect()
		if err != nil {
			log.Fatal("failed to connect to clientElastic: %v", err)
		}

		h := Handler{
			ClientElastic:    elasticMock,
			ClientAggregator: ca.(pb.AggregatorClient),
			Log:              log,
			RedisRepo:        redisRepoMock,
		}

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			req := &pb.SearchMovieRequest{Query: "harr potr"}
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
	redisRepoMock := redisMock.NewMockRepo(ctrl)

	h := Handler{
		ClientElastic:    ce.(pb.GatewayClient),
		ClientAggregator: ca.(pb.AggregatorClient),
		Log:              log,
		RedisRepo:        redisRepoMock,
	}

	t.Run("Searching in database", func(t *testing.T) {
		t.Run("set cache", func(t *testing.T) {
			redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(false, nil).Times(1)

			_, err := h.SearchMovie(ctx, &pb.SearchMovieRequest{
				Query: "Harry Potter",
				Cache: true,
			})
			assert.Nil(t, err)
		})

		t.Run("search elastic", func(t *testing.T) {
			redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(false, nil).Times(1)

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
			redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(true, nil).Times(1)

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
