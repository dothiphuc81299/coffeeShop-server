package app

import (
	"context"
	"log"
	"net/http"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/infra/storage/mongodb"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/category/categoryimpl"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/config"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/drink/drinkimpl"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/protocol/rest"
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

	categorySrv := categoryimpl.NewService(categoryStore, drinkStore)
	drinkSrv := drinkimpl.NewService(drinkStore, categoryStore)

	restServer := rest.NewServer(&rest.Dependences{
		CategorySrv: categorySrv,
		DrinkSrv:    drinkSrv,
	}, cfg)

	go func() {
		if err := restServer.Run(context.Background()); err != nil && err != http.ErrServerClosed {
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
