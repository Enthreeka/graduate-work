package client

import (
	"github.com/Enthreeka/proxy-service/pkg/logger"
	pb "github.com/Entreeka/proto-proxy/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayClient struct {
	log     *logger.Logger
	address string

	connection *grpc.ClientConn
}

func NewGrpcClient(log *logger.Logger, address string) *GatewayClient {
	return &GatewayClient{
		log:     log,
		address: address,
	}
}

//func (g *GatewayClient) initConnect(ctx context.Context) {
//	for {
//		select {
//		case <-ctx.Done():
//			g.log.Info("context cancel grpc connection")
//			return
//		default:
//			conn, err := grpc.Dial(
//				g.address,
//				grpc.WithTransportCredentials(insecure.NewCredentials()),
//			)
//			if err != nil {
//				g.log.Fatal("failed to dial to grpc: error = %v, address = ", err, g.address)
//			}
//			g.connection = conn
//		}
//	}
//}

func (g *GatewayClient) Connect() (interface{}, error) {
	g.log.Info("Established connections to elasticsearch-service")

	conn, err := grpc.Dial(
		g.address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		g.log.Error("failed to dial to grpc: error = %v, address = ", err, g.address)
		return nil, err
	}

	g.connection = conn

	client := pb.NewGatewayClient(conn)

	g.log.Info("Connect to grpc address = %s", g.address)
	return client, nil
}

func (g *GatewayClient) Close() {
	if g.connection != nil {
		if err := g.connection.Close(); err != nil {
			g.log.Error("failed while close dial in grpc reverse-proxy: %v", err)
		}
	}
}
