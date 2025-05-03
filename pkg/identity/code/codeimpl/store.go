package codeimpl

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/code"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const codecol = "code-register"

type CodedRegisterDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

func NewCodedRegisterDAO(db *mongo.Database) code.CodedRegisterDAO {
	return &CodedRegisterDAO{
		DB:  db,
		Col: db.Collection(codecol),
	}
}

func (d *CodedRegisterDAO) InsertOne(ctx context.Context, u code.CodedRegisterRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

func (d *CodedRegisterDAO) DeleteOne(ctx context.Context, u string) error {
	_, err := d.Col.DeleteOne(ctx, bson.M{"email": u})
	return err
}

func (w *CodedRegisterDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc code.CodedRegisterRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
