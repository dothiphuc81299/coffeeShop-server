package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const codecol = "code-register"

// UserDAO ....
type CodedRegisterDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewUserDAO ...
func NewCodedRegisterDAO(db *mongo.Database) model.CodedRegisterDAO {
	return &CodedRegisterDAO{
		DB:  db,
		Col: db.Collection(codecol),
	}
}

// InsertOne ...
func (d *CodedRegisterDAO) InsertOne(ctx context.Context, u model.CodedRegisterRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

func (d *CodedRegisterDAO) DeleteOne(ctx context.Context, u string) error {
	_, err := d.Col.DeleteOne(ctx, bson.M{"email": u})
	return err
}

func (w *CodedRegisterDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.CodedRegisterRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}


