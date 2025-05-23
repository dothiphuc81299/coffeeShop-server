package app

import (
	"context"
	"log"
	"net/http"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/infra/storage/mongodb"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/category/categoryimpl"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/config"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/drink/drinkimpl"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/order/orderimpl"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/protocol/grpc"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/protocol/rest"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/shippingaddress/shippingaddressimpl"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/useraccount/useraccountimpl"
)

type Server struct {
	MongoDB    *mongodb.Database
	cfg        *config.Config
	RestServer *rest.Server
}

const serviceName string = "order"

func NewServer() (*Server, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	mongoDB, err := mongodb.New(cfg.MongoDB.URI, cfg.MongoDB.DBName)
	if err != nil {
		return nil, err
	}

	categoryStore := categoryimpl.NewStore(mongoDB)
	drinkStore := drinkimpl.NewStore(mongoDB)
	userStore := useraccountimpl.NewStore(mongoDB)
	shippingStore := shippingaddressimpl.NewStore(mongoDB)
	orderStore := orderimpl.NewStore(mongoDB)

	categorySrv := categoryimpl.NewService(categoryStore, drinkStore)
	drinkSrv := drinkimpl.NewService(drinkStore, categoryStore)
	userSrv := useraccountimpl.NewService(userStore)
	ordersrv := orderimpl.NewService(orderStore, drinkStore, userStore, shippingStore)
	shippingSrv := shippingaddressimpl.NewService(shippingStore, userStore)

	restServer := rest.NewServer(&rest.Dependences{
		CategorySrv: categorySrv,
		DrinkSrv:    drinkSrv,
		OrderSrv:    ordersrv,
		ShippingSrv: shippingSrv,
	}, cfg)

	grpcServer := grpc.NewServer(&grpc.Dependencies{
		Cfg:            cfg,
		UserAccountSrv: userSrv,
	})

	go func() {
		if err := restServer.Run(context.Background()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Failed to start server: %v", err)
		}
	}()

	go func() {
		if err := grpcServer.Run(context.Background()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Failed to start server: %v", err)
		}
	}()

	return &Server{
		MongoDB: mongoDB,
		cfg:     cfg,
	}, nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.RestServer != nil {
		return s.RestServer.Shutdown(ctx)
	}
	return nil
}
