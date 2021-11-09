package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	pb "github.com/felipemarinho97/e-commerce/rest/api" // Update
	"github.com/felipemarinho97/e-commerce/rest/config"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	cfg := config.Get()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterEcommerceServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%d", cfg.GRPCHost, cfg.GRPCPort), opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), mux)
}

func main() {
	cfg := config.Get()

	log.Printf("server listening at %v", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err := run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
