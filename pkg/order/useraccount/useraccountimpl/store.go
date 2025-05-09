package useraccountimpl

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/infra/storage/mongodb"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/useraccount"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const usercol = "user_account"

type store struct {
	Col *mongo.Collection
}

func NewStore(db *mongodb.Database) *store {
	return &store{
		Col: db.GetCollection(usercol),
	}
}

func (s *store) InsertOne(ctx context.Context, u useraccount.UserAccountRaw) error {
	_, err := s.Col.InsertOne(ctx, u)
	return err
}

func (s *store) FindOneByCondition(ctx context.Context, cond interface{}) (u useraccount.UserAccountRaw, err error) {
	err = s.Col.FindOne(ctx, cond).Decode(&u)
	return u, err
}

func (s *store) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []useraccount.UserAccountRaw, err error) {
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

func (s *store) UpdateOne(ctx context.Context, filter interface{}, update interface{}) error {
	_, err := s.Col.UpdateOne(ctx, filter, update)
	return err
}
