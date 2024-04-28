package main

import (
	"context"
	"fmt"
	pb "github.com/Entreeka/proto-proxy/go"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedGatewayServer
}

func (s *server) GetSearchInfo(context.Context, *pb.GetSearchInfoRequest) (*pb.GetSearchInfoResponse, error) {
	fmt.Println("TEST")
	return &pb.GetSearchInfoResponse{Movie: nil}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGatewayServer(s, &server{})

	fmt.Println("Server listening on port 9001")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
