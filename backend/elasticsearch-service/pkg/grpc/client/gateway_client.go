package client

import (
	"context"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type client func(cc *grpc.ClientConn) interface{}

type GatewayClient struct {
	log     *logger.Logger
	address string

	connection *grpc.ClientConn
	client     client
}

func NewGrpcClient(log *logger.Logger, address string, client client) *GatewayClient {
	return &GatewayClient{
		log:     log,
		address: address,
		client:  client,
	}
}

func (g *GatewayClient) Connect() (interface{}, error) {
	g.log.Info("Established connections to grpc service")

	conn, err := grpc.Dial(
		g.address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		g.log.Error("failed to dial to grpc: error = %v, address = ", err, g.address)
		return nil, err
	}

	g.connection = conn

	c := g.client(conn)

	g.log.Info("Connect to grpc address = [http://%s]", g.address)
	return c, nil
}

func (g *GatewayClient) Close() {
	if g.connection != nil {
		if err := g.connection.Close(); err != nil {
			g.log.Error("failed while close dial in grpc reverse-proxy: %v", err)
		}
	}
}

func (g *GatewayClient) Ping(ctx context.Context) {
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Minute)
		for {
			select {
			case <-ctx.Done():
				g.log.Error("context canceled while pinging grpc service")
				return
			default:
				select {
				case <-ticker.C:
					state := g.connection.GetState()
					if state != connectivity.Ready {
						g.log.Error("grpc connection is not ready: [%s], state - %s", g.address, state.String())
					}
				}
			}
		}
	}(ctx)
}
