package grpc

import (
	"context"

	"time"

	orderapi "github.com/dothiphuc81299/coffeeShop-server/pkg/apis/order"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/useraccount"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	GrpcRequestTimeout = 5 * time.Second
)

func (s *Server) CreateUserAccount(ctx context.Context, req *orderapi.CreateUserAccountCommand) (*orderapi.CreateUserAccountResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, GrpcRequestTimeout)
	defer cancel()

	userID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, err
	}

	cmd := &useraccount.CreateUserAccountCommand{
		UserID:    userID,
		LoginName: req.LoginName,
		Active:    req.Active,
	}

	err = s.dependencies.UserAccountSrv.Create(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &orderapi.CreateUserAccountResponse{}, nil
}
