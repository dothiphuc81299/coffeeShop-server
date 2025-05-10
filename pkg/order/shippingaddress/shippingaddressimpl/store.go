package shippingaddressimpl

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/infra/storage/mongodb"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/shippingaddress"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s *store) UpdateOne(ctx context.Context, filter interface{}, update interface{}) error {
	_, err := s.Col.UpdateOne(ctx, filter, update)
	return err
}

func (s *store) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	_, err := s.Col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (s *store) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []shippingaddress.UserShippingAddressRaw, err error) {
	cursor, err := s.Col.Find(ctx, cond, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	return docs, nil
}

func (s *store) CountByCondition(ctx context.Context, cond interface{}) (int64, error) {
	return s.Col.CountDocuments(ctx, cond)
}
