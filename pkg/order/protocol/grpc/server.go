package grpc

import (
	"context"
	"log"
	"net"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/apis/order"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/config"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/useraccount"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type Server struct {
	order.UnimplementedOrderServer
	dependencies *Dependencies
}

type Dependencies struct {
	UserAccountSrv useraccount.Service
	Cfg            *config.Config
}

func NewServer(deps *Dependencies) *Server {
	return &Server{
		dependencies: deps,
	}
}

func (s *Server) Run(ctx context.Context) error {
	var err error
	listen, err := net.Listen("tcp", ":"+s.dependencies.Cfg.Server.GRPCPort)
	if err != nil {
		log.Printf("Failed to create listener for grpc endpoint: %s", err.Error())
		return err
	}

	opts := []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	}

	server := grpc.NewServer(opts...)

	order.RegisterOrderServer(server, s)

	go func() {
		<-ctx.Done()
		server.GracefulStop()
		log.Println("Shutting down RPC server")
	}()

	log.Printf("GRPC server started on port %s", s.dependencies.Cfg.Server.GRPCPort)
	if err := server.Serve(listen); err != nil {
		log.Printf("Starting GRPC server failed: %s", err.Error())
	}

	return err
}
