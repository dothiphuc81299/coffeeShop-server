package codeimpl

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/code"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/infra/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const codecol = "code-register"

type store struct {
	Col *mongo.Collection
}

func NewStore(db mongodb.DBConnector) *store {
	return &store{
		Col: db.GetCollection(codecol),
	}
}

func (s *store) InsertOne(ctx context.Context, u code.CodedRegisterRaw) error {
	_, err := s.Col.InsertOne(ctx, u)
	return err
}

func (s *store) DeleteOne(ctx context.Context, u string) error {
	_, err := s.Col.DeleteOne(ctx, bson.M{"email": u})
	return err
}

func (s *store) FindOneByCondition(ctx context.Context, cond interface{}) (doc code.CodedRegisterRaw, err error) {
	err = s.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
