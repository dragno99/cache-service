package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	proto "github.com/dragno99/cache-service/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAddServiceServer
}

func (s *server) Get(ctx context.Context, req *proto.Key) (*proto.Value, error) {

	return &proto.Value{Value: []byte("Suryansh will implement me")}, nil
}

func (s *server) Set(ctx context.Context, req *proto.KeyVal) (*proto.Empty, error) {

	return &proto.Empty{}, errors.New("not implemented yet. Suryansh will implement me")
}

func StartServer() {
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	proto.RegisterAddServiceServer(srv, &server{})

	fmt.Println("server started...")

	if err := srv.Serve(listner); err != nil {
		log.Fatal("failed to serve: %v", err.Error())
	}

}
