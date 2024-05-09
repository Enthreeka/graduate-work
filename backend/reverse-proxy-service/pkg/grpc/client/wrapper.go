package client

import (
	pb "github.com/Entreeka/proto-proxy/go"
	"google.golang.org/grpc"
)

func NewGatewayClientWrapper(cc *grpc.ClientConn) interface{} {
	return pb.NewGatewayClient(cc)
}

func NewAggregatorClientWrapper(cc *grpc.ClientConn) interface{} {
	return pb.NewAggregatorClient(cc)
}
