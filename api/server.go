package api

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/felipemarinho97/e-commerce/common"
	"github.com/felipemarinho97/e-commerce/config"
	"github.com/felipemarinho97/e-commerce/db"
	pb "github.com/felipemarinho97/e-commerce/examples/go/protos/api"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEcommerceServiceServer
	db db.Database
}

func Server(ctx context.Context) {
	cfg := config.Get()

	database, err := db.New()
	if err != nil {
		common.Logger.LogFatal("Server", err.Error())
		os.Exit(-1)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		common.Logger.LogFatal(fmt.Sprintf("failed to listen: %v", err))
		os.Exit(-1)
	}
	srv := &server{
		db: database,
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	pb.RegisterEcommerceServiceServer(s, srv)
	common.Logger.LogInfo(fmt.Sprintf("server listening at %v", lis.Addr()))

	if err := s.Serve(lis); err != nil {
		common.Logger.LogFatal(fmt.Sprintf("failed to serve: %v", err))
		os.Exit(-1)
	}
}
