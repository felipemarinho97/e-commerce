package api

import (
	"context"

	pb "github.com/felipemarinho97/e-commerce/examples/go/protos/api"
)

func (s server) Checkout(ctx context.Context, in *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	return &pb.CheckoutResponse{}, nil
}
