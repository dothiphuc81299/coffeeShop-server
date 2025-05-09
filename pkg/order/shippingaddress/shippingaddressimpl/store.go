package shippingaddressimpl

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/infra/storage/mongodb"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/shippingaddress"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	shippingaddresscol = "shipping_addresses"
)

type store struct {
	Col *mongo.Collection
}

func NewStore(db *mongodb.Database) *store {
	return &store{
		Col: db.GetCollection(shippingaddresscol),
	}
}

func (s *store) InsertOne(ctx context.Context, u shippingaddress.UserShippingAddressRaw) error {
	_, err := s.Col.InsertOne(ctx, u)
	return err
}

func (s *store) FindOneByCondition(ctx context.Context, cond interface{}) (u shippingaddress.UserShippingAddressRaw, err error) {
	err = s.Col.FindOne(ctx, cond).Decode(&u)
	return u, err
}
