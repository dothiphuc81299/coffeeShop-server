package app

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/config"

	"google.golang.org/grpc"
	orderapi "github.com/dothiphuc81299/coffeeShop-server/pkg/apis/order"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GrpcRequestTimeout = 5 * time.Second
)

func getOrderClient(cfg *config.Config) (orderapi.OrderClient, error) {
	conn, err := grpc.NewClient(cfg.GrpcClient.Order, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return orderapi.NewOrderClient(conn), nil
}
