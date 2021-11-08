package api

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/felipemarinho97/e-commerce/config"
	pb "github.com/felipemarinho97/e-commerce/examples/go/protos/api"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEcommerceServiceServer
}

func Server(ctx context.Context) {
	cfg := config.Get()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := &server{}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	pb.RegisterEcommerceServiceServer(s, srv)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
