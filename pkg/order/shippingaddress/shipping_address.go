package shippingaddress

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store interface {
	InsertOne(ctx context.Context, u UserShippingAddressRaw) error
	FindOneByCondition(ctx context.Context, cond interface{}) (u UserShippingAddressRaw, err error)
}

type Service interface {
	Create(ctx context.Context, cmd *CreateShippingAddressCommand) error
	Search(ctx context.Context, query *SearchShippingAddressQuery) ([]UserShippingAddressRaw, int64, error)
	GetDetail(ctx context.Context, id primitive.ObjectID) (UserShippingAddressRaw, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	Update(ctx context.Context, cmd *UpdateShippingAddressCommand) error
}
