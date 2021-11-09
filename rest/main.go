package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	// Update.
	pb "github.com/felipemarinho97/e-commerce/rest/api"
	"github.com/felipemarinho97/e-commerce/rest/config"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// run starts the HTTP server and blocks until the context is cancelled.
func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Get()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register REST gateway
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
