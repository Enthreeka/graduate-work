package handler

import (
	"context"
	redisMock "github.com/Enthreeka/reverse-proxy-service/internal/repo/redis/mocks"
	"github.com/Enthreeka/reverse-proxy-service/pkg/grpc/client"
	"github.com/Enthreeka/reverse-proxy-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	pbMock "github.com/Entreeka/proto-proxy/mock"
	"go.uber.org/mock/gomock"
	"sync"
	"testing"
)

func BenchmarkSearch_Test(b *testing.B) {
	ctrl := gomock.NewController(b)

	mockAggClient := pbMock.NewMockAggregatorClient(ctrl)
	redisRepoMock := redisMock.NewMockRepo(ctrl)
	redisRepoMock.EXPECT().IsExist(gomock.Any(), gomock.Any()).Return(false, nil).AnyTimes()

	log := logger.New("[backend-reverse-proxy-service]")

	clientElastic := client.NewGrpcClient(log, "localhost:9001", client.NewGatewayClientWrapper)
	ce, err := clientElastic.Connect()
	if err != nil {
		log.Fatal("failed to connect to clientElastic: %v", err)
	}

	h := Handler{
		ClientElastic:    ce.(pb.GatewayClient),
		ClientAggregator: mockAggClient,
		Log:              log,
		RedisRepo:        redisRepoMock,
	}

	b.Run("with sync pool", func(b *testing.B) {
		var movieRequestPool = sync.Pool{
			New: func() interface{} {
				return &pb.SearchMovieRequest{}
			},
		}
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			req := movieRequestPool.Get().(*pb.SearchMovieRequest)
			req.Query = "harr"

			_, _ = h.SearchMovie(context.Background(), req)
			movieRequestPool.Put(req)
		}
		b.StopTimer()
		b.ReportAllocs()
	})

	b.Run("without sync pool", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			req := &pb.SearchMovieRequest{Query: "harr"}
			_, _ = h.SearchMovie(context.Background(), req)
		}
		b.StopTimer()
		b.ReportAllocs()
	})
}
